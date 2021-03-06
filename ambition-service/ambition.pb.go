// Code generated by protoc-gen-go.
// source: ambition.proto
// DO NOT EDIT!

/*
Package ambition is a generated protocol buffer package.

It is generated from these files:
	ambition.proto

It has these top-level messages:
	OccurrencesByDateReq
	Action
	CreateOccurrenceRequest
	Occurrence
	User
	ActionsResponse
	OccurrencesResponse
*/
package ambition

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/TuneLab/go-truss/deftree/googlethirdparty"

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

type OccurrencesByDateReq struct {
	ActionID  int64  `protobuf:"varint,1,opt,name=ActionID" json:"ActionID,omitempty"`
	StartDate string `protobuf:"bytes,2,opt,name=StartDate" json:"StartDate,omitempty"`
	EndDate   string `protobuf:"bytes,3,opt,name=EndDate" json:"EndDate,omitempty"`
}

func (m *OccurrencesByDateReq) Reset()                    { *m = OccurrencesByDateReq{} }
func (m *OccurrencesByDateReq) String() string            { return proto.CompactTextString(m) }
func (*OccurrencesByDateReq) ProtoMessage()               {}
func (*OccurrencesByDateReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OccurrencesByDateReq) GetActionID() int64 {
	if m != nil {
		return m.ActionID
	}
	return 0
}

func (m *OccurrencesByDateReq) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *OccurrencesByDateReq) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

type Action struct {
	ID   int64  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	// TODO: Think about moving this to ambition-users
	// with a UserAction table
	UserID int64 `protobuf:"varint,3,opt,name=UserID" json:"UserID,omitempty"`
}

func (m *Action) Reset()                    { *m = Action{} }
func (m *Action) String() string            { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()               {}
func (*Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Action) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Action) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Action) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type CreateOccurrenceRequest struct {
	UserID     int64       `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	Occurrence *Occurrence `protobuf:"bytes,2,opt,name=Occurrence" json:"Occurrence,omitempty"`
}

func (m *CreateOccurrenceRequest) Reset()                    { *m = CreateOccurrenceRequest{} }
func (m *CreateOccurrenceRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateOccurrenceRequest) ProtoMessage()               {}
func (*CreateOccurrenceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateOccurrenceRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CreateOccurrenceRequest) GetOccurrence() *Occurrence {
	if m != nil {
		return m.Occurrence
	}
	return nil
}

type Occurrence struct {
	ID       int64  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	ActionID int64  `protobuf:"varint,2,opt,name=ActionID" json:"ActionID,omitempty"`
	Datetime string `protobuf:"bytes,3,opt,name=Datetime" json:"Datetime,omitempty"`
	Data     string `protobuf:"bytes,4,opt,name=Data" json:"Data,omitempty"`
}

func (m *Occurrence) Reset()                    { *m = Occurrence{} }
func (m *Occurrence) String() string            { return proto.CompactTextString(m) }
func (*Occurrence) ProtoMessage()               {}
func (*Occurrence) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Occurrence) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Occurrence) GetActionID() int64 {
	if m != nil {
		return m.ActionID
	}
	return 0
}

func (m *Occurrence) GetDatetime() string {
	if m != nil {
		return m.Datetime
	}
	return ""
}

func (m *Occurrence) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type User struct {
	UserID int64 `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *User) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type ActionsResponse struct {
	Actions []*Action `protobuf:"bytes,1,rep,name=Actions" json:"Actions,omitempty"`
}

func (m *ActionsResponse) Reset()                    { *m = ActionsResponse{} }
func (m *ActionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ActionsResponse) ProtoMessage()               {}
func (*ActionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ActionsResponse) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

type OccurrencesResponse struct {
	Occurrences []*Occurrence `protobuf:"bytes,1,rep,name=Occurrences" json:"Occurrences,omitempty"`
}

func (m *OccurrencesResponse) Reset()                    { *m = OccurrencesResponse{} }
func (m *OccurrencesResponse) String() string            { return proto.CompactTextString(m) }
func (*OccurrencesResponse) ProtoMessage()               {}
func (*OccurrencesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *OccurrencesResponse) GetOccurrences() []*Occurrence {
	if m != nil {
		return m.Occurrences
	}
	return nil
}

