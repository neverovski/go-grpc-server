// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/user/pb/user.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	api/user/pb/user.proto

It has these top-level messages:
	User
	GetUserRequest
	GetUserResponse
	PostUserRequest
	PostUserResponse
	UpdateUserRequest
	UpdateUserResponse
	DeleteUserRequest
	DeleteUserResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

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

type User struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	// @inject_tag: sql:"type:timestamptz,default:now()"
	CreatedAt *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	// @inject_tag: sql:"type:timestamptz"
	UpdatedAt *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type GetUserRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetUserRequest) Reset()                    { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()               {}
func (*GetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetUserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *GetUserResponse) Reset()                    { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()               {}
func (*GetUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type PostUserRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *PostUserRequest) Reset()                    { *m = PostUserRequest{} }
func (m *PostUserRequest) String() string            { return proto.CompactTextString(m) }
func (*PostUserRequest) ProtoMessage()               {}
func (*PostUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PostUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *PostUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type PostUserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *PostUserResponse) Reset()                    { *m = PostUserResponse{} }
func (m *PostUserResponse) String() string            { return proto.CompactTextString(m) }
func (*PostUserResponse) ProtoMessage()               {}
func (*PostUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PostUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserRequest struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *UpdateUserRequest) Reset()                    { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()               {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UpdateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UpdateUserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *UpdateUserResponse) Reset()                    { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()               {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteUserRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteUserRequest) Reset()                    { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()               {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteUserResponse struct {
}

func (m *DeleteUserResponse) Reset()                    { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()               {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*GetUserRequest)(nil), "pb.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "pb.GetUserResponse")
	proto.RegisterType((*PostUserRequest)(nil), "pb.PostUserRequest")
	proto.RegisterType((*PostUserResponse)(nil), "pb.PostUserResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "pb.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "pb.UpdateUserResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "pb.DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "pb.DeleteUserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	PostUser(ctx context.Context, in *PostUserRequest, opts ...grpc.CallOption) (*PostUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PostUser(ctx context.Context, in *PostUserRequest, opts ...grpc.CallOption) (*PostUserResponse, error) {
	out := new(PostUserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/PostUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	PostUser(context.Context, *PostUserRequest) (*PostUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PostUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PostUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/PostUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PostUser(ctx, req.(*PostUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "PostUser",
			Handler:    _UserService_PostUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/pb/user.proto",
}

func init() { proto.RegisterFile("api/user/pb/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x55, 0xc2, 0x16, 0xb6, 0xb3, 0xa8, 0x65, 0xa7, 0x4b, 0x59, 0xac, 0x4a, 0xac, 0xcc, 0xa5,
	0xea, 0x21, 0x81, 0xe5, 0x04, 0xb7, 0x0a, 0x24, 0x4e, 0x20, 0xb4, 0x74, 0xb9, 0x22, 0xa7, 0x19,
	0x2a, 0x4b, 0x9b, 0xd8, 0xc4, 0x4e, 0x2f, 0x88, 0x0b, 0x7f, 0x81, 0xff, 0xc3, 0x9f, 0xe0, 0xc6,
	0x99, 0x1f, 0x82, 0x62, 0x27, 0xdd, 0x7c, 0x88, 0x52, 0x4e, 0x91, 0x9f, 0xdf, 0xbc, 0x99, 0xf7,
	0xc6, 0x81, 0x43, 0xa1, 0x65, 0x5c, 0x1a, 0x2a, 0x62, 0x9d, 0xb8, 0x6f, 0xa4, 0x0b, 0x65, 0x15,
	0x86, 0x3a, 0x61, 0x47, 0x17, 0x4a, 0x5d, 0x6c, 0x28, 0xae, 0x28, 0x22, 0xcf, 0x95, 0x15, 0x56,
	0xaa, 0xdc, 0x78, 0x06, 0x7b, 0x54, 0xdf, 0xba, 0x53, 0x52, 0x7e, 0x8a, 0xad, 0xcc, 0xc8, 0x58,
	0x91, 0x69, 0x4f, 0xe0, 0x3f, 0x02, 0x18, 0xad, 0x0d, 0x15, 0xb8, 0x07, 0xa1, 0x4c, 0xe7, 0xc1,
	0x22, 0x38, 0xde, 0x5d, 0x85, 0x32, 0x45, 0x06, 0xe3, 0xaa, 0x53, 0x2e, 0x32, 0x9a, 0x87, 0x0e,
	0xbd, 0x3a, 0xe3, 0x0c, 0x76, 0x28, 0x13, 0x72, 0x33, 0xbf, 0xe5, 0x2e, 0xfc, 0x01, 0x9f, 0x03,
	0x9c, 0x17, 0x24, 0x2c, 0xa5, 0x1f, 0x85, 0x9d, 0x8f, 0x16, 0xc1, 0xf1, 0x64, 0xc9, 0x22, 0x3f,
	0x40, 0xd4, 0x0c, 0x10, 0x9d, 0x35, 0x03, 0xac, 0x76, 0x6b, 0xf6, 0xa9, 0xad, 0x4a, 0x4b, 0x9d,
	0x36, 0xa5, 0x3b, 0xff, 0x2e, 0xad, 0xd9, 0xa7, 0x96, 0x2f, 0x60, 0xef, 0x35, 0xd9, 0xca, 0xc2,
	0x8a, 0x3e, 0x97, 0x64, 0x6c, 0xdf, 0x09, 0x8f, 0x61, 0xff, 0x8a, 0x61, 0xb4, 0xca, 0x0d, 0xe1,
	0x11, 0x8c, 0x2a, 0x33, 0x8e, 0x34, 0x59, 0x8e, 0x23, 0x9d, 0x44, 0xee, 0xde, 0xa1, 0xfc, 0x25,
	0xec, 0xbf, 0x53, 0xa6, 0xa3, 0xd9, 0x4e, 0x23, 0xf8, 0x5b, 0x1a, 0x61, 0x2b, 0x0d, 0xfe, 0x04,
	0xee, 0x6d, 0x45, 0x6e, 0xd4, 0x76, 0x0d, 0xd3, 0xb5, 0xb3, 0x75, 0x8d, 0x99, 0xff, 0x5f, 0x0b,
	0x5f, 0x02, 0xb6, 0x65, 0x6f, 0x34, 0xca, 0x63, 0x98, 0xbe, 0xa2, 0x0d, 0x5d, 0x3b, 0x0a, 0x9f,
	0x01, 0xb6, 0x49, 0x5e, 0x78, 0xf9, 0x2b, 0x84, 0x49, 0x05, 0xbc, 0xa7, 0xe2, 0x52, 0x9e, 0x13,
	0xbe, 0x81, 0x3b, 0x75, 0xfa, 0x88, 0x55, 0x97, 0xee, 0xb2, 0xd8, 0x41, 0x07, 0xf3, 0x1a, 0xfc,
	0xe1, 0xb7, 0x9f, 0xbf, 0xbf, 0x87, 0x07, 0x38, 0x75, 0xaf, 0xfa, 0xf2, 0xa9, 0x7f, 0xfb, 0x5f,
	0x64, 0xfa, 0x15, 0xdf, 0xc2, 0xb8, 0x89, 0x15, 0x5d, 0x6d, 0x6f, 0x53, 0x6c, 0xd6, 0x05, 0x6b,
	0xc5, 0x07, 0x4e, 0x71, 0xca, 0xef, 0xb6, 0x15, 0x5f, 0x04, 0x27, 0x78, 0x06, 0xb0, 0x4d, 0x07,
	0xef, 0xbb, 0x1c, 0xfa, 0x4b, 0x60, 0x87, 0x7d, 0xb8, 0xab, 0xca, 0x06, 0xaa, 0x1f, 0x00, 0xb6,
	0xd1, 0x78, 0xd5, 0x41, 0x9e, 0x5e, 0x75, 0x98, 0x60, 0xe3, 0xfe, 0x64, 0xe8, 0x3e, 0xb9, 0xed,
	0xfe, 0x85, 0x67, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x57, 0xbb, 0xd1, 0xb6, 0x11, 0x04, 0x00,
	0x00,
}
