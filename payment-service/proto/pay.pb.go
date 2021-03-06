package paymentservice

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//PauRequestの定義から自動的に生成

type PayRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Amount               int64    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayRequest) Reset()         { *m = PayRequest{} }
func (m *PayRequest) String() string { return proto.CompactTextString(m) }
func (*PayRequest) ProtoMessage()    {}
func (*PayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pay_a9772073fe55a113, []int{0}
}
func (m *PayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayRequest.Unmarshal(m, b)
}
func (m *PayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayRequest.Marshal(b, m, deterministic)
}
func (dst *PayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayRequest.Merge(dst, src)
}
func (m *PayRequest) XXX_Size() int {
	return xxx_messageInfo_PayRequest.Size(m)
}
func (m *PayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PayRequest proto.InternalMessageInfo

func (m *PayRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PayRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PayRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PayRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PayRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

//PauRequestの定義から自動的に生成
type PayResponse struct {
	Paid                 bool     `protobuf:"varint,1,opt,name=paid,proto3" json:"paid,omitempty"`
	Captured             bool     `protobuf:"varint,3,opt,name=captured,proto3" json:"captured,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayResponse) Reset()         { *m = PayResponse{} }
func (m *PayResponse) String() string { return proto.CompactTextString(m) }
func (*PayResponse) ProtoMessage()    {}
func (*PayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pay_a9772073fe55a113, []int{1}
}
func (m *PayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayResponse.Unmarshal(m, b)
}
func (m *PayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayResponse.Marshal(b, m, deterministic)
}
func (dst *PayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayResponse.Merge(dst, src)
}
func (m *PayResponse) XXX_Size() int {
	return xxx_messageInfo_PayResponse.Size(m)
}
func (m *PayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PayResponse proto.InternalMessageInfo

func (m *PayResponse) GetPaid() bool {
	if m != nil {
		return m.Paid
	}
	return false
}

func (m *PayResponse) GetCaptured() bool {
	if m != nil {
		return m.Captured
	}
	return false
}

func (m *PayResponse) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*PayRequest)(nil), "paymentservice.PayRequest")
	proto.RegisterType((*PayResponse)(nil), "paymentservice.PayResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

//定義したserviceから生成されたinterface
type PayManagerClient interface {
	Charge(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error)
}

type payManagerClient struct {
	cc *grpc.ClientConn
}

func NewPayManagerClient(cc *grpc.ClientConn) PayManagerClient {
	return &payManagerClient{cc}
}

func (c *payManagerClient) Charge(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error) {
	out := new(PayResponse)
	err := c.cc.Invoke(ctx, "/paymentservice.PayManager/Charge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayManagerServer is the server API for PayManager service.
type PayManagerServer interface {
	Charge(context.Context, *PayRequest) (*PayResponse, error)
}

func RegisterPayManagerServer(s *grpc.Server, srv PayManagerServer) {
	s.RegisterService(&_PayManager_serviceDesc, srv)
}

func _PayManager_Charge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayManagerServer).Charge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentservice.PayManager/Charge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayManagerServer).Charge(ctx, req.(*PayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PayManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "paymentservice.PayManager",
	HandlerType: (*PayManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Charge",
			Handler:    _PayManager_Charge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pay.proto",
}

func init() { proto.RegisterFile("proto/pay.proto", fileDescriptor_pay_a9772073fe55a113) }

var fileDescriptor_pay_a9772073fe55a113 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x6d, 0x77, 0xb7, 0xd4, 0x59, 0x58, 0x61, 0x10, 0x09, 0xf5, 0x52, 0x7a, 0xda, 0x53,
	0x05, 0xfd, 0x09, 0x7b, 0x16, 0x34, 0xe0, 0x0f, 0x18, 0xdb, 0x61, 0x0d, 0xd2, 0x24, 0x26, 0xa9,
	0xd0, 0x9b, 0x3f, 0x5d, 0x9c, 0x2d, 0x5a, 0x61, 0x6f, 0xef, 0x65, 0x86, 0xcc, 0xf7, 0x1e, 0x5c,
	0xf9, 0xe0, 0x92, 0xbb, 0xf3, 0x34, 0xb5, 0xa2, 0x70, 0xe7, 0x69, 0x1a, 0xd8, 0xa6, 0xc8, 0xe1,
	0xd3, 0x74, 0xdc, 0x7c, 0x65, 0x00, 0x4f, 0x34, 0x69, 0xfe, 0x18, 0x39, 0x26, 0xdc, 0x41, 0x6e,
	0x7a, 0x95, 0xd5, 0xd9, 0x7e, 0xa5, 0x73, 0xd3, 0xe3, 0x35, 0x6c, 0x92, 0x7b, 0x67, 0xab, 0xf2,
	0x3a, 0xdb, 0x5f, 0xea, 0x93, 0xc1, 0x1b, 0x28, 0x68, 0x70, 0xa3, 0x4d, 0x6a, 0x25, 0x9b, 0xb3,
	0x43, 0x84, 0xb5, 0xa5, 0x81, 0xd5, 0x5a, 0x96, 0x45, 0x63, 0x0d, 0xdb, 0x9e, 0x63, 0x17, 0x8c,
	0x4f, 0xc6, 0x59, 0xb5, 0x91, 0xd1, 0xf2, 0xa9, 0x79, 0x81, 0xad, 0x10, 0x44, 0xef, 0x6c, 0xe4,
	0x9f, 0x4f, 0x3c, 0xcd, 0x10, 0xa5, 0x16, 0x8d, 0x15, 0x94, 0x1d, 0xf9, 0x34, 0x06, 0xee, 0xe5,
	0x64, 0xa9, 0x7f, 0xfd, 0x02, 0x26, 0x5f, 0xc2, 0xdc, 0x3f, 0x4b, 0xb0, 0x47, 0xb2, 0x74, 0xe4,
	0x80, 0x07, 0x28, 0x0e, 0x6f, 0x14, 0x8e, 0x8c, 0x55, 0xfb, 0xbf, 0x82, 0xf6, 0x2f, 0x7e, 0x75,
	0x7b, 0x76, 0x76, 0x02, 0x6b, 0x2e, 0x5e, 0x0b, 0xe9, 0xf0, 0xe1, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xde, 0x90, 0x28, 0xa8, 0x56, 0x01, 0x00, 0x00,
}
