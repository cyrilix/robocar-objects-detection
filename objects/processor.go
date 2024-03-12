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

func NewFilter(imgWidth int, imgHeight int, sizeThreshold float64, enableBigObjectsRemove bool, enableGroupObjects bool,
	enableBottomFilter bool, enableNearest bool) *ObjectFilter {
	return &ObjectFilter{
		imgWidth:               imgWidth,
		imgHeight:              imgHeight,
		sizeThreshold:          sizeThreshold,
		enableBigObjectsRemove: enableBigObjectsRemove,
		enableGroupObjects:     enableGroupObjects,
		enableBottomFilter:     enableBottomFilter,
		enableNearest:          enableNearest,
	}
}

type ObjectFilter struct {
	imgWidth               int
	imgHeight              int
	sizeThreshold          float64
	enableBigObjectsRemove bool
	enableGroupObjects     bool
	enableBottomFilter     bool
	enableNearest          bool
}

func (o *ObjectFilter) Process(objs []*events.Object, _ *Disparity) ([]*events.Object, error) {

	objects := objs
	if o.enableBigObjectsRemove {
		zap.S().Debugf("%v objects to filter", len(objects))
		objects = o.filterBigObjects(objects)
		zap.S().Debugf("%v objects after removing big objects", len(objects))
	}

	if o.enableBottomFilter {
		zap.S().Debugf("%v objects before removing bottom object", len(objects))
		objects = o.filterBottomImages(objects)
		zap.S().Debugf("%v objects after removing bottom object", len(objects))
	}

	if o.enableGroupObjects {
		zap.S().Debugf("%v objects to avoid before grouping", len(objects))
		if len(objects) > 1 {
			objects = GroupObjects(objects, o.imgWidth, o.imgHeight)
		}
		zap.S().Debugf("%v objects after objects grouping", len(objects))
	}

	// get nearest object
	if o.enableNearest {
		nearest, err := o.nearObject(objects)
		if err != nil {
			return nil, fmt.Errorf("unexpected error on nearest search object, ignore objects: %v", err)
		}
		objects = []*events.Object{nearest}
	}
	return objects, nil
}

func (o *ObjectFilter) filterBigObjects(objects []*events.Object) []*events.Object {
	objectFiltered := make([]*events.Object, 0, len(objects))
	sizeLimit := float64(o.imgWidth*o.imgHeight) * o.sizeThreshold
	for _, obj := range objects {
		if sizeObject(obj, o.imgWidth, o.imgHeight) < sizeLimit {
			objCpy := obj
			objectFiltered = append(objectFiltered, objCpy)
		}
	}
	return objectFiltered
}

func (o *ObjectFilter) filterBottomImages(objects []*events.Object) []*events.Object {
	objectFiltered := make([]*events.Object, 0, len(objects))
	for _, obj := range objects {
		if obj.Top > 0.90 {
			oCpy := obj
			objectFiltered = append(objectFiltered, oCpy)
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
