// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product/product.proto

package product

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// ProductRequest
type ProductRequest struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Product              *ProductSchema `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ProductRequest) Reset()         { *m = ProductRequest{} }
func (m *ProductRequest) String() string { return proto.CompactTextString(m) }
func (*ProductRequest) ProtoMessage()    {}
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_5ba75c655102b35a, []int{0}
}
func (m *ProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductRequest.Unmarshal(m, b)
}
func (m *ProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductRequest.Marshal(b, m, deterministic)
}
func (dst *ProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductRequest.Merge(dst, src)
}
func (m *ProductRequest) XXX_Size() int {
	return xxx_messageInfo_ProductRequest.Size(m)
}
func (m *ProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProductRequest proto.InternalMessageInfo

func (m *ProductRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductRequest) GetProduct() *ProductSchema {
	if m != nil {
		return m.Product
	}
	return nil
}

// The product interface schema
type ProductSchema struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	MerchantId           int32    `protobuf:"varint,4,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductSchema) Reset()         { *m = ProductSchema{} }
func (m *ProductSchema) String() string { return proto.CompactTextString(m) }
func (*ProductSchema) ProtoMessage()    {}
func (*ProductSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_5ba75c655102b35a, []int{1}
}
func (m *ProductSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductSchema.Unmarshal(m, b)
}
func (m *ProductSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductSchema.Marshal(b, m, deterministic)
}
func (dst *ProductSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductSchema.Merge(dst, src)
}
func (m *ProductSchema) XXX_Size() int {
	return xxx_messageInfo_ProductSchema.Size(m)
}
func (m *ProductSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductSchema.DiscardUnknown(m)
}

var xxx_messageInfo_ProductSchema proto.InternalMessageInfo

func (m *ProductSchema) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSchema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductSchema) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ProductSchema) GetMerchantId() int32 {
	if m != nil {
		return m.MerchantId
	}
	return 0
}

// ProductsRequest describes a request to get count products starting at id start
type ProductsRequest struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Start                int32    `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductsRequest) Reset()         { *m = ProductsRequest{} }
func (m *ProductsRequest) String() string { return proto.CompactTextString(m) }
func (*ProductsRequest) ProtoMessage()    {}
func (*ProductsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_5ba75c655102b35a, []int{2}
}
func (m *ProductsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductsRequest.Unmarshal(m, b)
}
func (m *ProductsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductsRequest.Marshal(b, m, deterministic)
}
func (dst *ProductsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductsRequest.Merge(dst, src)
}
func (m *ProductsRequest) XXX_Size() int {
	return xxx_messageInfo_ProductsRequest.Size(m)
}
func (m *ProductsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProductsRequest proto.InternalMessageInfo

func (m *ProductsRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ProductsRequest) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

// ProductsResponse is a list of products
type ProductsResponse struct {
	Products             []*ProductSchema `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ProductsResponse) Reset()         { *m = ProductsResponse{} }
