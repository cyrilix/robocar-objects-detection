// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/events.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DriveMode int32

const (
	DriveMode_INVALID DriveMode = 0
	DriveMode_USER    DriveMode = 1
	DriveMode_PILOT   DriveMode = 2
)

var DriveMode_name = map[int32]string{
	0: "INVALID",
	1: "USER",
	2: "PILOT",
}

var DriveMode_value = map[string]int32{
	"INVALID": 0,
	"USER":    1,
	"PILOT":   2,
}

func (x DriveMode) String() string {
	return proto.EnumName(DriveMode_name, int32(x))
}

func (DriveMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{0}
}

type TypeObject int32

const (
	TypeObject_ANY  TypeObject = 0
	TypeObject_CAR  TypeObject = 1
	TypeObject_BUMP TypeObject = 2
	TypeObject_PLOT TypeObject = 3
)

var TypeObject_name = map[int32]string{
	0: "ANY",
	1: "CAR",
	2: "BUMP",
	3: "PLOT",
}

var TypeObject_value = map[string]int32{
	"ANY":  0,
	"CAR":  1,
	"BUMP": 2,
	"PLOT": 3,
}

func (x TypeObject) String() string {
	return proto.EnumName(TypeObject_name, int32(x))
}

func (TypeObject) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{1}
}

type FrameRef struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrameRef) Reset()         { *m = FrameRef{} }
func (m *FrameRef) String() string { return proto.CompactTextString(m) }
func (*FrameRef) ProtoMessage()    {}
func (*FrameRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{0}
}

func (m *FrameRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrameRef.Unmarshal(m, b)
}
func (m *FrameRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrameRef.Marshal(b, m, deterministic)
}
func (m *FrameRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrameRef.Merge(m, src)
}
func (m *FrameRef) XXX_Size() int {
	return xxx_messageInfo_FrameRef.Size(m)
}
func (m *FrameRef) XXX_DiscardUnknown() {
	xxx_messageInfo_FrameRef.DiscardUnknown(m)
}

var xxx_messageInfo_FrameRef proto.InternalMessageInfo

