// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: responses.proto

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
	return file_responses_proto_enumTypes[0].Descriptor()
}

func (ServerStatusResponse_Status) Type() protoreflect.EnumType {
	return &file_responses_proto_enumTypes[0]
}

func (x ServerStatusResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerStatusResponse_Status.Descriptor instead.
func (ServerStatusResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{0, 0}
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
		mi := &file_responses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatusResponse) ProtoMessage() {}

func (x *ServerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[0]
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
	return file_responses_proto_rawDescGZIP(), []int{0}
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
		mi := &file_responses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginResponse) ProtoMessage() {}

func (x *PluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[1]
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
	return file_responses_proto_rawDescGZIP(), []int{1}
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
		mi := &file_responses_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginListResponse) ProtoMessage() {}

func (x *PluginListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[2]
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
	return file_responses_proto_rawDescGZIP(), []int{2}
}

func (x *PluginListResponse) GetPlugins() []string {
	if x != nil {
		return x.Plugins
	}
	return nil
}

var File_responses_proto protoreflect.FileDescriptor

var file_responses_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
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
	0x4f, 0x57, 0x4e, 0x10, 0x02, 0x22, 0x82, 0x01, 0x0a, 0x0e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x75, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0a, 0x75, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x0d, 0x73, 0x79, 0x73, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x53, 0x79, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52,
	0x0b, 0x73, 0x79, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0c, 0x0a, 0x0a,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2e, 0x0a, 0x12, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6e, 0x6b, 0x69, 0x73, 0x6f, 0x66,
	0x74, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_responses_proto_rawDescOnce sync.Once
	file_responses_proto_rawDescData = file_responses_proto_rawDesc
)

func file_responses_proto_rawDescGZIP() []byte {
	file_responses_proto_rawDescOnce.Do(func() {
		file_responses_proto_rawDescData = protoimpl.X.CompressGZIP(file_responses_proto_rawDescData)
	})
	return file_responses_proto_rawDescData
}

var file_responses_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_responses_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_responses_proto_goTypes = []interface{}{
	(ServerStatusResponse_Status)(0), // 0: ServerStatusResponse.Status
	(*ServerStatusResponse)(nil),     // 1: ServerStatusResponse
	(*PluginResponse)(nil),           // 2: PluginResponse
	(*PluginListResponse)(nil),       // 3: PluginListResponse
	(*timestamppb.Timestamp)(nil),    // 4: google.protobuf.Timestamp
	(*UptimeInfo)(nil),               // 5: UptimeInfo
	(*SysTimeInfo)(nil),              // 6: SysTimeInfo
}
var file_responses_proto_depIdxs = []int32{
	0, // 0: ServerStatusResponse.status:type_name -> ServerStatusResponse.Status
	4, // 1: ServerStatusResponse.last_updated:type_name -> google.protobuf.Timestamp
	5, // 2: PluginResponse.uptime_info:type_name -> UptimeInfo
	6, // 3: PluginResponse.sys_time_info:type_name -> SysTimeInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_responses_proto_init() }
func file_responses_proto_init() {
	if File_responses_proto != nil {
		return
	}
	file_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_responses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_responses_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_responses_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	file_responses_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PluginResponse_UptimeInfo)(nil),
		(*PluginResponse_SysTimeInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_responses_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_responses_proto_goTypes,
		DependencyIndexes: file_responses_proto_depIdxs,
		EnumInfos:         file_responses_proto_enumTypes,
		MessageInfos:      file_responses_proto_msgTypes,
	}.Build()
	File_responses_proto = out.File
	file_responses_proto_rawDesc = nil
	file_responses_proto_goTypes = nil
	file_responses_proto_depIdxs = nil
}