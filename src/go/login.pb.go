// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/login.proto

package protocol

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

// 登陆 (post)
type LoginReq struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b26af552c9148a54, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResp struct {
	Code                 Code            `protobuf:"varint,1,opt,name=code,proto3,enum=protocol.Code" json:"code,omitempty"`
	Token                string          `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	User                 *UserInfo       `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	SubmitRecords        []*SubmitRecord `protobuf:"bytes,4,rep,name=submit_records,json=submitRecords,proto3" json:"submit_records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b26af552c9148a54, []int{1}
}

func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_OK
}

func (m *LoginResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResp) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *LoginResp) GetSubmitRecords() []*SubmitRecord {
	if m != nil {
		return m.SubmitRecords
	}
	return nil
}

// 登出 (get)
type LogOut struct {
	Code                 Code     `protobuf:"varint,1,opt,name=code,proto3,enum=protocol.Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogOut) Reset()         { *m = LogOut{} }
func (m *LogOut) String() string { return proto.CompactTextString(m) }
func (*LogOut) ProtoMessage()    {}
func (*LogOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_b26af552c9148a54, []int{2}
}

func (m *LogOut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogOut.Unmarshal(m, b)
}
func (m *LogOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogOut.Marshal(b, m, deterministic)
}
func (m *LogOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogOut.Merge(m, src)
}
func (m *LogOut) XXX_Size() int {
	return xxx_messageInfo_LogOut.Size(m)
}
func (m *LogOut) XXX_DiscardUnknown() {
	xxx_messageInfo_LogOut.DiscardUnknown(m)
}

var xxx_messageInfo_LogOut proto.InternalMessageInfo

func (m *LogOut) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_OK
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "protocol.LoginReq")
	proto.RegisterType((*LoginResp)(nil), "protocol.LoginResp")
	proto.RegisterType((*LogOut)(nil), "protocol.LogOut")
}

func init() { proto.RegisterFile("proto/login.proto", fileDescriptor_b26af552c9148a54) }

var fileDescriptor_b26af552c9148a54 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8e, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x8d, 0x35, 0x9d, 0xd0, 0xa0, 0x8b, 0xc8, 0x92, 0x53, 0xc8, 0x41, 0x72, 0x90,
	0x08, 0xf1, 0x28, 0x82, 0xe8, 0x49, 0x28, 0x08, 0x2b, 0x5e, 0xbc, 0x48, 0xbb, 0x59, 0x43, 0xb1,
	0xc9, 0xc4, 0x9d, 0x0d, 0xbe, 0x94, 0x0f, 0x29, 0x9d, 0x24, 0xc6, 0x63, 0x4f, 0x3b, 0xdf, 0x7e,
	0xfb, 0xff, 0x3b, 0x70, 0xd6, 0x5a, 0x74, 0x78, 0xbd, 0xc3, 0x6a, 0xdb, 0xe4, 0x3c, 0x8b, 0x80,
	0x0f, 0x8d, 0xbb, 0xf8, 0xb4, 0x97, 0x1a, 0x4b, 0xd3, 0xbb, 0x58, 0x8c, 0x37, 0x75, 0x8d, 0xc3,
	0xfb, 0xf4, 0x1e, 0x82, 0xd5, 0x3e, 0xae, 0xcc, 0x97, 0x90, 0x70, 0xb2, 0xd6, 0x1a, 0xbb, 0xc6,
	0x49, 0x2f, 0xf1, 0xb2, 0x85, 0x1a, 0x51, 0xc4, 0x10, 0xb4, 0x6b, 0xa2, 0x6f, 0xb4, 0xa5, 0x3c,
	0x62, 0xf5, 0xc7, 0xe9, 0x8f, 0x07, 0x8b, 0xa1, 0x82, 0x5a, 0x91, 0x82, 0xbf, 0xff, 0x91, 0x0b,
	0xa2, 0x22, 0xca, 0xc7, 0x75, 0xf2, 0x47, 0x2c, 0x8d, 0x62, 0x27, 0xce, 0xe1, 0xd8, 0xe1, 0xa7,
	0x69, 0x86, 0xaa, 0x1e, 0xc4, 0x25, 0xf8, 0x1d, 0x19, 0x2b, 0x67, 0x89, 0x97, 0x85, 0x85, 0x98,
	0x92, 0xaf, 0x64, 0xec, 0x53, 0xf3, 0x81, 0x8a, 0xbd, 0xb8, 0x83, 0x88, 0xba, 0x4d, 0xbd, 0x75,
	0xef, 0xd6, 0x68, 0xb4, 0x25, 0x49, 0x3f, 0x99, 0x65, 0x61, 0x71, 0x31, 0x25, 0x5e, 0xd8, 0x2b,
	0xd6, 0x6a, 0x49, 0xff, 0x88, 0xd2, 0x2b, 0x98, 0xaf, 0xb0, 0x7a, 0xee, 0xdc, 0x21, 0xab, 0x3e,
	0x2c, 0xdf, 0xc2, 0x0a, 0x6f, 0x47, 0xb3, 0x99, 0xf3, 0x74, 0xf3, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0xef, 0x0d, 0x9c, 0x49, 0x79, 0x01, 0x00, 0x00,
}
