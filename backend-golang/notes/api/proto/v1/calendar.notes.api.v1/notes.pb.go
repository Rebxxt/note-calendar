// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.9
// source: notes.proto

package notes

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

type NoteGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoteGetRequest) Reset() {
	*x = NoteGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteGetRequest) ProtoMessage() {}

func (x *NoteGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteGetRequest.ProtoReflect.Descriptor instead.
func (*NoteGetRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{0}
}

type NoteGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoteGetResponse) Reset() {
	*x = NoteGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteGetResponse) ProtoMessage() {}

func (x *NoteGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteGetResponse.ProtoReflect.Descriptor instead.
func (*NoteGetResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{1}
}

type NoteListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoteListRequest) Reset() {
	*x = NoteListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteListRequest) ProtoMessage() {}

func (x *NoteListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteListRequest.ProtoReflect.Descriptor instead.
func (*NoteListRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{2}
}

type NoteListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoteListResponse) Reset() {
	*x = NoteListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteListResponse) ProtoMessage() {}

func (x *NoteListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteListResponse.ProtoReflect.Descriptor instead.
func (*NoteListResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{3}
}

type NoteCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoteCreateRequest) Reset() {
	*x = NoteCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteCreateRequest) ProtoMessage() {}

func (x *NoteCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteCreateRequest.ProtoReflect.Descriptor instead.
func (*NoteCreateRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{4}
}

type NoteCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoteCreateResponse) Reset() {
	*x = NoteCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteCreateResponse) ProtoMessage() {}

func (x *NoteCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteCreateResponse.ProtoReflect.Descriptor instead.
func (*NoteCreateResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{5}
}

type NoteDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoteDeleteRequest) Reset() {
	*x = NoteDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteDeleteRequest) ProtoMessage() {}

func (x *NoteDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteDeleteRequest.ProtoReflect.Descriptor instead.
func (*NoteDeleteRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{6}
}

type NoteDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoteDeleteResponse) Reset() {
	*x = NoteDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteDeleteResponse) ProtoMessage() {}

func (x *NoteDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteDeleteResponse.ProtoReflect.Descriptor instead.
func (*NoteDeleteResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{7}
}

type NoteUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoteUpdateRequest) Reset() {
	*x = NoteUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteUpdateRequest) ProtoMessage() {}

func (x *NoteUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteUpdateRequest.ProtoReflect.Descriptor instead.
func (*NoteUpdateRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{8}
}

type NoteUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoteUpdateResponse) Reset() {
	*x = NoteUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteUpdateResponse) ProtoMessage() {}

func (x *NoteUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteUpdateResponse.ProtoReflect.Descriptor instead.
func (*NoteUpdateResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{9}
}

var File_notes_proto protoreflect.FileDescriptor

var file_notes_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x22, 0x10, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x65, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x65, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x4e, 0x6f, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10,
	0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x4e,
	0x6f, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x14, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x4e,
	0x6f, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xdd, 0x03, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x63, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x3b, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notes_proto_rawDescOnce sync.Once
	file_notes_proto_rawDescData = file_notes_proto_rawDesc
)

func file_notes_proto_rawDescGZIP() []byte {
	file_notes_proto_rawDescOnce.Do(func() {
		file_notes_proto_rawDescData = protoimpl.X.CompressGZIP(file_notes_proto_rawDescData)
	})
	return file_notes_proto_rawDescData
}

var file_notes_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_notes_proto_goTypes = []interface{}{
	(*NoteGetRequest)(nil),     // 0: calendar.api.notes.v1.NoteGetRequest
	(*NoteGetResponse)(nil),    // 1: calendar.api.notes.v1.NoteGetResponse
	(*NoteListRequest)(nil),    // 2: calendar.api.notes.v1.NoteListRequest
	(*NoteListResponse)(nil),   // 3: calendar.api.notes.v1.NoteListResponse
	(*NoteCreateRequest)(nil),  // 4: calendar.api.notes.v1.NoteCreateRequest
	(*NoteCreateResponse)(nil), // 5: calendar.api.notes.v1.NoteCreateResponse
	(*NoteDeleteRequest)(nil),  // 6: calendar.api.notes.v1.NoteDeleteRequest
	(*NoteDeleteResponse)(nil), // 7: calendar.api.notes.v1.NoteDeleteResponse
	(*NoteUpdateRequest)(nil),  // 8: calendar.api.notes.v1.NoteUpdateRequest
	(*NoteUpdateResponse)(nil), // 9: calendar.api.notes.v1.NoteUpdateResponse
}
var file_notes_proto_depIdxs = []int32{
	0, // 0: calendar.api.notes.v1.Notes.Get:input_type -> calendar.api.notes.v1.NoteGetRequest
	2, // 1: calendar.api.notes.v1.Notes.List:input_type -> calendar.api.notes.v1.NoteListRequest
	4, // 2: calendar.api.notes.v1.Notes.Create:input_type -> calendar.api.notes.v1.NoteCreateRequest
	6, // 3: calendar.api.notes.v1.Notes.Delete:input_type -> calendar.api.notes.v1.NoteDeleteRequest
	8, // 4: calendar.api.notes.v1.Notes.Update:input_type -> calendar.api.notes.v1.NoteUpdateRequest
	1, // 5: calendar.api.notes.v1.Notes.Get:output_type -> calendar.api.notes.v1.NoteGetResponse
	3, // 6: calendar.api.notes.v1.Notes.List:output_type -> calendar.api.notes.v1.NoteListResponse
	5, // 7: calendar.api.notes.v1.Notes.Create:output_type -> calendar.api.notes.v1.NoteCreateResponse
	7, // 8: calendar.api.notes.v1.Notes.Delete:output_type -> calendar.api.notes.v1.NoteDeleteResponse
	9, // 9: calendar.api.notes.v1.Notes.Update:output_type -> calendar.api.notes.v1.NoteUpdateResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notes_proto_init() }
func file_notes_proto_init() {
	if File_notes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteGetRequest); i {
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
		file_notes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteGetResponse); i {
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
		file_notes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteListRequest); i {
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
		file_notes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteListResponse); i {
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
		file_notes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteCreateRequest); i {
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
		file_notes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteCreateResponse); i {
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
		file_notes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteDeleteRequest); i {
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
		file_notes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteDeleteResponse); i {
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
		file_notes_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteUpdateRequest); i {
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
		file_notes_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteUpdateResponse); i {
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
			RawDescriptor: file_notes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notes_proto_goTypes,
		DependencyIndexes: file_notes_proto_depIdxs,
		MessageInfos:      file_notes_proto_msgTypes,
	}.Build()
	File_notes_proto = out.File
	file_notes_proto_rawDesc = nil
	file_notes_proto_goTypes = nil
	file_notes_proto_depIdxs = nil
}
