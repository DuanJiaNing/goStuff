// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto3.proto

package example

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Person3 struct {
	Value                float32               `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	Value2               *wrappers.FloatValue  `protobuf:"bytes,2,opt,name=value2,proto3" json:"value2,omitempty"`
	StrValue             string                `protobuf:"bytes,3,opt,name=str_value,json=strValue,proto3" json:"str_value,omitempty"`
	StrValue2            *wrappers.StringValue `protobuf:"bytes,4,opt,name=str_value2,json=strValue2,proto3" json:"str_value2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person3) Reset()         { *m = Person3{} }
func (m *Person3) String() string { return proto.CompactTextString(m) }
func (*Person3) ProtoMessage()    {}
func (*Person3) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fee6d65e34a64b6, []int{0}
}

func (m *Person3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person3.Unmarshal(m, b)
}
func (m *Person3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person3.Marshal(b, m, deterministic)
}
func (m *Person3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person3.Merge(m, src)
}
func (m *Person3) XXX_Size() int {
	return xxx_messageInfo_Person3.Size(m)
}
func (m *Person3) XXX_DiscardUnknown() {
	xxx_messageInfo_Person3.DiscardUnknown(m)
}

var xxx_messageInfo_Person3 proto.InternalMessageInfo

func (m *Person3) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Person3) GetValue2() *wrappers.FloatValue {
	if m != nil {
		return m.Value2
	}
	return nil
}

func (m *Person3) GetStrValue() string {
	if m != nil {
		return m.StrValue
	}
	return ""
}

func (m *Person3) GetStrValue2() *wrappers.StringValue {
	if m != nil {
		return m.StrValue2
	}
	return nil
}

// test3: p.value=1, p.value2=nil               -> json: {"p":{"value":1}}
// test4: p.value=0, p.value2.value=0           -> json: {"p":{"value2":0}}
// test5: p.value=1, p.value2.value=0           -> json: {"p":{"value":1,"value2":0}}
// test6: p.str_value="", p.str_value2.value="" -> {"p":{"strValue2":""}}
type Data3 struct {
	P                    *Person3 `protobuf:"bytes,1,opt,name=p,proto3" json:"p,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data3) Reset()         { *m = Data3{} }
func (m *Data3) String() string { return proto.CompactTextString(m) }
func (*Data3) ProtoMessage()    {}
func (*Data3) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fee6d65e34a64b6, []int{1}
}

func (m *Data3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data3.Unmarshal(m, b)
}
func (m *Data3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data3.Marshal(b, m, deterministic)
}
func (m *Data3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data3.Merge(m, src)
}
func (m *Data3) XXX_Size() int {
	return xxx_messageInfo_Data3.Size(m)
}
func (m *Data3) XXX_DiscardUnknown() {
	xxx_messageInfo_Data3.DiscardUnknown(m)
}

var xxx_messageInfo_Data3 proto.InternalMessageInfo

func (m *Data3) GetP() *Person3 {
	if m != nil {
		return m.P
	}
	return nil
}

func init() {
	proto.RegisterType((*Person3)(nil), "example.Person3")
	proto.RegisterType((*Data3)(nil), "example.Data3")
}

func init() { proto.RegisterFile("proto3.proto", fileDescriptor_4fee6d65e34a64b6) }

var fileDescriptor_4fee6d65e34a64b6 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0xd6, 0x03, 0x53, 0x42, 0xec, 0xa9, 0x15, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0x52, 0x72,
	0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0xe1, 0xa4, 0xd2, 0x34, 0xfd, 0xf2, 0xa2, 0xc4,
	0x82, 0x82, 0xd4, 0xa2, 0x62, 0x88, 0x42, 0xa5, 0x75, 0x8c, 0x5c, 0xec, 0x01, 0xa9, 0x45, 0xc5,
	0xf9, 0x79, 0xc6, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x4c, 0x41, 0x10, 0x8e, 0x90, 0x31, 0x17, 0x1b, 0x98, 0x61, 0x24, 0xc1, 0xa4, 0xc0, 0xa8,
	0xc1, 0x6d, 0x24, 0xad, 0x07, 0x31, 0x52, 0x0f, 0x66, 0xa4, 0x9e, 0x5b, 0x4e, 0x7e, 0x62, 0x49,
	0x18, 0x48, 0x4d, 0x10, 0x54, 0xa9, 0x90, 0x34, 0x17, 0x67, 0x71, 0x49, 0x51, 0x3c, 0xc4, 0x38,
	0x66, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x8e, 0xe2, 0x92, 0x22, 0xb0, 0x22, 0x21, 0x6b, 0x2e, 0x2e,
	0xb8, 0xa4, 0x91, 0x04, 0x0b, 0xd8, 0x54, 0x19, 0x0c, 0x53, 0x83, 0x4b, 0x8a, 0x32, 0xf3, 0xd2,
	0x21, 0xc6, 0x72, 0xc2, 0xf4, 0x1a, 0x29, 0xa9, 0x73, 0xb1, 0xba, 0x24, 0x96, 0x24, 0x1a, 0x0b,
	0xc9, 0x71, 0x31, 0x16, 0x80, 0x5d, 0xca, 0x6d, 0x24, 0xa0, 0x07, 0xf5, 0xae, 0x1e, 0xd4, 0x2b,
	0x41, 0x8c, 0x05, 0x49, 0x6c, 0x90, 0x00, 0x01, 0x04, 0x00, 0x00, 0xff, 0xff, 0x29, 0x47, 0x05,
	0xc0, 0x19, 0x01, 0x00, 0x00,
}