func (m *FrameRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FrameRef) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type FrameMessage struct {
	Id                   *FrameRef `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Frame                []byte    `protobuf:"bytes,2,opt,name=frame,proto3" json:"frame,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FrameMessage) Reset()         { *m = FrameMessage{} }
func (m *FrameMessage) String() string { return proto.CompactTextString(m) }
func (*FrameMessage) ProtoMessage()    {}
func (*FrameMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{1}
}

func (m *FrameMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrameMessage.Unmarshal(m, b)
}
func (m *FrameMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrameMessage.Marshal(b, m, deterministic)
}
func (m *FrameMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrameMessage.Merge(m, src)
}
func (m *FrameMessage) XXX_Size() int {
	return xxx_messageInfo_FrameMessage.Size(m)
}
func (m *FrameMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_FrameMessage.DiscardUnknown(m)
}

var xxx_messageInfo_FrameMessage proto.InternalMessageInfo

func (m *FrameMessage) GetId() *FrameRef {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *FrameMessage) GetFrame() []byte {
	if m != nil {
		return m.Frame
	}
	return nil
}

type SteeringMessage struct {
	Steering             float32   `protobuf:"fixed32,1,opt,name=steering,proto3" json:"steering,omitempty"`
	Confidence           float32   `protobuf:"fixed32,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
	FrameRef             *FrameRef `protobuf:"bytes,3,opt,name=frame_ref,json=frameRef,proto3" json:"frame_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SteeringMessage) Reset()         { *m = SteeringMessage{} }
func (m *SteeringMessage) String() string { return proto.CompactTextString(m) }
func (*SteeringMessage) ProtoMessage()    {}
func (*SteeringMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{2}
}

func (m *SteeringMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SteeringMessage.Unmarshal(m, b)
}
func (m *SteeringMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SteeringMessage.Marshal(b, m, deterministic)
}
func (m *SteeringMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SteeringMessage.Merge(m, src)
}
func (m *SteeringMessage) XXX_Size() int {
	return xxx_messageInfo_SteeringMessage.Size(m)
}
func (m *SteeringMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SteeringMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SteeringMessage proto.InternalMessageInfo

func (m *SteeringMessage) GetSteering() float32 {
	if m != nil {
		return m.Steering
	}
	return 0
}

func (m *SteeringMessage) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

func (m *SteeringMessage) GetFrameRef() *FrameRef {
	if m != nil {
		return m.FrameRef
	}
	return nil
}

type ThrottleMessage struct {
	Throttle             float32   `protobuf:"fixed32,1,opt,name=throttle,proto3" json:"throttle,omitempty"`
	Confidence           float32   `protobuf:"fixed32,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
	FrameRef             *FrameRef `protobuf:"bytes,3,opt,name=frame_ref,json=frameRef,proto3" json:"frame_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ThrottleMessage) Reset()         { *m = ThrottleMessage{} }
func (m *ThrottleMessage) String() string { return proto.CompactTextString(m) }
func (*ThrottleMessage) ProtoMessage()    {}
func (*ThrottleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{3}
}

func (m *ThrottleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThrottleMessage.Unmarshal(m, b)
}
func (m *ThrottleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThrottleMessage.Marshal(b, m, deterministic)
}
func (m *ThrottleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThrottleMessage.Merge(m, src)
}
func (m *ThrottleMessage) XXX_Size() int {
	return xxx_messageInfo_ThrottleMessage.Size(m)
}
func (m *ThrottleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ThrottleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ThrottleMessage proto.InternalMessageInfo

func (m *ThrottleMessage) GetThrottle() float32 {
	if m != nil {
		return m.Throttle
	}
	return 0
}

func (m *ThrottleMessage) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

func (m *ThrottleMessage) GetFrameRef() *FrameRef {
	if m != nil {
		return m.FrameRef
	}
	return nil
}

type DriveModeMessage struct {
	DriveMode            DriveMode `protobuf:"varint,1,opt,name=drive_mode,json=driveMode,proto3,enum=robocar.events.DriveMode" json:"drive_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DriveModeMessage) Reset()         { *m = DriveModeMessage{} }
func (m *DriveModeMessage) String() string { return proto.CompactTextString(m) }
func (*DriveModeMessage) ProtoMessage()    {}
func (*DriveModeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{4}
}

func (m *DriveModeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriveModeMessage.Unmarshal(m, b)
}
func (m *DriveModeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriveModeMessage.Marshal(b, m, deterministic)
}
func (m *DriveModeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriveModeMessage.Merge(m, src)
}
func (m *DriveModeMessage) XXX_Size() int {
	return xxx_messageInfo_DriveModeMessage.Size(m)
}
func (m *DriveModeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DriveModeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DriveModeMessage proto.InternalMessageInfo

func (m *DriveModeMessage) GetDriveMode() DriveMode {
	if m != nil {
		return m.DriveMode
	}
	return DriveMode_INVALID
}

type ObjectsMessage struct {
	Objects              []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	FrameRef             *FrameRef `protobuf:"bytes,2,opt,name=frame_ref,json=frameRef,proto3" json:"frame_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ObjectsMessage) Reset()         { *m = ObjectsMessage{} }
func (m *ObjectsMessage) String() string { return proto.CompactTextString(m) }
func (*ObjectsMessage) ProtoMessage()    {}
func (*ObjectsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{5}
}

func (m *ObjectsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectsMessage.Unmarshal(m, b)
}
func (m *ObjectsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectsMessage.Marshal(b, m, deterministic)
}
func (m *ObjectsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectsMessage.Merge(m, src)
}
func (m *ObjectsMessage) XXX_Size() int {
	return xxx_messageInfo_ObjectsMessage.Size(m)
}
func (m *ObjectsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectsMessage proto.InternalMessageInfo

func (m *ObjectsMessage) GetObjects() []*Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

func (m *ObjectsMessage) GetFrameRef() *FrameRef {
	if m != nil {
		return m.FrameRef
	}
	return nil
}

// BoundingBox that contains an object
type Object struct {
	Type                 TypeObject `protobuf:"varint,1,opt,name=type,proto3,enum=robocar.events.TypeObject" json:"type,omitempty"`
	Left                 int32      `protobuf:"varint,2,opt,name=left,proto3" json:"left,omitempty"`
	Top                  int32      `protobuf:"varint,3,opt,name=top,proto3" json:"top,omitempty"`
	Right                int32      `protobuf:"varint,4,opt,name=right,proto3" json:"right,omitempty"`
	Bottom               int32      `protobuf:"varint,5,opt,name=bottom,proto3" json:"bottom,omitempty"`
	Confidence           float32    `protobuf:"fixed32,6,opt,name=confidence,proto3" json:"confidence,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{6}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetType() TypeObject {
	if m != nil {
		return m.Type
	}
	return TypeObject_ANY
}

func (m *Object) GetLeft() int32 {
	if m != nil {
		return m.Left
	}
	return 0
}

func (m *Object) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *Object) GetRight() int32 {
	if m != nil {
		return m.Right
	}
	return 0
}

func (m *Object) GetBottom() int32 {
	if m != nil {
		return m.Bottom
	}
	return 0
}

func (m *Object) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

type SwitchRecordMessage struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwitchRecordMessage) Reset()         { *m = SwitchRecordMessage{} }
func (m *SwitchRecordMessage) String() string { return proto.CompactTextString(m) }
func (*SwitchRecordMessage) ProtoMessage()    {}
func (*SwitchRecordMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{7}
}

func (m *SwitchRecordMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwitchRecordMessage.Unmarshal(m, b)
}
func (m *SwitchRecordMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwitchRecordMessage.Marshal(b, m, deterministic)
}
func (m *SwitchRecordMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwitchRecordMessage.Merge(m, src)
}
func (m *SwitchRecordMessage) XXX_Size() int {
	return xxx_messageInfo_SwitchRecordMessage.Size(m)
}
func (m *SwitchRecordMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SwitchRecordMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SwitchRecordMessage proto.InternalMessageInfo

func (m *SwitchRecordMessage) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func init() {
	proto.RegisterEnum("robocar.events.DriveMode", DriveMode_name, DriveMode_value)
	proto.RegisterEnum("robocar.events.TypeObject", TypeObject_name, TypeObject_value)
	proto.RegisterType((*FrameRef)(nil), "robocar.events.FrameRef")
	proto.RegisterType((*FrameMessage)(nil), "robocar.events.FrameMessage")
	proto.RegisterType((*SteeringMessage)(nil), "robocar.events.SteeringMessage")
	proto.RegisterType((*ThrottleMessage)(nil), "robocar.events.ThrottleMessage")
	proto.RegisterType((*DriveModeMessage)(nil), "robocar.events.DriveModeMessage")
	proto.RegisterType((*ObjectsMessage)(nil), "robocar.events.ObjectsMessage")
	proto.RegisterType((*Object)(nil), "robocar.events.Object")
	proto.RegisterType((*SwitchRecordMessage)(nil), "robocar.events.SwitchRecordMessage")
}

func init() { proto.RegisterFile("events/events.proto", fileDescriptor_8ec31f2d2a3db598) }

var fileDescriptor_8ec31f2d2a3db598 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x8d, 0x64, 0x4b, 0x96, 0xc6, 0xc1, 0x15, 0x9b, 0x12, 0xd4, 0x1c, 0x4a, 0xd0, 0xc9, 0x04,
	0xaa, 0x14, 0x97, 0x42, 0xaf, 0x4e, 0xd3, 0x82, 0xc1, 0x76, 0xcc, 0xda, 0x29, 0xb4, 0x97, 0xa0,
	0x8f, 0x91, 0xad, 0x62, 0x6b, 0xcd, 0x6a, 0x49, 0xf1, 0xbd, 0x3f, 0xa7, 0x3f, 0xb2, 0xec, 0xae,
	0xd6, 0x4d, 0x74, 0x69, 0x2f, 0x3d, 0xf9, 0xbd, 0xe7, 0x99, 0xb7, 0x6f, 0x67, 0x47, 0x70, 0x86,
	0x8f, 0x58, 0x89, 0xfa, 0x5a, 0xff, 0xc4, 0x7b, 0xce, 0x04, 0x23, 0x03, 0xce, 0x52, 0x96, 0x25,
	0x3c, 0xd6, 0x6a, 0x14, 0x83, 0xf7, 0x99, 0x27, 0x3b, 0xa4, 0x58, 0x10, 0x02, 0xdd, 0x2a, 0xd9,
	0x61, 0x68, 0x5d, 0x5a, 0x43, 0x9f, 0x2a, 0x4c, 0x06, 0x60, 0x97, 0x79, 0x68, 0x2b, 0xc5, 0x2e,
	0xf3, 0x68, 0x0e, 0xa7, 0xaa, 0x7e, 0x86, 0x75, 0x9d, 0xac, 0x91, 0x0c, 0xd5, 0xff, 0xb2, 0xa3,
	0x3f, 0x0a, 0xe3, 0xe7, 0xe6, 0xb1, 0x71, 0x96, 0x9d, 0xe4, 0x25, 0x38, 0x85, 0xe4, 0xca, 0xec,
	0x94, 0x6a, 0x12, 0xfd, 0xb4, 0xe0, 0xc5, 0x52, 0x20, 0xf2, 0xb2, 0x5a, 0x1b, 0xcf, 0x0b, 0xf0,
	0xea, 0x46, 0x52, 0xce, 0x36, 0x3d, 0x72, 0xf2, 0x1a, 0x20, 0x63, 0x55, 0x51, 0xe6, 0x58, 0x65,
	0xda, 0xca, 0xa6, 0x4f, 0x14, 0xf2, 0x1e, 0x7c, 0x65, 0xfc, 0xc0, 0xb1, 0x08, 0x3b, 0x7f, 0x89,
	0xe5, 0x15, 0x0d, 0x52, 0x31, 0x56, 0x1b, 0xce, 0x84, 0xd8, 0xe2, 0x93, 0x18, 0xa2, 0x91, 0x4c,
	0x0c, 0xc3, 0xff, 0x57, 0x8c, 0x29, 0x04, 0xb7, 0xbc, 0x7c, 0xc4, 0x19, 0xcb, 0x8f, 0x31, 0x3e,
	0x00, 0xe4, 0x52, 0x7b, 0xd8, 0xb1, 0x5c, 0x07, 0x19, 0x8c, 0x5e, 0xb5, 0xbd, 0x8e, 0x5d, 0xd4,
	0xcf, 0x0d, 0x8c, 0x0e, 0x30, 0xb8, 0x4b, 0xbf, 0x63, 0x26, 0x6a, 0xe3, 0xf5, 0x16, 0x7a, 0x4c,
	0x2b, 0xa1, 0x75, 0xd9, 0x19, 0xf6, 0x47, 0xe7, 0x6d, 0x23, 0xdd, 0x40, 0x4d, 0xd9, 0xf3, 0x8b,
	0xd8, 0xff, 0x7c, 0x91, 0x5f, 0x16, 0xb8, 0xda, 0x8a, 0xc4, 0xd0, 0x15, 0x87, 0xbd, 0x49, 0x7e,
	0xd1, 0x6e, 0x5e, 0x1d, 0xf6, 0xd8, 0x1c, 0xaa, 0xea, 0xe4, 0x16, 0x6e, 0xb1, 0x10, 0xea, 0x30,
	0x87, 0x2a, 0x4c, 0x02, 0xe8, 0x08, 0xb6, 0x57, 0x83, 0x74, 0xa8, 0x84, 0x72, 0x9b, 0x78, 0xb9,
	0xde, 0x88, 0xb0, 0xab, 0x34, 0x4d, 0xc8, 0x39, 0xb8, 0x29, 0x13, 0x82, 0xed, 0x42, 0x47, 0xc9,
	0x0d, 0x6b, 0x3d, 0x97, 0xdb, 0x7e, 0xae, 0xe8, 0x1a, 0xce, 0x96, 0x3f, 0x4a, 0x91, 0x6d, 0x28,
	0x66, 0x8c, 0xe7, 0x66, 0x5c, 0x21, 0xf4, 0xb0, 0x4a, 0xd2, 0x2d, 0xea, 0x0d, 0xf7, 0xa8, 0xa1,
	0x57, 0x6f, 0xc0, 0x3f, 0x8e, 0x9c, 0xf4, 0xa1, 0x37, 0x99, 0x7f, 0x19, 0x4f, 0x27, 0xb7, 0xc1,
	0x09, 0xf1, 0xa0, 0x7b, 0xbf, 0xfc, 0x44, 0x03, 0x8b, 0xf8, 0xe0, 0x2c, 0x26, 0xd3, 0xbb, 0x55,
	0x60, 0x5f, 0x8d, 0x00, 0xfe, 0xdc, 0x93, 0xf4, 0xa0, 0x33, 0x9e, 0x7f, 0x0d, 0x4e, 0x24, 0xf8,
	0x38, 0x96, 0xa5, 0x1e, 0x74, 0x6f, 0xee, 0x67, 0x8b, 0xc0, 0x96, 0x68, 0x21, 0x7b, 0x3a, 0x37,
	0xde, 0x37, 0x57, 0x8f, 0x28, 0x75, 0xd5, 0xa7, 0xfb, 0xee, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xac, 0xa9, 0xc0, 0x00, 0xd1, 0x03, 0x00, 0x00,
}
