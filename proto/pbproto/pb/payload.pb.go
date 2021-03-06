// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payload.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Payload struct {
	Seq                  int32    `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Mtype                int32    `protobuf:"varint,2,opt,name=mtype,proto3" json:"mtype,omitempty"`
	ServiceMethod        string   `protobuf:"bytes,3,opt,name=serviceMethod,proto3" json:"serviceMethod,omitempty"`
	Status               []byte   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Meta                 []byte   `protobuf:"bytes,5,opt,name=meta,proto3" json:"meta,omitempty"`
	BodyCodec            int32    `protobuf:"varint,6,opt,name=bodyCodec,proto3" json:"bodyCodec,omitempty"`
	Body                 []byte   `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_payload_05c44227d6505218, []int{0}
}
func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (dst *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(dst, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Payload) GetMtype() int32 {
	if m != nil {
		return m.Mtype
	}
	return 0
}

func (m *Payload) GetServiceMethod() string {
	if m != nil {
		return m.ServiceMethod
	}
	return ""
}

func (m *Payload) GetStatus() []byte {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Payload) GetMeta() []byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Payload) GetBodyCodec() int32 {
	if m != nil {
		return m.BodyCodec
	}
	return 0
}

func (m *Payload) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Payload)(nil), "pb.payload")
}

func init() { proto.RegisterFile("payload.proto", fileDescriptor_payload_05c44227d6505218) }

var fileDescriptor_payload_05c44227d6505218 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0xac, 0xcc,
	0xc9, 0x4f, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xda, 0xca,
	0xc8, 0xc5, 0x0e, 0x15, 0x15, 0x12, 0xe0, 0x62, 0x2e, 0x4e, 0x2d, 0x94, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x0d, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0x73, 0x4b, 0x2a, 0x0b, 0x52, 0x25, 0x98, 0xc0,
	0x62, 0x10, 0x8e, 0x90, 0x0a, 0x17, 0x6f, 0x71, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x6f, 0x6a,
	0x49, 0x46, 0x7e, 0x8a, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0xaa, 0xa0, 0x90, 0x18, 0x17,
	0x5b, 0x71, 0x49, 0x62, 0x49, 0x69, 0xb1, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x94, 0x27,
	0x24, 0xc4, 0xc5, 0x92, 0x9b, 0x5a, 0x92, 0x28, 0xc1, 0x0a, 0x16, 0x05, 0xb3, 0x85, 0x64, 0xb8,
	0x38, 0x93, 0xf2, 0x53, 0x2a, 0x9d, 0xf3, 0x53, 0x52, 0x93, 0x25, 0xd8, 0xc0, 0x76, 0x21, 0x04,
	0x40, 0x3a, 0x40, 0x1c, 0x09, 0x76, 0x88, 0x0e, 0x10, 0x3b, 0x89, 0x0d, 0xec, 0x05, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0x7a, 0x4c, 0x10, 0xd3, 0x00, 0x00, 0x00,
}
