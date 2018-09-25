// Code generated by protoc-gen-go. DO NOT EDIT.
// source: environment.proto

package environment

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

type GeneralRequest struct {
	VariableType         string   `protobuf:"bytes,1,opt,name=variableType,proto3" json:"variableType,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	VariableName         string   `protobuf:"bytes,3,opt,name=variableName,proto3" json:"variableName,omitempty"`
	ReplaceExisting      bool     `protobuf:"varint,4,opt,name=replaceExisting,proto3" json:"replaceExisting,omitempty"`
	Path                 string   `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralRequest) Reset()         { *m = GeneralRequest{} }
func (m *GeneralRequest) String() string { return proto.CompactTextString(m) }
func (*GeneralRequest) ProtoMessage()    {}
func (*GeneralRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{0}
}

func (m *GeneralRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralRequest.Unmarshal(m, b)
}
func (m *GeneralRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralRequest.Marshal(b, m, deterministic)
}
func (m *GeneralRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralRequest.Merge(m, src)
}
func (m *GeneralRequest) XXX_Size() int {
	return xxx_messageInfo_GeneralRequest.Size(m)
}
func (m *GeneralRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralRequest proto.InternalMessageInfo

func (m *GeneralRequest) GetVariableType() string {
	if m != nil {
		return m.VariableType
	}
	return ""
}

func (m *GeneralRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *GeneralRequest) GetVariableName() string {
	if m != nil {
		return m.VariableName
	}
	return ""
}

func (m *GeneralRequest) GetReplaceExisting() bool {
	if m != nil {
		return m.ReplaceExisting
	}
	return false
}

func (m *GeneralRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type GeneralResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralResponse) Reset()         { *m = GeneralResponse{} }
func (m *GeneralResponse) String() string { return proto.CompactTextString(m) }
func (*GeneralResponse) ProtoMessage()    {}
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{1}
}

func (m *GeneralResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralResponse.Unmarshal(m, b)
}
func (m *GeneralResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralResponse.Marshal(b, m, deterministic)
}
func (m *GeneralResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralResponse.Merge(m, src)
}
func (m *GeneralResponse) XXX_Size() int {
	return xxx_messageInfo_GeneralResponse.Size(m)
}
func (m *GeneralResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralResponse proto.InternalMessageInfo

func (m *GeneralResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GeneralRequest)(nil), "environment.GeneralRequest")
	proto.RegisterType((*GeneralResponse)(nil), "environment.GeneralResponse")
}

func init() { proto.RegisterFile("environment.proto", fileDescriptor_64e647b85623514a) }

var fileDescriptor_64e647b85623514a = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xcd, 0x2b, 0xcb,
	0x2c, 0xca, 0xcf, 0xcb, 0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x52, 0x5a, 0xc3, 0xc8, 0xc5, 0xe7, 0x9e, 0x9a, 0x97, 0x5a, 0x94, 0x98, 0x13, 0x94, 0x5a,
	0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xa4, 0xc4, 0xc5, 0x53, 0x96, 0x58, 0x94, 0x99, 0x98, 0x94, 0x93,
	0x1a, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x22, 0x26, 0x24, 0xc2,
	0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0x04, 0x96, 0x84, 0x70, 0x90, 0x75, 0xfa, 0x25,
	0xe6, 0xa6, 0x4a, 0x30, 0xa3, 0xea, 0x04, 0x89, 0x09, 0x69, 0x70, 0xf1, 0x17, 0xa5, 0x16, 0xe4,
	0x24, 0x26, 0xa7, 0xba, 0x56, 0x64, 0x16, 0x97, 0x64, 0xe6, 0xa5, 0x4b, 0xb0, 0x28, 0x30, 0x6a,
	0x70, 0x04, 0xa1, 0x0b, 0x0b, 0x09, 0x71, 0xb1, 0x14, 0x24, 0x96, 0x64, 0x48, 0xb0, 0x82, 0x4d,
	0x01, 0xb3, 0x95, 0xb4, 0xb9, 0xf8, 0xe1, 0xae, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x92,
	0xe0, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x87, 0xb9, 0x14, 0xc6, 0x35, 0x8a, 0xe1, 0xe2,
	0x76, 0x45, 0x78, 0x55, 0xc8, 0x97, 0x8b, 0xcf, 0x2f, 0xb5, 0x1c, 0x59, 0x44, 0x5a, 0x0f, 0x39,
	0x74, 0x50, 0x83, 0x41, 0x4a, 0x06, 0xbb, 0x24, 0xc4, 0x56, 0x25, 0x86, 0x24, 0x36, 0x70, 0x68,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x21, 0xa9, 0x7b, 0x24, 0x62, 0x01, 0x00, 0x00,
}
