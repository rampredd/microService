// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

package gateway

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

type QueryRequest struct {
	Req                  *query.Request `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway_ceed824dd376ca4c, []int{0}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (dst *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(dst, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetReq() *query.Request {
	if m != nil {
		return m.Req
	}
	return nil
}

type QueryResponse struct {
	Rsp                  *query.Response `protobuf:"bytes,1,opt,name=rsp,proto3" json:"rsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway_ceed824dd376ca4c, []int{1}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (dst *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(dst, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetRsp() *query.Response {
	if m != nil {
		return m.Rsp
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "QueryResponse")
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor_gateway_ceed824dd376ca4c) }

var fileDescriptor_gateway_ceed824dd376ca4c = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4f, 0x2c, 0x49,
	0x2d, 0x4f, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xd2, 0xc9, 0xcd, 0x4c, 0x2e, 0xca,
	0x0f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x75, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x07, 0xcb,
	0xe8, 0x17, 0x96, 0xa6, 0x16, 0x55, 0xea, 0x17, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x40, 0x55,
	0xeb, 0x12, 0x54, 0x5d, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x0a, 0x51, 0xae, 0xe4, 0xc9, 0xc5, 0x13,
	0x08, 0x12, 0x0f, 0x82, 0x18, 0x22, 0x64, 0xc9, 0xc5, 0x5c, 0x94, 0x5a, 0x28, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x6d, 0xa4, 0xae, 0x87, 0x69, 0x18, 0x44, 0x9f, 0x1e, 0xd8, 0x30, 0x3d, 0xa8, 0xae,
	0x20, 0x90, 0x1e, 0x25, 0x6f, 0x2e, 0x5e, 0xa8, 0x51, 0x10, 0x1b, 0x84, 0xac, 0xb8, 0x98, 0x8b,
	0x8a, 0x0b, 0xa0, 0x66, 0x69, 0x10, 0x36, 0x0b, 0xa2, 0x2d, 0x08, 0xa4, 0xc9, 0xc8, 0x84, 0x8b,
	0xdd, 0x1d, 0x12, 0x0a, 0x42, 0x9a, 0x5c, 0x6c, 0xc1, 0xa9, 0x89, 0x45, 0xc9, 0x19, 0x42, 0xbc,
	0x7a, 0xc8, 0x6e, 0x95, 0xe2, 0xd3, 0x43, 0xb1, 0x4f, 0x89, 0x21, 0x89, 0x0d, 0x6c, 0xa0, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x78, 0x16, 0x7d, 0x7a, 0x42, 0x01, 0x00, 0x00,
}
