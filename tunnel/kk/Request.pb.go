// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Request.proto

package kk

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Request struct {
	// *
	// 请求ID
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// *
	// URI
	Uri string `protobuf:"bytes,2,opt,name=uri" json:"uri,omitempty"`
	// *
	// 数据类型
	Type string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	// *
	// 数据
	// @type {bytes}
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// *
	// 头
	Header map[string]string `protobuf:"bytes,5,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Request) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Request) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Request) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Request) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "kk.Request")
}

func init() { proto.RegisterFile("Request.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0d, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xca, 0xce, 0x56, 0xda, 0xcf,
	0xc8, 0xc5, 0x0e, 0x15, 0x15, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x0e, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe0, 0x62, 0x2e, 0x2d, 0xca, 0x94, 0x60, 0x52, 0x60, 0xd4,
	0xe0, 0x0c, 0x02, 0x31, 0x85, 0x84, 0xb8, 0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x98, 0xc1, 0x42,
	0x60, 0x36, 0x48, 0x2c, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x45, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc,
	0x16, 0xd2, 0xe7, 0x62, 0xcb, 0x48, 0x4d, 0x4c, 0x49, 0x2d, 0x92, 0x60, 0x55, 0x60, 0xd6, 0xe0,
	0x36, 0x12, 0xd7, 0xcb, 0xce, 0xd6, 0x83, 0x59, 0xee, 0x01, 0x96, 0x71, 0xcd, 0x2b, 0x29, 0xaa,
	0x0c, 0x82, 0x2a, 0x93, 0xb2, 0xe4, 0xe2, 0x46, 0x12, 0x06, 0xd9, 0x9c, 0x9d, 0x5a, 0x09, 0x76,
	0x0a, 0x67, 0x10, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x0a, 0x75, 0x0d,
	0x84, 0x63, 0xc5, 0x64, 0xc1, 0x98, 0xc4, 0x06, 0xf6, 0x8c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xff, 0xaf, 0x51, 0x48, 0xdd, 0x00, 0x00, 0x00,
}
