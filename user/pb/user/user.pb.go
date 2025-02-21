// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Base struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *Base) Reset() {
	*x = Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base) ProtoMessage() {}

func (x *Base) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base.ProtoReflect.Descriptor instead.
func (*Base) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *Base) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Base) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type AddUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Account  string `protobuf:"bytes,2,opt,name=Account,proto3" json:"Account,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Role     string `protobuf:"bytes,4,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *AddUserReq) Reset() {
	*x = AddUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserReq) ProtoMessage() {}

func (x *AddUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserReq.ProtoReflect.Descriptor instead.
func (*AddUserReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *AddUserReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddUserReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *AddUserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddUserReq) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type AddUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *Base           `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *AddUserResData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddUserRes) Reset() {
	*x = AddUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRes) ProtoMessage() {}

func (x *AddUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRes.ProtoReflect.Descriptor instead.
func (*AddUserRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *AddUserRes) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *AddUserRes) GetData() *AddUserResData {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddUserResData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Account  string `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"`
}

func (x *AddUserResData) Reset() {
	*x = AddUserResData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserResData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserResData) ProtoMessage() {}

func (x *AddUserResData) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserResData.ProtoReflect.Descriptor instead.
func (*AddUserResData) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *AddUserResData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddUserResData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddUserResData) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type UpdateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Motto   string `protobuf:"bytes,2,opt,name=Motto,proto3" json:"Motto,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	Tel     string `protobuf:"bytes,4,opt,name=Tel,proto3" json:"Tel,omitempty"`
	Email   string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	QQ      string `protobuf:"bytes,6,opt,name=QQ,proto3" json:"QQ,omitempty"`
	Wechat  string `protobuf:"bytes,7,opt,name=Wechat,proto3" json:"Wechat,omitempty"`
	GitHub  string `protobuf:"bytes,8,opt,name=GitHub,proto3" json:"GitHub,omitempty"`
	Role    string `protobuf:"bytes,9,opt,name=Role,proto3" json:"Role,omitempty"`
	Avatar  string `protobuf:"bytes,10,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Cover   string `protobuf:"bytes,11,opt,name=Cover,proto3" json:"Cover,omitempty"`
}

func (x *UpdateUserReq) Reset() {
	*x = UpdateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserReq) ProtoMessage() {}

func (x *UpdateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserReq.ProtoReflect.Descriptor instead.
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUserReq) GetMotto() string {
	if x != nil {
		return x.Motto
	}
	return ""
}

func (x *UpdateUserReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateUserReq) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *UpdateUserReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserReq) GetQQ() string {
	if x != nil {
		return x.QQ
	}
	return ""
}

func (x *UpdateUserReq) GetWechat() string {
	if x != nil {
		return x.Wechat
	}
	return ""
}

func (x *UpdateUserReq) GetGitHub() string {
	if x != nil {
		return x.GitHub
	}
	return ""
}

func (x *UpdateUserReq) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UpdateUserReq) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UpdateUserReq) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

type UpdateUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *Base `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *UpdateUserRes) Reset() {
	*x = UpdateUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRes) ProtoMessage() {}

func (x *UpdateUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRes.ProtoReflect.Descriptor instead.
func (*UpdateUserRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserRes) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

type UserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *UserInfoReq) Reset() {
	*x = UserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReq) ProtoMessage() {}

func (x *UserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReq.ProtoReflect.Descriptor instead.
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserInfoReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UserInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *Base            `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *UserInfoResData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UserInfoRes) Reset() {
	*x = UserInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRes) ProtoMessage() {}

