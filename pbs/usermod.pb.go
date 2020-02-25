// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usermod.proto

package pbs

import (
	fmt "fmt"
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

// option go_package="17jzh.com/user-service/models";
type UsersMod struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Member               int64    `protobuf:"varint,2,opt,name=Member,proto3" json:"Member,omitempty"`
	Realname             string   `protobuf:"bytes,3,opt,name=Realname,proto3" json:"Realname,omitempty"`
	Headimg              string   `protobuf:"bytes,4,opt,name=Headimg,proto3" json:"Headimg,omitempty"`
	Headimg2             string   `protobuf:"bytes,5,opt,name=Headimg2,proto3" json:"Headimg2,omitempty"`
	Mobile               string   `protobuf:"bytes,6,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	RoleId               int32    `protobuf:"varint,7,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	Cid                  int32    `protobuf:"varint,8,opt,name=Cid,proto3" json:"Cid,omitempty"`
	IsVip                int32    `protobuf:"varint,10,opt,name=IsVip,proto3" json:"IsVip,omitempty"`
	Status               int32    `protobuf:"varint,11,opt,name=Status,proto3" json:"Status,omitempty"`
	EduType              int32    `protobuf:"varint,12,opt,name=EduType,proto3" json:"EduType,omitempty"`
	EduYear              int32    `protobuf:"varint,13,opt,name=EduYear,proto3" json:"EduYear,omitempty"`
	Exp                  int32    `protobuf:"varint,14,opt,name=Exp,proto3" json:"Exp,omitempty"`
	LoginAt              string   `protobuf:"bytes,15,opt,name=LoginAt,proto3" json:"LoginAt,omitempty"`
	DeviceId             string   `protobuf:"bytes,16,opt,name=DeviceId,proto3" json:"DeviceId,omitempty"`
	ClientType           int32    `protobuf:"varint,17,opt,name=ClientType,proto3" json:"ClientType,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,18,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	CreatedAt            string   `protobuf:"bytes,19,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	DeletedAt            string   `protobuf:"bytes,20,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	SchoolId             int32    `protobuf:"varint,21,opt,name=SchoolId,proto3" json:"SchoolId,omitempty"`
	ClassroomId          int32    `protobuf:"varint,22,opt,name=ClassroomId,proto3" json:"ClassroomId,omitempty"`
	PlateformId          int32    `protobuf:"varint,23,opt,name=PlateformId,proto3" json:"PlateformId,omitempty"`
	IsNotify             int32    `protobuf:"varint,24,opt,name=IsNotify,proto3" json:"IsNotify,omitempty"`
	Openid               string   `protobuf:"bytes,25,opt,name=Openid,proto3" json:"Openid,omitempty"`
	Nickname             string   `protobuf:"bytes,26,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersMod) Reset()         { *m = UsersMod{} }
func (m *UsersMod) String() string { return proto.CompactTextString(m) }
func (*UsersMod) ProtoMessage()    {}
func (*UsersMod) Descriptor() ([]byte, []int) {
	return fileDescriptor_18846f0aec365394, []int{0}
}

func (m *UsersMod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersMod.Unmarshal(m, b)
}
func (m *UsersMod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersMod.Marshal(b, m, deterministic)
}
func (m *UsersMod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersMod.Merge(m, src)
}
func (m *UsersMod) XXX_Size() int {
	return xxx_messageInfo_UsersMod.Size(m)
}
func (m *UsersMod) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersMod.DiscardUnknown(m)
}

var xxx_messageInfo_UsersMod proto.InternalMessageInfo

func (m *UsersMod) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UsersMod) GetMember() int64 {
	if m != nil {
		return m.Member
	}
	return 0
}

func (m *UsersMod) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *UsersMod) GetHeadimg() string {
	if m != nil {
		return m.Headimg
	}
	return ""
}

func (m *UsersMod) GetHeadimg2() string {
	if m != nil {
		return m.Headimg2
	}
	return ""
}

func (m *UsersMod) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UsersMod) GetRoleId() int32 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *UsersMod) GetCid() int32 {
	if m != nil {
		return m.Cid
	}
	return 0
}

func (m *UsersMod) GetIsVip() int32 {
	if m != nil {
		return m.IsVip
	}
	return 0
}

func (m *UsersMod) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UsersMod) GetEduType() int32 {
	if m != nil {
		return m.EduType
	}
	return 0
}

func (m *UsersMod) GetEduYear() int32 {
	if m != nil {
		return m.EduYear
	}
	return 0
}

func (m *UsersMod) GetExp() int32 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *UsersMod) GetLoginAt() string {
	if m != nil {
		return m.LoginAt
	}
	return ""
}

func (m *UsersMod) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *UsersMod) GetClientType() int32 {
	if m != nil {
		return m.ClientType
	}
	return 0
}

func (m *UsersMod) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *UsersMod) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UsersMod) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

func (m *UsersMod) GetSchoolId() int32 {
	if m != nil {
		return m.SchoolId
	}
	return 0
}

func (m *UsersMod) GetClassroomId() int32 {
	if m != nil {
		return m.ClassroomId
	}
	return 0
}

func (m *UsersMod) GetPlateformId() int32 {
	if m != nil {
		return m.PlateformId
	}
	return 0
}

func (m *UsersMod) GetIsNotify() int32 {
	if m != nil {
		return m.IsNotify
	}
	return 0
}

func (m *UsersMod) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *UsersMod) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

type UserId struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserId) Reset()         { *m = UserId{} }
func (m *UserId) String() string { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()    {}
func (*UserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_18846f0aec365394, []int{1}
}

func (m *UserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserId.Unmarshal(m, b)
}
func (m *UserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserId.Marshal(b, m, deterministic)
}
func (m *UserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserId.Merge(m, src)
}
func (m *UserId) XXX_Size() int {
	return xxx_messageInfo_UserId.Size(m)
}
func (m *UserId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserId.DiscardUnknown(m)
}

var xxx_messageInfo_UserId proto.InternalMessageInfo

func (m *UserId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

//根据名称参数处理业务
type UserName struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserName) Reset()         { *m = UserName{} }
func (m *UserName) String() string { return proto.CompactTextString(m) }
func (*UserName) ProtoMessage()    {}
func (*UserName) Descriptor() ([]byte, []int) {
	return fileDescriptor_18846f0aec365394, []int{2}
}

func (m *UserName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserName.Unmarshal(m, b)
}
func (m *UserName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserName.Marshal(b, m, deterministic)
}
func (m *UserName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserName.Merge(m, src)
}
func (m *UserName) XXX_Size() int {
	return xxx_messageInfo_UserName.Size(m)
}
func (m *UserName) XXX_DiscardUnknown() {
	xxx_messageInfo_UserName.DiscardUnknown(m)
}

var xxx_messageInfo_UserName proto.InternalMessageInfo

func (m *UserName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//处理学校相关的参数
type UserEdu struct {
	EduType              int32    `protobuf:"varint,1,opt,name=EduType,proto3" json:"EduType,omitempty"`
	EduYear              int32    `protobuf:"varint,2,opt,name=EduYear,proto3" json:"EduYear,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserEdu) Reset()         { *m = UserEdu{} }
func (m *UserEdu) String() string { return proto.CompactTextString(m) }
func (*UserEdu) ProtoMessage()    {}
func (*UserEdu) Descriptor() ([]byte, []int) {
	return fileDescriptor_18846f0aec365394, []int{3}
}

func (m *UserEdu) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEdu.Unmarshal(m, b)
}
func (m *UserEdu) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEdu.Marshal(b, m, deterministic)
}
func (m *UserEdu) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEdu.Merge(m, src)
}
func (m *UserEdu) XXX_Size() int {
	return xxx_messageInfo_UserEdu.Size(m)
}
func (m *UserEdu) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEdu.DiscardUnknown(m)
}

