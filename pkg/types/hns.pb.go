// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hns.proto

package types

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type HnsGetNetworkRequest struct {
	// Types that are valid to be assigned to Options:
	//	*HnsGetNetworkRequest_Address
	//	*HnsGetNetworkRequest_Name
	Options isHnsGetNetworkRequest_Options `protobuf_oneof:"Options"`
}

func (m *HnsGetNetworkRequest) Reset()         { *m = HnsGetNetworkRequest{} }
func (m *HnsGetNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*HnsGetNetworkRequest) ProtoMessage()    {}
func (*HnsGetNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecdc1a541a08ef53, []int{0}
}
func (m *HnsGetNetworkRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HnsGetNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HnsGetNetworkRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HnsGetNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HnsGetNetworkRequest.Merge(m, src)
}
func (m *HnsGetNetworkRequest) XXX_Size() int {
	return m.Size()
}
func (m *HnsGetNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HnsGetNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HnsGetNetworkRequest proto.InternalMessageInfo

type isHnsGetNetworkRequest_Options interface {
	isHnsGetNetworkRequest_Options()
	MarshalTo([]byte) (int, error)
	Size() int
}

type HnsGetNetworkRequest_Address struct {
	Address string `protobuf:"bytes,1,opt,name=Address,proto3,oneof"`
}
type HnsGetNetworkRequest_Name struct {
	Name string `protobuf:"bytes,2,opt,name=Name,proto3,oneof"`
}

func (*HnsGetNetworkRequest_Address) isHnsGetNetworkRequest_Options() {}
func (*HnsGetNetworkRequest_Name) isHnsGetNetworkRequest_Options()    {}

func (m *HnsGetNetworkRequest) GetOptions() isHnsGetNetworkRequest_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *HnsGetNetworkRequest) GetAddress() string {
	if x, ok := m.GetOptions().(*HnsGetNetworkRequest_Address); ok {
		return x.Address
	}
	return ""
}

func (m *HnsGetNetworkRequest) GetName() string {
	if x, ok := m.GetOptions().(*HnsGetNetworkRequest_Name); ok {
		return x.Name
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HnsGetNetworkRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HnsGetNetworkRequest_OneofMarshaler, _HnsGetNetworkRequest_OneofUnmarshaler, _HnsGetNetworkRequest_OneofSizer, []interface{}{
		(*HnsGetNetworkRequest_Address)(nil),
		(*HnsGetNetworkRequest_Name)(nil),
	}
}

func _HnsGetNetworkRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HnsGetNetworkRequest)
	// Options
	switch x := m.Options.(type) {
	case *HnsGetNetworkRequest_Address:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Address)
	case *HnsGetNetworkRequest_Name:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Name)
	case nil:
	default:
		return fmt.Errorf("HnsGetNetworkRequest.Options has unexpected type %T", x)
	}
	return nil
}

func _HnsGetNetworkRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HnsGetNetworkRequest)
	switch tag {
	case 1: // Options.Address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Options = &HnsGetNetworkRequest_Address{x}
		return true, err
	case 2: // Options.Name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Options = &HnsGetNetworkRequest_Name{x}
		return true, err
	default:
		return false, nil
	}
}

