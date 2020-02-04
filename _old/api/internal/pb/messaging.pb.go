// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/internal/pb/messaging.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// ExecuteCommandRequest is the request to a Messaging.ExecuteCommand() RPC.
type ExecuteCommandRequest struct {
	// Application identifies which application the command is targetting.
	Application *Identity `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	// Envelope is a container for the command message.
	Envelope             *MessageEnvelope `protobuf:"bytes,2,opt,name=envelope,proto3" json:"envelope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ExecuteCommandRequest) Reset()         { *m = ExecuteCommandRequest{} }
func (m *ExecuteCommandRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteCommandRequest) ProtoMessage()    {}
func (*ExecuteCommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_141cdf861d5078a2, []int{0}
}

func (m *ExecuteCommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteCommandRequest.Unmarshal(m, b)
}
func (m *ExecuteCommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteCommandRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteCommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteCommandRequest.Merge(m, src)
}
func (m *ExecuteCommandRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteCommandRequest.Size(m)
}
func (m *ExecuteCommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteCommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteCommandRequest proto.InternalMessageInfo

func (m *ExecuteCommandRequest) GetApplication() *Identity {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *ExecuteCommandRequest) GetEnvelope() *MessageEnvelope {
	if m != nil {
		return m.Envelope
	}
	return nil
}

// ExecuteCommandResponse is the response from a Messaging.ExecuteCommand() RPC.
type ExecuteCommandResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteCommandResponse) Reset()         { *m = ExecuteCommandResponse{} }
func (m *ExecuteCommandResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteCommandResponse) ProtoMessage()    {}
func (*ExecuteCommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_141cdf861d5078a2, []int{1}
}

func (m *ExecuteCommandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteCommandResponse.Unmarshal(m, b)
}
func (m *ExecuteCommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteCommandResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteCommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteCommandResponse.Merge(m, src)
}
func (m *ExecuteCommandResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteCommandResponse.Size(m)
}
func (m *ExecuteCommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteCommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteCommandResponse proto.InternalMessageInfo

// ConsumeEventsRequest is the request to a Messaging.ConsumeEvents() RPC.
type ConsumeEventsRequest struct {
	// Application identifies which application's stream to consume from.
	Application *Identity `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	// Offset is the offset of the first event to consume.
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Events is a list of protocol-names for foreign events that the caller
	// wishes to consume.
	Events               []string `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeEventsRequest) Reset()         { *m = ConsumeEventsRequest{} }
func (m *ConsumeEventsRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumeEventsRequest) ProtoMessage()    {}
func (*ConsumeEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_141cdf861d5078a2, []int{2}
}

func (m *ConsumeEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeEventsRequest.Unmarshal(m, b)
}
func (m *ConsumeEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeEventsRequest.Marshal(b, m, deterministic)
}
func (m *ConsumeEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeEventsRequest.Merge(m, src)
}
func (m *ConsumeEventsRequest) XXX_Size() int {
	return xxx_messageInfo_ConsumeEventsRequest.Size(m)
}
func (m *ConsumeEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeEventsRequest proto.InternalMessageInfo

func (m *ConsumeEventsRequest) GetApplication() *Identity {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *ConsumeEventsRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ConsumeEventsRequest) GetEvents() []string {
	if m != nil {
		return m.Events
	}
	return nil
}

// ConsumeEventsResponse is the response from a Messaging.ConsumeEvents() RPC.
type ConsumeEventsResponse struct {
	// Offset is the offset of message within the source application's stream.
	Offset uint64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// Envelope is a container for the message.
	Envelope             *MessageEnvelope `protobuf:"bytes,2,opt,name=envelope,proto3" json:"envelope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ConsumeEventsResponse) Reset()         { *m = ConsumeEventsResponse{} }
func (m *ConsumeEventsResponse) String() string { return proto.CompactTextString(m) }
func (*ConsumeEventsResponse) ProtoMessage()    {}
func (*ConsumeEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_141cdf861d5078a2, []int{3}
}

func (m *ConsumeEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeEventsResponse.Unmarshal(m, b)
}
func (m *ConsumeEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeEventsResponse.Marshal(b, m, deterministic)
}
func (m *ConsumeEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeEventsResponse.Merge(m, src)
}
func (m *ConsumeEventsResponse) XXX_Size() int {
	return xxx_messageInfo_ConsumeEventsResponse.Size(m)
}
func (m *ConsumeEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeEventsResponse proto.InternalMessageInfo

func (m *ConsumeEventsResponse) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ConsumeEventsResponse) GetEnvelope() *MessageEnvelope {
	if m != nil {
		return m.Envelope
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecuteCommandRequest)(nil), "dogma.api.messaging.v1.ExecuteCommandRequest")
	proto.RegisterType((*ExecuteCommandResponse)(nil), "dogma.api.messaging.v1.ExecuteCommandResponse")
	proto.RegisterType((*ConsumeEventsRequest)(nil), "dogma.api.messaging.v1.ConsumeEventsRequest")
	proto.RegisterType((*ConsumeEventsResponse)(nil), "dogma.api.messaging.v1.ConsumeEventsResponse")
}

func init() { proto.RegisterFile("api/internal/pb/messaging.proto", fileDescriptor_141cdf861d5078a2) }

var fileDescriptor_141cdf861d5078a2 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0xb3, 0x62, 0x88, 0x2c, 0xd1, 0x43, 0x23, 0xa4, 0x69, 0xa2, 0x92, 0x5e, 0xc4, 0x28,
	0x5b, 0xc5, 0x37, 0x80, 0x70, 0xf0, 0xc0, 0xa5, 0x47, 0x6f, 0x0b, 0x0c, 0x75, 0x13, 0x76, 0x76,
	0x61, 0xa7, 0x04, 0x5f, 0xc1, 0xa7, 0xf0, 0x01, 0x7d, 0x08, 0x63, 0x5b, 0x15, 0x9a, 0x92, 0x70,
	0xe0, 0xb8, 0x3b, 0xff, 0xff, 0xef, 0xb7, 0x33, 0xc3, 0x6f, 0xa4, 0x55, 0x91, 0x42, 0x82, 0x15,
	0xca, 0x45, 0x64, 0x27, 0x91, 0x06, 0xe7, 0x64, 0xa2, 0x30, 0x11, 0x76, 0x65, 0xc8, 0x78, 0xed,
	0x99, 0x49, 0xb4, 0x14, 0xd2, 0x2a, 0xf1, 0x5f, 0x5a, 0x3f, 0x05, 0x57, 0xd5, 0x46, 0xc8, 0x6d,
	0xc1, 0x75, 0xb9, 0xac, 0x66, 0x80, 0xa4, 0xe8, 0x3d, 0xaf, 0x87, 0x9f, 0x8c, 0xb7, 0x46, 0x1b,
	0x98, 0xa6, 0x04, 0x43, 0xa3, 0xb5, 0xc4, 0x59, 0x0c, 0xcb, 0x14, 0x1c, 0x79, 0x03, 0xde, 0x94,
	0xd6, 0x2e, 0xd4, 0x54, 0x92, 0x32, 0xe8, 0xb3, 0x0e, 0xeb, 0x36, 0xfb, 0x1d, 0x51, 0x8d, 0x21,
	0x5e, 0x8a, 0xd8, 0x78, 0xdb, 0xe4, 0x0d, 0xf9, 0x19, 0xe0, 0x1a, 0x16, 0xc6, 0x82, 0x7f, 0x92,
	0x05, 0xdc, 0xee, 0x0b, 0x18, 0xe7, 0xd8, 0xa3, 0x42, 0x1e, 0xff, 0x19, 0x43, 0x9f, 0xb7, 0xcb,
	0x84, 0xce, 0x1a, 0x74, 0x10, 0x7e, 0x30, 0x7e, 0x39, 0x34, 0xe8, 0x52, 0x0d, 0xa3, 0x35, 0x20,
	0xb9, 0x63, 0xb2, 0xb7, 0x79, 0xdd, 0xcc, 0xe7, 0x0e, 0x28, 0x23, 0x3f, 0x8d, 0x8b, 0xd3, 0xcf,
	0x3d, 0x64, 0x8f, 0xf9, 0xb5, 0x4e, 0xad, 0xdb, 0x88, 0x8b, 0x53, 0x48, 0xbc, 0x55, 0x62, 0xc9,
	0x29, 0xb7, 0x82, 0xd8, 0x4e, 0xd0, 0x31, 0x9a, 0xd3, 0xff, 0x62, 0xbc, 0x31, 0xfe, 0x95, 0x7a,
	0x86, 0x5f, 0xec, 0xb6, 0xca, 0xeb, 0xed, 0x8b, 0xac, 0x1c, 0x7a, 0x20, 0x0e, 0x95, 0x17, 0x7f,
	0x43, 0x7e, 0xbe, 0xf3, 0x69, 0xef, 0x61, 0x5f, 0x40, 0xd5, 0x9c, 0x82, 0xde, 0x81, 0xea, 0xfc,
	0xb5, 0x47, 0x36, 0xb8, 0x7f, 0xbd, 0x4b, 0x14, 0xbd, 0xa5, 0x13, 0x31, 0x35, 0x3a, 0xca, 0xcc,
	0xa4, 0x96, 0x91, 0xc2, 0xb9, 0xda, 0x44, 0xa5, 0x55, 0x9f, 0xd4, 0xb3, 0x15, 0x7f, 0xfe, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x07, 0x24, 0x46, 0x07, 0x5c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessagingClient is the client API for Messaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessagingClient interface {
	// ExecuteCommand executes a command message.
	ExecuteCommand(ctx context.Context, in *ExecuteCommandRequest, opts ...grpc.CallOption) (*ExecuteCommandResponse, error)
	// ConsumesEvents streams events from an applications event stream.
	ConsumeEvents(ctx context.Context, in *ConsumeEventsRequest, opts ...grpc.CallOption) (Messaging_ConsumeEventsClient, error)
}

type messagingClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingClient(cc grpc.ClientConnInterface) MessagingClient {
	return &messagingClient{cc}
}

func (c *messagingClient) ExecuteCommand(ctx context.Context, in *ExecuteCommandRequest, opts ...grpc.CallOption) (*ExecuteCommandResponse, error) {
	out := new(ExecuteCommandResponse)
	err := c.cc.Invoke(ctx, "/dogma.api.messaging.v1.Messaging/ExecuteCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) ConsumeEvents(ctx context.Context, in *ConsumeEventsRequest, opts ...grpc.CallOption) (Messaging_ConsumeEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Messaging_serviceDesc.Streams[0], "/dogma.api.messaging.v1.Messaging/ConsumeEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &messagingConsumeEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messaging_ConsumeEventsClient interface {
	Recv() (*ConsumeEventsResponse, error)
	grpc.ClientStream
}

type messagingConsumeEventsClient struct {
	grpc.ClientStream
}

func (x *messagingConsumeEventsClient) Recv() (*ConsumeEventsResponse, error) {
	m := new(ConsumeEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessagingServer is the server API for Messaging service.
type MessagingServer interface {
	// ExecuteCommand executes a command message.
	ExecuteCommand(context.Context, *ExecuteCommandRequest) (*ExecuteCommandResponse, error)
	// ConsumesEvents streams events from an applications event stream.
	ConsumeEvents(*ConsumeEventsRequest, Messaging_ConsumeEventsServer) error
}

// UnimplementedMessagingServer can be embedded to have forward compatible implementations.
type UnimplementedMessagingServer struct {
}

func (*UnimplementedMessagingServer) ExecuteCommand(ctx context.Context, req *ExecuteCommandRequest) (*ExecuteCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCommand not implemented")
}
func (*UnimplementedMessagingServer) ConsumeEvents(req *ConsumeEventsRequest, srv Messaging_ConsumeEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsumeEvents not implemented")
}

func RegisterMessagingServer(s *grpc.Server, srv MessagingServer) {
	s.RegisterService(&_Messaging_serviceDesc, srv)
}

func _Messaging_ExecuteCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).ExecuteCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dogma.api.messaging.v1.Messaging/ExecuteCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).ExecuteCommand(ctx, req.(*ExecuteCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_ConsumeEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagingServer).ConsumeEvents(m, &messagingConsumeEventsServer{stream})
}

type Messaging_ConsumeEventsServer interface {
	Send(*ConsumeEventsResponse) error
	grpc.ServerStream
}

type messagingConsumeEventsServer struct {
	grpc.ServerStream
}

func (x *messagingConsumeEventsServer) Send(m *ConsumeEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Messaging_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dogma.api.messaging.v1.Messaging",
	HandlerType: (*MessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteCommand",
			Handler:    _Messaging_ExecuteCommand_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConsumeEvents",
			Handler:       _Messaging_ConsumeEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/internal/pb/messaging.proto",
}