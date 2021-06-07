// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: judger.proto

package pb

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

type CompileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source   []byte `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Language string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *CompileRequest) Reset() {
	*x = CompileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompileRequest) ProtoMessage() {}

func (x *CompileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_judger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompileRequest.ProtoReflect.Descriptor instead.
func (*CompileRequest) Descriptor() ([]byte, []int) {
	return file_judger_proto_rawDescGZIP(), []int{0}
}

func (x *CompileRequest) GetSource() []byte {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *CompileRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type CompileResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exec           map[string][]byte `protobuf:"bytes,1,rep,name=exec,proto3" json:"exec,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Args           []string          `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	Env            []string          `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	ProcLimit      uint64            `protobuf:"varint,4,opt,name=procLimit,proto3" json:"procLimit,omitempty"`
	CompileMessage []byte            `protobuf:"bytes,5,opt,name=compileMessage,proto3" json:"compileMessage,omitempty"`
}

func (x *CompileResult) Reset() {
	*x = CompileResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompileResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompileResult) ProtoMessage() {}

func (x *CompileResult) ProtoReflect() protoreflect.Message {
	mi := &file_judger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompileResult.ProtoReflect.Descriptor instead.
func (*CompileResult) Descriptor() ([]byte, []int) {
	return file_judger_proto_rawDescGZIP(), []int{1}
}

func (x *CompileResult) GetExec() map[string][]byte {
	if x != nil {
		return x.Exec
	}
	return nil
}

func (x *CompileResult) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *CompileResult) GetEnv() []string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *CompileResult) GetProcLimit() uint64 {
	if x != nil {
		return x.ProcLimit
	}
	return 0
}

func (x *CompileResult) GetCompileMessage() []byte {
	if x != nil {
		return x.CompileMessage
	}
	return nil
}

var File_judger_proto protoreflect.FileDescriptor

var file_judger_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x44, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0xe5, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x65, 0x78,
	0x65, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x65, 0x78, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e,
	0x76, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x37, 0x0a, 0x09, 0x45, 0x78, 0x65, 0x63, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x32, 0x3b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x43,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x69, 0x79,
	0x6c, 0x65, 0x2f, 0x55, 0x4f, 0x4a, 0x2d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x72, 0x32, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_judger_proto_rawDescOnce sync.Once
	file_judger_proto_rawDescData = file_judger_proto_rawDesc
)

func file_judger_proto_rawDescGZIP() []byte {
	file_judger_proto_rawDescOnce.Do(func() {
		file_judger_proto_rawDescData = protoimpl.X.CompressGZIP(file_judger_proto_rawDescData)
	})
	return file_judger_proto_rawDescData
}

var file_judger_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_judger_proto_goTypes = []interface{}{
	(*CompileRequest)(nil), // 0: pb.CompileRequest
	(*CompileResult)(nil),  // 1: pb.CompileResult
	nil,                    // 2: pb.CompileResult.ExecEntry
}
var file_judger_proto_depIdxs = []int32{
	2, // 0: pb.CompileResult.exec:type_name -> pb.CompileResult.ExecEntry
	0, // 1: pb.Compile.Compile:input_type -> pb.CompileRequest
	1, // 2: pb.Compile.Compile:output_type -> pb.CompileResult
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_judger_proto_init() }
func file_judger_proto_init() {
	if File_judger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_judger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompileRequest); i {
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
		file_judger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompileResult); i {
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
			RawDescriptor: file_judger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_judger_proto_goTypes,
		DependencyIndexes: file_judger_proto_depIdxs,
		MessageInfos:      file_judger_proto_msgTypes,
	}.Build()
	File_judger_proto = out.File
	file_judger_proto_rawDesc = nil
	file_judger_proto_goTypes = nil
	file_judger_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CompileClient is the client API for Compile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompileClient interface {
	// Compile compiles source code into executables
	Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResult, error)
}

type compileClient struct {
	cc grpc.ClientConnInterface
}

func NewCompileClient(cc grpc.ClientConnInterface) CompileClient {
	return &compileClient{cc}
}

func (c *compileClient) Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResult, error) {
	out := new(CompileResult)
	err := c.cc.Invoke(ctx, "/pb.Compile/Compile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompileServer is the server API for Compile service.
type CompileServer interface {
	// Compile compiles source code into executables
	Compile(context.Context, *CompileRequest) (*CompileResult, error)
}

// UnimplementedCompileServer can be embedded to have forward compatible implementations.
type UnimplementedCompileServer struct {
}

func (*UnimplementedCompileServer) Compile(context.Context, *CompileRequest) (*CompileResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compile not implemented")
}

func RegisterCompileServer(s *grpc.Server, srv CompileServer) {
	s.RegisterService(&_Compile_serviceDesc, srv)
}

func _Compile_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompileServer).Compile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Compile/Compile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompileServer).Compile(ctx, req.(*CompileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Compile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Compile",
	HandlerType: (*CompileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compile",
			Handler:    _Compile_Compile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "judger.proto",
}