func _HnsGetNetworkRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HnsGetNetworkRequest)
	// Options
	switch x := m.Options.(type) {
	case *HnsGetNetworkRequest_Address:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Address)))
		n += len(x.Address)
	case *HnsGetNetworkRequest_Name:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HnsGetNetworkResponse struct {
	Data *HnsNetwork `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *HnsGetNetworkResponse) Reset()         { *m = HnsGetNetworkResponse{} }
func (m *HnsGetNetworkResponse) String() string { return proto.CompactTextString(m) }
func (*HnsGetNetworkResponse) ProtoMessage()    {}
func (*HnsGetNetworkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecdc1a541a08ef53, []int{1}
}
func (m *HnsGetNetworkResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HnsGetNetworkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HnsGetNetworkResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HnsGetNetworkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HnsGetNetworkResponse.Merge(m, src)
}
func (m *HnsGetNetworkResponse) XXX_Size() int {
	return m.Size()
}
func (m *HnsGetNetworkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HnsGetNetworkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HnsGetNetworkResponse proto.InternalMessageInfo

func (m *HnsGetNetworkResponse) GetData() *HnsNetwork {
	if m != nil {
		return m.Data
	}
	return nil
}

type HnsNetwork struct {
	ID           string              `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Type         string              `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Subnets      []*HnsNetworkSubnet `protobuf:"bytes,3,rep,name=Subnets,proto3" json:"Subnets,omitempty"`
	ManagementIP string              `protobuf:"bytes,4,opt,name=ManagementIP,proto3" json:"ManagementIP,omitempty"`
}

func (m *HnsNetwork) Reset()         { *m = HnsNetwork{} }
func (m *HnsNetwork) String() string { return proto.CompactTextString(m) }
func (*HnsNetwork) ProtoMessage()    {}
func (*HnsNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecdc1a541a08ef53, []int{2}
}
func (m *HnsNetwork) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HnsNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HnsNetwork.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HnsNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HnsNetwork.Merge(m, src)
}
func (m *HnsNetwork) XXX_Size() int {
	return m.Size()
}
func (m *HnsNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_HnsNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_HnsNetwork proto.InternalMessageInfo

func (m *HnsNetwork) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *HnsNetwork) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *HnsNetwork) GetSubnets() []*HnsNetworkSubnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

func (m *HnsNetwork) GetManagementIP() string {
	if m != nil {
		return m.ManagementIP
	}
	return ""
}

type HnsNetworkSubnet struct {
	AddressCIDR    string `protobuf:"bytes,1,opt,name=AddressCIDR,proto3" json:"AddressCIDR,omitempty"`
	GatewayAddress string `protobuf:"bytes,2,opt,name=GatewayAddress,proto3" json:"GatewayAddress,omitempty"`
}

func (m *HnsNetworkSubnet) Reset()         { *m = HnsNetworkSubnet{} }
func (m *HnsNetworkSubnet) String() string { return proto.CompactTextString(m) }
func (*HnsNetworkSubnet) ProtoMessage()    {}
func (*HnsNetworkSubnet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecdc1a541a08ef53, []int{3}
}
func (m *HnsNetworkSubnet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HnsNetworkSubnet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HnsNetworkSubnet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HnsNetworkSubnet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HnsNetworkSubnet.Merge(m, src)
}
func (m *HnsNetworkSubnet) XXX_Size() int {
	return m.Size()
}
func (m *HnsNetworkSubnet) XXX_DiscardUnknown() {
	xxx_messageInfo_HnsNetworkSubnet.DiscardUnknown(m)
}

var xxx_messageInfo_HnsNetworkSubnet proto.InternalMessageInfo

func (m *HnsNetworkSubnet) GetAddressCIDR() string {
	if m != nil {
		return m.AddressCIDR
	}
	return ""
}

func (m *HnsNetworkSubnet) GetGatewayAddress() string {
	if m != nil {
		return m.GatewayAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*HnsGetNetworkRequest)(nil), "wins.HnsGetNetworkRequest")
	proto.RegisterType((*HnsGetNetworkResponse)(nil), "wins.HnsGetNetworkResponse")
	proto.RegisterType((*HnsNetwork)(nil), "wins.HnsNetwork")
	proto.RegisterType((*HnsNetworkSubnet)(nil), "wins.HnsNetworkSubnet")
}

func init() { proto.RegisterFile("hns.proto", fileDescriptor_ecdc1a541a08ef53) }

var fileDescriptor_ecdc1a541a08ef53 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4f, 0xfa, 0x40,
	0x10, 0xc5, 0x5b, 0xe8, 0xff, 0x4f, 0x18, 0x0c, 0x21, 0x13, 0x34, 0x0d, 0x26, 0x95, 0x34, 0xc6,
	0x70, 0x22, 0x06, 0xcf, 0x1e, 0xc4, 0x26, 0xd0, 0x83, 0x68, 0x8a, 0x5e, 0x8c, 0x97, 0x45, 0x26,
	0x4a, 0x0c, 0xdb, 0xda, 0x59, 0x24, 0x7c, 0x01, 0xcf, 0x7e, 0x2c, 0x8f, 0x1c, 0x3d, 0x1a, 0xf8,
	0x22, 0xc6, 0x6d, 0x2b, 0x5a, 0xbd, 0xed, 0xbe, 0x37, 0x79, 0xfb, 0xdb, 0x37, 0x50, 0xbe, 0x97,
	0xdc, 0x8e, 0xe2, 0x50, 0x85, 0x68, 0xcd, 0x27, 0x92, 0xdd, 0x21, 0xd4, 0xfb, 0x92, 0x7b, 0xa4,
	0x06, 0xa4, 0xe6, 0x61, 0xfc, 0x10, 0xd0, 0xe3, 0x8c, 0x58, 0x61, 0x03, 0x4a, 0x27, 0xe3, 0x71,
	0x4c, 0xcc, 0xb6, 0xd9, 0x34, 0x5b, 0xe5, 0xbe, 0x11, 0x64, 0x02, 0xd6, 0xc1, 0x1a, 0x88, 0x29,
	0xd9, 0x85, 0xd4, 0xd0, 0xb7, 0x6e, 0x19, 0x4a, 0xe7, 0x91, 0x9a, 0x84, 0x92, 0xdd, 0x63, 0xd8,
	0xce, 0x85, 0x72, 0x14, 0x4a, 0x26, 0xdc, 0x07, 0xcb, 0x13, 0x4a, 0xe8, 0xc8, 0x4a, 0xa7, 0xd6,
	0xfe, 0x44, 0x68, 0xf7, 0x25, 0x67, 0x73, 0xda, 0x75, 0x9f, 0x4d, 0x80, 0x8d, 0x88, 0x55, 0x28,
	0xf8, 0x5e, 0x42, 0x11, 0x14, 0x7c, 0x0f, 0x11, 0xac, 0xcb, 0x45, 0x94, 0x3e, 0x1f, 0xe8, 0x33,
	0x1e, 0x42, 0x69, 0x38, 0x1b, 0x49, 0x52, 0x6c, 0x17, 0x9b, 0xc5, 0x56, 0xa5, 0xb3, 0x93, 0xcf,
	0x4e, 0xec, 0x20, 0x1b, 0x43, 0x17, 0xb6, 0xce, 0x84, 0x14, 0x77, 0x34, 0x25, 0xa9, 0xfc, 0x0b,
	0xdb, 0xd2, 0x69, 0x3f, 0x34, 0xf7, 0x06, 0x6a, 0xf9, 0x00, 0x6c, 0x42, 0x25, 0xed, 0xe1, 0xd4,
	0xf7, 0x82, 0x14, 0xeb, 0xbb, 0x84, 0x07, 0x50, 0xed, 0x09, 0x45, 0x73, 0xb1, 0xc8, 0x1a, 0x4c,
	0x48, 0x73, 0x6a, 0xe7, 0x4a, 0xff, 0x72, 0x48, 0xf1, 0xd3, 0xe4, 0x96, 0xb0, 0x07, 0xb0, 0x29,
	0x0c, 0x1b, 0x5f, 0xf8, 0xbf, 0x56, 0xd3, 0xd8, 0xfd, 0xd3, 0x4b, 0x1a, 0x76, 0x8d, 0xee, 0xde,
	0xeb, 0xca, 0x31, 0x97, 0x2b, 0xc7, 0x7c, 0x5f, 0x39, 0xe6, 0xcb, 0xda, 0x31, 0x96, 0x6b, 0xc7,
	0x78, 0x5b, 0x3b, 0xc6, 0xf5, 0x3f, 0xb5, 0x88, 0x88, 0x47, 0xff, 0xf5, 0xfe, 0x8f, 0x3e, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x2c, 0xf4, 0x42, 0x31, 0x0c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HnsServiceClient is the client API for HnsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HnsServiceClient interface {
	GetNetwork(ctx context.Context, in *HnsGetNetworkRequest, opts ...grpc.CallOption) (*HnsGetNetworkResponse, error)
}

type hnsServiceClient struct {
	cc *grpc.ClientConn
}

func NewHnsServiceClient(cc *grpc.ClientConn) HnsServiceClient {
	return &hnsServiceClient{cc}
}

func (c *hnsServiceClient) GetNetwork(ctx context.Context, in *HnsGetNetworkRequest, opts ...grpc.CallOption) (*HnsGetNetworkResponse, error) {
	out := new(HnsGetNetworkResponse)
	err := c.cc.Invoke(ctx, "/wins.HnsService/GetNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HnsServiceServer is the server API for HnsService service.
type HnsServiceServer interface {
	GetNetwork(context.Context, *HnsGetNetworkRequest) (*HnsGetNetworkResponse, error)
}

func RegisterHnsServiceServer(s *grpc.Server, srv HnsServiceServer) {
	s.RegisterService(&_HnsService_serviceDesc, srv)
}

func _HnsService_GetNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HnsGetNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HnsServiceServer).GetNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wins.HnsService/GetNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HnsServiceServer).GetNetwork(ctx, req.(*HnsGetNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HnsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wins.HnsService",
	HandlerType: (*HnsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNetwork",
			Handler:    _HnsService_GetNetwork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hns.proto",
}

func (m *HnsGetNetworkRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HnsGetNetworkRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Options != nil {
		nn1, err := m.Options.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *HnsGetNetworkRequest_Address) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintHns(dAtA, i, uint64(len(m.Address)))
	i += copy(dAtA[i:], m.Address)
	return i, nil
}
func (m *HnsGetNetworkRequest_Name) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x12
	i++
	i = encodeVarintHns(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	return i, nil
}
func (m *HnsGetNetworkResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HnsGetNetworkResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHns(dAtA, i, uint64(m.Data.Size()))
		n2, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *HnsNetwork) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HnsNetwork) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHns(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHns(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.Subnets) > 0 {
		for _, msg := range m.Subnets {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintHns(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.ManagementIP) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintHns(dAtA, i, uint64(len(m.ManagementIP)))
		i += copy(dAtA[i:], m.ManagementIP)
	}
	return i, nil
}

func (m *HnsNetworkSubnet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HnsNetworkSubnet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AddressCIDR) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHns(dAtA, i, uint64(len(m.AddressCIDR)))
		i += copy(dAtA[i:], m.AddressCIDR)
	}
	if len(m.GatewayAddress) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHns(dAtA, i, uint64(len(m.GatewayAddress)))
		i += copy(dAtA[i:], m.GatewayAddress)
	}
	return i, nil
}

func encodeVarintHns(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HnsGetNetworkRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Options != nil {
		n += m.Options.Size()
	}
	return n
}

func (m *HnsGetNetworkRequest_Address) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	n += 1 + l + sovHns(uint64(l))
	return n
}
func (m *HnsGetNetworkRequest_Name) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovHns(uint64(l))
	return n
}
func (m *HnsGetNetworkResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovHns(uint64(l))
	}
	return n
}

func (m *HnsNetwork) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovHns(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovHns(uint64(l))
	}
	if len(m.Subnets) > 0 {
		for _, e := range m.Subnets {
			l = e.Size()
			n += 1 + l + sovHns(uint64(l))
		}
	}
	l = len(m.ManagementIP)
	if l > 0 {
		n += 1 + l + sovHns(uint64(l))
	}
	return n
}

func (m *HnsNetworkSubnet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AddressCIDR)
	if l > 0 {
		n += 1 + l + sovHns(uint64(l))
	}
	l = len(m.GatewayAddress)
	if l > 0 {
		n += 1 + l + sovHns(uint64(l))
	}
	return n
}

func sovHns(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHns(x uint64) (n int) {
	return sovHns(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HnsGetNetworkRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHns
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HnsGetNetworkRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HnsGetNetworkRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = &HnsGetNetworkRequest_Address{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = &HnsGetNetworkRequest_Name{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHns
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHns
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HnsGetNetworkResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHns
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HnsGetNetworkResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HnsGetNetworkResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHns
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &HnsNetwork{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHns
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHns
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HnsNetwork) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHns
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HnsNetwork: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HnsNetwork: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subnets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHns
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subnets = append(m.Subnets, &HnsNetworkSubnet{})
			if err := m.Subnets[len(m.Subnets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ManagementIP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ManagementIP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHns
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHns
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HnsNetworkSubnet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHns
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HnsNetworkSubnet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HnsNetworkSubnet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressCIDR", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressCIDR = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GatewayAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GatewayAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHns
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHns
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHns(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHns
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHns
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHns
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthHns
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHns
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHns
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHns(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHns
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHns = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHns   = fmt.Errorf("proto: integer overflow")
)
