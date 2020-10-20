// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: api.proto

package uinput

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UinputTriggerMessage_PluginMode int32

const (
	UinputTriggerMessage_POD       UinputTriggerMessage_PluginMode = 0
	UinputTriggerMessage_CONTAINER UinputTriggerMessage_PluginMode = 1
)

// Enum value maps for UinputTriggerMessage_PluginMode.
var (
	UinputTriggerMessage_PluginMode_name = map[int32]string{
		0: "POD",
		1: "CONTAINER",
	}
	UinputTriggerMessage_PluginMode_value = map[string]int32{
		"POD":       0,
		"CONTAINER": 1,
	}
)

func (x UinputTriggerMessage_PluginMode) Enum() *UinputTriggerMessage_PluginMode {
	p := new(UinputTriggerMessage_PluginMode)
	*p = x
	return p
}

func (x UinputTriggerMessage_PluginMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UinputTriggerMessage_PluginMode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (UinputTriggerMessage_PluginMode) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x UinputTriggerMessage_PluginMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UinputTriggerMessage_PluginMode.Descriptor instead.
func (UinputTriggerMessage_PluginMode) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0, 0}
}

type UinputTriggerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodName    string                          `protobuf:"bytes,1,opt,name=podName,proto3" json:"podName,omitempty"`
	PluginMode UinputTriggerMessage_PluginMode `protobuf:"varint,2,opt,name=pluginMode,proto3,enum=uinput.UinputTriggerMessage_PluginMode" json:"pluginMode,omitempty"`
}

func (x *UinputTriggerMessage) Reset() {
	*x = UinputTriggerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UinputTriggerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UinputTriggerMessage) ProtoMessage() {}

func (x *UinputTriggerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UinputTriggerMessage.ProtoReflect.Descriptor instead.
func (*UinputTriggerMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *UinputTriggerMessage) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *UinputTriggerMessage) GetPluginMode() UinputTriggerMessage_PluginMode {
	if x != nil {
		return x.PluginMode
	}
	return UinputTriggerMessage_POD
}

type UinputTriggerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UinputTriggerResponse) Reset() {
	*x = UinputTriggerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UinputTriggerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UinputTriggerResponse) ProtoMessage() {}

func (x *UinputTriggerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UinputTriggerResponse.ProtoReflect.Descriptor instead.
func (*UinputTriggerResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x75, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x22, 0x9f, 0x01, 0x0a, 0x14, 0x55, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x75, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x2e, 0x55, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x22,
	0x24, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x07, 0x0a,
	0x03, 0x50, 0x4f, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49,
	0x4e, 0x45, 0x52, 0x10, 0x01, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb2,
	0x01, 0x0a, 0x0b, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50,
	0x0a, 0x11, 0x55, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4f,
	0x70, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x75, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x55, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x1d, 0x2e, 0x75, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x55, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x51, 0x0a, 0x12, 0x55, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x1c, 0x2e, 0x75, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e,
	0x55, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x1d, 0x2e, 0x75, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x55, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_goTypes = []interface{}{
	(UinputTriggerMessage_PluginMode)(0), // 0: uinput.UinputTriggerMessage.PluginMode
	(*UinputTriggerMessage)(nil),         // 1: uinput.UinputTriggerMessage
	(*UinputTriggerResponse)(nil),        // 2: uinput.UinputTriggerResponse
}
var file_api_proto_depIdxs = []int32{
	0, // 0: uinput.UinputTriggerMessage.pluginMode:type_name -> uinput.UinputTriggerMessage.PluginMode
	1, // 1: uinput.HostService.UinputTriggerOpen:input_type -> uinput.UinputTriggerMessage
	1, // 2: uinput.HostService.UinputTriggerClose:input_type -> uinput.UinputTriggerMessage
	2, // 3: uinput.HostService.UinputTriggerOpen:output_type -> uinput.UinputTriggerResponse
	2, // 4: uinput.HostService.UinputTriggerClose:output_type -> uinput.UinputTriggerResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UinputTriggerMessage); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UinputTriggerResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HostServiceClient is the client API for HostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HostServiceClient interface {
	UinputTriggerOpen(ctx context.Context, in *UinputTriggerMessage, opts ...grpc.CallOption) (*UinputTriggerResponse, error)
	UinputTriggerClose(ctx context.Context, in *UinputTriggerMessage, opts ...grpc.CallOption) (*UinputTriggerResponse, error)
}

type hostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostServiceClient(cc grpc.ClientConnInterface) HostServiceClient {
	return &hostServiceClient{cc}
}

func (c *hostServiceClient) UinputTriggerOpen(ctx context.Context, in *UinputTriggerMessage, opts ...grpc.CallOption) (*UinputTriggerResponse, error) {
	out := new(UinputTriggerResponse)
	err := c.cc.Invoke(ctx, "/uinput.HostService/UinputTriggerOpen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostServiceClient) UinputTriggerClose(ctx context.Context, in *UinputTriggerMessage, opts ...grpc.CallOption) (*UinputTriggerResponse, error) {
	out := new(UinputTriggerResponse)
	err := c.cc.Invoke(ctx, "/uinput.HostService/UinputTriggerClose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostServiceServer is the server API for HostService service.
type HostServiceServer interface {
	UinputTriggerOpen(context.Context, *UinputTriggerMessage) (*UinputTriggerResponse, error)
	UinputTriggerClose(context.Context, *UinputTriggerMessage) (*UinputTriggerResponse, error)
}

// UnimplementedHostServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHostServiceServer struct {
}

func (*UnimplementedHostServiceServer) UinputTriggerOpen(context.Context, *UinputTriggerMessage) (*UinputTriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UinputTriggerOpen not implemented")
}
func (*UnimplementedHostServiceServer) UinputTriggerClose(context.Context, *UinputTriggerMessage) (*UinputTriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UinputTriggerClose not implemented")
}

func RegisterHostServiceServer(s *grpc.Server, srv HostServiceServer) {
	s.RegisterService(&_HostService_serviceDesc, srv)
}

func _HostService_UinputTriggerOpen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UinputTriggerMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).UinputTriggerOpen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uinput.HostService/UinputTriggerOpen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).UinputTriggerOpen(ctx, req.(*UinputTriggerMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostService_UinputTriggerClose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UinputTriggerMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).UinputTriggerClose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uinput.HostService/UinputTriggerClose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).UinputTriggerClose(ctx, req.(*UinputTriggerMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _HostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "uinput.HostService",
	HandlerType: (*HostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UinputTriggerOpen",
			Handler:    _HostService_UinputTriggerOpen_Handler,
		},
		{
			MethodName: "UinputTriggerClose",
			Handler:    _HostService_UinputTriggerClose_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}