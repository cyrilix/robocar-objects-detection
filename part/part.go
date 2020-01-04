package part

import (
	"github.com/cyrilix/robocar-base/service"
	"github.com/cyrilix/robocar-protobuf/go/events"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"gocv.io/x/gocv"
)

func New(client mqtt.Client, cameraTopic, objectsTopic string, detector ObjectDetector) *ObjectDetectPart {
	return &ObjectDetectPart{
		client:         client,
		frameChan:      make(chan frameToProcess),
		readyForNext:   make(chan interface{}, 1),
		cancel:         make(chan interface{}),
		objectDetector: detector,
		cameraTopic:    cameraTopic,
		objectsTopic:   objectsTopic,
	}
}

type ObjectDetectPart struct {
	client       mqtt.Client
	frameChan    chan frameToProcess
	readyForNext chan interface{}
	cancel       chan interface{}

	objectDetector ObjectDetector

	cameraTopic, objectsTopic string
}

func (o *ObjectDetectPart) OnFrame(_ mqtt.Client, msg mqtt.Message) {
	var frameMsg events.FrameMessage
	err := proto.Unmarshal(msg.Payload(), &frameMsg)
	if err != nil {
		log.Errorf("unable to unmarshal %T message: %v", frameMsg, err)
		return
	}

	img, err := gocv.IMDecode(frameMsg.GetFrame(), gocv.IMReadUnchanged)
	if err != nil {
		log.Errorf("unable to decode image: %v", err)
		return
	}
	frame := frameToProcess{
		ref: frameMsg.GetId(),
		Mat: img,
	}
	o.frameChan <- frame
}

type frameToProcess struct {
	ref *events.FrameRef
	gocv.Mat
}

func (o *ObjectDetectPart) Start() error {
	registerCallBacks(o)

	ready := true
	var frame = frameToProcess{}
	defer func() {
		if err := frame.Close(); err != nil {
			log.Errorf("unable to close frame: %v", err)
		}
	}()

	for {
		select {
		case f := <-o.frameChan:
			log.Debug("new frame")
			oldFrame := frame
			frame = f
			if err := oldFrame.Close(); err != nil {
				log.Errorf("unable to close frame: %v", err)
			}
			if ready {
				log.Debug("process frame")
				go o.processFrame(frame)
				ready = false
			}
		case <-o.readyForNext:
			ready = true
		case <-o.cancel:
			log.Infof("Stop service")
			return nil
		}
	}
}

func (o *ObjectDetectPart) Stop() {
	defer func() {
		if err := o.objectDetector.Close(); err != nil {
			log.Errorf("unable to close objectDetector: %v", err)
		}
	}()
	close(o.readyForNext)
	close(o.cancel)
	service.StopService("object-detection", o.client, o.cameraTopic)
}

func (o *ObjectDetectPart) processFrame(frame frameToProcess) {
	defer func() { o.readyForNext <- struct{}{} }()

	bboxes, err := o.objectDetector.DetectObjects(&frame.Mat)
	if err != nil {
		log.Errorf("unable to detect object into image: %v", err)
		return
	}

	msg := events.ObjectsMessage{
		Objects:              *bboxes,
		FrameRef:             frame.ref,
	}

	payload, err := proto.Marshal(&msg)
	if err != nil {
		log.Errorf("unable to mashal protobuf %T message: %v", msg, err)
		return
	}
	publish(o.client, o.objectsTopic, &payload)
}

var registerCallBacks = func(o *ObjectDetectPart) {
	err := service.RegisterCallback(o.client, o.cameraTopic, o.OnFrame)
	if err != nil {
		log.Panicf("unable to register callback to topic %v:%v", o.cameraTopic, err)
	}
}

var publish = func(client mqtt.Client, topic string, payload *[]byte) {
	client.Publish(topic, 0, false, *payload)
}
