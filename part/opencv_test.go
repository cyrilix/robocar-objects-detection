package part

import (
	"github.com/cyrilix/robocar-protobuf/go/events"
	log "github.com/sirupsen/logrus"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"io/ioutil"
	"testing"
)

func TestCarDetector_SearchCar(t *testing.T) {

	detector := NewCarDetector("testdata/ssd_mobilenet_v2_coco_2018_03_29/frozen_inference_graph.pb", "testdata/ssd_mobilenet_v2_coco_2018_03_29.pbtxt")
	defer func() {
		if err := detector.Close(); err != nil {
			t.Errorf("unable to close CarDetector: %v", err)
		}
	}()

	cases := []struct {
		imgPath         string
		expectedObjects []events.Object
	}{
		{"testdata/image.jpg", []events.Object{},
		},
		{"testdata/image2.jpg", []events.Object{
			{Type: events.TypeObject_ANY, Left: 54, Top: 5, Right: 117, Bottom: 55, Confidence: 0.70918894},},
		},
		{"testdata/image3.jpg", []events.Object{
			{Type: events.TypeObject_ANY, Left: 15, Top: 0, Right: 59, Bottom: 34, Confidence: 0.25008404},
			{Type: events.TypeObject_ANY, Left: 92, Top: 8, Right: 108, Bottom: 19, Confidence: 0.5126402},
			{Type: events.TypeObject_ANY, Left: 116, Top: 9, Right: 141, Bottom: 32, Confidence: 0.4450239},},
		},
		{"testdata/image4.jpg", []events.Object{
			{Type: events.TypeObject_ANY, Left: 49, Top: 2, Right: 153, Bottom: 46, Confidence: 0.63352036},},
		},
		{"testdata/image5.jpg", []events.Object{
			{Type: events.TypeObject_ANY, Left: -21, Top: -13, Right: 136, Bottom: 104, Confidence: 0.64845103},
			{Type: events.TypeObject_ANY, Left: -21, Top: -13, Right: 136, Bottom: 104, Confidence: 0.4342566},
			{Type: events.TypeObject_ANY, Left: -21, Top: -13, Right: 136, Bottom: 104, Confidence: 0.47660118},},
		},
	}

	for _, c := range cases {
		jpegContent, err := ioutil.ReadFile(c.imgPath)
		if err != nil {
			t.Errorf("unable to load image: %v", err)
			continue
		}

		img, err := gocv.IMDecode(jpegContent, gocv.IMReadUnchanged)
		if err != nil {
			t.Fatalf("unable to decode test image %v: %v", c.imgPath, err)
		}

		objects, err := detector.DetectObjects(&img)
		if err != nil {
			t.Errorf("unexpected error for image %v: %v", c.imgPath, err)
			t.Fail()
		}
		log.Printf("objects: %v", *objects)

		if len(*objects) != len(c.expectedObjects) {
			t.Errorf("bad number of boundingBox found for %v: %v, wants %v", c.imgPath, len(*objects), len(c.expectedObjects))
		}
		for idx, obj := range *objects {
			expectedObj := c.expectedObjects[idx]
			if (*obj).String() != expectedObj.String() {
				t.Errorf("bad boundingBox for %v: %v, wants %v", c.imgPath, obj.String(), expectedObj.String())
			}
		}

		for _, bb := range *objects {
			gocv.Rectangle(&img, image.Rect(int(bb.Left), int(bb.Top), int(bb.Right), int(bb.Bottom)), color.RGBA{0, 255, 0, 0}, 2)
		}
		gocv.IMWrite("/tmp/"+c.imgPath, img)

		if err := img.Close(); err != nil {
			log.Errorf("unable to close image: %v", err)
		}
	}
}
