// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/problem_manage.proto

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

// 获取全部题目信息（只下发 id & title & difficulty & pass_rate）
type GetProblemsReq struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProblemsReq) Reset()         { *m = GetProblemsReq{} }
func (m *GetProblemsReq) String() string { return proto.CompactTextString(m) }
func (*GetProblemsReq) ProtoMessage()    {}
func (*GetProblemsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{0}
}

func (m *GetProblemsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProblemsReq.Unmarshal(m, b)
}
func (m *GetProblemsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProblemsReq.Marshal(b, m, deterministic)
}
func (m *GetProblemsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProblemsReq.Merge(m, src)
}
func (m *GetProblemsReq) XXX_Size() int {
	return xxx_messageInfo_GetProblemsReq.Size(m)
}
func (m *GetProblemsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProblemsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetProblemsReq proto.InternalMessageInfo

func (m *GetProblemsReq) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type GetProblemsResp struct {
	Code                 Code       `protobuf:"varint,1,opt,name=code,proto3,enum=protocol.Code" json:"code,omitempty"`
	Problems             []*Problem `protobuf:"bytes,2,rep,name=problems,proto3" json:"problems,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetProblemsResp) Reset()         { *m = GetProblemsResp{} }
func (m *GetProblemsResp) String() string { return proto.CompactTextString(m) }
func (*GetProblemsResp) ProtoMessage()    {}
func (*GetProblemsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{1}
}

func (m *GetProblemsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProblemsResp.Unmarshal(m, b)
}
func (m *GetProblemsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProblemsResp.Marshal(b, m, deterministic)
}
func (m *GetProblemsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProblemsResp.Merge(m, src)
}
func (m *GetProblemsResp) XXX_Size() int {
	return xxx_messageInfo_GetProblemsResp.Size(m)
}
func (m *GetProblemsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProblemsResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetProblemsResp proto.InternalMessageInfo

func (m *GetProblemsResp) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_OK
}

func (m *GetProblemsResp) GetProblems() []*Problem {
	if m != nil {
		return m.Problems
	}
	return nil
}

// 根据ID获得题目具体信息
type GetProblemByIDReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProblemByIDReq) Reset()         { *m = GetProblemByIDReq{} }
func (m *GetProblemByIDReq) String() string { return proto.CompactTextString(m) }
func (*GetProblemByIDReq) ProtoMessage()    {}
func (*GetProblemByIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{2}
}

func (m *GetProblemByIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProblemByIDReq.Unmarshal(m, b)
}
func (m *GetProblemByIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProblemByIDReq.Marshal(b, m, deterministic)
}
func (m *GetProblemByIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProblemByIDReq.Merge(m, src)
}
func (m *GetProblemByIDReq) XXX_Size() int {
	return xxx_messageInfo_GetProblemByIDReq.Size(m)
}
func (m *GetProblemByIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProblemByIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetProblemByIDReq proto.InternalMessageInfo

func (m *GetProblemByIDReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetProblemByIDResp struct {
	Code                 Code     `protobuf:"varint,1,opt,name=code,proto3,enum=protocol.Code" json:"code,omitempty"`
	Problem              *Problem `protobuf:"bytes,2,opt,name=problem,proto3" json:"problem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProblemByIDResp) Reset()         { *m = GetProblemByIDResp{} }
func (m *GetProblemByIDResp) String() string { return proto.CompactTextString(m) }
func (*GetProblemByIDResp) ProtoMessage()    {}
func (*GetProblemByIDResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{3}
}

func (m *GetProblemByIDResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProblemByIDResp.Unmarshal(m, b)
}
func (m *GetProblemByIDResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProblemByIDResp.Marshal(b, m, deterministic)
}
func (m *GetProblemByIDResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProblemByIDResp.Merge(m, src)
}
func (m *GetProblemByIDResp) XXX_Size() int {
	return xxx_messageInfo_GetProblemByIDResp.Size(m)
}
func (m *GetProblemByIDResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProblemByIDResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetProblemByIDResp proto.InternalMessageInfo

func (m *GetProblemByIDResp) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_OK
}

func (m *GetProblemByIDResp) GetProblem() *Problem {
	if m != nil {
		return m.Problem
	}
	return nil
}

// 新增题目
type AddProblemReq struct {
	Problem              *Problem `protobuf:"bytes,1,opt,name=problem,proto3" json:"problem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddProblemReq) Reset()         { *m = AddProblemReq{} }
func (m *AddProblemReq) String() string { return proto.CompactTextString(m) }
func (*AddProblemReq) ProtoMessage()    {}
func (*AddProblemReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{4}
}

func (m *AddProblemReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProblemReq.Unmarshal(m, b)
}
func (m *AddProblemReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProblemReq.Marshal(b, m, deterministic)
}
func (m *AddProblemReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProblemReq.Merge(m, src)
}
func (m *AddProblemReq) XXX_Size() int {
	return xxx_messageInfo_AddProblemReq.Size(m)
}
func (m *AddProblemReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProblemReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddProblemReq proto.InternalMessageInfo

func (m *AddProblemReq) GetProblem() *Problem {
	if m != nil {
		return m.Problem
	}
	return nil
}

type AddProblemResp struct {
	Code                 Code     `protobuf:"varint,1,opt,name=code,proto3,enum=protocol.Code" json:"code,omitempty"`
	IsSuccess            bool     `protobuf:"varint,2,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddProblemResp) Reset()         { *m = AddProblemResp{} }
func (m *AddProblemResp) String() string { return proto.CompactTextString(m) }
func (*AddProblemResp) ProtoMessage()    {}
func (*AddProblemResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{5}
}

func (m *AddProblemResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProblemResp.Unmarshal(m, b)
}
func (m *AddProblemResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProblemResp.Marshal(b, m, deterministic)
}
func (m *AddProblemResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProblemResp.Merge(m, src)
}
func (m *AddProblemResp) XXX_Size() int {
	return xxx_messageInfo_AddProblemResp.Size(m)
}
func (m *AddProblemResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProblemResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddProblemResp proto.InternalMessageInfo

func (m *AddProblemResp) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_OK
}

func (m *AddProblemResp) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func init() {
	proto.RegisterType((*GetProblemsReq)(nil), "protocol.GetProblemsReq")
	proto.RegisterType((*GetProblemsResp)(nil), "protocol.GetProblemsResp")
	proto.RegisterType((*GetProblemByIDReq)(nil), "protocol.GetProblemByIDReq")
	proto.RegisterType((*GetProblemByIDResp)(nil), "protocol.GetProblemByIDResp")
	proto.RegisterType((*AddProblemReq)(nil), "protocol.AddProblemReq")
	proto.RegisterType((*AddProblemResp)(nil), "protocol.AddProblemResp")
}

func init() { proto.RegisterFile("proto/problem_manage.proto", fileDescriptor_a2e4af9cb4edd521) }

var fileDescriptor_a2e4af9cb4edd521 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8e, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0x22, 0x9a, 0x4e, 0xe9, 0xda, 0xee, 0x29, 0x04, 0x84, 0xb0, 0x5e, 0x0a, 0x62,
	0x84, 0x7a, 0xd4, 0x8b, 0x55, 0x10, 0x6f, 0xb2, 0xbd, 0x79, 0x29, 0x69, 0x76, 0x08, 0x81, 0xa6,
	0x93, 0x76, 0xe2, 0xc1, 0x6f, 0x2f, 0xd9, 0x64, 0x6d, 0x15, 0xa1, 0x3d, 0xed, 0xf2, 0xe6, 0xf7,
	0xfe, 0x40, 0x5c, 0xef, 0xa8, 0xa1, 0xbb, 0x7a, 0x47, 0xab, 0x35, 0x56, 0xcb, 0x2a, 0xdb, 0x64,
	0x05, 0xa6, 0x56, 0x94, 0xa1, 0x7d, 0x72, 0x5a, 0xc7, 0xe3, 0x8e, 0xca, 0xc9, 0xf4, 0xb7, 0x58,
	0x3a, 0xa5, 0xaa, 0x68, 0xd3, 0x69, 0x4a, 0x81, 0x78, 0xc5, 0xe6, 0xbd, 0x8b, 0x62, 0x8d, 0x5b,
	0x39, 0x86, 0xa0, 0xc9, 0x8a, 0xc8, 0x4b, 0xbc, 0xe9, 0x40, 0xb7, 0x5f, 0x65, 0xe0, 0xf2, 0x17,
	0xc3, 0xb5, 0x54, 0x70, 0xd6, 0x06, 0x5b, 0x4a, 0xcc, 0x44, 0xea, 0x5a, 0xd3, 0x67, 0x32, 0xa8,
	0xed, 0x4d, 0xde, 0x42, 0xd8, 0x4f, 0xe4, 0xc8, 0x4f, 0x82, 0xe9, 0x70, 0x36, 0xd9, 0x73, 0x7d,
	0x9a, 0xfe, 0x41, 0xd4, 0x35, 0x4c, 0xf6, 0x2d, 0xf3, 0xaf, 0xb7, 0x97, 0x76, 0x8c, 0x00, 0xbf,
	0x34, 0xb6, 0x25, 0xd0, 0x7e, 0x69, 0x14, 0x82, 0xfc, 0x0b, 0x9d, 0xb8, 0xe6, 0x06, 0x2e, 0xfa,
	0xaa, 0xc8, 0x4f, 0xbc, 0xff, 0xc7, 0x38, 0x42, 0x3d, 0xc2, 0xe8, 0xc9, 0x18, 0x27, 0xe3, 0xf6,
	0xd0, 0xed, 0x1d, 0x75, 0x2f, 0x40, 0x1c, 0xba, 0x4f, 0x1c, 0x78, 0x05, 0x50, 0xf2, 0x92, 0x3f,
	0xf3, 0x1c, 0x99, 0xed, 0xc6, 0x50, 0x0f, 0x4a, 0x5e, 0x74, 0xc2, 0x7c, 0xf4, 0x31, 0x2c, 0xe8,
	0xc1, 0x19, 0x57, 0xe7, 0xf6, 0x77, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xc6, 0x92, 0xc1,
	0x0c, 0x02, 0x00, 0x00,
}