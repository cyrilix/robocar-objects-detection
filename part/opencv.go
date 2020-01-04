package part

import (
	"fmt"
	"github.com/cyrilix/robocar-protobuf/go/events"
	log "github.com/sirupsen/logrus"
	"gocv.io/x/gocv"
	"image"
	"io"
)

type ObjectDetector interface {
	DetectObjects(img *gocv.Mat) (*[]*events.Object, error)
	io.Closer
}

func NewCarDetector(model, config string) *CarDetector {
	net := gocv.ReadNet(model, config)
	return &CarDetector{net: net}
}

type CarDetector struct {
	net gocv.Net
}

func (d *CarDetector) Close() error {
	if err := d.net.Close(); err != nil {
		return fmt.Errorf("unable to close CarDetector: %v", err)
	}
	return nil
}

func (d *CarDetector) DetectObjects(img *gocv.Mat) (*[]*events.Object, error) {

	var blob = gocv.BlobFromImage(*img, 1.0, image.Pt(224, 244), gocv.NewScalar(0, 0, 0, 0), true, false)
	defer func() {
		if err := blob.Close(); err != nil {
			log.Errorf("unable to close bob: %v", err)
		}
	}()

	d.net.SetInput(blob, "")
	results := d.net.Forward("detection_out")
	defer func() {
		if err := results.Close(); err != nil {
			log.Errorf("unable to close Mat result: %v", err)
		}
	}()

	return performDetection(&results, img.Rows(), img.Cols()), nil
}

// performDetection analyzes the results from the detector network,
// which produces an output blob with a shape 1x1xNx7
// where N is the number of detections, and each detection
// is a vector of float values
// [batchId, classId, confidence, left, top, right, bottom]
func performDetection(results *gocv.Mat, nbRows, nbCols int) *[]*events.Object {
	objects := make([]*events.Object, 0, results.Total())

	for i := 0; i < results.Total(); i += 7 {
		confidence := results.GetFloatAt(0, i+2)
		if confidence > 0.25 {
			log.Debugf("classId: %v\n", results.GetFloatAt(0, i+1))

			objects = append(objects, &events.Object{
				Type:  events.TypeObject_ANY,
				Left:  int32(results.GetFloatAt(0, i+3) * float32(nbCols)),
				Top:   int32(results.GetFloatAt(0, i+4) * float32(nbRows)),
				Right: int32(results.GetFloatAt(0, i+5) * float32(nbCols)),
				Bottom:     int32(results.GetFloatAt(0, i+6) * float32(nbRows)),
				Confidence: confidence,
			})
		}
	}
	return &objects
}
