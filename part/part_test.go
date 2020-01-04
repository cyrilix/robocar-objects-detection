package part

import (
	"github.com/cyrilix/robocar-base/testtools"
	"github.com/cyrilix/robocar-protobuf/go/events"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"gocv.io/x/gocv"
	"io/ioutil"
	"sync"
	"testing"
	"time"
)

type FakeObjectDetector struct {
	muResult sync.Mutex
	result   *[]*events.Object
}

func (f *FakeObjectDetector) DetectObjects(_ *gocv.Mat) (*[]*events.Object, error) {
	f.muResult.Lock()
	defer f.muResult.Unlock()
	return f.result, nil
}

func (f *FakeObjectDetector) SetResult(r *[]*events.Object) {
	f.muResult.Lock()
	defer f.muResult.Unlock()
	f.result = r
}

func (f *FakeObjectDetector) Close() error {
	return nil
}

func TestObjectDetectPart(t *testing.T) {
	oldRegister := registerCallBacks
	oldPublish := publish
	defer func() {
		registerCallBacks = oldRegister
		publish = oldPublish
	}()

	registerCallBacks = func(_ *ObjectDetectPart) {}

	var muEventsPublished sync.Mutex
	eventsPublished := make(map[string][]byte)
	publish = func(client mqtt.Client, topic string, payload *[]byte) {
		muEventsPublished.Lock()
		defer muEventsPublished.Unlock()
		eventsPublished[topic] = *payload
	}


	cameraTopic := "topic/camera"
	objectsTopic := "topic/objects"
	detector := FakeObjectDetector{}
	part := New(nil, cameraTopic, objectsTopic, &detector)

	expectedObjects := []*events.Object{{Type: events.TypeObject_CAR, Left: 1, Top: 2, Right: 3, Bottom: 4, Confidence: 0.5},}

	msg := testtools.NewFakeMessageFromProtobuf("topic/camera", loadImage(t, "testdata/image.jpg"))
	detector.SetResult(&expectedObjects)

	go func() {
		if err := part.Start(); err != nil {
			t.Fatalf("unable to start instance: %v", err)
		}
	}()
	//defer part.Stop()
	time.Sleep(50 * time.Millisecond)
	part.OnFrame(nil, msg)
	time.Sleep(100 * time.Millisecond)

	var objects events.ObjectsMessage
	muEventsPublished.Lock()
	payload := eventsPublished[objectsTopic]
	muEventsPublished.Unlock()

	err := proto.Unmarshal(payload, &objects)
	if err != nil {
		t.Errorf("unable to unmarshall %T msg: %v", objects, err)
	}
	if len(objects.GetObjects()) != 1 {
		t.Errorf("incorrect object number found: %v, wants %v", len(objects.GetObjects()), len(expectedObjects))
	}

	for idx, object := range objects.GetObjects() {
		if (*object).String() != expectedObjects[idx].String() {
			t.Errorf("invalid object at %v/%v: %v, wants %v", idx, len(expectedObjects), object, expectedObjects[idx])
		}
	}
}

func loadImage(t *testing.T, imgPath string) *events.FrameMessage {
	jpegContent, err := ioutil.ReadFile(imgPath)
	if err != nil {
		t.Fatalf("unable to load image: %v", err)
	}
	msg := &events.FrameMessage{
		Id:                   &events.FrameRef{
			Name:                 imgPath,
			Id:                   "01",
		},
		Frame:                jpegContent,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	return msg
}
