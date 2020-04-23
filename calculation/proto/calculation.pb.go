// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: calculation/proto/calculation.proto

package proto

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

type Factors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstFactor  float64 `protobuf:"fixed64,1,opt,name=first_factor,json=firstFactor,proto3" json:"first_factor,omitempty"`
	SecondFactor float64 `protobuf:"fixed64,2,opt,name=second_factor,json=secondFactor,proto3" json:"second_factor,omitempty"`
}

func (x *Factors) Reset() {
	*x = Factors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculation_proto_calculation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Factors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Factors) ProtoMessage() {}

func (x *Factors) ProtoReflect() protoreflect.Message {
	mi := &file_calculation_proto_calculation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Factors.ProtoReflect.Descriptor instead.
func (*Factors) Descriptor() ([]byte, []int) {
	return file_calculation_proto_calculation_proto_rawDescGZIP(), []int{0}
}

func (x *Factors) GetFirstFactor() float64 {
	if x != nil {
		return x.FirstFactor
	}
	return 0
}

func (x *Factors) GetSecondFactor() float64 {
	if x != nil {
		return x.SecondFactor
	}
	return 0
}

type FactorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Factors *Factors `protobuf:"bytes,1,opt,name=factors,proto3" json:"factors,omitempty"`
}

func (x *FactorsRequest) Reset() {
	*x = FactorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculation_proto_calculation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FactorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FactorsRequest) ProtoMessage() {}

func (x *FactorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculation_proto_calculation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FactorsRequest.ProtoReflect.Descriptor instead.
func (*FactorsRequest) Descriptor() ([]byte, []int) {
	return file_calculation_proto_calculation_proto_rawDescGZIP(), []int{1}
}

func (x *FactorsRequest) GetFactors() *Factors {
	if x != nil {
		return x.Factors
	}
	return nil
}

type FactorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product float64 `protobuf:"fixed64,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *FactorsResponse) Reset() {
	*x = FactorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculation_proto_calculation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FactorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FactorsResponse) ProtoMessage() {}

func (x *FactorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculation_proto_calculation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FactorsResponse.ProtoReflect.Descriptor instead.
func (*FactorsResponse) Descriptor() ([]byte, []int) {
	return file_calculation_proto_calculation_proto_rawDescGZIP(), []int{2}
}

func (x *FactorsResponse) GetProduct() float64 {
	if x != nil {
		return x.Product
	}
	return 0
}

var File_calculation_proto_calculation_proto protoreflect.FileDescriptor

var file_calculation_proto_calculation_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x07, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x40, 0x0a, 0x0e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x07,
	0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x2b, 0x0a, 0x0f, 0x46, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x32, 0x5d, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculation_proto_calculation_proto_rawDescOnce sync.Once
	file_calculation_proto_calculation_proto_rawDescData = file_calculation_proto_calculation_proto_rawDesc
)

func file_calculation_proto_calculation_proto_rawDescGZIP() []byte {
	file_calculation_proto_calculation_proto_rawDescOnce.Do(func() {
		file_calculation_proto_calculation_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculation_proto_calculation_proto_rawDescData)
	})
	return file_calculation_proto_calculation_proto_rawDescData
}

var file_calculation_proto_calculation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_calculation_proto_calculation_proto_goTypes = []interface{}{
	(*Factors)(nil),         // 0: calculation.Factors
	(*FactorsRequest)(nil),  // 1: calculation.FactorsRequest
	(*FactorsResponse)(nil), // 2: calculation.FactorsResponse
}
var file_calculation_proto_calculation_proto_depIdxs = []int32{
	0, // 0: calculation.FactorsRequest.factors:type_name -> calculation.Factors
	1, // 1: calculation.CalculationService.Multiply:input_type -> calculation.FactorsRequest
	2, // 2: calculation.CalculationService.Multiply:output_type -> calculation.FactorsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_calculation_proto_calculation_proto_init() }
func file_calculation_proto_calculation_proto_init() {
	if File_calculation_proto_calculation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculation_proto_calculation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Factors); i {
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
		file_calculation_proto_calculation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FactorsRequest); i {
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
		file_calculation_proto_calculation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FactorsResponse); i {
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
			RawDescriptor: file_calculation_proto_calculation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculation_proto_calculation_proto_goTypes,
		DependencyIndexes: file_calculation_proto_calculation_proto_depIdxs,
		MessageInfos:      file_calculation_proto_calculation_proto_msgTypes,
	}.Build()
	File_calculation_proto_calculation_proto = out.File
	file_calculation_proto_calculation_proto_rawDesc = nil
	file_calculation_proto_calculation_proto_goTypes = nil
	file_calculation_proto_calculation_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculationServiceClient is the client API for CalculationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculationServiceClient interface {
	// Unary
	Multiply(ctx context.Context, in *FactorsRequest, opts ...grpc.CallOption) (*FactorsResponse, error)
}

type calculationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculationServiceClient(cc grpc.ClientConnInterface) CalculationServiceClient {
	return &calculationServiceClient{cc}
}

func (c *calculationServiceClient) Multiply(ctx context.Context, in *FactorsRequest, opts ...grpc.CallOption) (*FactorsResponse, error) {
	out := new(FactorsResponse)
	err := c.cc.Invoke(ctx, "/calculation.CalculationService/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculationServiceServer is the server API for CalculationService service.
type CalculationServiceServer interface {
	// Unary
	Multiply(context.Context, *FactorsRequest) (*FactorsResponse, error)
}

// UnimplementedCalculationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculationServiceServer struct {
}

func (*UnimplementedCalculationServiceServer) Multiply(context.Context, *FactorsRequest) (*FactorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}

func RegisterCalculationServiceServer(s *grpc.Server, srv CalculationServiceServer) {
	s.RegisterService(&_CalculationService_serviceDesc, srv)
}

func _CalculationService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FactorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculationServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculation.CalculationService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculationServiceServer).Multiply(ctx, req.(*FactorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculation.CalculationService",
	HandlerType: (*CalculationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Multiply",
			Handler:    _CalculationService_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculation/proto/calculation.proto",
}
