// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/permissions/v1beta1/resources.proto

package permissionsv1beta1

import (
	fmt "fmt"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/identity/group/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	proto "github.com/golang/protobuf/proto"
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

// SubjectReference references either a user or a group by id.
type SubjectReference struct {
	// Types that are valid to be assigned to Spec:
	//	*SubjectReference_UserId
	//	*SubjectReference_GroupId
	Spec                 isSubjectReference_Spec `protobuf_oneof:"spec"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SubjectReference) Reset()         { *m = SubjectReference{} }
func (m *SubjectReference) String() string { return proto.CompactTextString(m) }
func (*SubjectReference) ProtoMessage()    {}
func (*SubjectReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_3530957b3a793cab, []int{0}
}

func (m *SubjectReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubjectReference.Unmarshal(m, b)
}
func (m *SubjectReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubjectReference.Marshal(b, m, deterministic)
}
func (m *SubjectReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubjectReference.Merge(m, src)
}
func (m *SubjectReference) XXX_Size() int {
	return xxx_messageInfo_SubjectReference.Size(m)
}
func (m *SubjectReference) XXX_DiscardUnknown() {
	xxx_messageInfo_SubjectReference.DiscardUnknown(m)
}

var xxx_messageInfo_SubjectReference proto.InternalMessageInfo

type isSubjectReference_Spec interface {
	isSubjectReference_Spec()
}

type SubjectReference_UserId struct {
	UserId *v1beta1.UserId `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3,oneof"`
}

type SubjectReference_GroupId struct {
	GroupId *v1beta11.GroupId `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3,oneof"`
}

func (*SubjectReference_UserId) isSubjectReference_Spec() {}

func (*SubjectReference_GroupId) isSubjectReference_Spec() {}

func (m *SubjectReference) GetSpec() isSubjectReference_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *SubjectReference) GetUserId() *v1beta1.UserId {
	if x, ok := m.GetSpec().(*SubjectReference_UserId); ok {
		return x.UserId
	}
	return nil
}

func (m *SubjectReference) GetGroupId() *v1beta11.GroupId {
	if x, ok := m.GetSpec().(*SubjectReference_GroupId); ok {
		return x.GroupId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubjectReference) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubjectReference_UserId)(nil),
		(*SubjectReference_GroupId)(nil),
	}
}

func init() {
	proto.RegisterType((*SubjectReference)(nil), "cs3.permissions.v1beta1.SubjectReference")
}

func init() {
	proto.RegisterFile("cs3/permissions/v1beta1/resources.proto", fileDescriptor_3530957b3a793cab)
}

var fileDescriptor_3530957b3a793cab = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0x2e, 0x36, 0xd6,
	0x2f, 0x48, 0x2d, 0xca, 0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4f, 0x2e, 0x36, 0xd6, 0x43, 0x52, 0xa8, 0x07, 0x55, 0x28, 0xa5,
	0x05, 0x32, 0x21, 0x33, 0x25, 0x35, 0xaf, 0x24, 0xb3, 0xa4, 0x52, 0x3f, 0xbd, 0x28, 0xbf, 0xb4,
	0x00, 0x97, 0x21, 0x52, 0x9a, 0x28, 0x6a, 0x4b, 0x8b, 0x53, 0x8b, 0x70, 0x29, 0x55, 0x9a, 0xc5,
	0xc8, 0x25, 0x10, 0x5c, 0x9a, 0x94, 0x95, 0x9a, 0x5c, 0x12, 0x94, 0x9a, 0x96, 0x5a, 0x94, 0x9a,
	0x97, 0x9c, 0x2a, 0x64, 0xc3, 0xc5, 0x0e, 0xd2, 0x14, 0x9f, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x6d, 0xa4, 0xa8, 0x07, 0x72, 0x16, 0xcc, 0x44, 0x3d, 0x90, 0x24, 0xcc, 0x61, 0x7a, 0xa1,
	0xc5, 0xa9, 0x45, 0x9e, 0x29, 0x1e, 0x0c, 0x41, 0x6c, 0xa5, 0x60, 0x96, 0x90, 0x03, 0x17, 0x07,
	0xd8, 0x79, 0x20, 0xed, 0x4c, 0x60, 0xed, 0xca, 0xa8, 0xda, 0xc1, 0xb2, 0x70, 0xfd, 0xee, 0x20,
	0x1e, 0xd8, 0x00, 0xf6, 0x74, 0x08, 0xd3, 0x89, 0x8d, 0x8b, 0xa5, 0xb8, 0x20, 0x35, 0xd9, 0xa9,
	0x96, 0x4b, 0x3a, 0x39, 0x3f, 0x57, 0x0f, 0x47, 0x90, 0x38, 0xf1, 0x05, 0xc1, 0x3c, 0x13, 0x00,
	0xf2, 0x4b, 0x00, 0x63, 0x94, 0x10, 0x92, 0x32, 0xa8, 0xaa, 0x45, 0x4c, 0xcc, 0xce, 0x01, 0x11,
	0xab, 0x98, 0xc4, 0x9d, 0x8b, 0x8d, 0xf5, 0x02, 0x90, 0x4c, 0x09, 0x33, 0x74, 0x02, 0xc9, 0x9f,
	0x02, 0xcb, 0xc4, 0x20, 0xc9, 0xc4, 0x40, 0x65, 0x92, 0xd8, 0xc0, 0x41, 0x64, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xe2, 0x1c, 0xb5, 0x05, 0xbd, 0x01, 0x00, 0x00,
}
