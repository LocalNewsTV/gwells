// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/lithology/lithology.proto

package grpc_lithology_v1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LithologyRequest struct {
	Description          string   `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LithologyRequest) Reset()         { *m = LithologyRequest{} }
func (m *LithologyRequest) String() string { return proto.CompactTextString(m) }
func (*LithologyRequest) ProtoMessage()    {}
func (*LithologyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd082db2a96c71d4, []int{0}
}

func (m *LithologyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LithologyRequest.Unmarshal(m, b)
}
func (m *LithologyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LithologyRequest.Marshal(b, m, deterministic)
}
func (m *LithologyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LithologyRequest.Merge(m, src)
}
func (m *LithologyRequest) XXX_Size() int {
	return xxx_messageInfo_LithologyRequest.Size(m)
}
func (m *LithologyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LithologyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LithologyRequest proto.InternalMessageInfo

func (m *LithologyRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type LithologyResponse struct {
	Soils                []string `protobuf:"bytes,1,rep,name=soils,proto3" json:"soils,omitempty"`
	Moisture             string   `protobuf:"bytes,2,opt,name=moisture,proto3" json:"moisture,omitempty"`
	Consistency          string   `protobuf:"bytes,3,opt,name=consistency,proto3" json:"consistency,omitempty"`
	Colour               string   `protobuf:"bytes,4,opt,name=colour,proto3" json:"colour,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LithologyResponse) Reset()         { *m = LithologyResponse{} }
func (m *LithologyResponse) String() string { return proto.CompactTextString(m) }
func (*LithologyResponse) ProtoMessage()    {}
func (*LithologyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd082db2a96c71d4, []int{1}
}

func (m *LithologyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LithologyResponse.Unmarshal(m, b)
}
func (m *LithologyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LithologyResponse.Marshal(b, m, deterministic)
}
func (m *LithologyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LithologyResponse.Merge(m, src)
}
func (m *LithologyResponse) XXX_Size() int {
	return xxx_messageInfo_LithologyResponse.Size(m)
}
func (m *LithologyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LithologyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LithologyResponse proto.InternalMessageInfo

func (m *LithologyResponse) GetSoils() []string {
	if m != nil {
		return m.Soils
	}
	return nil
}

func (m *LithologyResponse) GetMoisture() string {
	if m != nil {
		return m.Moisture
	}
	return ""
}

func (m *LithologyResponse) GetConsistency() string {
	if m != nil {
		return m.Consistency
	}
	return ""
}

func (m *LithologyResponse) GetColour() string {
	if m != nil {
		return m.Colour
	}
	return ""
}

func init() {
	proto.RegisterType((*LithologyRequest)(nil), "grpc.lithology.v1.LithologyRequest")
	proto.RegisterType((*LithologyResponse)(nil), "grpc.lithology.v1.LithologyResponse")
}

func init() { proto.RegisterFile("proto/lithology/lithology.proto", fileDescriptor_dd082db2a96c71d4) }

var fileDescriptor_dd082db2a96c71d4 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xad, 0xab, 0x8b, 0x3b, 0x82, 0xb8, 0x41, 0x24, 0xec, 0xc5, 0x52, 0x3d, 0xec, 0x29,
	0xe2, 0x9f, 0x8f, 0xe1, 0x41, 0xea, 0xd9, 0x83, 0xc6, 0xa1, 0x06, 0x62, 0x26, 0xcd, 0xa4, 0x85,
	0x9e, 0xfc, 0xea, 0x62, 0xaa, 0x6d, 0x50, 0xd8, 0x5b, 0x7e, 0xef, 0x25, 0x79, 0xf3, 0x06, 0x2e,
	0x7c, 0xa0, 0x48, 0xd7, 0xd6, 0xc4, 0x77, 0xb2, 0xd4, 0x0c, 0xf3, 0x49, 0x25, 0x47, 0xac, 0x9b,
	0xe0, 0xb5, 0x9a, 0xd5, 0xfe, 0xa6, 0xba, 0x87, 0xd3, 0x87, 0x5f, 0xae, 0xb1, 0xed, 0x90, 0xa3,
	0x28, 0xe1, 0xf8, 0x0d, 0x59, 0x07, 0xe3, 0xa3, 0x21, 0x27, 0x8b, 0xb2, 0xd8, 0xae, 0xea, 0x5c,
	0xaa, 0x3e, 0x61, 0x9d, 0xbd, 0x62, 0x4f, 0x8e, 0x51, 0x9c, 0xc1, 0x21, 0x93, 0xb1, 0x2c, 0x8b,
	0x72, 0xb1, 0x5d, 0xd5, 0x23, 0x88, 0x0d, 0x1c, 0x7d, 0x90, 0xe1, 0xd8, 0x05, 0x94, 0xfb, 0xe9,
	0xa7, 0x89, 0xbf, 0x83, 0x34, 0x39, 0x36, 0x1c, 0xd1, 0xe9, 0x41, 0x2e, 0xc6, 0xa0, 0x4c, 0x12,
	0xe7, 0xb0, 0xd4, 0x64, 0xa9, 0x0b, 0xf2, 0x20, 0x99, 0x3f, 0x74, 0xdb, 0x66, 0x63, 0x3f, 0x61,
	0xe8, 0x8d, 0x46, 0xf1, 0x0c, 0x27, 0x8f, 0x2f, 0x81, 0x71, 0x32, 0xc4, 0xa5, 0xfa, 0x57, 0x58,
	0xfd, 0x6d, 0xbb, 0xb9, 0xda, 0x7d, 0x69, 0x2c, 0x57, 0xed, 0xbd, 0x2e, 0xd3, 0x0e, 0xef, 0xbe,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x65, 0x91, 0x7f, 0x59, 0x66, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LithologyServiceClient is the client API for LithologyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LithologyServiceClient interface {
	ParseLithology(ctx context.Context, in *LithologyRequest, opts ...grpc.CallOption) (*LithologyResponse, error)
}

type lithologyServiceClient struct {
	cc *grpc.ClientConn
}

func NewLithologyServiceClient(cc *grpc.ClientConn) LithologyServiceClient {
	return &lithologyServiceClient{cc}
}

func (c *lithologyServiceClient) ParseLithology(ctx context.Context, in *LithologyRequest, opts ...grpc.CallOption) (*LithologyResponse, error) {
	out := new(LithologyResponse)
	err := c.cc.Invoke(ctx, "/grpc.lithology.v1.LithologyService/ParseLithology", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LithologyServiceServer is the server API for LithologyService service.
type LithologyServiceServer interface {
	ParseLithology(context.Context, *LithologyRequest) (*LithologyResponse, error)
}

func RegisterLithologyServiceServer(s *grpc.Server, srv LithologyServiceServer) {
	s.RegisterService(&_LithologyService_serviceDesc, srv)
}

func _LithologyService_ParseLithology_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LithologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LithologyServiceServer).ParseLithology(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.lithology.v1.LithologyService/ParseLithology",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LithologyServiceServer).ParseLithology(ctx, req.(*LithologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LithologyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.lithology.v1.LithologyService",
	HandlerType: (*LithologyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParseLithology",
			Handler:    _LithologyService_ParseLithology_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/lithology/lithology.proto",
}
