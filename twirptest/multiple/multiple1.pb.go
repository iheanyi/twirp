// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multiple1.proto

// test to make sure that multiple proto files in one package works

package multiple

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

type Msg1 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg1) Reset()         { *m = Msg1{} }
func (m *Msg1) String() string { return proto.CompactTextString(m) }
func (*Msg1) ProtoMessage()    {}
func (*Msg1) Descriptor() ([]byte, []int) {
	return fileDescriptor_78303a5590f73d4e, []int{0}
}

func (m *Msg1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg1.Unmarshal(m, b)
}
func (m *Msg1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg1.Marshal(b, m, deterministic)
}
func (m *Msg1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg1.Merge(m, src)
}
func (m *Msg1) XXX_Size() int {
	return xxx_messageInfo_Msg1.Size(m)
}
func (m *Msg1) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg1.DiscardUnknown(m)
}

var xxx_messageInfo_Msg1 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Msg1)(nil), "twirp.twirptest.multiple.Msg1")
}

func init() { proto.RegisterFile("multiple1.proto", fileDescriptor_78303a5590f73d4e) }

var fileDescriptor_78303a5590f73d4e = []byte{
	// 104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x2d, 0xcd, 0x29,
	0xc9, 0x2c, 0xc8, 0x49, 0x35, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x28, 0x29, 0xcf,
	0x2c, 0x2a, 0xd0, 0x03, 0x93, 0x25, 0xa9, 0xc5, 0x25, 0x7a, 0x30, 0x05, 0x4a, 0x6c, 0x5c, 0x2c,
	0xbe, 0xc5, 0xe9, 0x86, 0x46, 0x7e, 0x5c, 0x2c, 0xc1, 0x65, 0xc9, 0x86, 0x42, 0x6e, 0x5c, 0x2c,
	0xc1, 0xa9, 0x79, 0x29, 0x42, 0x72, 0x7a, 0xb8, 0xb4, 0xe8, 0x81, 0xd4, 0x4b, 0x11, 0x90, 0x77,
	0xe2, 0x8a, 0xe2, 0x80, 0x09, 0x24, 0xb1, 0x81, 0x1d, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x01, 0xb7, 0x59, 0x2f, 0x97, 0x00, 0x00, 0x00,
}
