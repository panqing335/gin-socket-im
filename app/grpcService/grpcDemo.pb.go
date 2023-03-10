// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/grpcDemo.proto

package grpcService

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DemoRequest struct {
	Json                 string   `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DemoRequest) Reset()         { *m = DemoRequest{} }
func (m *DemoRequest) String() string { return proto.CompactTextString(m) }
func (*DemoRequest) ProtoMessage()    {}
func (*DemoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_779dfbe0e9c70f7e, []int{0}
}
func (m *DemoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DemoRequest.Unmarshal(m, b)
}
func (m *DemoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DemoRequest.Marshal(b, m, deterministic)
}
func (m *DemoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemoRequest.Merge(m, src)
}
func (m *DemoRequest) XXX_Size() int {
	return xxx_messageInfo_DemoRequest.Size(m)
}
func (m *DemoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DemoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DemoRequest proto.InternalMessageInfo

func (m *DemoRequest) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

type DemoReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DemoReply) Reset()         { *m = DemoReply{} }
func (m *DemoReply) String() string { return proto.CompactTextString(m) }
func (*DemoReply) ProtoMessage()    {}
func (*DemoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_779dfbe0e9c70f7e, []int{1}
}
func (m *DemoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DemoReply.Unmarshal(m, b)
}
func (m *DemoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DemoReply.Marshal(b, m, deterministic)
}
func (m *DemoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemoReply.Merge(m, src)
}
func (m *DemoReply) XXX_Size() int {
	return xxx_messageInfo_DemoReply.Size(m)
}
func (m *DemoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DemoReply.DiscardUnknown(m)
}

var xxx_messageInfo_DemoReply proto.InternalMessageInfo

func (m *DemoReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*DemoRequest)(nil), "grpcService.DemoRequest")
	proto.RegisterType((*DemoReply)(nil), "grpcService.DemoReply")
}

func init() { proto.RegisterFile("protos/grpcDemo.proto", fileDescriptor_779dfbe0e9c70f7e) }

var fileDescriptor_779dfbe0e9c70f7e = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x2f, 0x2a, 0x48, 0x76, 0x49, 0xcd, 0xcd, 0xd7, 0x03, 0xf3, 0x85, 0xb8,
	0x41, 0xfc, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x25, 0x45, 0x2e, 0x6e, 0x90, 0x54, 0x50,
	0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x56, 0x71, 0x7e, 0x9e, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0xa4, 0xca, 0xc5, 0x09, 0x51, 0x52, 0x90, 0x53, 0x29,
	0x24, 0xc1, 0xc5, 0x9e, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x0a, 0x55, 0x03, 0xe3, 0x1a, 0xb9,
	0x72, 0xb1, 0x80, 0x94, 0x09, 0xd9, 0x72, 0x71, 0x86, 0xe6, 0x25, 0x16, 0x55, 0x3a, 0x27, 0xe6,
	0xe4, 0x08, 0x49, 0xe8, 0x21, 0x59, 0xa6, 0x87, 0x64, 0x93, 0x94, 0x18, 0x16, 0x99, 0x82, 0x9c,
	0x4a, 0x27, 0xc5, 0x28, 0x79, 0x3d, 0xfd, 0xc4, 0x82, 0x02, 0x7d, 0x24, 0x69, 0x7d, 0x6b, 0x05,
	0x24, 0x5e, 0x12, 0x1b, 0xd8, 0x1f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x10, 0xb9,
	0xe9, 0xe0, 0x00, 0x00, 0x00,
}
