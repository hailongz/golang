// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kk.Response.proto

package kk

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Response struct {
	// *
	// 跟踪ID
	Trace string `protobuf:"bytes,1,opt,name=trace" json:"trace,omitempty"`
	// *
	// 类型
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// *
	// 数据
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Response) GetTrace() string {
	if m != nil {
		return m.Trace
	}
	return ""
}

func (m *Response) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Response) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*Response)(nil), "kk.Response")
}

func init() { proto.RegisterFile("kk.Response.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0xce, 0xd6, 0x0b,
	0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xca,
	0xce, 0x56, 0xf2, 0xe3, 0xe2, 0x80, 0x89, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0x14, 0x25, 0x26, 0xa7,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x42, 0x5c, 0x2c, 0x25, 0x95, 0x05,
	0xa9, 0x12, 0x4c, 0x60, 0x41, 0x30, 0x5b, 0x48, 0x82, 0x8b, 0x3d, 0x39, 0x3f, 0xaf, 0x24, 0x35,
	0xaf, 0x44, 0x82, 0x59, 0x81, 0x51, 0x83, 0x27, 0x08, 0xc6, 0x4d, 0x62, 0x03, 0x1b, 0x6d, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x0b, 0x8d, 0xe4, 0x6f, 0x00, 0x00, 0x00,
}
