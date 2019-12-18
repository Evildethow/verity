// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/internal/pb/message.proto

package pb

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

// MessageSource describes the source of a message.
type MessageSource struct {
	Application          *Identity `protobuf:"bytes,4,opt,name=application,proto3" json:"application,omitempty"`
	Handler              *Identity `protobuf:"bytes,5,opt,name=handler,proto3" json:"handler,omitempty"`
	InstanceId           string    `protobuf:"bytes,6,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MessageSource) Reset()         { *m = MessageSource{} }
func (m *MessageSource) String() string { return proto.CompactTextString(m) }
func (*MessageSource) ProtoMessage()    {}
func (*MessageSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f7c7c952ed9e0f9, []int{0}
}

func (m *MessageSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageSource.Unmarshal(m, b)
}
func (m *MessageSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageSource.Marshal(b, m, deterministic)
}
func (m *MessageSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageSource.Merge(m, src)
}
func (m *MessageSource) XXX_Size() int {
	return xxx_messageInfo_MessageSource.Size(m)
}
func (m *MessageSource) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageSource.DiscardUnknown(m)
}

var xxx_messageInfo_MessageSource proto.InternalMessageInfo

func (m *MessageSource) GetApplication() *Identity {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *MessageSource) GetHandler() *Identity {
	if m != nil {
		return m.Handler
	}
	return nil
}

func (m *MessageSource) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

// MessageMetaData contains meta-data about a message.
type MessageMetaData struct {
	MessageId            string         `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	CausationId          string         `protobuf:"bytes,2,opt,name=causation_id,json=causationId,proto3" json:"causation_id,omitempty"`
	CorrelationId        string         `protobuf:"bytes,3,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	Source               *MessageSource `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	CreatedAt            string         `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ScheduledFor         string         `protobuf:"bytes,6,opt,name=scheduled_for,json=scheduledFor,proto3" json:"scheduled_for,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MessageMetaData) Reset()         { *m = MessageMetaData{} }
func (m *MessageMetaData) String() string { return proto.CompactTextString(m) }
func (*MessageMetaData) ProtoMessage()    {}
func (*MessageMetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f7c7c952ed9e0f9, []int{1}
}

func (m *MessageMetaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageMetaData.Unmarshal(m, b)
}
func (m *MessageMetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageMetaData.Marshal(b, m, deterministic)
}
func (m *MessageMetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageMetaData.Merge(m, src)
}
func (m *MessageMetaData) XXX_Size() int {
	return xxx_messageInfo_MessageMetaData.Size(m)
}
func (m *MessageMetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageMetaData.DiscardUnknown(m)
}

var xxx_messageInfo_MessageMetaData proto.InternalMessageInfo

func (m *MessageMetaData) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *MessageMetaData) GetCausationId() string {
	if m != nil {
		return m.CausationId
	}
	return ""
}

func (m *MessageMetaData) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *MessageMetaData) GetSource() *MessageSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *MessageMetaData) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *MessageMetaData) GetScheduledFor() string {
	if m != nil {
		return m.ScheduledFor
	}
	return ""
}

