// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jsonmultiplex/iam/v0/iam.proto

package jsonmultiplex_iam_v0

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Account struct {
	Name       string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DeleteTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// Write-only field, required only for CreateAccount. Never returned to the
	// client.
	RootPassword         string   `protobuf:"bytes,5,opt,name=root_password,json=rootPassword,proto3" json:"root_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b7e58ef322f61c, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Account) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Account) GetDeleteTime() *timestamp.Timestamp {
	if m != nil {
		return m.DeleteTime
	}
	return nil
}

func (m *Account) GetRootPassword() string {
	if m != nil {
		return m.RootPassword
	}
	return ""
}

type User struct {
	Name       string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DeleteTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// Write-only field, required only for CreateUser and optional for UpdateUser.
	// Never returned to the client.
	Password             string   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b7e58ef322f61c, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *User) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *User) GetDeleteTime() *timestamp.Timestamp {
	if m != nil {
		return m.DeleteTime
	}
	return nil
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Session struct {
	Name       string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	ExpireTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	Account    string               `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	User       string               `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	// Write-only field, required only for CreateSession. Never returned to the
	// client.
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	// Read-only field, returned only from CreateSession.
	Token                string   `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b7e58ef322f61c, []int{2}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Session) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Session) GetExpireTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

func (m *Session) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Session) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Session) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Session) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CreateAccountRequest struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b7e58ef322f61c, []int{3}
}

