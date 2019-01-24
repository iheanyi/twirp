// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto.proto

// Test to make sure that a package named proto doesn't break

package proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Msg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Msg)(nil), "twirp.twirptest.proto.Msg")
}

func init() { proto.RegisterFile("proto.proto", fileDescriptor_2fcc84b9998d60d8) }

var fileDescriptor_2fcc84b9998d60d8 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x03, 0x93, 0x42, 0xa2, 0x25, 0xe5, 0x99, 0x45, 0x05, 0x7a, 0x60, 0xb2, 0x24, 0xb5,
	0xb8, 0x04, 0x22, 0xac, 0xc4, 0xca, 0xc5, 0xec, 0x5b, 0x9c, 0x6e, 0xe4, 0xca, 0xc5, 0x1c, 0x5c,
	0x96, 0x2c, 0x64, 0xc7, 0xc5, 0x12, 0x9c, 0x9a, 0x97, 0x22, 0x24, 0xa5, 0x87, 0x55, 0xb5, 0x9e,
	0x6f, 0x71, 0xba, 0x14, 0x1e, 0x39, 0x27, 0xf6, 0x28, 0x56, 0x30, 0x27, 0x89, 0x0d, 0x4c, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x37, 0xae, 0xb5, 0x83, 0x00, 0x00, 0x00,
}
