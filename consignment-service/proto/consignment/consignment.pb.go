// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package consignment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_7dbbcf7507b410e8, []int{0}
}
func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (dst *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(dst, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_7dbbcf7507b410e8, []int{1}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (dst *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(dst, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_7dbbcf7507b410e8, []int{2}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_7dbbcf7507b410e8, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "consignment.Consignment")
	proto.RegisterType((*Container)(nil), "consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "consignment.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ConsignmentService service

type ConsignmentServiceClient interface {
	Create(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type consignmentServiceClient struct {
	c           client.Client
	serviceName string
}

func NewConsignmentServiceClient(serviceName string, c client.Client) ConsignmentServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "consignment"
	}
	return &consignmentServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *consignmentServiceClient) Create(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ConsignmentService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consignmentServiceClient) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ConsignmentService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConsignmentService service

type ConsignmentServiceHandler interface {
	Create(context.Context, *Consignment, *Response) error
	Get(context.Context, *GetRequest, *Response) error
}

func RegisterConsignmentServiceHandler(s server.Server, hdlr ConsignmentServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ConsignmentService{hdlr}, opts...))
}

type ConsignmentService struct {
	ConsignmentServiceHandler
}

func (h *ConsignmentService) Create(ctx context.Context, in *Consignment, out *Response) error {
	return h.ConsignmentServiceHandler.Create(ctx, in, out)
}

func (h *ConsignmentService) Get(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ConsignmentServiceHandler.Get(ctx, in, out)
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_consignment_7dbbcf7507b410e8)
}

var fileDescriptor_consignment_7dbbcf7507b410e8 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x4e, 0x02, 0x31,
	0x14, 0x85, 0x1d, 0x86, 0xbf, 0xb9, 0x43, 0x5c, 0xdc, 0x44, 0x68, 0x74, 0xe1, 0x64, 0xdc, 0xb0,
	0xc2, 0x04, 0x13, 0x4d, 0xd4, 0x1d, 0x0b, 0xc2, 0xb6, 0x3e, 0x80, 0xc1, 0xf6, 0x06, 0x9b, 0x48,
	0x8b, 0x6d, 0xc1, 0x77, 0xf0, 0x21, 0x7c, 0x02, 0x1f, 0xd2, 0xd0, 0x61, 0xa4, 0x68, 0xd8, 0xf5,
	0x9c, 0x73, 0xcf, 0xcc, 0xd7, 0x9b, 0xc2, 0xd5, 0xca, 0x1a, 0x6f, 0xae, 0x85, 0xd1, 0x4e, 0x2d,
	0xf4, 0x92, 0xb4, 0x8f, 0xcf, 0xa3, 0x90, 0x62, 0x1e, 0x59, 0xe5, 0x77, 0x02, 0xf9, 0x64, 0xaf,
	0xf1, 0x14, 0x1a, 0x4a, 0xb2, 0xa4, 0x48, 0x86, 0x19, 0x6f, 0x28, 0x89, 0x05, 0xe4, 0x92, 0x9c,
	0xb0, 0x6a, 0xe5, 0x95, 0xd1, 0xac, 0x11, 0x82, 0xd8, 0xc2, 0x3e, 0xb4, 0x3f, 0x48, 0x2d, 0x5e,
	0x3d, 0x4b, 0x8b, 0x64, 0xd8, 0xe2, 0x3b, 0x85, 0xb7, 0x00, 0xc2, 0x68, 0x3f, 0x57, 0x9a, 0xac,
	0x63, 0xcd, 0x22, 0x1d, 0xe6, 0xe3, 0xfe, 0x28, 0xc6, 0x99, 0xd4, 0x31, 0x8f, 0x26, 0xf1, 0x02,
	0xb2, 0x0d, 0x39, 0x47, 0x6f, 0xcf, 0x4a, 0xb2, 0x56, 0xf8, 0x5f, 0xb7, 0x32, 0x66, 0xb2, 0x5c,
	0x42, 0xf6, 0xdb, 0xfa, 0xc7, 0x7a, 0x09, 0xb9, 0x58, 0x3b, 0x6f, 0x96, 0x64, 0xb7, 0xdd, 0x8a,
	0x15, 0x6a, 0x6b, 0x26, 0xb7, 0xa8, 0xc6, 0xaa, 0x85, 0xd2, 0x01, 0x35, 0xe3, 0x3b, 0x85, 0x03,
	0xe8, 0xac, 0x5d, 0x55, 0x6a, 0x56, 0xc1, 0x56, 0xce, 0x64, 0xd9, 0x03, 0x98, 0x92, 0xe7, 0xf4,
	0xbe, 0x26, 0xe7, 0xcb, 0xaf, 0x04, 0xba, 0x9c, 0xdc, 0xca, 0x68, 0x47, 0xc8, 0xa0, 0x23, 0x2c,
	0xcd, 0x3d, 0x55, 0x04, 0x5d, 0x5e, 0x4b, 0xbc, 0x87, 0x78, 0xc3, 0x01, 0x23, 0x1f, 0xb3, 0xbf,
	0x37, 0xaf, 0xcf, 0x3c, 0x1e, 0xc6, 0x47, 0xe8, 0x45, 0xd2, 0xb1, 0x34, 0xac, 0xed, 0x78, 0xf9,
	0x60, 0x7a, 0xfc, 0x99, 0x00, 0x46, 0xe9, 0x13, 0xd9, 0x8d, 0x12, 0x84, 0x0f, 0xd0, 0x9e, 0x04,
	0x36, 0x3c, 0xfa, 0xa1, 0xf3, 0xb3, 0x83, 0xa4, 0xbe, 0x65, 0x79, 0x82, 0x77, 0x90, 0x4e, 0xc9,
	0xe3, 0xe0, 0x20, 0xdf, 0x2f, 0xe5, 0x68, 0xf1, 0xa5, 0x1d, 0x5e, 0xdb, 0xcd, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x5e, 0x6c, 0x0b, 0xd7, 0x94, 0x02, 0x00, 0x00,
}
