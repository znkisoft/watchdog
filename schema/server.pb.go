// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.3
// source: server.proto

package schema

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PluginType int32

const (
	PluginType_PLUGIN_TYPE_UNKNOWN PluginType = 0
	PluginType_UPTIME              PluginType = 1
	PluginType_MEMORY              PluginType = 2
	PluginType_CPU                 PluginType = 3
	PluginType_DISK                PluginType = 4
	PluginType_NETWORK             PluginType = 5
	PluginType_SYS_TIME            PluginType = 6
)

// Enum value maps for PluginType.
var (
	PluginType_name = map[int32]string{
		0: "PLUGIN_TYPE_UNKNOWN",
		1: "UPTIME",
		2: "MEMORY",
		3: "CPU",
		4: "DISK",
		5: "NETWORK",
		6: "SYS_TIME",
	}
	PluginType_value = map[string]int32{
		"PLUGIN_TYPE_UNKNOWN": 0,
		"UPTIME":              1,
		"MEMORY":              2,
		"CPU":                 3,
		"DISK":                4,
		"NETWORK":             5,
		"SYS_TIME":            6,
	}
)

func (x PluginType) Enum() *PluginType {
	p := new(PluginType)
	*p = x
	return p
}

func (x PluginType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginType) Descriptor() protoreflect.EnumDescriptor {
	return file_server_proto_enumTypes[0].Descriptor()
}

func (PluginType) Type() protoreflect.EnumType {
	return &file_server_proto_enumTypes[0]
}

func (x PluginType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginType.Descriptor instead.
func (PluginType) EnumDescriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

type ServerStatusResponse_Status int32

const (
	ServerStatusResponse_SERVER_STATUS_UNKNOWN ServerStatusResponse_Status = 0
	ServerStatusResponse_UP                    ServerStatusResponse_Status = 1
	ServerStatusResponse_DOWN                  ServerStatusResponse_Status = 2
)

// Enum value maps for ServerStatusResponse_Status.
var (
	ServerStatusResponse_Status_name = map[int32]string{
		0: "SERVER_STATUS_UNKNOWN",
		1: "UP",
		2: "DOWN",
	}
	ServerStatusResponse_Status_value = map[string]int32{
		"SERVER_STATUS_UNKNOWN": 0,
		"UP":                    1,
		"DOWN":                  2,
	}
)

func (x ServerStatusResponse_Status) Enum() *ServerStatusResponse_Status {
	p := new(ServerStatusResponse_Status)
	*p = x
	return p
}

func (x ServerStatusResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerStatusResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_server_proto_enumTypes[1].Descriptor()
}

func (ServerStatusResponse_Status) Type() protoreflect.EnumType {
	return &file_server_proto_enumTypes[1]
}

func (x ServerStatusResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerStatusResponse_Status.Descriptor instead.
func (ServerStatusResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{5, 0}
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ip            string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port          string `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	Protocol      string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	CheckInterval int32  `protobuf:"varint,5,opt,name=check_interval,json=checkInterval,proto3" json:"check_interval,omitempty"`
	Timeout       int32  `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *ServerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServerInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ServerInfo) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ServerInfo) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ServerInfo) GetCheckInterval() int32 {
	if x != nil {
		return x.CheckInterval
	}
	return 0
}

func (x *ServerInfo) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type UptimeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds int32 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
}

func (x *UptimeInfo) Reset() {
	*x = UptimeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UptimeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UptimeInfo) ProtoMessage() {}

func (x *UptimeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UptimeInfo.ProtoReflect.Descriptor instead.
func (*UptimeInfo) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

func (x *UptimeInfo) GetSeconds() int32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

type SysTimeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *SysTimeInfo) Reset() {
	*x = SysTimeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysTimeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysTimeInfo) ProtoMessage() {}

func (x *SysTimeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysTimeInfo.ProtoReflect.Descriptor instead.
func (*SysTimeInfo) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3}
}

func (x *SysTimeInfo) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type ServerStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *ServerStatusRequest) Reset() {
	*x = ServerStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatusRequest) ProtoMessage() {}