func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (m *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(m, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type ListUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b7e58ef322f61c, []int{4}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

type ListUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b7e58ef322f61c, []int{5}
}

func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (m *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(m, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type GetUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b7e58ef322f61c, []int{6}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b7e58ef322f61c, []int{7}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateSessionRequest struct {
	Session              *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSessionRequest) Reset()         { *m = CreateSessionRequest{} }
func (m *CreateSessionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSessionRequest) ProtoMessage()    {}
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b7e58ef322f61c, []int{8}
}

func (m *CreateSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSessionRequest.Unmarshal(m, b)
}
func (m *CreateSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSessionRequest.Marshal(b, m, deterministic)
}
func (m *CreateSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSessionRequest.Merge(m, src)
}
func (m *CreateSessionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSessionRequest.Size(m)
}
func (m *CreateSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSessionRequest proto.InternalMessageInfo

func (m *CreateSessionRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "jsonmultiplex.iam.v0.Account")
	proto.RegisterType((*User)(nil), "jsonmultiplex.iam.v0.User")
	proto.RegisterType((*Session)(nil), "jsonmultiplex.iam.v0.Session")
	proto.RegisterType((*CreateAccountRequest)(nil), "jsonmultiplex.iam.v0.CreateAccountRequest")
	proto.RegisterType((*ListUsersRequest)(nil), "jsonmultiplex.iam.v0.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "jsonmultiplex.iam.v0.ListUsersResponse")
	proto.RegisterType((*GetUserRequest)(nil), "jsonmultiplex.iam.v0.GetUserRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "jsonmultiplex.iam.v0.CreateUserRequest")
	proto.RegisterType((*CreateSessionRequest)(nil), "jsonmultiplex.iam.v0.CreateSessionRequest")
}

func init() { proto.RegisterFile("jsonmultiplex/iam/v0/iam.proto", fileDescriptor_73b7e58ef322f61c) }

var fileDescriptor_73b7e58ef322f61c = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0xd6, 0x76, 0xa5, 0xaf, 0x1b, 0xa2, 0x56, 0xa5, 0x45, 0x11, 0x83, 0xca, 0x4c, 0x6c,
	0xda, 0x21, 0xa9, 0xc6, 0x01, 0xa9, 0x88, 0xc3, 0x34, 0x21, 0x84, 0x04, 0x02, 0x05, 0x38, 0x4f,
	0x59, 0x6b, 0xaa, 0x8c, 0x26, 0x0e, 0xb1, 0x33, 0x86, 0x10, 0x17, 0xfe, 0x02, 0xbf, 0x88, 0xdf,
	0xc0, 0x5f, 0xe0, 0xca, 0x19, 0x71, 0x43, 0x7e, 0xb6, 0xd3, 0x06, 0x85, 0x74, 0x17, 0x2e, 0x9c,
	0x5a, 0xdb, 0xdf, 0xf7, 0x3d, 0x7f, 0xef, 0x7b, 0x31, 0xdc, 0x3a, 0x17, 0x3c, 0x4d, 0x8a, 0x85,
	0x8c, 0xb3, 0x05, 0xbb, 0x0c, 0xe2, 0x28, 0x09, 0x2e, 0xc6, 0xea, 0xc7, 0xcf, 0x72, 0x2e, 0x39,
	0x19, 0x56, 0xce, 0x7d, 0x75, 0x70, 0x31, 0xf6, 0x6e, 0xce, 0x39, 0x9f, 0x2f, 0x58, 0x10, 0x65,
	0x71, 0x10, 0xa5, 0x29, 0x97, 0x91, 0x8c, 0x79, 0x2a, 0x34, 0xc7, 0xbb, 0x6d, 0x4e, 0x71, 0x75,
	0x56, 0xbc, 0x09, 0x64, 0x9c, 0x30, 0x21, 0xa3, 0x24, 0xd3, 0x00, 0xfa, 0xcb, 0x81, 0xee, 0xf1,
	0x74, 0xca, 0x8b, 0x54, 0x12, 0x02, 0xed, 0x34, 0x4a, 0x98, 0xeb, 0x8c, 0x9c, 0x83, 0x5e, 0x88,
	0xff, 0xc9, 0x03, 0xe8, 0x4f, 0x73, 0x16, 0x49, 0x76, 0xaa, 0x98, 0xee, 0xc6, 0xc8, 0x39, 0xe8,
	0x1f, 0x79, 0xbe, 0x96, 0xf5, 0xad, 0xac, 0xff, 0xca, 0xca, 0x86, 0xa0, 0xe1, 0x6a, 0x43, 0x91,
	0x8b, 0x6c, 0x56, 0x92, 0x5b, 0xeb, 0xc9, 0x1a, 0x6e, 0xc9, 0x33, 0xb6, 0x60, 0x96, 0xdc, 0x5e,
	0x4f, 0xd6, 0x70, 0x24, 0xdf, 0x81, 0xed, 0x9c, 0x73, 0x79, 0x9a, 0x45, 0x42, 0xbc, 0xe7, 0xf9,
	0xcc, 0xed, 0xa0, 0xa7, 0x2d, 0xb5, 0xf9, 0xc2, 0xec, 0xd1, 0x1f, 0x0e, 0xb4, 0x5f, 0x0b, 0x96,
	0xff, 0x4f, 0xc6, 0x3d, 0xb8, 0x56, 0x7a, 0xde, 0x44, 0x3b, 0xe5, 0x9a, 0xfe, 0x74, 0xa0, 0xfb,
	0x92, 0x09, 0x11, 0xf3, 0xf4, 0x9f, 0x58, 0x66, 0x97, 0x59, 0x9c, 0x5f, 0xdd, 0xb2, 0x86, 0x23,
	0xd9, 0x85, 0x6e, 0xa4, 0x87, 0x10, 0xed, 0xf6, 0x42, 0xbb, 0x54, 0xf7, 0x2c, 0x04, 0xcb, 0x4d,
	0x7e, 0xf8, 0xbf, 0xc9, 0x23, 0x19, 0x42, 0x47, 0xf2, 0xb7, 0x2c, 0x75, 0xbb, 0x78, 0xa0, 0x17,
	0xf4, 0x39, 0x0c, 0x4f, 0xf0, 0xaa, 0x66, 0xd4, 0x43, 0xf6, 0xae, 0x60, 0x42, 0x92, 0xfb, 0xcb,
	0xba, 0x0e, 0x5e, 0x78, 0xd7, 0xaf, 0xfb, 0xc8, 0x7c, 0x4b, 0xb3, 0x68, 0x4a, 0xe0, 0xc6, 0xd3,
	0x58, 0x48, 0x35, 0x3d, 0xc2, 0x88, 0xd1, 0x47, 0x30, 0x58, 0xd9, 0x13, 0x19, 0x4f, 0x05, 0x23,
	0x63, 0xe8, 0xa8, 0x3b, 0x0b, 0xd7, 0x19, 0xb5, 0xb0, 0x21, 0xb5, 0xfa, 0x8a, 0x13, 0x6a, 0x20,
	0xdd, 0x83, 0xeb, 0x8f, 0x19, 0xaa, 0xd8, 0x5b, 0xd6, 0x64, 0x45, 0x4f, 0x60, 0xa0, 0x1d, 0xad,
	0x02, 0x7d, 0xd3, 0x2c, 0xc7, 0x34, 0xff, 0xef, 0xb5, 0x10, 0xb7, 0x6c, 0x8b, 0x99, 0x8a, 0x95,
	0xb6, 0x08, 0xbd, 0xd3, 0xdc, 0x16, 0x4b, 0xb3, 0xe8, 0xa3, 0xaf, 0x6d, 0x68, 0x3d, 0x39, 0x7e,
	0x46, 0x3e, 0xc0, 0x76, 0xa5, 0xdf, 0xe4, 0xb0, 0x5e, 0xa0, 0x2e, 0x14, 0xaf, 0x39, 0x03, 0xba,
	0xfb, 0xf9, 0xdb, 0xf7, 0x2f, 0x1b, 0x3b, 0x74, 0x4b, 0xbd, 0x8e, 0x26, 0x0f, 0x31, 0x29, 0x07,
	0x86, 0x43, 0xaf, 0x4c, 0x81, 0xdc, 0xad, 0x97, 0xfa, 0x33, 0x3a, 0x6f, 0x7f, 0x2d, 0x4e, 0xc7,
	0x49, 0x07, 0x58, 0xbc, 0x4f, 0x7a, 0xaa, 0x38, 0xe6, 0x45, 0xe6, 0xd0, 0x35, 0x79, 0x91, 0xbd,
	0x7a, 0x99, 0x6a, 0x9c, 0x5e, 0x43, 0x2e, 0xd4, 0x43, 0xfd, 0x21, 0x21, 0x4a, 0xff, 0xa3, 0x4a,
	0xfa, 0x21, 0x56, 0x09, 0x0e, 0x3f, 0x91, 0x73, 0x80, 0x65, 0xe4, 0x64, 0xbf, 0xa9, 0xa3, 0x57,
	0x2d, 0xb7, 0x83, 0xe5, 0x06, 0x74, 0x69, 0x67, 0xa2, 0x3f, 0xb1, 0x32, 0x40, 0xfb, 0x5e, 0x34,
	0x06, 0x58, 0x1d, 0x1f, 0xaf, 0x79, 0x5a, 0xaa, 0x01, 0x9a, 0xc9, 0x11, 0x13, 0x3b, 0x43, 0x67,
	0x9b, 0xf8, 0x56, 0xdc, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x5b, 0xff, 0xa9, 0x0f, 0x07,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IAMClient is the client API for IAM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IAMClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*Account, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Session, error)
}

type iAMClient struct {
	cc *grpc.ClientConn
}

func NewIAMClient(cc *grpc.ClientConn) IAMClient {
	return &iAMClient{cc}
}

func (c *iAMClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/jsonmultiplex.iam.v0.IAM/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/jsonmultiplex.iam.v0.IAM/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/jsonmultiplex.iam.v0.IAM/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/jsonmultiplex.iam.v0.IAM/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/jsonmultiplex.iam.v0.IAM/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAMServer is the server API for IAM service.
type IAMServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*Account, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	CreateSession(context.Context, *CreateSessionRequest) (*Session, error)
}

// UnimplementedIAMServer can be embedded to have forward compatible implementations.
type UnimplementedIAMServer struct {
}

func (*UnimplementedIAMServer) CreateAccount(ctx context.Context, req *CreateAccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (*UnimplementedIAMServer) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (*UnimplementedIAMServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedIAMServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedIAMServer) CreateSession(ctx context.Context, req *CreateSessionRequest) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}

func RegisterIAMServer(s *grpc.Server, srv IAMServer) {
	s.RegisterService(&_IAM_serviceDesc, srv)
}

func _IAM_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jsonmultiplex.iam.v0.IAM/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAM_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jsonmultiplex.iam.v0.IAM/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAM_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jsonmultiplex.iam.v0.IAM/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAM_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jsonmultiplex.iam.v0.IAM/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAM_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jsonmultiplex.iam.v0.IAM/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IAM_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jsonmultiplex.iam.v0.IAM",
	HandlerType: (*IAMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _IAM_CreateAccount_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _IAM_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _IAM_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _IAM_CreateUser_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _IAM_CreateSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jsonmultiplex/iam/v0/iam.proto",
}