func init() {
	proto.RegisterType((*OccurrencesByDateReq)(nil), "ambition.OccurrencesByDateReq")
	proto.RegisterType((*Action)(nil), "ambition.Action")
	proto.RegisterType((*CreateOccurrenceRequest)(nil), "ambition.CreateOccurrenceRequest")
	proto.RegisterType((*Occurrence)(nil), "ambition.Occurrence")
	proto.RegisterType((*User)(nil), "ambition.User")
	proto.RegisterType((*ActionsResponse)(nil), "ambition.ActionsResponse")
	proto.RegisterType((*OccurrencesResponse)(nil), "ambition.OccurrencesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Ambition service

type AmbitionClient interface {
	// CreateAction requires a UserID and a Name
	CreateAction(ctx context.Context, in *Action, opts ...grpc.CallOption) (*Action, error)
	// CreateOccurrence requires a UserID and Occurrence.ActionID
	// TODO: If Datetime is provided it will be used
	// TODO: If Data is provided it will be stored
	CreateOccurrence(ctx context.Context, in *CreateOccurrenceRequest, opts ...grpc.CallOption) (*Occurrence, error)
	// ReadAction requires either an ID, or BOTH a UserId and Name
	ReadAction(ctx context.Context, in *Action, opts ...grpc.CallOption) (*Action, error)
	// ReadAction
	// TODO:
	ReadActions(ctx context.Context, in *User, opts ...grpc.CallOption) (*ActionsResponse, error)
	ReadOccurrencesByDate(ctx context.Context, in *OccurrencesByDateReq, opts ...grpc.CallOption) (*OccurrencesResponse, error)
	// ReadOccurrences takes an action which must be populated with a
	// UserID and an ActionID which must match the values for that action
	// TODO:
	ReadOccurrences(ctx context.Context, in *Action, opts ...grpc.CallOption) (*OccurrencesResponse, error)
}

type ambitionClient struct {
	cc *grpc.ClientConn
}

func NewAmbitionClient(cc *grpc.ClientConn) AmbitionClient {
	return &ambitionClient{cc}
}

func (c *ambitionClient) CreateAction(ctx context.Context, in *Action, opts ...grpc.CallOption) (*Action, error) {
	out := new(Action)
	err := grpc.Invoke(ctx, "/ambition.Ambition/CreateAction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ambitionClient) CreateOccurrence(ctx context.Context, in *CreateOccurrenceRequest, opts ...grpc.CallOption) (*Occurrence, error) {
	out := new(Occurrence)
	err := grpc.Invoke(ctx, "/ambition.Ambition/CreateOccurrence", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ambitionClient) ReadAction(ctx context.Context, in *Action, opts ...grpc.CallOption) (*Action, error) {
	out := new(Action)
	err := grpc.Invoke(ctx, "/ambition.Ambition/ReadAction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ambitionClient) ReadActions(ctx context.Context, in *User, opts ...grpc.CallOption) (*ActionsResponse, error) {
	out := new(ActionsResponse)
	err := grpc.Invoke(ctx, "/ambition.Ambition/ReadActions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ambitionClient) ReadOccurrencesByDate(ctx context.Context, in *OccurrencesByDateReq, opts ...grpc.CallOption) (*OccurrencesResponse, error) {
	out := new(OccurrencesResponse)
	err := grpc.Invoke(ctx, "/ambition.Ambition/ReadOccurrencesByDate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ambitionClient) ReadOccurrences(ctx context.Context, in *Action, opts ...grpc.CallOption) (*OccurrencesResponse, error) {
	out := new(OccurrencesResponse)
	err := grpc.Invoke(ctx, "/ambition.Ambition/ReadOccurrences", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ambition service

type AmbitionServer interface {
	// CreateAction requires a UserID and a Name
	CreateAction(context.Context, *Action) (*Action, error)
	// CreateOccurrence requires a UserID and Occurrence.ActionID
	// TODO: If Datetime is provided it will be used
	// TODO: If Data is provided it will be stored
	CreateOccurrence(context.Context, *CreateOccurrenceRequest) (*Occurrence, error)
	// ReadAction requires either an ID, or BOTH a UserId and Name
	ReadAction(context.Context, *Action) (*Action, error)
	// ReadAction
	// TODO:
	ReadActions(context.Context, *User) (*ActionsResponse, error)
	ReadOccurrencesByDate(context.Context, *OccurrencesByDateReq) (*OccurrencesResponse, error)
	// ReadOccurrences takes an action which must be populated with a
	// UserID and an ActionID which must match the values for that action
	// TODO:
	ReadOccurrences(context.Context, *Action) (*OccurrencesResponse, error)
}

func RegisterAmbitionServer(s *grpc.Server, srv AmbitionServer) {
	s.RegisterService(&_Ambition_serviceDesc, srv)
}

func _Ambition_CreateAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Action)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbitionServer).CreateAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ambition.Ambition/CreateAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbitionServer).CreateAction(ctx, req.(*Action))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ambition_CreateOccurrence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOccurrenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbitionServer).CreateOccurrence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ambition.Ambition/CreateOccurrence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbitionServer).CreateOccurrence(ctx, req.(*CreateOccurrenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ambition_ReadAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Action)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbitionServer).ReadAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ambition.Ambition/ReadAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbitionServer).ReadAction(ctx, req.(*Action))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ambition_ReadActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbitionServer).ReadActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ambition.Ambition/ReadActions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbitionServer).ReadActions(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ambition_ReadOccurrencesByDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OccurrencesByDateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbitionServer).ReadOccurrencesByDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ambition.Ambition/ReadOccurrencesByDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbitionServer).ReadOccurrencesByDate(ctx, req.(*OccurrencesByDateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ambition_ReadOccurrences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Action)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbitionServer).ReadOccurrences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ambition.Ambition/ReadOccurrences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbitionServer).ReadOccurrences(ctx, req.(*Action))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ambition_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ambition.Ambition",
	HandlerType: (*AmbitionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAction",
			Handler:    _Ambition_CreateAction_Handler,
		},
		{
			MethodName: "CreateOccurrence",
			Handler:    _Ambition_CreateOccurrence_Handler,
		},
		{
			MethodName: "ReadAction",
			Handler:    _Ambition_ReadAction_Handler,
		},
		{
			MethodName: "ReadActions",
			Handler:    _Ambition_ReadActions_Handler,
		},
		{
			MethodName: "ReadOccurrencesByDate",
			Handler:    _Ambition_ReadOccurrencesByDate_Handler,
		},
		{
			MethodName: "ReadOccurrences",
			Handler:    _Ambition_ReadOccurrences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ambition.proto",
}

