// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/ocm/core/v1beta1/ocm_core_api.proto

package corev1beta1

import (
	context "context"
	fmt "fmt"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
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

// https://rawgit.com/GEANT/OCM-API/v1/docs.html#null%2Fpaths%2F~1shares%2Fpost
type CreateOCMCoreShareRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// OPTIONAL.
	// Description for the share.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// REQUIRED.
	// Name of the resource (file or folder).
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// REQUIRED.
	// Identifier to identify the resource at the provider side. This is unique per provider.
	ProviderId string `protobuf:"bytes,4,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	// REQUIRED.
	// Provider specific identifier of the user that wants to share the resource.
	Owner *v1beta11.UserId `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	// REQUIRED.
	// Consumer specific identifier of the user or group the provider wants to share the resource with.
	// This is known in advance, for example using the OCM invitation flow.
	// Please note that the consumer service endpoint is known in advance as well, so this is no part of the request body.
	// TODO: this field needs to represent either a user or group in the future, not only a user.
	ShareWith *v1beta11.UserId `protobuf:"bytes,6,opt,name=share_with,json=shareWith,proto3" json:"share_with,omitempty"`
	// REQUIRED.
	// The protocol which is used to establish synchronisation.
	Protocol             *Protocol `protobuf:"bytes,7,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateOCMCoreShareRequest) Reset()         { *m = CreateOCMCoreShareRequest{} }
func (m *CreateOCMCoreShareRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOCMCoreShareRequest) ProtoMessage()    {}
func (*CreateOCMCoreShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e8f411283db0fc7, []int{0}
}

func (m *CreateOCMCoreShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOCMCoreShareRequest.Unmarshal(m, b)
}
func (m *CreateOCMCoreShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOCMCoreShareRequest.Marshal(b, m, deterministic)
}
func (m *CreateOCMCoreShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOCMCoreShareRequest.Merge(m, src)
}
func (m *CreateOCMCoreShareRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOCMCoreShareRequest.Size(m)
}
func (m *CreateOCMCoreShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOCMCoreShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOCMCoreShareRequest proto.InternalMessageInfo

func (m *CreateOCMCoreShareRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *CreateOCMCoreShareRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateOCMCoreShareRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateOCMCoreShareRequest) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *CreateOCMCoreShareRequest) GetOwner() *v1beta11.UserId {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *CreateOCMCoreShareRequest) GetShareWith() *v1beta11.UserId {
	if m != nil {
		return m.ShareWith
	}
	return nil
}

func (m *CreateOCMCoreShareRequest) GetProtocol() *Protocol {
	if m != nil {
		return m.Protocol
	}
	return nil
}

type CreateOCMCoreShareResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// Unique ID to identify the share at the consumer side.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// REQUIRED.
	Created              *v1beta1.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateOCMCoreShareResponse) Reset()         { *m = CreateOCMCoreShareResponse{} }
func (m *CreateOCMCoreShareResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOCMCoreShareResponse) ProtoMessage()    {}
func (*CreateOCMCoreShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e8f411283db0fc7, []int{1}
}

func (m *CreateOCMCoreShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOCMCoreShareResponse.Unmarshal(m, b)
}
func (m *CreateOCMCoreShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOCMCoreShareResponse.Marshal(b, m, deterministic)
}
func (m *CreateOCMCoreShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOCMCoreShareResponse.Merge(m, src)
}
func (m *CreateOCMCoreShareResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOCMCoreShareResponse.Size(m)
}
func (m *CreateOCMCoreShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOCMCoreShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOCMCoreShareResponse proto.InternalMessageInfo

func (m *CreateOCMCoreShareResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreateOCMCoreShareResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *CreateOCMCoreShareResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateOCMCoreShareResponse) GetCreated() *v1beta1.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateOCMCoreShareRequest)(nil), "cs3.ocm.core.v1beta1.CreateOCMCoreShareRequest")
	proto.RegisterType((*CreateOCMCoreShareResponse)(nil), "cs3.ocm.core.v1beta1.CreateOCMCoreShareResponse")
}

func init() {
	proto.RegisterFile("cs3/ocm/core/v1beta1/ocm_core_api.proto", fileDescriptor_4e8f411283db0fc7)
}

var fileDescriptor_4e8f411283db0fc7 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x26, 0x69, 0xbb, 0xb5, 0x27, 0xa0, 0x30, 0x14, 0x4c, 0x97, 0xaa, 0x6b, 0x11, 0xac, 0x37,
	0x13, 0x77, 0x17, 0x14, 0xbc, 0xd2, 0xcd, 0x55, 0x2f, 0x24, 0x4b, 0xea, 0x0f, 0xc8, 0xc2, 0x92,
	0x4e, 0x0e, 0xec, 0x80, 0xc9, 0x4c, 0x67, 0x26, 0x5d, 0xfa, 0x00, 0xbe, 0x88, 0x97, 0x3e, 0x89,
	0xf8, 0x0c, 0x3e, 0x8c, 0xcc, 0x64, 0x12, 0x8a, 0x46, 0xd8, 0xbb, 0xe4, 0xfb, 0x39, 0x39, 0xe7,
	0xfb, 0x02, 0xcf, 0x99, 0x9e, 0x27, 0x82, 0x55, 0x09, 0x13, 0x0a, 0x93, 0x9b, 0xe9, 0x15, 0x9a,
	0x62, 0x6a, 0x81, 0xb5, 0x05, 0xd6, 0x85, 0xe4, 0x54, 0x2a, 0x61, 0x04, 0x39, 0x66, 0x7a, 0x4e,
	0x05, 0xab, 0xa8, 0xc5, 0xa9, 0x17, 0x8e, 0x5f, 0x58, 0x3b, 0x2f, 0xb1, 0x36, 0xdc, 0xdc, 0x26,
	0x8d, 0x46, 0xd5, 0xcf, 0x50, 0xa8, 0x45, 0xa3, 0x18, 0xea, 0x76, 0xc0, 0xf8, 0xd9, 0xe0, 0x97,
	0xfe, 0x56, 0x9d, 0x5a, 0x95, 0x92, 0xac, 0x17, 0x68, 0x53, 0x98, 0xa6, 0x63, 0x1f, 0x59, 0xd6,
	0xdc, 0x4a, 0xd4, 0x3d, 0xef, 0xde, 0x5a, 0xfa, 0xec, 0x77, 0x08, 0x27, 0xa9, 0xc2, 0xc2, 0x60,
	0x96, 0xbe, 0x4f, 0x85, 0xc2, 0xcb, 0x4d, 0xa1, 0x30, 0xc7, 0xeb, 0x06, 0xb5, 0x21, 0x53, 0x18,
	0x09, 0x59, 0x5c, 0x37, 0x18, 0x07, 0x93, 0xe0, 0x3c, 0x9a, 0x9d, 0x50, 0x7b, 0x52, 0xeb, 0xf7,
	0xd3, 0x68, 0xe6, 0x04, 0xb9, 0x17, 0x92, 0x09, 0x44, 0x25, 0x6a, 0xa6, 0xb8, 0x34, 0x5c, 0xd4,
	0x71, 0x38, 0x09, 0xce, 0x8f, 0xf2, 0xbb, 0x10, 0x21, 0xb0, 0x5f, 0x17, 0x15, 0xc6, 0x7b, 0x8e,
	0x72, 0xcf, 0xe4, 0x09, 0x44, 0x52, 0x89, 0x1b, 0x5e, 0xa2, 0x5a, 0xf3, 0x32, 0xde, 0x77, 0x14,
	0x74, 0xd0, 0x45, 0x49, 0x5e, 0xc3, 0x81, 0xd8, 0xd6, 0xa8, 0xe2, 0x03, 0xb7, 0xc8, 0x53, 0xb7,
	0x48, 0x97, 0x22, 0xb5, 0x29, 0xf6, 0x0b, 0x7d, 0xd4, 0xd6, 0x91, 0xb7, 0x7a, 0xf2, 0x16, 0x40,
	0xdb, 0x93, 0xd6, 0x5b, 0x6e, 0x36, 0xf1, 0x68, 0x57, 0xf7, 0x91, 0x33, 0x7d, 0xe6, 0x66, 0x43,
	0xde, 0xc0, 0x3d, 0x97, 0x15, 0x13, 0x5f, 0xe3, 0x43, 0xe7, 0x7f, 0x4c, 0x87, 0x9a, 0xa5, 0x4b,
	0xaf, 0xca, 0x7b, 0xfd, 0xd9, 0xcf, 0x00, 0xc6, 0x43, 0xf1, 0x6a, 0x29, 0x6a, 0x8d, 0x24, 0x81,
	0x51, 0x5b, 0x96, 0xcf, 0xf7, 0xa1, 0x1b, 0xac, 0x24, 0xeb, 0x67, 0x5e, 0x3a, 0x3a, 0xf7, 0xb2,
	0x3b, 0x85, 0x84, 0xbb, 0x16, 0x72, 0x1f, 0x42, 0x5e, 0xfa, 0xb0, 0x43, 0x5e, 0x92, 0x57, 0x70,
	0xc8, 0xdc, 0x46, 0x6d, 0xcc, 0xd1, 0xec, 0x74, 0x60, 0xc6, 0x07, 0x5e, 0xa1, 0x36, 0x45, 0x25,
	0xf3, 0x4e, 0x3c, 0xfb, 0x16, 0x00, 0x64, 0xac, 0xb2, 0x47, 0xbc, 0x5b, 0x5e, 0x90, 0x2d, 0x90,
	0x7f, 0x0f, 0x23, 0xc9, 0x70, 0x32, 0xff, 0xfd, 0xc3, 0xc6, 0x2f, 0x77, 0x37, 0xb4, 0x99, 0x2d,
	0x6a, 0x88, 0x99, 0xa8, 0x06, 0x6d, 0x8b, 0x07, 0xdd, 0x82, 0x92, 0xbb, 0x32, 0x96, 0xc1, 0x97,
	0xc8, 0x0a, 0x3c, 0xff, 0x3d, 0xdc, 0x4b, 0xb3, 0xf4, 0x47, 0x78, 0x9c, 0xea, 0x39, 0xcd, 0x58,
	0x45, 0xad, 0x96, 0x7e, 0x9a, 0x2e, 0x2c, 0xf9, 0xcb, 0xc1, 0xab, 0x8c, 0x55, 0x2b, 0x0b, 0xaf,
	0x3c, 0x7c, 0x35, 0x72, 0x65, 0xce, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x27, 0xaa, 0xb1,
	0xf7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OcmCoreAPIClient is the client API for OcmCoreAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OcmCoreAPIClient interface {
	// Creates a new ocm share.
	CreateOCMCoreShare(ctx context.Context, in *CreateOCMCoreShareRequest, opts ...grpc.CallOption) (*CreateOCMCoreShareResponse, error)
}

type ocmCoreAPIClient struct {
	cc *grpc.ClientConn
}

func NewOcmCoreAPIClient(cc *grpc.ClientConn) OcmCoreAPIClient {
	return &ocmCoreAPIClient{cc}
}

func (c *ocmCoreAPIClient) CreateOCMCoreShare(ctx context.Context, in *CreateOCMCoreShareRequest, opts ...grpc.CallOption) (*CreateOCMCoreShareResponse, error) {
	out := new(CreateOCMCoreShareResponse)
	err := c.cc.Invoke(ctx, "/cs3.ocm.core.v1beta1.OcmCoreAPI/CreateOCMCoreShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcmCoreAPIServer is the server API for OcmCoreAPI service.
type OcmCoreAPIServer interface {
	// Creates a new ocm share.
	CreateOCMCoreShare(context.Context, *CreateOCMCoreShareRequest) (*CreateOCMCoreShareResponse, error)
}

// UnimplementedOcmCoreAPIServer can be embedded to have forward compatible implementations.
type UnimplementedOcmCoreAPIServer struct {
}

func (*UnimplementedOcmCoreAPIServer) CreateOCMCoreShare(ctx context.Context, req *CreateOCMCoreShareRequest) (*CreateOCMCoreShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOCMCoreShare not implemented")
}

func RegisterOcmCoreAPIServer(s *grpc.Server, srv OcmCoreAPIServer) {
	s.RegisterService(&_OcmCoreAPI_serviceDesc, srv)
}

func _OcmCoreAPI_CreateOCMCoreShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOCMCoreShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcmCoreAPIServer).CreateOCMCoreShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.ocm.core.v1beta1.OcmCoreAPI/CreateOCMCoreShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcmCoreAPIServer).CreateOCMCoreShare(ctx, req.(*CreateOCMCoreShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OcmCoreAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.ocm.core.v1beta1.OcmCoreAPI",
	HandlerType: (*OcmCoreAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOCMCoreShare",
			Handler:    _OcmCoreAPI_CreateOCMCoreShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/ocm/core/v1beta1/ocm_core_api.proto",
}