// MessageEnvelope is a container for a message and its meta-data.
type MessageEnvelope struct {
	MetaData             *MessageMetaData `protobuf:"bytes,1,opt,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty"`
	Packet               *Packet          `protobuf:"bytes,2,opt,name=packet,proto3" json:"packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MessageEnvelope) Reset()         { *m = MessageEnvelope{} }
func (m *MessageEnvelope) String() string { return proto.CompactTextString(m) }
func (*MessageEnvelope) ProtoMessage()    {}
func (*MessageEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f7c7c952ed9e0f9, []int{2}
}

func (m *MessageEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageEnvelope.Unmarshal(m, b)
}
func (m *MessageEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageEnvelope.Marshal(b, m, deterministic)
}
func (m *MessageEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageEnvelope.Merge(m, src)
}
func (m *MessageEnvelope) XXX_Size() int {
	return xxx_messageInfo_MessageEnvelope.Size(m)
}
func (m *MessageEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_MessageEnvelope proto.InternalMessageInfo

func (m *MessageEnvelope) GetMetaData() *MessageMetaData {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func (m *MessageEnvelope) GetPacket() *Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageSource)(nil), "dogma.api.messaging.v1.MessageSource")
	proto.RegisterType((*MessageMetaData)(nil), "dogma.api.messaging.v1.MessageMetaData")
	proto.RegisterType((*MessageEnvelope)(nil), "dogma.api.messaging.v1.MessageEnvelope")
}

func init() { proto.RegisterFile("api/internal/pb/message.proto", fileDescriptor_8f7c7c952ed9e0f9) }

var fileDescriptor_8f7c7c952ed9e0f9 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x5f, 0x8b, 0xd3, 0x40,
	0x14, 0xc5, 0x89, 0x7f, 0xa2, 0x99, 0x6c, 0x15, 0xe6, 0x41, 0xca, 0xe2, 0xae, 0xb5, 0xb2, 0xb8,
	0x22, 0x24, 0x58, 0xc1, 0x07, 0xc1, 0x07, 0x97, 0x55, 0xc8, 0xc3, 0x82, 0xc4, 0x37, 0x5f, 0xc2,
	0xed, 0xcc, 0xdd, 0x74, 0x30, 0x99, 0x19, 0x27, 0x37, 0x45, 0x3f, 0x85, 0x5f, 0xc5, 0x2f, 0x28,
	0x48, 0x27, 0x93, 0x5a, 0x8b, 0x45, 0x5f, 0xcf, 0x9c, 0xdf, 0xdc, 0x7b, 0x0e, 0x97, 0x9d, 0x80,
	0x55, 0xb9, 0xd2, 0x84, 0x4e, 0x43, 0x93, 0xdb, 0x65, 0xde, 0x62, 0xd7, 0x41, 0x8d, 0x99, 0x75,
	0x86, 0x0c, 0x7f, 0x20, 0x4d, 0xdd, 0x42, 0x06, 0x56, 0x65, 0xc3, 0x83, 0xd2, 0x75, 0xb6, 0x7e,
	0x71, 0x7c, 0xba, 0x8f, 0x29, 0x89, 0x9a, 0x14, 0x7d, 0x1b, 0xb8, 0xe3, 0x87, 0xfb, 0xef, 0x16,
	0xc4, 0x67, 0xa4, 0xe1, 0x75, 0xfe, 0x23, 0x62, 0x93, 0xab, 0x61, 0xce, 0x47, 0xd3, 0x3b, 0x81,
	0xfc, 0x82, 0xa5, 0x60, 0x6d, 0xa3, 0x04, 0x90, 0x32, 0x7a, 0x7a, 0x6b, 0x16, 0x9d, 0xa7, 0x8b,
	0x59, 0xf6, 0xf7, 0xe9, 0x59, 0x11, 0x86, 0x95, 0xbb, 0x10, 0x7f, 0xcd, 0xee, 0xac, 0x40, 0xcb,
	0x06, 0xdd, 0xf4, 0xf6, 0x7f, 0xf2, 0x23, 0xc0, 0x1f, 0xb1, 0x54, 0xe9, 0x8e, 0x40, 0x0b, 0xac,
	0x94, 0x9c, 0xc6, 0xb3, 0xe8, 0x3c, 0x29, 0xd9, 0x28, 0x15, 0x72, 0xfe, 0x33, 0x62, 0xf7, 0xc3,
	0xca, 0x57, 0x48, 0x70, 0x09, 0x04, 0xfc, 0x84, 0xb1, 0xd0, 0xd6, 0x86, 0x89, 0x3c, 0x93, 0x04,
	0xa5, 0x90, 0xfc, 0x31, 0x3b, 0x12, 0xd0, 0x77, 0x7e, 0xb9, 0x8d, 0xe1, 0x86, 0x37, 0xa4, 0x5b,
	0xad, 0x90, 0xfc, 0x8c, 0xdd, 0x13, 0xc6, 0x39, 0x6c, 0xb6, 0xa6, 0x9b, 0xde, 0x34, 0xd9, 0x51,
	0x0b, 0xc9, 0xdf, 0xb0, 0xb8, 0xf3, 0x3d, 0x85, 0x62, 0xce, 0x0e, 0x05, 0xfb, 0xa3, 0xd4, 0x32,
	0x40, 0x9b, 0x3d, 0x85, 0x43, 0x20, 0x94, 0x15, 0x90, 0xef, 0x26, 0x29, 0x93, 0xa0, 0xbc, 0x25,
	0xfe, 0x84, 0x4d, 0x3a, 0xb1, 0x42, 0xd9, 0x37, 0x28, 0xab, 0x6b, 0xe3, 0x42, 0xfa, 0xa3, 0xad,
	0xf8, 0xde, 0xb8, 0xf9, 0xf7, 0xdf, 0xf9, 0xdf, 0xe9, 0x35, 0x36, 0xc6, 0x22, 0xbf, 0x64, 0x49,
	0x8b, 0x04, 0x95, 0x04, 0x02, 0x1f, 0x3f, 0x5d, 0x3c, 0xfd, 0xc7, 0x66, 0x63, 0x77, 0xe5, 0xdd,
	0x76, 0x6c, 0xf1, 0x15, 0x8b, 0x87, 0xe3, 0xf0, 0x05, 0xa5, 0x8b, 0xd3, 0x43, 0x5f, 0x7c, 0xf0,
	0xae, 0x32, 0xb8, 0x2f, 0x9e, 0x7f, 0x7a, 0x56, 0x2b, 0x5a, 0xf5, 0xcb, 0x4c, 0x98, 0x36, 0xf7,
	0x0c, 0xa9, 0x2f, 0xb9, 0xd2, 0xd7, 0xea, 0x6b, 0xbe, 0x77, 0x7e, 0xcb, 0xd8, 0x1f, 0xde, 0xcb,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xfe, 0xc7, 0x0a, 0xef, 0x02, 0x00, 0x00,
}
