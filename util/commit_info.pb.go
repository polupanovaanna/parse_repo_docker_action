// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commit_info.proto

package util

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CommitInfo struct {
	HeadHash             string   `protobuf:"bytes,1,opt,name=headHash,proto3" json:"headHash,omitempty"`
	CommitDiff           string   `protobuf:"bytes,2,opt,name=commitDiff,proto3" json:"commitDiff,omitempty"`
	CommandLine          string   `protobuf:"bytes,3,opt,name=commandLine,proto3" json:"commandLine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitInfo) Reset()         { *m = CommitInfo{} }
func (m *CommitInfo) String() string { return proto.CompactTextString(m) }
func (*CommitInfo) ProtoMessage()    {}
func (*CommitInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ab387fc4cd736a4, []int{0}
}

func (m *CommitInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitInfo.Unmarshal(m, b)
}
func (m *CommitInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitInfo.Marshal(b, m, deterministic)
}
func (m *CommitInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitInfo.Merge(m, src)
}
func (m *CommitInfo) XXX_Size() int {
	return xxx_messageInfo_CommitInfo.Size(m)
}
func (m *CommitInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CommitInfo proto.InternalMessageInfo

func (m *CommitInfo) GetHeadHash() string {
	if m != nil {
		return m.HeadHash
	}
	return ""
}

func (m *CommitInfo) GetCommitDiff() string {
	if m != nil {
		return m.CommitDiff
	}
	return ""
}

func (m *CommitInfo) GetCommandLine() string {
	if m != nil {
		return m.CommandLine
	}
	return ""
}

type ServerResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerResponse) Reset()         { *m = ServerResponse{} }
func (m *ServerResponse) String() string { return proto.CompactTextString(m) }
func (*ServerResponse) ProtoMessage()    {}
func (*ServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ab387fc4cd736a4, []int{1}
}

func (m *ServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerResponse.Unmarshal(m, b)
}
func (m *ServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerResponse.Marshal(b, m, deterministic)
}
func (m *ServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerResponse.Merge(m, src)
}
func (m *ServerResponse) XXX_Size() int {
	return xxx_messageInfo_ServerResponse.Size(m)
}
func (m *ServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerResponse proto.InternalMessageInfo

func (m *ServerResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*CommitInfo)(nil), "util.CommitInfo")
	proto.RegisterType((*ServerResponse)(nil), "util.ServerResponse")
}

func init() { proto.RegisterFile("commit_info.proto", fileDescriptor_6ab387fc4cd736a4) }

var fileDescriptor_6ab387fc4cd736a4 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xce, 0xcf, 0xcd,
	0xcd, 0x2c, 0x89, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29,
	0x2d, 0xc9, 0xcc, 0x51, 0xca, 0xe2, 0xe2, 0x72, 0x06, 0x4b, 0x79, 0xe6, 0xa5, 0xe5, 0x0b, 0x49,
	0x71, 0x71, 0x64, 0xa4, 0x26, 0xa6, 0x78, 0x24, 0x16, 0x67, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0xc1, 0xf9, 0x42, 0x72, 0x5c, 0x5c, 0x10, 0x43, 0x5c, 0x32, 0xd3, 0xd2, 0x24, 0x98, 0xc0,
	0xb2, 0x48, 0x22, 0x42, 0x0a, 0x5c, 0xdc, 0x20, 0x5e, 0x62, 0x5e, 0x8a, 0x4f, 0x66, 0x5e, 0xaa,
	0x04, 0x33, 0x58, 0x01, 0xb2, 0x90, 0x92, 0x0e, 0x17, 0x5f, 0x70, 0x6a, 0x51, 0x59, 0x6a, 0x51,
	0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0xc8, 0xbe, 0x22, 0x28, 0x1b, 0x66, 0x1f, 0x8c,
	0x6f, 0xe4, 0x0c, 0x73, 0x99, 0x4b, 0x62, 0x49, 0xa2, 0x90, 0x29, 0x17, 0x67, 0x48, 0x51, 0x62,
	0x5e, 0x71, 0x4e, 0x62, 0x49, 0xaa, 0x90, 0x80, 0x1e, 0xc8, 0xed, 0x7a, 0x08, 0x87, 0x4b, 0x89,
	0x40, 0x44, 0x50, 0x8d, 0x57, 0x62, 0x48, 0x62, 0x03, 0xfb, 0xd5, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x21, 0x54, 0x05, 0x2d, 0x00, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommitDataClient is the client API for CommitData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommitDataClient interface {
	Translate(ctx context.Context, in *CommitInfo, opts ...grpc.CallOption) (*ServerResponse, error)
}

type commitDataClient struct {
	cc *grpc.ClientConn
}

func NewCommitDataClient(cc *grpc.ClientConn) CommitDataClient {
	return &commitDataClient{cc}
}

func (c *commitDataClient) Translate(ctx context.Context, in *CommitInfo, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/util.CommitData/Translate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommitDataServer is the server API for CommitData service.
type CommitDataServer interface {
	Translate(context.Context, *CommitInfo) (*ServerResponse, error)
}

// UnimplementedCommitDataServer can be embedded to have forward compatible implementations.
type UnimplementedCommitDataServer struct {
}

func (*UnimplementedCommitDataServer) Translate(ctx context.Context, req *CommitInfo) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Translate not implemented")
}

func RegisterCommitDataServer(s *grpc.Server, srv CommitDataServer) {
	s.RegisterService(&_CommitData_serviceDesc, srv)
}

func _CommitData_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommitDataServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/util.CommitData/Translate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommitDataServer).Translate(ctx, req.(*CommitInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommitData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "util.CommitData",
	HandlerType: (*CommitDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _CommitData_Translate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commit_info.proto",
}