func init() { proto.RegisterFile("ambition.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0xed, 0x28, 0x4d, 0x27, 0x55, 0x5a, 0x0d, 0x01, 0x8c, 0x05, 0x55, 0xd8, 0xa7, 0x08,
	0x89, 0x58, 0x0a, 0x15, 0x0f, 0x48, 0x3c, 0x14, 0x0c, 0x52, 0x24, 0x0a, 0x92, 0x81, 0x03, 0x6c,
	0x9c, 0xc5, 0x31, 0x6a, 0xbc, 0xe9, 0xee, 0xfa, 0xa1, 0xaf, 0x5c, 0x81, 0xa3, 0x70, 0x14, 0xae,
	0xc0, 0x41, 0xd0, 0xfa, 0x6f, 0x57, 0x89, 0x2b, 0xc4, 0xdb, 0xce, 0xcc, 0x37, 0xdf, 0x7e, 0xf3,
	0x8d, 0x06, 0xc6, 0x74, 0xbb, 0xca, 0x54, 0xc6, 0xf3, 0xf9, 0x4e, 0x70, 0xc5, 0x71, 0xd8, 0xc4,
	0xc1, 0xfb, 0x34, 0x53, 0x9b, 0x62, 0x35, 0x4f, 0xf8, 0x36, 0xfc, 0x52, 0xe4, 0xec, 0x03, 0x5d,
	0x85, 0x29, 0x7f, 0xae, 0x44, 0x21, 0x65, 0xb8, 0x66, 0xdf, 0x94, 0x60, 0x2c, 0x4c, 0x39, 0x4f,
	0xaf, 0x99, 0xda, 0x64, 0x62, 0xbd, 0xa3, 0x42, 0xdd, 0x86, 0x34, 0xcf, 0xb9, 0xa2, 0x9a, 0x40,
	0x56, 0x8c, 0xe4, 0x3b, 0x4c, 0x3e, 0x25, 0x49, 0x21, 0x04, 0xcb, 0x13, 0x26, 0xdf, 0xdc, 0x46,
	0x54, 0xb1, 0x98, 0xdd, 0x60, 0x00, 0xc3, 0xcb, 0x44, 0x03, 0x97, 0x91, 0xef, 0x4c, 0x9d, 0x99,
	0x17, 0xb7, 0x31, 0x3e, 0x86, 0xe3, 0xcf, 0x8a, 0x0a, 0xa5, 0xb1, 0xbe, 0x3b, 0x75, 0x66, 0xc7,
	0xb1, 0x49, 0xa0, 0x0f, 0x47, 0xef, 0xf2, 0x75, 0x59, 0xf3, 0xca, 0x5a, 0x13, 0x92, 0x08, 0x06,
	0x15, 0x07, 0x8e, 0xc1, 0x6d, 0x79, 0xdd, 0x65, 0x84, 0x08, 0xfd, 0x8f, 0x74, 0xdb, 0x90, 0x95,
	0x6f, 0x7c, 0x00, 0x83, 0xaf, 0x92, 0x89, 0x65, 0x54, 0xd2, 0x78, 0x71, 0x1d, 0x91, 0x14, 0x1e,
	0xbe, 0x15, 0x8c, 0x2a, 0x66, 0x74, 0xc7, 0xec, 0xa6, 0x60, 0x52, 0x59, 0x2d, 0x8e, 0xdd, 0x82,
	0x17, 0x00, 0x06, 0x5c, 0x7e, 0x32, 0x5a, 0x4c, 0xe6, 0xad, 0xb7, 0x16, 0x91, 0x85, 0x23, 0x1b,
	0xbb, 0xeb, 0x40, 0xb2, 0x6d, 0x90, 0xbb, 0x67, 0x50, 0x00, 0x43, 0x3d, 0xb0, 0xca, 0xb6, 0x8d,
	0x07, 0x6d, 0xac, 0x47, 0x8d, 0xa8, 0xa2, 0x7e, 0xbf, 0x1a, 0x55, 0xbf, 0xc9, 0x39, 0xf4, 0xb5,
	0xd2, 0xbb, 0xf4, 0x93, 0xd7, 0x70, 0x5a, 0x71, 0xcb, 0x98, 0xc9, 0x1d, 0xcf, 0x25, 0xc3, 0x67,
	0x70, 0x54, 0xa7, 0x7c, 0x67, 0xea, 0xcd, 0x46, 0x8b, 0x33, 0x33, 0x4f, 0x55, 0x88, 0x1b, 0x00,
	0xb9, 0x82, 0x7b, 0xd6, 0x8e, 0x5b, 0x8a, 0x97, 0x30, 0xb2, 0xd2, 0x35, 0x4d, 0xb7, 0x2d, 0x36,
	0x70, 0xf1, 0xcb, 0x83, 0xe1, 0x65, 0x0d, 0xc2, 0x0b, 0x38, 0xa9, 0xb6, 0x51, 0x6f, 0xf6, 0x40,
	0x46, 0x70, 0x90, 0x21, 0x3d, 0xbc, 0x82, 0xb3, 0xfd, 0x1d, 0xe2, 0x53, 0x83, 0xbb, 0x63, 0xbf,
	0x41, 0xa7, 0x38, 0xd2, 0xc3, 0x05, 0x40, 0xcc, 0xe8, 0xfa, 0xbf, 0x24, 0xbc, 0x82, 0x91, 0xe9,
	0x91, 0x38, 0x36, 0x10, 0x6d, 0x7a, 0xf0, 0x68, 0xbf, 0xa5, 0xf5, 0x8d, 0xf4, 0xf0, 0x1a, 0xee,
	0xeb, 0xde, 0x83, 0xc3, 0xc1, 0xf3, 0x2e, 0x81, 0xe6, 0xaa, 0x82, 0x27, 0x9d, 0xf5, 0x96, 0x79,
	0xf2, 0xe3, 0xf7, 0x9f, 0x9f, 0xee, 0x18, 0x4f, 0x42, 0x6e, 0xaa, 0x18, 0xc1, 0xe9, 0xde, 0x6f,
	0x1d, 0x23, 0xfe, 0x83, 0xb9, 0xb7, 0x1a, 0x94, 0xf7, 0xfe, 0xe2, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x49, 0xc8, 0xf4, 0xaf, 0x53, 0x04, 0x00, 0x00,
}
