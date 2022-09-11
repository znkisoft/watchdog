// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: models.proto

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
	return file_models_proto_enumTypes[0].Descriptor()
}

func (PluginType) Type() protoreflect.EnumType {
	return &file_models_proto_enumTypes[0]
}

func (x PluginType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginType.Descriptor instead.
func (PluginType) EnumDescriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{0}
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
		mi := &file_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[0]
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
	return file_models_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Userver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO set primary key with option field
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Ip            string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port          string `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	Protocol      string `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
	CheckInterval int32  `protobuf:"varint,6,opt,name=check_interval,json=checkInterval,proto3" json:"check_interval,omitempty"`
	Timeout       int32  `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Userver) Reset() {
	*x = Userver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Userver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Userver) ProtoMessage() {}

func (x *Userver) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Userver.ProtoReflect.Descriptor instead.
func (*Userver) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{1}
}

func (x *Userver) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Userver) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Userver) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Userver) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *Userver) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Userver) GetCheckInterval() int32 {
	if x != nil {
		return x.CheckInterval
	}
	return 0
}

func (x *Userver) GetTimeout() int32 {
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
		mi := &file_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UptimeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UptimeInfo) ProtoMessage() {}

func (x *UptimeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[2]
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
	return file_models_proto_rawDescGZIP(), []int{2}
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
		mi := &file_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysTimeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysTimeInfo) ProtoMessage() {}

func (x *SysTimeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[3]
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
	return file_models_proto_rawDescGZIP(), []int{3}
}

func (x *SysTimeInfo) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_models_proto protoreflect.FileDescriptor

var file_models_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x20, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xae, 0x01, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x22, 0x26, 0x0a, 0x0a, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x3d, 0x0a, 0x0b, 0x53, 0x79,
	0x73, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x2a, 0x6b, 0x0a, 0x0a, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x4c, 0x55, 0x47, 0x49,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x10,
	0x03, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x49, 0x53, 0x4b, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x4e,
	0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x59, 0x53, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x10, 0x06, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6e, 0x6b, 0x69, 0x73, 0x6f, 0x66, 0x74, 0x2f, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_proto_rawDescOnce sync.Once
	file_models_proto_rawDescData = file_models_proto_rawDesc
)

func file_models_proto_rawDescGZIP() []byte {
	file_models_proto_rawDescOnce.Do(func() {
		file_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_proto_rawDescData)
	})
	return file_models_proto_rawDescData
}

var file_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_models_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_models_proto_goTypes = []interface{}{
	(PluginType)(0),               // 0: PluginType
	(*Ping)(nil),                  // 1: Ping
	(*Userver)(nil),               // 2: Userver
	(*UptimeInfo)(nil),            // 3: UptimeInfo
	(*SysTimeInfo)(nil),           // 4: SysTimeInfo
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_models_proto_depIdxs = []int32{
	5, // 0: SysTimeInfo.time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_proto_init() }
func file_models_proto_init() {
	if File_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Userver); i {
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
		file_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_proto_goTypes,
		DependencyIndexes: file_models_proto_depIdxs,
		EnumInfos:         file_models_proto_enumTypes,
		MessageInfos:      file_models_proto_msgTypes,
	}.Build()
	File_models_proto = out.File
	file_models_proto_rawDesc = nil
	file_models_proto_goTypes = nil
	file_models_proto_depIdxs = nil
}