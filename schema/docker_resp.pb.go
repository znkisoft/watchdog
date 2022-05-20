// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: docker_resp.proto

package schema

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

type ListContainersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	All     bool   `protobuf:"varint,1,opt,name=all,proto3" json:"all,omitempty"`
	Limit   int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Size    bool   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Filters string `protobuf:"bytes,4,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListContainersRequest) Reset() {
	*x = ListContainersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListContainersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContainersRequest) ProtoMessage() {}

func (x *ListContainersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContainersRequest.ProtoReflect.Descriptor instead.
func (*ListContainersRequest) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{0}
}

func (x *ListContainersRequest) GetAll() bool {
	if x != nil {
		return x.All
	}
	return false
}

func (x *ListContainersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListContainersRequest) GetSize() bool {
	if x != nil {
		return x.Size
	}
	return false
}

func (x *ListContainersRequest) GetFilters() string {
	if x != nil {
		return x.Filters
	}
	return ""
}

type ListContainersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListContainersResponse) Reset() {
	*x = ListContainersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListContainersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContainersResponse) ProtoMessage() {}

func (x *ListContainersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContainersResponse.ProtoReflect.Descriptor instead.
func (*ListContainersResponse) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{1}
}

type CreateContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateContainerRequest) Reset() {
	*x = CreateContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContainerRequest) ProtoMessage() {}

func (x *CreateContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContainerRequest.ProtoReflect.Descriptor instead.
func (*CreateContainerRequest) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{2}
}

type CreateContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateContainerResponse) Reset() {
	*x = CreateContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContainerResponse) ProtoMessage() {}

func (x *CreateContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContainerResponse.ProtoReflect.Descriptor instead.
func (*CreateContainerResponse) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{3}
}

type StartContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartContainerRequest) Reset() {
	*x = StartContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartContainerRequest) ProtoMessage() {}

func (x *StartContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartContainerRequest.ProtoReflect.Descriptor instead.
func (*StartContainerRequest) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{4}
}

type StartContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartContainerResponse) Reset() {
	*x = StartContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartContainerResponse) ProtoMessage() {}

func (x *StartContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartContainerResponse.ProtoReflect.Descriptor instead.
func (*StartContainerResponse) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{5}
}

type StopContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopContainerRequest) Reset() {
	*x = StopContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopContainerRequest) ProtoMessage() {}

func (x *StopContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopContainerRequest.ProtoReflect.Descriptor instead.
func (*StopContainerRequest) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{6}
}

type StopContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopContainerResponse) Reset() {
	*x = StopContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopContainerResponse) ProtoMessage() {}

func (x *StopContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopContainerResponse.ProtoReflect.Descriptor instead.
func (*StopContainerResponse) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{7}
}

type RestartContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RestartContainerRequest) Reset() {
	*x = RestartContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartContainerRequest) ProtoMessage() {}

func (x *RestartContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestartContainerRequest.ProtoReflect.Descriptor instead.
func (*RestartContainerRequest) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{8}
}

type RestartContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RestartContainerResponse) Reset() {
	*x = RestartContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartContainerResponse) ProtoMessage() {}

func (x *RestartContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestartContainerResponse.ProtoReflect.Descriptor instead.
func (*RestartContainerResponse) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{9}
}

type RemoveContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveContainerRequest) Reset() {
	*x = RemoveContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveContainerRequest) ProtoMessage() {}

func (x *RemoveContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveContainerRequest.ProtoReflect.Descriptor instead.
func (*RemoveContainerRequest) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{10}
}

type RemoveContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveContainerResponse) Reset() {
	*x = RemoveContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveContainerResponse) ProtoMessage() {}

func (x *RemoveContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveContainerResponse.ProtoReflect.Descriptor instead.
func (*RemoveContainerResponse) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{11}
}

type KillContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KillContainerRequest) Reset() {
	*x = KillContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KillContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillContainerRequest) ProtoMessage() {}

func (x *KillContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillContainerRequest.ProtoReflect.Descriptor instead.
func (*KillContainerRequest) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{12}
}

type KillContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KillContainerResponse) Reset() {
	*x = KillContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KillContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillContainerResponse) ProtoMessage() {}

func (x *KillContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillContainerResponse.ProtoReflect.Descriptor instead.
func (*KillContainerResponse) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{13}
}

// PruneContainerRequest is the request to delete a container which is not running.
type PruneContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PruneContainerRequest) Reset() {
	*x = PruneContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PruneContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PruneContainerRequest) ProtoMessage() {}

func (x *PruneContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PruneContainerRequest.ProtoReflect.Descriptor instead.
func (*PruneContainerRequest) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{14}
}

// PruneContainerResponse is the response to delete a container which is not running.
type PruneContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PruneContainerResponse) Reset() {
	*x = PruneContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_resp_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PruneContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PruneContainerResponse) ProtoMessage() {}

func (x *PruneContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docker_resp_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PruneContainerResponse.ProtoReflect.Descriptor instead.
func (*PruneContainerResponse) Descriptor() ([]byte, []int) {
	return file_docker_resp_proto_rawDescGZIP(), []int{15}
}

var File_docker_resp_proto protoreflect.FileDescriptor

var file_docker_resp_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x74, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15,
	0x53, 0x74, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x0a, 0x16,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x16, 0x0a, 0x14, 0x4b, 0x69, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x4b, 0x69, 0x6c,
	0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x50,
	0x72, 0x75, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6e, 0x6b, 0x69, 0x73, 0x6f, 0x66, 0x74, 0x2f, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_docker_resp_proto_rawDescOnce sync.Once
	file_docker_resp_proto_rawDescData = file_docker_resp_proto_rawDesc
)

func file_docker_resp_proto_rawDescGZIP() []byte {
	file_docker_resp_proto_rawDescOnce.Do(func() {
		file_docker_resp_proto_rawDescData = protoimpl.X.CompressGZIP(file_docker_resp_proto_rawDescData)
	})
	return file_docker_resp_proto_rawDescData
}

var file_docker_resp_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_docker_resp_proto_goTypes = []interface{}{
	(*ListContainersRequest)(nil),    // 0: ListContainersRequest
	(*ListContainersResponse)(nil),   // 1: ListContainersResponse
	(*CreateContainerRequest)(nil),   // 2: CreateContainerRequest
	(*CreateContainerResponse)(nil),  // 3: CreateContainerResponse
	(*StartContainerRequest)(nil),    // 4: StartContainerRequest
	(*StartContainerResponse)(nil),   // 5: StartContainerResponse
	(*StopContainerRequest)(nil),     // 6: StopContainerRequest
	(*StopContainerResponse)(nil),    // 7: StopContainerResponse
	(*RestartContainerRequest)(nil),  // 8: RestartContainerRequest
	(*RestartContainerResponse)(nil), // 9: RestartContainerResponse
	(*RemoveContainerRequest)(nil),   // 10: RemoveContainerRequest
	(*RemoveContainerResponse)(nil),  // 11: RemoveContainerResponse
	(*KillContainerRequest)(nil),     // 12: KillContainerRequest
	(*KillContainerResponse)(nil),    // 13: KillContainerResponse
	(*PruneContainerRequest)(nil),    // 14: PruneContainerRequest
	(*PruneContainerResponse)(nil),   // 15: PruneContainerResponse
}
var file_docker_resp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_docker_resp_proto_init() }
func file_docker_resp_proto_init() {
	if File_docker_resp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_docker_resp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListContainersRequest); i {
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
		file_docker_resp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListContainersResponse); i {
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
		file_docker_resp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContainerRequest); i {
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
		file_docker_resp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContainerResponse); i {
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
		file_docker_resp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartContainerRequest); i {
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
		file_docker_resp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartContainerResponse); i {
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
		file_docker_resp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopContainerRequest); i {
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
		file_docker_resp_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopContainerResponse); i {
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
		file_docker_resp_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestartContainerRequest); i {
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
		file_docker_resp_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestartContainerResponse); i {
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
		file_docker_resp_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveContainerRequest); i {
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
		file_docker_resp_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveContainerResponse); i {
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
		file_docker_resp_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KillContainerRequest); i {
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
		file_docker_resp_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KillContainerResponse); i {
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
		file_docker_resp_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PruneContainerRequest); i {
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
		file_docker_resp_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PruneContainerResponse); i {
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
			RawDescriptor: file_docker_resp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_docker_resp_proto_goTypes,
		DependencyIndexes: file_docker_resp_proto_depIdxs,
		MessageInfos:      file_docker_resp_proto_msgTypes,
	}.Build()
	File_docker_resp_proto = out.File
	file_docker_resp_proto_rawDesc = nil
	file_docker_resp_proto_goTypes = nil
	file_docker_resp_proto_depIdxs = nil
}
