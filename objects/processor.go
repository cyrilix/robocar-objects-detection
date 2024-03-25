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

func NewFilter(imgWidth int, imgHeight int, enableGroupObjects bool, enableNearest bool, filters ...Filter) *ObjectFilter {
	return &ObjectFilter{
		imgWidth:           imgWidth,
		imgHeight:          imgHeight,
		enableGroupObjects: enableGroupObjects,
		enableNearest:      enableNearest,
		filters:            filters,
	}
}

type ObjectFilter struct {
	imgWidth           int
	imgHeight          int
	enableGroupObjects bool
	enableNearest      bool
	filters            []Filter
}

func (o *ObjectFilter) Process(objs []*events.Object, _ *Disparity) ([]*events.Object, error) {

	objects := objs
	objectFiltereds := make([]*events.Object, 0, len(objects))

	if len(o.filters) > 0 {
		zap.S().Debugf("%v objects to avoid before filtering", len(objects))
		for _, obj := range objects {
			object := obj
			for _, filter := range o.filters {
				if filter.Filter(object) {
					objectFiltereds = append(objectFiltereds, object)
				}
			}
		}
		zap.S().Debugf("%v objects to avoid after filtering", len(objectFiltereds))
	} else {
		objectFiltereds = objects
	}

	if o.enableGroupObjects {
		zap.S().Debugf("%v objects to avoid before grouping", len(objectFiltereds))
		if len(objectFiltereds) > 1 {
			objectFiltereds = GroupObjects(objectFiltereds, o.imgWidth, o.imgHeight)
		}
		zap.S().Debugf("%v objects after objects grouping", len(objectFiltereds))
	}

	// get nearest object
	if o.enableNearest {
		nearest, err := o.nearObject(objectFiltereds)
		if err != nil {
			return nil, fmt.Errorf("unexpected error on nearest search object, ignore objects: %v", err)
		}
		objectFiltereds = []*events.Object{nearest}
	}
	return objectFiltereds, nil
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

type Filter interface {
	Filter(*events.Object) bool
}

func NewBigObjectFilter(sizeThreshold float64, imgWidth, imgHeight int) *BigObjectFilter {
	sizeLimit := float64(imgWidth*imgHeight) * sizeThreshold
	return &BigObjectFilter{
		imgWidth:  imgWidth,
		imgHeight: imgHeight,
		sizeLimit: sizeLimit,
	}
}

type BigObjectFilter struct {
	imgWidth, imgHeight int
	sizeLimit           float64
}

func (b BigObjectFilter) Filter(object *events.Object) bool {
	return sizeObject(object, b.imgWidth, b.imgHeight) < b.sizeLimit
}

type BottomFilter struct {
	bottomLimit float64
}

func NewBottomFilter(bottomLimit float64) *BottomFilter {
	return &BottomFilter{bottomLimit}
}

func (b BottomFilter) Filter(object *events.Object) bool {
	return object.Bottom < float32(1.-b.bottomLimit)
}
