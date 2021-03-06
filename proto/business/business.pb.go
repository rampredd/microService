// Code generated by protoc-gen-go. DO NOT EDIT.
// source: business.proto

package business

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import query "microService/proto/query"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BusinessRequest struct {
	Req                  *query.Request `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BusinessRequest) Reset()         { *m = BusinessRequest{} }
func (m *BusinessRequest) String() string { return proto.CompactTextString(m) }
func (*BusinessRequest) ProtoMessage()    {}
func (*BusinessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_business_ef9796052d467bbe, []int{0}
}
func (m *BusinessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BusinessRequest.Unmarshal(m, b)
}
func (m *BusinessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BusinessRequest.Marshal(b, m, deterministic)
}
func (dst *BusinessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BusinessRequest.Merge(dst, src)
}
func (m *BusinessRequest) XXX_Size() int {
	return xxx_messageInfo_BusinessRequest.Size(m)
}
func (m *BusinessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BusinessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BusinessRequest proto.InternalMessageInfo

func (m *BusinessRequest) GetReq() *query.Request {
	if m != nil {
		return m.Req
	}
	return nil
}

type BusinessResponse struct {
	Rsp                  *query.Response `protobuf:"bytes,1,opt,name=rsp,proto3" json:"rsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BusinessResponse) Reset()         { *m = BusinessResponse{} }
func (m *BusinessResponse) String() string { return proto.CompactTextString(m) }
func (*BusinessResponse) ProtoMessage()    {}
func (*BusinessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_business_ef9796052d467bbe, []int{1}
}
func (m *BusinessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BusinessResponse.Unmarshal(m, b)
}
func (m *BusinessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BusinessResponse.Marshal(b, m, deterministic)
}
func (dst *BusinessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BusinessResponse.Merge(dst, src)
}
func (m *BusinessResponse) XXX_Size() int {
	return xxx_messageInfo_BusinessResponse.Size(m)
}
func (m *BusinessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BusinessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BusinessResponse proto.InternalMessageInfo

func (m *BusinessResponse) GetRsp() *query.Response {
	if m != nil {
		return m.Rsp
	}
	return nil
}

func init() {
	proto.RegisterType((*BusinessRequest)(nil), "BusinessRequest")
	proto.RegisterType((*BusinessResponse)(nil), "BusinessResponse")
}

func init() { proto.RegisterFile("business.proto", fileDescriptor_business_ef9796052d467bbe) }

var fileDescriptor_business_ef9796052d467bbe = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2a, 0x2d, 0xce,
	0xcc, 0x4b, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xd2, 0xc9, 0xcd, 0x4c, 0x2e,
	0xca, 0x0f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x75, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x07,
	0xcb, 0xe8, 0x17, 0x96, 0xa6, 0x16, 0x55, 0xea, 0x17, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x40,
	0x55, 0xeb, 0x12, 0x54, 0x5d, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x0a, 0x51, 0xae, 0xe4, 0xc3, 0xc5,
	0xef, 0x04, 0xb5, 0x2e, 0x08, 0x62, 0x8e, 0x90, 0x25, 0x17, 0x73, 0x51, 0x6a, 0xa1, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0xb7, 0x91, 0xba, 0x1e, 0xa6, 0x79, 0x10, 0xad, 0x7a, 0x60, 0xf3, 0xf4, 0xa0,
	0xba, 0x82, 0x40, 0x7a, 0x94, 0xfc, 0xb8, 0x04, 0x10, 0xa6, 0x41, 0xec, 0x11, 0xb2, 0xe2, 0x62,
	0x2e, 0x2a, 0x2e, 0x80, 0x1a, 0xa7, 0x41, 0xd8, 0x38, 0x88, 0xb6, 0x20, 0x90, 0x26, 0x23, 0x6b,
	0x2e, 0x0e, 0x98, 0x79, 0x42, 0xfa, 0x5c, 0x6c, 0xc1, 0xa9, 0x89, 0x45, 0xc9, 0x19, 0x42, 0x02,
	0x7a, 0x68, 0x4e, 0x96, 0x12, 0xd4, 0x43, 0xb7, 0x56, 0x89, 0x21, 0x89, 0x0d, 0x6c, 0xae, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x46, 0x8c, 0x02, 0x4f, 0x50, 0x01, 0x00, 0x00,
}