var xxx_messageInfo_UserEdu proto.InternalMessageInfo

func (m *UserEdu) GetEduType() int32 {
	if m != nil {
		return m.EduType
	}
	return 0
}

func (m *UserEdu) GetEduYear() int32 {
	if m != nil {
		return m.EduYear
	}
	return 0
}

//根据unionid查询用户
type UserOpenId struct {
	OpenId               string   `protobuf:"bytes,1,opt,name=OpenId,proto3" json:"OpenId,omitempty"`
	UnionId              string   `protobuf:"bytes,2,opt,name=UnionId,proto3" json:"UnionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserOpenId) Reset()         { *m = UserOpenId{} }
func (m *UserOpenId) String() string { return proto.CompactTextString(m) }
func (*UserOpenId) ProtoMessage()    {}
func (*UserOpenId) Descriptor() ([]byte, []int) {
	return fileDescriptor_18846f0aec365394, []int{4}
}

func (m *UserOpenId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserOpenId.Unmarshal(m, b)
}
func (m *UserOpenId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserOpenId.Marshal(b, m, deterministic)
}
func (m *UserOpenId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOpenId.Merge(m, src)
}
func (m *UserOpenId) XXX_Size() int {
	return xxx_messageInfo_UserOpenId.Size(m)
}
func (m *UserOpenId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOpenId.DiscardUnknown(m)
}

var xxx_messageInfo_UserOpenId proto.InternalMessageInfo

func (m *UserOpenId) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *UserOpenId) GetUnionId() string {
	if m != nil {
		return m.UnionId
	}
	return ""
}

func init() {
	proto.RegisterType((*UsersMod)(nil), "pbs.UsersMod")
	proto.RegisterType((*UserId)(nil), "pbs.UserId")
	proto.RegisterType((*UserName)(nil), "pbs.UserName")
	proto.RegisterType((*UserEdu)(nil), "pbs.userEdu")
	proto.RegisterType((*UserOpenId)(nil), "pbs.UserOpenId")
}

func init() { proto.RegisterFile("usermod.proto", fileDescriptor_18846f0aec365394) }

var fileDescriptor_18846f0aec365394 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0x65, 0xa7, 0xce, 0xc7, 0x94, 0x96, 0xb2, 0x94, 0x32, 0x54, 0xa8, 0x8a, 0x7c, 0xca,
	0x89, 0x03, 0x9c, 0x41, 0xaa, 0xd2, 0x48, 0x58, 0xa2, 0x01, 0xb9, 0x04, 0x89, 0xa3, 0x93, 0x9d,
	0x96, 0x15, 0xb6, 0xd7, 0xb2, 0x1d, 0xd4, 0xfc, 0x15, 0x7e, 0x2d, 0x9a, 0x9d, 0xb5, 0x31, 0x07,
	0x4e, 0xd9, 0xe7, 0x7d, 0x77, 0x3e, 0x36, 0x33, 0x86, 0x93, 0x7d, 0x43, 0x75, 0x61, 0xf5, 0x9b,
	0xaa, 0xb6, 0xad, 0x55, 0xa3, 0x6a, 0xdb, 0xc4, 0xbf, 0x23, 0x98, 0x6e, 0x1a, 0xaa, 0x9b, 0x5b,
	0xab, 0xd5, 0x29, 0x84, 0x89, 0xc6, 0x60, 0x1e, 0x2c, 0x46, 0x69, 0x98, 0x68, 0x75, 0x01, 0xe3,
	0x5b, 0x2a, 0xb6, 0x54, 0x63, 0xe8, 0x34, 0x4f, 0xea, 0x12, 0xa6, 0x29, 0x65, 0x79, 0x99, 0x15,
	0x84, 0xa3, 0x79, 0xb0, 0x98, 0xa5, 0x3d, 0x2b, 0x84, 0xc9, 0x47, 0xca, 0xb4, 0x29, 0x1e, 0xf0,
	0xc8, 0x59, 0x1d, 0x72, 0x94, 0x3f, 0xbe, 0xc5, 0x48, 0xa2, 0x3a, 0x76, 0x95, 0xec, 0xd6, 0xe4,
	0x84, 0x63, 0xe7, 0x78, 0x62, 0x3d, 0xb5, 0x39, 0x25, 0x1a, 0x27, 0xf3, 0x60, 0x11, 0xa5, 0x9e,
	0xd4, 0x19, 0x8c, 0x96, 0x46, 0xe3, 0xd4, 0x89, 0x7c, 0x54, 0xe7, 0x10, 0x25, 0xcd, 0x37, 0x53,
	0x21, 0x38, 0x4d, 0x80, 0xe3, 0xef, 0xda, 0xac, 0xdd, 0x37, 0x78, 0x2c, 0xf1, 0x42, 0xdc, 0xe5,
	0x4a, 0xef, 0xbf, 0x1e, 0x2a, 0xc2, 0x27, 0xce, 0xe8, 0xd0, 0x3b, 0xdf, 0x29, 0xab, 0xf1, 0xa4,
	0x77, 0x18, 0xb9, 0xe6, 0xea, 0xb1, 0xc2, 0x53, 0xa9, 0xb9, 0x7a, 0xac, 0xf8, 0xee, 0x27, 0xfb,
	0x60, 0xca, 0xeb, 0x16, 0x9f, 0xca, 0x5b, 0x3d, 0xf2, 0x5b, 0x6f, 0xe8, 0x97, 0xd9, 0x71, 0xe7,
	0x67, 0xf2, 0xd6, 0x8e, 0xd5, 0x15, 0xc0, 0x32, 0x37, 0x54, 0xb6, 0xae, 0xfc, 0x33, 0x97, 0x6e,
	0xa0, 0xa8, 0xd7, 0x30, 0xdb, 0x54, 0x3a, 0x6b, 0x49, 0x5f, 0xb7, 0xa8, 0x5c, 0xf0, 0x5f, 0x81,
	0xdd, 0x65, 0x4d, 0xde, 0x7d, 0x2e, 0x6e, 0x2f, 0xb0, 0x7b, 0x43, 0x39, 0x89, 0x7b, 0x2e, 0x6e,
	0x2f, 0x70, 0x57, 0x77, 0xbb, 0x1f, 0xd6, 0xe6, 0x89, 0xc6, 0x17, 0xae, 0x6e, 0xcf, 0x6a, 0x0e,
	0xc7, 0xcb, 0x3c, 0x6b, 0x9a, 0xda, 0xda, 0x22, 0xd1, 0x78, 0xe1, 0xec, 0xa1, 0xc4, 0x37, 0xbe,
	0xe4, 0x59, 0x4b, 0xf7, 0xb6, 0xe6, 0x1b, 0x2f, 0xe5, 0xc6, 0x40, 0xe2, 0xfc, 0x49, 0xb3, 0xb6,
	0xad, 0xb9, 0x3f, 0x20, 0x4a, 0xfe, 0x8e, 0x79, 0x12, 0x9f, 0x2b, 0x2a, 0x8d, 0xc6, 0x57, 0x32,
	0x61, 0x21, 0x8e, 0x59, 0x9b, 0xdd, 0x4f, 0xb7, 0x4b, 0x97, 0xf2, 0x4f, 0x75, 0x1c, 0x23, 0x8c,
	0x79, 0x37, 0x93, 0xe1, 0x66, 0x46, 0xbc, 0x99, 0xf1, 0x95, 0x6c, 0xed, 0x9a, 0x37, 0x4e, 0xc1,
	0x11, 0xff, 0x3a, 0x77, 0x96, 0xba, 0x73, 0xfc, 0x1e, 0x26, 0xbc, 0xec, 0x2b, 0xbd, 0x1f, 0x8e,
	0x3a, 0xf8, 0xef, 0xa8, 0xc3, 0x7f, 0x46, 0x1d, 0x7f, 0x00, 0xe0, 0xf4, 0xdc, 0xa2, 0x7c, 0x06,
	0x72, 0xf2, 0x25, 0x3c, 0x71, 0xfc, 0xa6, 0x34, 0x96, 0x8d, 0x50, 0xc6, 0xef, 0x71, 0x3b, 0x76,
	0x5f, 0xd8, 0xbb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x14, 0x4e, 0x08, 0x72, 0x03, 0x00,
	0x00,
}
