package objects

import (
	"fmt"
	"github.com/cyrilix/robocar-protobuf/go/events"
)
import "gocv.io/x/gocv"
import "go.uber.org/zap"

type Processor interface {
	Process(objects []*events.Object, disparity *Disparity) ([]*events.Object, error)
}

type Disparity struct {
	Ref *events.FrameRef
	gocv.Mat
	Baseline            float64
	FocalLengthInPixels float64
}

func (d *Disparity) DistanceInMm(object *events.Object) int {
	disp := d.Mat.GetDoubleAt(
		// Box coordonnates in %
		int(float64(d.Mat.Rows())*float64(object.GetTop()+((object.GetBottom()-object.GetTop())/2))),
		int(float64(d.Mat.Cols())*float64(object.GetLeft()+((object.GetRight()-object.GetLeft())/2))),
	)
	distance := int(d.FocalLengthInPixels * d.Baseline / disp)
	return distance
}

type ObjectProcessor struct {
	minDistanceInMm, maxDistanceInMm int
}

func (o ObjectProcessor) Process(objects []*events.Object, disparity *Disparity) ([]*events.Object, error) {
	result := make([]*events.Object, 0, len(objects))

	for _, obj := range objects {
		// TODO: remove big object
		distance := disparity.DistanceInMm(obj)
		if distance < o.maxDistanceInMm && distance > o.minDistanceInMm {
			result = append(result, obj)
		}
	}
	// Todo: sort by distance
	return result, nil
}

func NewFilter(imgWidth int, imgHeight int, sizeThreshold float64) *ObjectFilter {
	return &ObjectFilter{
		imgWidth:      imgWidth,
		imgHeight:     imgHeight,
		sizeThreshold: sizeThreshold,
	}
}

type ObjectFilter struct {
	imgWidth      int
	imgHeight     int
	sizeThreshold float64
}

func (o *ObjectFilter) Process(objs []*events.Object, _ *Disparity) ([]*events.Object, error) {

	zap.S().Debugf("%v objects to filter", len(objs))
	objects := o.filterBigObjects(objs)
	zap.S().Debugf("%v objects after removing big objects", len(objects))

	//objects = o.filterBottomImages(objects)
	//zap.S().Debugf("%v objects after removing bottom object", len(objects))

	zap.S().Debugf("%v objects to avoid before grouping", len(objects))
	if len(objects) == 0 {
		return []*events.Object{}, nil
	}

	grpObjs := GroupObjects(objects, o.imgWidth, o.imgHeight)
	zap.S().Debugf("%v objects after objects grouping", len(objects))

	// get nearest object
	nearest, err := o.nearObject(grpObjs)
	if err != nil {
		return nil, fmt.Errorf("unexpected error on nearest search object, ignore objects: %v", err)
	}
	return []*events.Object{nearest}, nil
}

func (o *ObjectFilter) filterBigObjects(objects []*events.Object) []*events.Object {
	objectFiltered := make([]*events.Object, 0, len(objects))
	sizeLimit := float64(o.imgWidth*o.imgHeight) * o.sizeThreshold
	for _, obj := range objects {
		if sizeObject(obj, o.imgWidth, o.imgHeight) < sizeLimit {
			objectFiltered = append(objectFiltered, obj)
		}
	}
	return objectFiltered
}

func (o *ObjectFilter) filterBottomImages(objects []*events.Object) []*events.Object {
	objectFiltered := make([]*events.Object, 0, len(objects))
	for _, o := range objects {
		if o.Top > 0.90 {
			objectFiltered = append(objectFiltered, o)
		}
	}
	return objectFiltered
}

func (o *ObjectFilter) nearObject(objects []*events.Object) (*events.Object, error) {
	if len(objects) == 0 {
		return nil, fmt.Errorf("list objects must contain at least one object")
	}
	if len(objects) == 1 {
		return objects[0], nil
	}

	var result *events.Object
	for _, obj := range objects {
		objCpy := obj
		if result == nil || objCpy.Bottom > result.Bottom {
			result = objCpy
			continue
		}
	}
	return result, nil
}