func (x *UserInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRes.ProtoReflect.Descriptor instead.
func (*UserInfoRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserInfoRes) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UserInfoRes) GetData() *UserInfoResData {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserInfoResData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Account  string `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"`
	Motto    string `protobuf:"bytes,4,opt,name=Motto,proto3" json:"Motto,omitempty"`
	Address  string `protobuf:"bytes,5,opt,name=Address,proto3" json:"Address,omitempty"`
	Tel      string `protobuf:"bytes,6,opt,name=Tel,proto3" json:"Tel,omitempty"`
	Email    string `protobuf:"bytes,7,opt,name=Email,proto3" json:"Email,omitempty"`
	QQ       string `protobuf:"bytes,8,opt,name=QQ,proto3" json:"QQ,omitempty"`
	Wechat   string `protobuf:"bytes,9,opt,name=Wechat,proto3" json:"Wechat,omitempty"`
	GitHub   string `protobuf:"bytes,10,opt,name=GitHub,proto3" json:"GitHub,omitempty"`
	Role     string `protobuf:"bytes,11,opt,name=Role,proto3" json:"Role,omitempty"`
	Avatar   string `protobuf:"bytes,12,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Cover    string `protobuf:"bytes,13,opt,name=Cover,proto3" json:"Cover,omitempty"`
}

func (x *UserInfoResData) Reset() {
	*x = UserInfoResData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResData) ProtoMessage() {}

func (x *UserInfoResData) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResData.ProtoReflect.Descriptor instead.
func (*UserInfoResData) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *UserInfoResData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfoResData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfoResData) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserInfoResData) GetMotto() string {
	if x != nil {
		return x.Motto
	}
	return ""
}

func (x *UserInfoResData) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserInfoResData) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *UserInfoResData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfoResData) GetQQ() string {
	if x != nil {
		return x.QQ
	}
	return ""
}

func (x *UserInfoResData) GetWechat() string {
	if x != nil {
		return x.Wechat
	}
	return ""
}

func (x *UserInfoResData) GetGitHub() string {
	if x != nil {
		return x.GitHub
	}
	return ""
}

func (x *UserInfoResData) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserInfoResData) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserInfoResData) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *Base         `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *LoginResData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *LoginRes) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *LoginRes) GetData() *LoginResData {
	if x != nil {
		return x.Data
	}
	return nil
}

type LoginResData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string           `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Info  *UserInfoResData `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *LoginResData) Reset() {
	*x = LoginResData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResData) ProtoMessage() {}

func (x *LoginResData) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResData.ProtoReflect.Descriptor instead.
func (*LoginResData) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{11}
}

func (x *LoginResData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResData) GetInfo() *UserInfoResData {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x2c, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67,
	0x22, 0x72, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x22, 0x56, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a, 0x0e,
	0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xf9, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x74, 0x74, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x6f, 0x74, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x54, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x51, 0x51, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x51, 0x51, 0x12, 0x16,
	0x0a, 0x06, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x12, 0x12,
	0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72,
	0x22, 0x2f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64,
	0x22, 0x58, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12,
	0x1e, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb1, 0x02, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x74, 0x74, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x6f, 0x74, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x54, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02,
	0x51, 0x51, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x51, 0x51, 0x12, 0x16, 0x0a, 0x06,
	0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x65,
	0x63, 0x68, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x12, 0x12, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x22, 0x42,
	0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x52, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x1e,
	0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4f, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x29, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xc8, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x2d, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12,
	0x36, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_user_proto_goTypes = []interface{}{
	(*Base)(nil),            // 0: user.Base
	(*AddUserReq)(nil),      // 1: user.AddUserReq
	(*AddUserRes)(nil),      // 2: user.AddUserRes
	(*AddUserResData)(nil),  // 3: user.AddUserResData
	(*UpdateUserReq)(nil),   // 4: user.UpdateUserReq
	(*UpdateUserRes)(nil),   // 5: user.UpdateUserRes
	(*UserInfoReq)(nil),     // 6: user.UserInfoReq
	(*UserInfoRes)(nil),     // 7: user.UserInfoRes
	(*UserInfoResData)(nil), // 8: user.UserInfoResData
	(*LoginReq)(nil),        // 9: user.LoginReq
	(*LoginRes)(nil),        // 10: user.LoginRes
	(*LoginResData)(nil),    // 11: user.LoginResData
}
var file_user_proto_depIdxs = []int32{
	0,  // 0: user.AddUserRes.base:type_name -> user.Base
	3,  // 1: user.AddUserRes.data:type_name -> user.AddUserResData
	0,  // 2: user.UpdateUserRes.base:type_name -> user.Base
	0,  // 3: user.UserInfoRes.base:type_name -> user.Base
	8,  // 4: user.UserInfoRes.data:type_name -> user.UserInfoResData
	0,  // 5: user.LoginRes.base:type_name -> user.Base
	11, // 6: user.LoginRes.data:type_name -> user.LoginResData
	8,  // 7: user.LoginResData.info:type_name -> user.UserInfoResData
	1,  // 8: user.User.AddUser:input_type -> user.AddUserReq
	4,  // 9: user.User.UpdateUser:input_type -> user.UpdateUserReq
	6,  // 10: user.User.UserInfo:input_type -> user.UserInfoReq
	9,  // 11: user.User.Login:input_type -> user.LoginReq
	2,  // 12: user.User.AddUser:output_type -> user.AddUserRes
	5,  // 13: user.User.UpdateUser:output_type -> user.UpdateUserRes
	7,  // 14: user.User.UserInfo:output_type -> user.UserInfoRes
	10, // 15: user.User.Login:output_type -> user.LoginRes
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserResData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
