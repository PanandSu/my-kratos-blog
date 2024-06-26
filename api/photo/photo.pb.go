// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v5.26.0
// source: api/photo/photo.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CommonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Result string `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CommonReply) Reset() {
	*x = CommonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_photo_photo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonReply) ProtoMessage() {}

func (x *CommonReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_photo_photo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonReply.ProtoReflect.Descriptor instead.
func (*CommonReply) Descriptor() ([]byte, []int) {
	return file_api_photo_photo_proto_rawDescGZIP(), []int{0}
}

func (x *CommonReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonReply) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type PhotoData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Photo    string `protobuf:"bytes,2,opt,name=photo,proto3" json:"photo,omitempty"`
	Date     string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Title    string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Position string `protobuf:"bytes,5,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *PhotoData) Reset() {
	*x = PhotoData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_photo_photo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoData) ProtoMessage() {}

func (x *PhotoData) ProtoReflect() protoreflect.Message {
	mi := &file_api_photo_photo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoData.ProtoReflect.Descriptor instead.
func (*PhotoData) Descriptor() ([]byte, []int) {
	return file_api_photo_photo_proto_rawDescGZIP(), []int{1}
}

func (x *PhotoData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PhotoData) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *PhotoData) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *PhotoData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PhotoData) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

type CreatePhotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *PhotoData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreatePhotoRequest) Reset() {
	*x = CreatePhotoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_photo_photo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePhotoRequest) ProtoMessage() {}

func (x *CreatePhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_photo_photo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePhotoRequest.ProtoReflect.Descriptor instead.
func (*CreatePhotoRequest) Descriptor() ([]byte, []int) {
	return file_api_photo_photo_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePhotoRequest) GetData() *PhotoData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreatePhotoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common *CommonReply `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
}

func (x *CreatePhotoReply) Reset() {
	*x = CreatePhotoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_photo_photo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePhotoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePhotoReply) ProtoMessage() {}

func (x *CreatePhotoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_photo_photo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePhotoReply.ProtoReflect.Descriptor instead.
func (*CreatePhotoReply) Descriptor() ([]byte, []int) {
	return file_api_photo_photo_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePhotoReply) GetCommon() *CommonReply {
	if x != nil {
		return x.Common
	}
	return nil
}

type DeletePhotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePhotoRequest) Reset() {
	*x = DeletePhotoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_photo_photo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePhotoRequest) ProtoMessage() {}

func (x *DeletePhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_photo_photo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePhotoRequest.ProtoReflect.Descriptor instead.
func (*DeletePhotoRequest) Descriptor() ([]byte, []int) {
	return file_api_photo_photo_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePhotoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeletePhotoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common *CommonReply `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
}

func (x *DeletePhotoReply) Reset() {
	*x = DeletePhotoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_photo_photo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePhotoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePhotoReply) ProtoMessage() {}

func (x *DeletePhotoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_photo_photo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePhotoReply.ProtoReflect.Descriptor instead.
func (*DeletePhotoReply) Descriptor() ([]byte, []int) {
	return file_api_photo_photo_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePhotoReply) GetCommon() *CommonReply {
	if x != nil {
		return x.Common
	}
	return nil
}

type ListPhotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPhotoRequest) Reset() {
	*x = ListPhotoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_photo_photo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPhotoRequest) ProtoMessage() {}

func (x *ListPhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_photo_photo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPhotoRequest.ProtoReflect.Descriptor instead.
func (*ListPhotoRequest) Descriptor() ([]byte, []int) {
	return file_api_photo_photo_proto_rawDescGZIP(), []int{6}
}

type ListPhotoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common *CommonReply `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	Data   []*PhotoData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListPhotoReply) Reset() {
	*x = ListPhotoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_photo_photo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPhotoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPhotoReply) ProtoMessage() {}

func (x *ListPhotoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_photo_photo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPhotoReply.ProtoReflect.Descriptor instead.
func (*ListPhotoReply) Descriptor() ([]byte, []int) {
	return file_api_photo_photo_proto_rawDescGZIP(), []int{7}
}

func (x *ListPhotoReply) GetCommon() *CommonReply {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *ListPhotoReply) GetData() []*PhotoData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_photo_photo_proto protoreflect.FileDescriptor

var file_api_photo_photo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x77, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x38,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x28, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5e,
	0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x28, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x91,
	0x02, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x57, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a,
	0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x12, 0x5c, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x51, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_photo_photo_proto_rawDescOnce sync.Once
	file_api_photo_photo_proto_rawDescData = file_api_photo_photo_proto_rawDesc
)

func file_api_photo_photo_proto_rawDescGZIP() []byte {
	file_api_photo_photo_proto_rawDescOnce.Do(func() {
		file_api_photo_photo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_photo_photo_proto_rawDescData)
	})
	return file_api_photo_photo_proto_rawDescData
}

var file_api_photo_photo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_photo_photo_proto_goTypes = []interface{}{
	(*CommonReply)(nil),        // 0: api.CommonReply
	(*PhotoData)(nil),          // 1: api.PhotoData
	(*CreatePhotoRequest)(nil), // 2: api.CreatePhotoRequest
	(*CreatePhotoReply)(nil),   // 3: api.CreatePhotoReply
	(*DeletePhotoRequest)(nil), // 4: api.DeletePhotoRequest
	(*DeletePhotoReply)(nil),   // 5: api.DeletePhotoReply
	(*ListPhotoRequest)(nil),   // 6: api.ListPhotoRequest
	(*ListPhotoReply)(nil),     // 7: api.ListPhotoReply
}
var file_api_photo_photo_proto_depIdxs = []int32{
	1, // 0: api.CreatePhotoRequest.data:type_name -> api.PhotoData
	0, // 1: api.CreatePhotoReply.common:type_name -> api.CommonReply
	0, // 2: api.DeletePhotoReply.common:type_name -> api.CommonReply
	0, // 3: api.ListPhotoReply.common:type_name -> api.CommonReply
	1, // 4: api.ListPhotoReply.data:type_name -> api.PhotoData
	2, // 5: api.Photo.CreatePhoto:input_type -> api.CreatePhotoRequest
	4, // 6: api.Photo.DeletePhoto:input_type -> api.DeletePhotoRequest
	6, // 7: api.Photo.ListPhoto:input_type -> api.ListPhotoRequest
	3, // 8: api.Photo.CreatePhoto:output_type -> api.CreatePhotoReply
	5, // 9: api.Photo.DeletePhoto:output_type -> api.DeletePhotoReply
	7, // 10: api.Photo.ListPhoto:output_type -> api.ListPhotoReply
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_photo_photo_proto_init() }
func file_api_photo_photo_proto_init() {
	if File_api_photo_photo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_photo_photo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonReply); i {
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
		file_api_photo_photo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoData); i {
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
		file_api_photo_photo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePhotoRequest); i {
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
		file_api_photo_photo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePhotoReply); i {
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
		file_api_photo_photo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePhotoRequest); i {
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
		file_api_photo_photo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePhotoReply); i {
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
		file_api_photo_photo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPhotoRequest); i {
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
		file_api_photo_photo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPhotoReply); i {
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
			RawDescriptor: file_api_photo_photo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_photo_photo_proto_goTypes,
		DependencyIndexes: file_api_photo_photo_proto_depIdxs,
		MessageInfos:      file_api_photo_photo_proto_msgTypes,
	}.Build()
	File_api_photo_photo_proto = out.File
	file_api_photo_photo_proto_rawDesc = nil
	file_api_photo_photo_proto_goTypes = nil
	file_api_photo_photo_proto_depIdxs = nil
}
