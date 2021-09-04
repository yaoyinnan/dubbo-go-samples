// Code generated by protoc-gen-go. DO NOT EDIT.
// source: samples_api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

import (
	"dubbo.apache.org/dubbo-go/v3/protocol"
	dgrpc "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	"dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	"github.com/dubbogo/triple/pkg/common"
	tripleConstant "github.com/dubbogo/triple/pkg/common/constant"
	dubbo3 "github.com/dubbogo/triple/pkg/triple"
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

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d196fe99e714eb, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d196fe99e714eb, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "api.HelloRequest")
	proto.RegisterType((*User)(nil), "api.User")
}

func init() { proto.RegisterFile("samples_api.proto", fileDescriptor_67d196fe99e714eb) }

var fileDescriptor_67d196fe99e714eb = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4e, 0xcc, 0x2d,
	0xc8, 0x49, 0x2d, 0x8e, 0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e,
	0x2c, 0xc8, 0x54, 0x52, 0xe2, 0xe2, 0xf1, 0x48, 0xcd, 0xc9, 0xc9, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x95, 0x6c, 0xb8, 0x58, 0x42, 0x8b, 0x53, 0x8b, 0xb0, 0xc9, 0x09, 0xf1, 0x71,
	0x31, 0x65, 0xa6, 0x48, 0x30, 0x81, 0x45, 0x98, 0x32, 0x53, 0x84, 0x04, 0xb8, 0x98, 0x13, 0xd3,
	0x53, 0x25, 0x98, 0x15, 0x18, 0x35, 0x58, 0x83, 0x40, 0x4c, 0x23, 0x53, 0x2e, 0x76, 0xf7, 0xa2,
	0xd4, 0xd4, 0x92, 0xd4, 0x22, 0x21, 0x2d, 0x2e, 0x8e, 0xe0, 0xc4, 0x4a, 0xb0, 0x7d, 0x42, 0x82,
	0x7a, 0x20, 0x97, 0x20, 0xdb, 0x2d, 0xc5, 0x09, 0x16, 0x02, 0x59, 0xa5, 0xc4, 0x90, 0xc4, 0x06,
	0x76, 0xa4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x46, 0x82, 0xe3, 0x9f, 0xb9, 0x00, 0x00, 0x00,
}

type greeterDubbo3Client struct {
	cc *dubbo3.TripleConn
}

func NewGreeterDubbo3Client(cc *dubbo3.TripleConn) GreeterClient {
	return &greeterDubbo3Client{cc}
}
func (c *greeterDubbo3Client) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*User, common.ErrorWithAttachment) {
	out := new(User)
	interfaceKey := ctx.Value(tripleConstant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/SayHello", in, out)
}

// GreeterClientImpl is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*User, common.ErrorWithAttachment)
}

type GreeterClientImpl struct {
	// Sends a greeting
	SayHello func(ctx context.Context, in *HelloRequest) (*User, error)
}

type Greeter_SayHelloClient interface {
	CloseAndRecv() (*User, error)
	grpc.ClientStream
}

type greeterSayHelloClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloClient) CloseAndRecv() (*User, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *GreeterClientImpl) Reference() string {
	return "greeterImpl"
}

func (c *GreeterClientImpl) GetDubboStub(cc *dubbo3.TripleConn) GreeterClient {
	return NewGreeterDubbo3Client(cc)
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*User, error)
}

type GreeterProviderBase struct {
	proxyImpl protocol.Invoker
}

func (s *GreeterProviderBase) SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *GreeterProviderBase) GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (c *GreeterProviderBase) Reference() string {
	return "greeterImpl"
}

func _Triple_Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dgrpc.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	invo := invocation.NewRPCInvocation("SayHello", args, nil)
	if interceptor == nil {
		result := base.GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.GetProxyImpl().Invoke(context.Background(), invo)
		return result.Result(), result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func (s *GreeterProviderBase) ServiceDesc() *grpc.ServiceDesc {
	return &grpc.ServiceDesc{
		ServiceName: "api.Greeter",
		HandlerType: (*GreeterServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "SayHello",
				Handler:    _Triple_Greeter_SayHello_Handler,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "samples_api.proto",
	}
}
