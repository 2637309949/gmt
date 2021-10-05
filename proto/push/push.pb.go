// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push/push.proto

package push

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

// EventT1 defined TODO
type EventT1 struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventT1) Reset()         { *m = EventT1{} }
func (m *EventT1) String() string { return proto.CompactTextString(m) }
func (*EventT1) ProtoMessage()    {}
func (*EventT1) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae0042da44e9a7a7, []int{0}
}

func (m *EventT1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventT1.Unmarshal(m, b)
}
func (m *EventT1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventT1.Marshal(b, m, deterministic)
}
func (m *EventT1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventT1.Merge(m, src)
}
func (m *EventT1) XXX_Size() int {
	return xxx_messageInfo_EventT1.Size(m)
}
func (m *EventT1) XXX_DiscardUnknown() {
	xxx_messageInfo_EventT1.DiscardUnknown(m)
}

var xxx_messageInfo_EventT1 proto.InternalMessageInfo

func (m *EventT1) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EventT1) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *EventT1) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EventT1)(nil), "push.EventT1")
}

func init() { proto.RegisterFile("push/push.proto", fileDescriptor_ae0042da44e9a7a7) }

var fileDescriptor_ae0042da44e9a7a7 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0x2d, 0xce,
	0xd0, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x52, 0x20, 0x17,
	0xbb, 0x6b, 0x59, 0x6a, 0x5e, 0x49, 0x88, 0xa1, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x0c, 0x17, 0x67, 0x49, 0x66, 0x6e, 0x6a,
	0x71, 0x49, 0x62, 0x6e, 0x81, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x42, 0x40, 0x48, 0x82,
	0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x55, 0x82, 0x19, 0xac, 0x05, 0xc6, 0x4d, 0x62,
	0x03, 0x9b, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x2e, 0x9c, 0x99, 0x72, 0x00, 0x00,
	0x00,
}