func (x *ServerStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatusRequest.ProtoReflect.Descriptor instead.
func (*ServerStatusRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{4}
}

func (x *ServerStatusRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type ServerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      ServerStatusResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ServerStatusResponse_Status" json:"status,omitempty"`
	LastUpdated *timestamppb.Timestamp      `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *ServerStatusResponse) Reset() {
	*x = ServerStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatusResponse) ProtoMessage() {}

func (x *ServerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatusResponse.ProtoReflect.Descriptor instead.
func (*ServerStatusResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{5}
}

func (x *ServerStatusResponse) GetStatus() ServerStatusResponse_Status {
	if x != nil {
		return x.Status
	}
	return ServerStatusResponse_SERVER_STATUS_UNKNOWN
}

func (x *ServerStatusResponse) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

type PluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type PluginType `protobuf:"varint,1,opt,name=type,proto3,enum=PluginType" json:"type,omitempty"`
}

func (x *PluginRequest) Reset() {
	*x = PluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginRequest) ProtoMessage() {}

func (x *PluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginRequest.ProtoReflect.Descriptor instead.
func (*PluginRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{6}
}

func (x *PluginRequest) GetType() PluginType {
	if x != nil {
		return x.Type
	}
	return PluginType_PLUGIN_TYPE_UNKNOWN
}

type PluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to PluginInfo:
	//	*PluginResponse_UptimeInfo
	//	*PluginResponse_SysTimeInfo
	PluginInfo isPluginResponse_PluginInfo `protobuf_oneof:"PluginInfo"`
}

func (x *PluginResponse) Reset() {
	*x = PluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginResponse) ProtoMessage() {}

func (x *PluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginResponse.ProtoReflect.Descriptor instead.
func (*PluginResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{7}
}

func (m *PluginResponse) GetPluginInfo() isPluginResponse_PluginInfo {
	if m != nil {
		return m.PluginInfo
	}
	return nil
}

func (x *PluginResponse) GetUptimeInfo() *UptimeInfo {
	if x, ok := x.GetPluginInfo().(*PluginResponse_UptimeInfo); ok {
		return x.UptimeInfo
	}
	return nil
}

func (x *PluginResponse) GetSysTimeInfo() *SysTimeInfo {
	if x, ok := x.GetPluginInfo().(*PluginResponse_SysTimeInfo); ok {
		return x.SysTimeInfo
	}
	return nil
}

type isPluginResponse_PluginInfo interface {
	isPluginResponse_PluginInfo()
}

type PluginResponse_UptimeInfo struct {
	UptimeInfo *UptimeInfo `protobuf:"bytes,1,opt,name=uptime_info,json=uptimeInfo,proto3,oneof"`
}

type PluginResponse_SysTimeInfo struct {
	SysTimeInfo *SysTimeInfo `protobuf:"bytes,2,opt,name=sys_time_info,json=sysTimeInfo,proto3,oneof"`
}

func (*PluginResponse_UptimeInfo) isPluginResponse_PluginInfo() {}

func (*PluginResponse_SysTimeInfo) isPluginResponse_PluginInfo() {}

type PluginListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugins []string `protobuf:"bytes,1,rep,name=plugins,proto3" json:"plugins,omitempty"`
}

func (x *PluginListResponse) Reset() {
	*x = PluginListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginListResponse) ProtoMessage() {}

func (x *PluginListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginListResponse.ProtoReflect.Descriptor instead.
func (*PluginListResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{8}
}

func (x *PluginListResponse) GetPlugins() []string {
	if x != nil {
		return x.Plugins
	}
	return nil
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x20, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xa1, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x26, 0x0a, 0x0a, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x3d, 0x0a,
	0x0b, 0x53, 0x79, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x13,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x22, 0xc2, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x35,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45, 0x52, 0x56,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44,
	0x4f, 0x57, 0x4e, 0x10, 0x02, 0x22, 0x30, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x0e, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x75, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0a,
	0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x0d, 0x73, 0x79,
	0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x79, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x0b, 0x73, 0x79, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0c,
	0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2e, 0x0a, 0x12,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2a, 0x6b, 0x0a, 0x0a,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x4c,
	0x55, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x43,
	0x50, 0x55, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x49, 0x53, 0x4b, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x59, 0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x06, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6e, 0x6b, 0x69, 0x73, 0x6f, 0x66, 0x74,
	0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_server_proto_goTypes = []interface{}{
	(PluginType)(0),                  // 0: PluginType
	(ServerStatusResponse_Status)(0), // 1: ServerStatusResponse.Status
	(*Ping)(nil),                     // 2: Ping
	(*ServerInfo)(nil),               // 3: ServerInfo
	(*UptimeInfo)(nil),               // 4: UptimeInfo
	(*SysTimeInfo)(nil),              // 5: SysTimeInfo
	(*ServerStatusRequest)(nil),      // 6: ServerStatusRequest
	(*ServerStatusResponse)(nil),     // 7: ServerStatusResponse
	(*PluginRequest)(nil),            // 8: PluginRequest
	(*PluginResponse)(nil),           // 9: PluginResponse
	(*PluginListResponse)(nil),       // 10: PluginListResponse
	(*timestamppb.Timestamp)(nil),    // 11: google.protobuf.Timestamp
}
var file_server_proto_depIdxs = []int32{
	11, // 0: SysTimeInfo.time:type_name -> google.protobuf.Timestamp
	1,  // 1: ServerStatusResponse.status:type_name -> ServerStatusResponse.Status
	11, // 2: ServerStatusResponse.last_updated:type_name -> google.protobuf.Timestamp
	0,  // 3: PluginRequest.type:type_name -> PluginType
	4,  // 4: PluginResponse.uptime_info:type_name -> UptimeInfo
	5,  // 5: PluginResponse.sys_time_info:type_name -> SysTimeInfo
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
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
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UptimeInfo); i {
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
		file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysTimeInfo); i {
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
		file_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStatusRequest); i {
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
		file_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStatusResponse); i {
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
		file_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginRequest); i {
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
		file_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginResponse); i {
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
		file_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginListResponse); i {
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
	file_server_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*PluginResponse_UptimeInfo)(nil),
		(*PluginResponse_SysTimeInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		EnumInfos:         file_server_proto_enumTypes,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
