// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: upload.proto

package upload

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

	Code int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *Base) Reset() {
	*x = Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base) ProtoMessage() {}

func (x *Base) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[0]
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
	return file_upload_proto_rawDescGZIP(), []int{0}
}

func (x *Base) GetCode() int32 {
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

type FileUploadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File       []byte `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
	FileName   string `protobuf:"bytes,2,opt,name=FileName,proto3" json:"FileName,omitempty"`
	Path       string `protobuf:"bytes,3,opt,name=Path,proto3" json:"Path,omitempty"`
	BucketName string `protobuf:"bytes,4,opt,name=BucketName,proto3" json:"BucketName,omitempty"`
}

func (x *FileUploadReq) Reset() {
	*x = FileUploadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadReq) ProtoMessage() {}

func (x *FileUploadReq) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadReq.ProtoReflect.Descriptor instead.
func (*FileUploadReq) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{1}
}

func (x *FileUploadReq) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *FileUploadReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileUploadReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileUploadReq) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

type FileUploadRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *Base              `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *FileUploadResData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FileUploadRes) Reset() {
	*x = FileUploadRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadRes) ProtoMessage() {}

func (x *FileUploadRes) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadRes.ProtoReflect.Descriptor instead.
func (*FileUploadRes) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{2}
}

func (x *FileUploadRes) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *FileUploadRes) GetData() *FileUploadResData {
	if x != nil {
		return x.Data
	}
	return nil
}

type FileUploadResData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=FileName,proto3" json:"FileName,omitempty"`
	Path     string `protobuf:"bytes,2,opt,name=Path,proto3" json:"Path,omitempty"`
	ETag     string `protobuf:"bytes,3,opt,name=ETag,proto3" json:"ETag,omitempty"`
	Size     uint64 `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (x *FileUploadResData) Reset() {
	*x = FileUploadResData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadResData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadResData) ProtoMessage() {}

func (x *FileUploadResData) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadResData.ProtoReflect.Descriptor instead.
func (*FileUploadResData) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{3}
}

func (x *FileUploadResData) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileUploadResData) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileUploadResData) GetETag() string {
	if x != nil {
		return x.ETag
	}
	return ""
}

func (x *FileUploadResData) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ImageUploadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File       []byte `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
	FileName   string `protobuf:"bytes,2,opt,name=FileName,proto3" json:"FileName,omitempty"`
	Path       string `protobuf:"bytes,3,opt,name=Path,proto3" json:"Path,omitempty"`
	OriPath    string `protobuf:"bytes,4,opt,name=OriPath,proto3" json:"OriPath,omitempty"`
	Quality    uint32 `protobuf:"varint,5,opt,name=Quality,proto3" json:"Quality,omitempty"`
	BucketName string `protobuf:"bytes,6,opt,name=BucketName,proto3" json:"BucketName,omitempty"`
}

func (x *ImageUploadReq) Reset() {
	*x = ImageUploadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageUploadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUploadReq) ProtoMessage() {}

func (x *ImageUploadReq) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUploadReq.ProtoReflect.Descriptor instead.
func (*ImageUploadReq) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{4}
}

func (x *ImageUploadReq) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *ImageUploadReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ImageUploadReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ImageUploadReq) GetOriPath() string {
	if x != nil {
		return x.OriPath
	}
	return ""
}

func (x *ImageUploadReq) GetQuality() uint32 {
	if x != nil {
		return x.Quality
	}
	return 0
}

func (x *ImageUploadReq) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

type ImageUploadRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *Base               `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *ImageUploadResData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ImageUploadRes) Reset() {
	*x = ImageUploadRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageUploadRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUploadRes) ProtoMessage() {}

func (x *ImageUploadRes) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUploadRes.ProtoReflect.Descriptor instead.
func (*ImageUploadRes) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{5}
}

func (x *ImageUploadRes) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *ImageUploadRes) GetData() *ImageUploadResData {
	if x != nil {
		return x.Data
	}
	return nil
}

type ImageUploadResData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	OriPath string `protobuf:"bytes,2,opt,name=OriPath,proto3" json:"OriPath,omitempty"`
	ETag    string `protobuf:"bytes,3,opt,name=ETag,proto3" json:"ETag,omitempty"`
	OriETag string `protobuf:"bytes,4,opt,name=OriETag,proto3" json:"OriETag,omitempty"`
	Size    uint64 `protobuf:"varint,5,opt,name=Size,proto3" json:"Size,omitempty"`
	OriSize uint64 `protobuf:"varint,6,opt,name=OriSize,proto3" json:"OriSize,omitempty"`
}

func (x *ImageUploadResData) Reset() {
	*x = ImageUploadResData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageUploadResData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUploadResData) ProtoMessage() {}

func (x *ImageUploadResData) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUploadResData.ProtoReflect.Descriptor instead.
func (*ImageUploadResData) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{6}
}

func (x *ImageUploadResData) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ImageUploadResData) GetOriPath() string {
	if x != nil {
		return x.OriPath
	}
	return ""
}

func (x *ImageUploadResData) GetETag() string {
	if x != nil {
		return x.ETag
	}
	return ""
}

func (x *ImageUploadResData) GetOriETag() string {
	if x != nil {
		return x.OriETag
	}
	return ""
}

func (x *ImageUploadResData) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ImageUploadResData) GetOriSize() uint64 {
	if x != nil {
		return x.OriSize
	}
	return 0
}

var File_upload_proto protoreflect.FileDescriptor

var file_upload_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x2c, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4d, 0x73, 0x67, 0x22, 0x73, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x60, 0x0a, 0x0d, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6b, 0x0a, 0x11, 0x46,
	0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x45, 0x54, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x45, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x0e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x46,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x18, 0x0a, 0x07, 0x4f, 0x72, 0x69, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4f, 0x72, 0x69, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x51, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x51, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x62, 0x0a, 0x0e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9e, 0x01, 0x0a, 0x12, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x69, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x69, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x45, 0x54, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x45, 0x54, 0x61, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x69, 0x45, 0x54, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4f, 0x72, 0x69, 0x45, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x4f, 0x72, 0x69, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x4f, 0x72, 0x69, 0x53, 0x69, 0x7a, 0x65, 0x32, 0x83, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x15, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x12,
	0x3d, 0x0a, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16,
	0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_upload_proto_rawDescOnce sync.Once
	file_upload_proto_rawDescData = file_upload_proto_rawDesc
)

func file_upload_proto_rawDescGZIP() []byte {
	file_upload_proto_rawDescOnce.Do(func() {
		file_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_upload_proto_rawDescData)
	})
	return file_upload_proto_rawDescData
}

var file_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_upload_proto_goTypes = []interface{}{
	(*Base)(nil),               // 0: upload.Base
	(*FileUploadReq)(nil),      // 1: upload.FileUploadReq
	(*FileUploadRes)(nil),      // 2: upload.FileUploadRes
	(*FileUploadResData)(nil),  // 3: upload.FileUploadResData
	(*ImageUploadReq)(nil),     // 4: upload.ImageUploadReq
	(*ImageUploadRes)(nil),     // 5: upload.ImageUploadRes
	(*ImageUploadResData)(nil), // 6: upload.ImageUploadResData
}
var file_upload_proto_depIdxs = []int32{
	0, // 0: upload.FileUploadRes.base:type_name -> upload.Base
	3, // 1: upload.FileUploadRes.data:type_name -> upload.FileUploadResData
	0, // 2: upload.ImageUploadRes.base:type_name -> upload.Base
	6, // 3: upload.ImageUploadRes.data:type_name -> upload.ImageUploadResData
	1, // 4: upload.Upload.FileUpload:input_type -> upload.FileUploadReq
	4, // 5: upload.Upload.ImageUpload:input_type -> upload.ImageUploadReq
	2, // 6: upload.Upload.FileUpload:output_type -> upload.FileUploadRes
	5, // 7: upload.Upload.ImageUpload:output_type -> upload.ImageUploadRes
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_upload_proto_init() }
func file_upload_proto_init() {
	if File_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadReq); i {
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
		file_upload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadRes); i {
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
		file_upload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadResData); i {
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
		file_upload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageUploadReq); i {
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
		file_upload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageUploadRes); i {
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
		file_upload_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageUploadResData); i {
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
			RawDescriptor: file_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_upload_proto_goTypes,
		DependencyIndexes: file_upload_proto_depIdxs,
		MessageInfos:      file_upload_proto_msgTypes,
	}.Build()
	File_upload_proto = out.File
	file_upload_proto_rawDesc = nil
	file_upload_proto_goTypes = nil
	file_upload_proto_depIdxs = nil
}
