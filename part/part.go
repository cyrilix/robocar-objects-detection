package part

import (
	"github.com/cyrilix/robocar-base/service"
	"github.com/cyrilix/robocar-objects-detection/objects"
	"github.com/cyrilix/robocar-protobuf/go/events"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"gocv.io/x/gocv"
	"log"
	"sync"
)

func New(client mqtt.Client, disparityTopic, objectsTopic, objectCleanTopic string, processor objects.Processor) *ObjectDetectPart {
	return &ObjectDetectPart{
		client:            client,
		publishChan:       make(chan interface{}, 10),
		cancel:            make(chan interface{}),
		processor:         processor,
		disparityTopic:    disparityTopic,
		objectsTopic:      objectsTopic,
		objectsCleanTopic: objectCleanTopic,
	}
}

type ObjectDetectPart struct {
	client      mqtt.Client
	publishChan chan interface{}
	cancel      chan interface{}

	muObjects sync.RWMutex
	objects   []*events.Object
	frameRef  *events.FrameRef

	muDisparity sync.RWMutex
	disparity   *objects.Disparity

	processor objects.Processor

	disparityTopic, objectsTopic, objectsCleanTopic string
}

func (o *ObjectDetectPart) onObjects(_ mqtt.Client, message mqtt.Message) {
	var msg events.ObjectsMessage
	err := proto.Unmarshal(message.Payload(), &msg)
	if err != nil {
		zap.S().Errorf("unable to unmarshal protobuf %T message: %v", msg, err)
		return
	}

	o.muObjects.Lock()
	defer o.muObjects.Unlock()
	o.objects = msg.GetObjects()
	o.frameRef = msg.GetFrameRef()
	zap.S().Debugf("%v object(s) received", len(o.objects))
	o.publishChan <- struct{}{}
}

func (o *ObjectDetectPart) onDisparityMessage(_ mqtt.Client, msg mqtt.Message) {
	var disparityMsg events.DisparityMessage
	err := proto.Unmarshal(msg.Payload(), &disparityMsg)
	if err != nil {
		zap.S().Errorf("unable to unmarshal %T message: %v", disparityMsg, err)
		return
	}

	img, err := gocv.IMDecode(disparityMsg.GetDisparity(), gocv.IMReadUnchanged)
	if err != nil {
		zap.S().Errorf("unable to decode image: %v", err)
		return
	}
	disparity := objects.Disparity{
		Ref:                 disparityMsg.GetFrameRef(),
		Mat:                 img,
		Baseline:            disparityMsg.GetBaselineInMm(),
		FocalLengthInPixels: disparityMsg.GetFocalLengthInPixels(),
	}
	o.muDisparity.Lock()
	defer o.muDisparity.Unlock()
	if o.disparity != nil {
		o.disparity.Close()
	}
	o.disparity = &disparity
	o.publishChan <- struct{}{}
}

func (o *ObjectDetectPart) Start() error {
	registerCallBacks(o)
	for {
		select {
		case <-o.publishChan:
			o.publishObject()
		case <-o.cancel:
			zap.S().Infof("Stop service")
			return nil
		}
	}
}

func (o *ObjectDetectPart) Stop() {
	close(o.cancel)
	service.StopService("object-detection", o.client, o.disparityTopic)
}

func (o *ObjectDetectPart) publishObject() {
	o.muObjects.Lock()
	defer o.muObjects.Unlock()
	o.muDisparity.Lock()
	defer o.muDisparity.Unlock()

	objs, err := o.processor.Process(o.objects, o.disparity)
	message := events.ObjectsMessage{
		Objects:  objs,
		FrameRef: o.frameRef,
	}
	payload, err := proto.Marshal(&message)
	if err != nil {
		zap.S().Errorf("unable to unmarshal protobuf %T message: %v", payload, err)
		return
	}
	publish(o.client, o.objectsCleanTopic, payload)
}

var registerCallBacks = func(o *ObjectDetectPart) {
	err := service.RegisterCallback(o.client, o.disparityTopic, o.onDisparityMessage)
	if err != nil {
		log.Panicf("unable to register callback to topic %v:%v", o.disparityTopic, err)
	}
	err = service.RegisterCallback(o.client, o.objectsTopic, o.onObjects)
	if err != nil {
		log.Panicf("unable to register callback to topic %v:%v", o.objectsTopic, err)
	}
}

var publish = func(client mqtt.Client, topic string, payload []byte) {
	zap.S().Debugf("publish to %s", topic)
	client.Publish(topic, 0, false, payload)
}