func (m *ProductsResponse) String() string { return proto.CompactTextString(m) }
func (*ProductsResponse) ProtoMessage()    {}
func (*ProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_5ba75c655102b35a, []int{3}
}
func (m *ProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductsResponse.Unmarshal(m, b)
}
func (m *ProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductsResponse.Marshal(b, m, deterministic)
}
func (dst *ProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductsResponse.Merge(dst, src)
}
func (m *ProductsResponse) XXX_Size() int {
	return xxx_messageInfo_ProductsResponse.Size(m)
}
func (m *ProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductsResponse proto.InternalMessageInfo

func (m *ProductsResponse) GetProducts() []*ProductSchema {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*ProductRequest)(nil), "product.ProductRequest")
	proto.RegisterType((*ProductSchema)(nil), "product.ProductSchema")
	proto.RegisterType((*ProductsRequest)(nil), "product.ProductsRequest")
	proto.RegisterType((*ProductsResponse)(nil), "product.ProductsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	// Create a product
	CreateProduct(ctx context.Context, in *ProductSchema, opts ...grpc.CallOption) (*ProductSchema, error)
	// Get a product by id
	GetProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductSchema, error)
	// Update a product by id
	UpdateProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductSchema, error)
	// Delete a product by id
	DeleteProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*wrappers.StringValue, error)
	// GetProducts gets a list of products by start and count
	GetProducts(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *ProductSchema, opts ...grpc.CallOption) (*ProductSchema, error) {
	out := new(ProductSchema)
	err := c.cc.Invoke(ctx, "/product.ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductSchema, error) {
	out := new(ProductSchema)
	err := c.cc.Invoke(ctx, "/product.ProductService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductSchema, error) {
	out := new(ProductSchema)
	err := c.cc.Invoke(ctx, "/product.ProductService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*wrappers.StringValue, error) {
	out := new(wrappers.StringValue)
	err := c.cc.Invoke(ctx, "/product.ProductService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProducts(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error) {
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	// Create a product
	CreateProduct(context.Context, *ProductSchema) (*ProductSchema, error)
	// Get a product by id
	GetProduct(context.Context, *ProductRequest) (*ProductSchema, error)
	// Update a product by id
	UpdateProduct(context.Context, *ProductRequest) (*ProductSchema, error)
	// Delete a product by id
	DeleteProduct(context.Context, *ProductRequest) (*wrappers.StringValue, error)
	// GetProducts gets a list of products by start and count
	GetProducts(context.Context, *ProductsRequest) (*ProductsResponse, error)
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*ProductSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProduct(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProducts(ctx, req.(*ProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _ProductService_GetProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product/product.proto",
}

func init() { proto.RegisterFile("product/product.proto", fileDescriptor_product_5ba75c655102b35a) }

var fileDescriptor_product_5ba75c655102b35a = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x5f, 0x4f, 0xf2, 0x30,
	0x14, 0xc6, 0xd9, 0x60, 0xef, 0x9f, 0x43, 0xc6, 0xfb, 0xa6, 0x41, 0x9d, 0xc4, 0xe8, 0xb2, 0x2b,
	0xae, 0x86, 0xc1, 0x6b, 0x63, 0x54, 0xe2, 0x9f, 0x3b, 0x33, 0xa2, 0xb7, 0xa6, 0x6c, 0x47, 0x98,
	0x81, 0xb5, 0xb6, 0x9d, 0x7e, 0x65, 0x3f, 0x86, 0x61, 0x6d, 0x41, 0x51, 0x88, 0xf1, 0x6a, 0x3b,
	0x4f, 0x9f, 0xf3, 0x3b, 0x7d, 0xda, 0xc2, 0x16, 0x17, 0x2c, 0x2b, 0x53, 0xd5, 0x33, 0xdf, 0x98,
	0x0b, 0xa6, 0x18, 0xf9, 0x6d, 0xca, 0xce, 0xfe, 0x98, 0xb1, 0xf1, 0x14, 0x7b, 0x95, 0x3c, 0x2a,
	0x1f, 0x7a, 0x2f, 0x82, 0x72, 0x8e, 0x42, 0x6a, 0x63, 0x94, 0x40, 0xeb, 0x46, 0x5b, 0x13, 0x7c,
	0x2a, 0x51, 0x2a, 0xd2, 0x02, 0x37, 0xcf, 0x02, 0x27, 0x74, 0xba, 0x5e, 0xe2, 0xe6, 0x19, 0x39,
	0x04, 0x0b, 0x0b, 0xdc, 0xd0, 0xe9, 0x36, 0xfb, 0xdb, 0xb1, 0x9d, 0x65, 0x3a, 0x87, 0xe9, 0x04,
	0x67, 0x34, 0xb1, 0xb6, 0xe8, 0x11, 0xfc, 0x0f, 0x2b, 0x9f, 0x90, 0x04, 0x1a, 0x05, 0x9d, 0x61,
	0xc5, 0xfb, 0x9b, 0x54, 0xff, 0xa4, 0x0d, 0x1e, 0x17, 0x79, 0x8a, 0x41, 0x3d, 0x74, 0xba, 0x4e,
	0xa2, 0x0b, 0x72, 0x00, 0xcd, 0x19, 0x8a, 0x74, 0x42, 0x0b, 0x75, 0x9f, 0x67, 0x41, 0xa3, 0x42,
	0x80, 0x95, 0xae, 0xb3, 0xe8, 0x18, 0xfe, 0x99, 0x59, 0xd2, 0x06, 0x68, 0x83, 0x97, 0xb2, 0xb2,
	0x50, 0x66, 0xa0, 0x2e, 0xe6, 0xaa, 0x54, 0x54, 0xe8, 0x10, 0x5e, 0xa2, 0x8b, 0xe8, 0x02, 0xfe,
	0x2f, 0xdb, 0x25, 0x67, 0x85, 0x44, 0xd2, 0x87, 0x3f, 0x26, 0x89, 0x0c, 0x9c, 0xb0, 0xbe, 0x21,
	0xf1, 0xc2, 0xd7, 0x7f, 0x75, 0x17, 0xe7, 0x38, 0x44, 0xf1, 0x3c, 0xdf, 0xfa, 0x29, 0xf8, 0xe7,
	0x02, 0xa9, 0x42, 0xa3, 0x93, 0x35, 0x94, 0xce, 0x1a, 0x3d, 0xaa, 0x91, 0x13, 0x80, 0x4b, 0x54,
	0xb6, 0x7f, 0x67, 0xd5, 0x67, 0x02, 0x6f, 0x00, 0x9c, 0x81, 0x7f, 0xcb, 0xb3, 0x77, 0x7b, 0xf8,
	0x01, 0xe3, 0x0a, 0xfc, 0x01, 0x4e, 0xf1, 0x1b, 0x8c, 0xbd, 0x58, 0x3f, 0xb6, 0xd8, 0x3e, 0xb6,
	0x78, 0xa8, 0x44, 0x5e, 0x8c, 0xef, 0xe8, 0xb4, 0xc4, 0xa8, 0x46, 0x06, 0xd0, 0x5c, 0xc6, 0x91,
	0x24, 0x58, 0xe5, 0xd8, 0x1b, 0xec, 0xec, 0x7e, 0xb1, 0xa2, 0x2f, 0x27, 0xaa, 0x8d, 0x7e, 0x55,
	0xf4, 0xa3, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xf5, 0x3b, 0x86, 0xfa, 0x02, 0x00, 0x00,
}
