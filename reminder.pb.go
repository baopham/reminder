// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reminder.proto

/*
Package reminder is a generated protocol buffer package.

It is generated from these files:
	reminder.proto

It has these top-level messages:
	GetRequest
	GetResponse
	CreateRequest
	CreateResponse
	Reminder
*/
package reminder

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetResponse struct {
	Reminder *Reminder `protobuf:"bytes,1,opt,name=reminder" json:"reminder,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetResponse) GetReminder() *Reminder {
	if m != nil {
		return m.Reminder
	}
	return nil
}

type CreateRequest struct {
	UserId      string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Url         string `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type CreateResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Reminder struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=userId" json:"userId,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Url         string `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
}

func (m *Reminder) Reset()                    { *m = Reminder{} }
func (m *Reminder) String() string            { return proto.CompactTextString(m) }
func (*Reminder) ProtoMessage()               {}
func (*Reminder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Reminder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Reminder) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Reminder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Reminder) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Reminder) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "reminder.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "reminder.GetResponse")
	proto.RegisterType((*CreateRequest)(nil), "reminder.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "reminder.CreateResponse")
	proto.RegisterType((*Reminder)(nil), "reminder.Reminder")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ReminderService service

type ReminderServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type reminderServiceClient struct {
	cc *grpc.ClientConn
}

func NewReminderServiceClient(cc *grpc.ClientConn) ReminderServiceClient {
	return &reminderServiceClient{cc}
}

func (c *reminderServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/reminder.ReminderService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reminderServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/reminder.ReminderService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ReminderService service

type ReminderServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
}

func RegisterReminderServiceServer(s *grpc.Server, srv ReminderServiceServer) {
	s.RegisterService(&_ReminderService_serviceDesc, srv)
}

func _ReminderService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReminderServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reminder.ReminderService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReminderServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReminderService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReminderServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reminder.ReminderService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReminderServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReminderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reminder.ReminderService",
	HandlerType: (*ReminderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ReminderService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ReminderService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reminder.proto",
}

func init() { proto.RegisterFile("reminder.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x3b, 0x4f, 0xc3, 0x30,
	0x10, 0xc7, 0xc9, 0x83, 0xa8, 0x5c, 0x44, 0x40, 0x27, 0x1e, 0x56, 0x01, 0x29, 0xca, 0xc4, 0x94,
	0x4a, 0x85, 0xb5, 0x0b, 0x0c, 0x15, 0x6b, 0xd8, 0xd8, 0xf2, 0x38, 0x51, 0x4b, 0x24, 0x0e, 0xb6,
	0xc3, 0xc0, 0x07, 0xe0, 0x73, 0x23, 0xdc, 0x38, 0x29, 0x81, 0x6e, 0x77, 0xf7, 0xbf, 0xc7, 0xef,
	0xce, 0x86, 0x48, 0x52, 0xcd, 0x9b, 0x8a, 0x64, 0xda, 0x4a, 0xa1, 0x05, 0xce, 0xac, 0x9f, 0x5c,
	0x03, 0xac, 0x49, 0x67, 0xf4, 0xde, 0x91, 0xd2, 0x18, 0x81, 0xcb, 0x2b, 0xe6, 0xc4, 0xce, 0xed,
	0x51, 0xe6, 0xf2, 0x2a, 0x59, 0x41, 0x68, 0x54, 0xd5, 0x8a, 0x46, 0x11, 0xa6, 0x30, 0x14, 0x9a,
	0xa4, 0x70, 0x89, 0xe9, 0xd0, 0x39, 0xeb, 0x8d, 0x6c, 0x6c, 0x2e, 0xe0, 0xf8, 0x51, 0x52, 0xae,
	0xc9, 0xf6, 0xbf, 0x80, 0xa0, 0x53, 0x24, 0x9f, 0xec, 0x8c, 0xde, 0x43, 0x04, 0xbf, 0xc9, 0x6b,
	0x62, 0xae, 0x89, 0x1a, 0x1b, 0x63, 0x08, 0x2b, 0x52, 0xa5, 0xe4, 0xad, 0xe6, 0xa2, 0x61, 0x9e,
	0x91, 0x76, 0x43, 0x78, 0x0a, 0x5e, 0x27, 0xdf, 0x98, 0x6f, 0x94, 0x1f, 0x33, 0x89, 0x21, 0xb2,
	0x03, 0x7b, 0xe4, 0xe9, 0x46, 0x9f, 0x30, 0xb3, 0xa0, 0x53, 0x6d, 0x87, 0xce, 0xfd, 0x97, 0xce,
	0xdb, 0x4f, 0xe7, 0xef, 0xa5, 0x3b, 0x1c, 0xe8, 0x96, 0x5f, 0x0e, 0x9c, 0xd8, 0xe1, 0xcf, 0x24,
	0x3f, 0x78, 0x49, 0x78, 0x0f, 0xde, 0x9a, 0x34, 0x9e, 0x8d, 0x77, 0x1c, 0x9f, 0x63, 0x7e, 0x3e,
	0x89, 0x6e, 0x77, 0x4a, 0x0e, 0x70, 0x05, 0xc1, 0x76, 0x4f, 0xbc, 0x1c, 0x53, 0x7e, 0x9d, 0x7a,
	0xce, 0xfe, 0x0a, 0xb6, 0xfc, 0xe1, 0xe6, 0xe5, 0xea, 0x95, 0xeb, 0x4d, 0x57, 0xa4, 0xa5, 0xa8,
	0x17, 0x45, 0x2e, 0xda, 0x4d, 0x5e, 0x2f, 0x6c, 0x7e, 0x11, 0x98, 0x4f, 0x72, 0xf7, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x59, 0xb7, 0x5d, 0x03, 0x36, 0x02, 0x00, 0x00,
}
