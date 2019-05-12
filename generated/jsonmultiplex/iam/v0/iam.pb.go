// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jsonmultiplex/iam/v0/iam.proto

package jsonmultiplex_iam_v0

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type DeleteUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b7e58ef322f61c, []int{8}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
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
	return fileDescriptor_73b7e58ef322f61c, []int{9}
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
	proto.RegisterType((*DeleteUserRequest)(nil), "jsonmultiplex.iam.v0.DeleteUserRequest")
	proto.RegisterType((*CreateSessionRequest)(nil), "jsonmultiplex.iam.v0.CreateSessionRequest")
}

func init() { proto.RegisterFile("jsonmultiplex/iam/v0/iam.proto", fileDescriptor_73b7e58ef322f61c) }

var fileDescriptor_73b7e58ef322f61c = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0x95, 0xdb, 0xa6, 0x79, 0xb9, 0x69, 0x9f, 0xc8, 0x28, 0xa2, 0x96, 0xa1, 0x10, 0x0d, 0x15,
	0xa9, 0xb2, 0xb0, 0xa3, 0xb2, 0x40, 0x0a, 0x62, 0x51, 0x95, 0x0a, 0x21, 0x81, 0x40, 0x01, 0xd6,
	0x95, 0x9b, 0x0c, 0x91, 0x4b, 0xec, 0x31, 0x9e, 0x71, 0x69, 0x85, 0xd8, 0xf0, 0x0b, 0xfc, 0x06,
	0x7f, 0xc3, 0x2f, 0xb0, 0x65, 0x8d, 0xd8, 0xa1, 0xb9, 0x33, 0xe3, 0x24, 0xad, 0xeb, 0x74, 0xc3,
	0x86, 0x55, 0x3c, 0x33, 0xe7, 0xdc, 0x3b, 0xe7, 0x9e, 0x63, 0x07, 0xee, 0x9c, 0x08, 0x9e, 0xc4,
	0xf9, 0x54, 0x46, 0xe9, 0x94, 0x9d, 0x05, 0x51, 0x18, 0x07, 0xa7, 0x7d, 0xf5, 0xe3, 0xa7, 0x19,
	0x97, 0x9c, 0xb4, 0x17, 0xce, 0x7d, 0x75, 0x70, 0xda, 0xf7, 0x6e, 0x4f, 0x38, 0x9f, 0x4c, 0x59,
	0x10, 0xa6, 0x51, 0x10, 0x26, 0x09, 0x97, 0xa1, 0x8c, 0x78, 0x22, 0x34, 0xc7, 0xbb, 0x6b, 0x4e,
	0x71, 0x75, 0x9c, 0xbf, 0x0b, 0x64, 0x14, 0x33, 0x21, 0xc3, 0x38, 0x35, 0x80, 0x5b, 0x17, 0x01,
	0x2c, 0x4e, 0xe5, 0xb9, 0x3e, 0xa4, 0xbf, 0x1d, 0xa8, 0xef, 0x8f, 0x46, 0x3c, 0x4f, 0x24, 0x21,
	0xb0, 0x96, 0x84, 0x31, 0x73, 0x9d, 0x8e, 0xb3, 0xdb, 0x18, 0xe2, 0x33, 0x79, 0x04, 0xcd, 0x51,
	0xc6, 0x42, 0xc9, 0x8e, 0x54, 0x59, 0x77, 0xa5, 0xe3, 0xec, 0x36, 0xf7, 0x3c, 0x5f, 0x97, 0xf4,
	0x6d, 0x49, 0xff, 0x8d, 0xed, 0x39, 0x04, 0x0d, 0x57, 0x1b, 0x8a, 0x9c, 0xa7, 0xe3, 0x82, 0xbc,
	0xba, 0x9c, 0xac, 0xe1, 0x96, 0x3c, 0x66, 0x53, 0x66, 0xc9, 0x6b, 0xcb, 0xc9, 0x1a, 0x8e, 0xe4,
	0x7b, 0xb0, 0x99, 0x71, 0x2e, 0x8f, 0xd2, 0x50, 0x88, 0x8f, 0x3c, 0x1b, 0xbb, 0x35, 0xd4, 0xb4,
	0xa1, 0x36, 0x5f, 0x99, 0x3d, 0xfa, 0xd3, 0x81, 0xb5, 0xb7, 0x82, 0x65, 0xff, 0x92, 0x70, 0x0f,
	0xfe, 0x2b, 0x34, 0xaf, 0xa3, 0x9c, 0x62, 0x4d, 0x7f, 0x39, 0x50, 0x7f, 0xcd, 0x84, 0x88, 0x78,
	0xf2, 0x57, 0x24, 0xb3, 0xb3, 0x34, 0xca, 0xae, 0x2f, 0x59, 0xc3, 0x91, 0xec, 0x42, 0x3d, 0xd4,
	0x21, 0x44, 0xb9, 0x8d, 0xa1, 0x5d, 0xaa, 0x7b, 0xe6, 0x82, 0x65, 0xc6, 0x3f, 0x7c, 0xae, 0xd2,
	0x48, 0xda, 0x50, 0x93, 0xfc, 0x3d, 0x4b, 0xdc, 0x3a, 0x1e, 0xe8, 0x05, 0x7d, 0x09, 0xed, 0x03,
	0xbc, 0xaa, 0x89, 0xfa, 0x90, 0x7d, 0xc8, 0x99, 0x90, 0xe4, 0xe1, 0xac, 0xaf, 0x83, 0x17, 0xde,
	0xf6, 0xcb, 0xde, 0x40, 0xdf, 0xd2, 0x2c, 0x9a, 0x12, 0xb8, 0xf1, 0x3c, 0x12, 0x52, 0xa5, 0x47,
	0x98, 0x62, 0xf4, 0x10, 0x5a, 0x73, 0x7b, 0x22, 0xe5, 0x89, 0x60, 0xa4, 0x0f, 0x35, 0x75, 0x67,
	0xe1, 0x3a, 0x9d, 0x55, 0x1c, 0x48, 0x69, 0x7d, 0xc5, 0x19, 0x6a, 0x20, 0xdd, 0x81, 0xff, 0x9f,
	0x32, 0xac, 0x62, 0x6f, 0x59, 0xe2, 0x15, 0x3d, 0x80, 0x96, 0x56, 0x34, 0x0f, 0xf4, 0xcd, 0xb0,
	0x1c, 0x33, 0xfc, 0xab, 0x7b, 0x21, 0x8e, 0x76, 0xa1, 0xf5, 0x04, 0xa3, 0xb3, 0xac, 0x5b, 0x31,
	0x3f, 0x13, 0x9f, 0xb9, 0xf9, 0x09, 0xbd, 0x53, 0x3d, 0x3f, 0x4b, 0xb3, 0xe8, 0xbd, 0x6f, 0x35,
	0x58, 0x7d, 0xb6, 0xff, 0x82, 0x9c, 0xc3, 0xe6, 0x82, 0x31, 0xa4, 0x57, 0x5e, 0xa0, 0xcc, 0x3d,
	0xaf, 0xda, 0x2c, 0xba, 0xfd, 0xe5, 0xfb, 0x8f, 0xaf, 0x2b, 0x5b, 0x74, 0x43, 0x7d, 0x63, 0x8d,
	0x71, 0x62, 0x50, 0x24, 0x8b, 0x43, 0xa3, 0xb0, 0x8b, 0xdc, 0x2f, 0x2f, 0x75, 0xd1, 0x63, 0xaf,
	0xbb, 0x14, 0xa7, 0x7d, 0xa7, 0x2d, 0x6c, 0xde, 0x24, 0x0d, 0xd5, 0x1c, 0x8d, 0x25, 0x13, 0xa8,
	0x1b, 0x63, 0xc9, 0x4e, 0x79, 0x99, 0x45, 0xdf, 0xbd, 0x0a, 0x03, 0xa9, 0x87, 0xf5, 0xdb, 0x84,
	0xa8, 0xfa, 0x9f, 0x94, 0x49, 0x8f, 0xb1, 0x4b, 0xd0, 0xfb, 0x4c, 0x4e, 0x00, 0x66, 0xd9, 0x20,
	0xdd, 0xaa, 0x89, 0x5e, 0xb7, 0xdd, 0x16, 0xb6, 0x6b, 0xd1, 0x99, 0x9c, 0x81, 0x7e, 0x17, 0x23,
	0x80, 0x59, 0x84, 0xae, 0xea, 0x75, 0x29, 0x64, 0xde, 0xcd, 0x4b, 0x1f, 0x86, 0x43, 0xf5, 0xa7,
	0x64, 0x65, 0xf5, 0xca, 0x64, 0x15, 0x59, 0xb1, 0xdf, 0xb0, 0xca, 0xac, 0x2c, 0x26, 0xd5, 0xab,
	0x0e, 0xe6, 0x62, 0x56, 0x4c, 0x48, 0xc5, 0xc0, 0xc6, 0xf5, 0x78, 0x1d, 0xaf, 0xf9, 0xe0, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x02, 0x18, 0x90, 0xc0, 0x07, 0x00, 0x00,
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
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
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

func (c *iAMClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/jsonmultiplex.iam.v0.IAM/DeleteUser", in, out, opts...)
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
	DeleteUser(context.Context, *DeleteUserRequest) (*empty.Empty, error)
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
func (*UnimplementedIAMServer) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
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

func _IAM_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jsonmultiplex.iam.v0.IAM/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).DeleteUser(ctx, req.(*DeleteUserRequest))
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
			MethodName: "DeleteUser",
			Handler:    _IAM_DeleteUser_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _IAM_CreateSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jsonmultiplex/iam/v0/iam.proto",
}
