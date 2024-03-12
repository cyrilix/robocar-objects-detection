package objects

import (
	"github.com/cyrilix/robocar-protobuf/go/events"
	"reflect"
	"testing"
)

var (
	objectOnMiddleDistant = events.Object{
		Type:       events.TypeObject_ANY,
		Left:       0.4,
		Top:        0.1,
		Right:      0.6,
		Bottom:     0.2,
		Confidence: 0.9,
	}
	objectOnLeftDistant = events.Object{
		Type:       events.TypeObject_ANY,
		Left:       0.1,
		Top:        0.09,
		Right:      0.3,
		Bottom:     0.19,
		Confidence: 0.9,
	}
	objectOnRightDistant = events.Object{
		Type:       events.TypeObject_ANY,
		Left:       0.7,
		Top:        0.21,
		Right:      0.9,
		Bottom:     0.11,
		Confidence: 0.9,
	}
	objectOnMiddleNear = events.Object{
		Type:       events.TypeObject_ANY,
		Left:       0.4,
		Top:        0.8,
		Right:      0.6,
		Bottom:     0.9,
		Confidence: 0.9,
	}
	objectOnRightNear = events.Object{
		Type:       events.TypeObject_ANY,
		Left:       0.7,
		Top:        0.8,
		Right:      0.9,
		Bottom:     0.9,
		Confidence: 0.9,
	}
	objectOnLeftNear = events.Object{
		Type:       events.TypeObject_ANY,
		Left:       0.1,
		Top:        0.8,
		Right:      0.3,
		Bottom:     0.9,
		Confidence: 0.9,
	}
)

func TestNewFilter(t *testing.T) {
	type args struct {
		imgWidth      int
		imgHeight     int
		sizeThreshold float64
	}
	tests := []struct {
		name string
		args args
		want *ObjectFilter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilter(tt.args.imgWidth, tt.args.imgHeight, tt.args.sizeThreshold, false,
				false, false, false); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectFilter_Process(t *testing.T) {
	type fields struct {
		imgWidth      int
		imgHeight     int
		sizeThreshold float64
	}
	type args struct {
		objs []*events.Object
		in1  *Disparity
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*events.Object
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ObjectFilter{
				imgWidth:      tt.fields.imgWidth,
				imgHeight:     tt.fields.imgHeight,
				sizeThreshold: tt.fields.sizeThreshold,
			}
			got, err := o.Process(tt.args.objs, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectFilter_filterBigObjects(t *testing.T) {
	type fields struct {
		imgWidth      int
		imgHeight     int
		sizeThreshold float64
	}
	type args struct {
		objects []*events.Object
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*events.Object
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ObjectFilter{
				imgWidth:      tt.fields.imgWidth,
				imgHeight:     tt.fields.imgHeight,
				sizeThreshold: tt.fields.sizeThreshold,
			}
			if got := o.filterBigObjects(tt.args.objects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterBigObjects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectFilter_filterBottomImages(t *testing.T) {
	type fields struct {
		imgWidth      int
		imgHeight     int
		sizeThreshold float64
	}
	type args struct {
		objects []*events.Object
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*events.Object
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ObjectFilter{
				imgWidth:      tt.fields.imgWidth,
				imgHeight:     tt.fields.imgHeight,
				sizeThreshold: tt.fields.sizeThreshold,
			}
			if got := o.filterBottomImages(tt.args.objects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterBottomImages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectFilter_nearObject(t *testing.T) {
	type args struct {
		objects []*events.Object
	}
	tests := []struct {
		name    string
		args    args
		want    *events.Object
		wantErr bool
	}{
		{
			name: "List object is empty",
			args: args{
				objects: []*events.Object{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List with only one object",
			args: args{
				objects: []*events.Object{&objectOnMiddleNear},
			},
			want:    &objectOnMiddleNear,
			wantErr: false,
		},
		{
			name: "List with many objects",
			args: args{
				objects: []*events.Object{&objectOnLeftDistant, &objectOnMiddleNear, &objectOnRightDistant, &objectOnMiddleDistant},
			},
			want:    &objectOnMiddleNear,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ObjectFilter{}
			got, err := c.nearObject(tt.args.objects)
			if (err != nil) != tt.wantErr {
				t.Errorf("nearObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nearObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCorrector_nearObject(t *testing.T) {
	type args struct {
		objects []*events.Object
	}
	tests := []struct {
		name    string
		args    args
		want    *events.Object
		wantErr bool
	}{
		{
			name: "List object is empty",
			args: args{
				objects: []*events.Object{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List with only one object",
			args: args{
				objects: []*events.Object{&objectOnMiddleNear},
			},
			want:    &objectOnMiddleNear,
			wantErr: false,
		},
		{
			name: "List with many objects",
			args: args{
				objects: []*events.Object{&objectOnLeftDistant, &objectOnMiddleNear, &objectOnRightDistant, &objectOnMiddleDistant},
			},
			want:    &objectOnMiddleNear,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ObjectFilter{}
			got, err := c.nearObject(tt.args.objects)
			if (err != nil) != tt.wantErr {
				t.Errorf("nearObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nearObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestObjectProcessor_Process(t *testing.T) {
	type fields struct {
		minDistanceInMm int
		maxDistanceInMm int
	}
	type args struct {
		objects   []*events.Object
		disparity *Disparity
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*events.Object
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := ObjectProcessor{
				minDistanceInMm: tt.fields.minDistanceInMm,
				maxDistanceInMm: tt.fields.maxDistanceInMm,
			}
			got, err := o.Process(tt.args.objects, tt.args.disparity)
			if (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process() got = %v, want %v", got, tt.want)
			}
		})
	}
}
