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

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
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

	File       []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	FileName   string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FileType   string `protobuf:"bytes,3,opt,name=fileType,proto3" json:"fileType,omitempty"`
	OriFolder  string `protobuf:"bytes,4,opt,name=oriFolder,proto3" json:"oriFolder,omitempty"`
	Folder     string `protobuf:"bytes,5,opt,name=Folder,proto3" json:"Folder,omitempty"`
	BucketName string `protobuf:"bytes,6,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	Quality    uint32 `protobuf:"varint,7,opt,name=quality,proto3" json:"quality,omitempty"`
	W          uint32 `protobuf:"varint,8,opt,name=w,proto3" json:"w,omitempty"`
	H          uint32 `protobuf:"varint,9,opt,name=h,proto3" json:"h,omitempty"`
	Type       string `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	UserId     string `protobuf:"bytes,11,opt,name=userId,proto3" json:"userId,omitempty"`
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

func (x *FileUploadReq) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *FileUploadReq) GetOriFolder() string {
	if x != nil {
		return x.OriFolder
	}
	return ""
}

func (x *FileUploadReq) GetFolder() string {
	if x != nil {
		return x.Folder
	}
	return ""
}

func (x *FileUploadReq) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *FileUploadReq) GetQuality() uint32 {
	if x != nil {
		return x.Quality
	}
	return 0
}

func (x *FileUploadReq) GetW() uint32 {
	if x != nil {
		return x.W
	}
	return 0
}

func (x *FileUploadReq) GetH() uint32 {
	if x != nil {
		return x.H
	}
	return 0
}

func (x *FileUploadReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FileUploadReq) GetUserId() string {
	if x != nil {
		return x.UserId
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

	FileName    string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	OriFilePath string `protobuf:"bytes,2,opt,name=oriFilePath,proto3" json:"oriFilePath,omitempty"`
	FilePath    string `protobuf:"bytes,3,opt,name=filePath,proto3" json:"filePath,omitempty"`
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

func (x *FileUploadResData) GetOriFilePath() string {
	if x != nil {
		return x.OriFilePath
	}
	return ""
}

func (x *FileUploadResData) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

var File_upload_proto protoreflect.FileDescriptor

var file_upload_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x2c, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x93, 0x02, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x69, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x69, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x77, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x77,
	0x12, 0x0c, 0x0a, 0x01, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x0d, 0x46, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6d, 0x0a, 0x11,
	0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x6f, 0x72, 0x69, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x32, 0x44, 0x0a, 0x06, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x15, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_upload_proto_goTypes = []interface{}{
	(*Base)(nil),              // 0: upload.Base
	(*FileUploadReq)(nil),     // 1: upload.FileUploadReq
	(*FileUploadRes)(nil),     // 2: upload.FileUploadRes
	(*FileUploadResData)(nil), // 3: upload.FileUploadResData
}
var file_upload_proto_depIdxs = []int32{
	0, // 0: upload.FileUploadRes.base:type_name -> upload.Base
	3, // 1: upload.FileUploadRes.data:type_name -> upload.FileUploadResData
	1, // 2: upload.Upload.FileUpload:input_type -> upload.FileUploadReq
	2, // 3: upload.Upload.FileUpload:output_type -> upload.FileUploadRes
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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
