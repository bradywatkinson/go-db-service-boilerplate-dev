// Code generated by protoc-gen-go. DO NOT EDIT.
// source: review/review.proto

package review

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

// ReviewRequest
type ReviewRequest struct {
	Id                   int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Review               *ReviewSchema `protobuf:"bytes,2,opt,name=review,proto3" json:"review,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReviewRequest) Reset()         { *m = ReviewRequest{} }
func (m *ReviewRequest) String() string { return proto.CompactTextString(m) }
func (*ReviewRequest) ProtoMessage()    {}
func (*ReviewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_review_f618f0afd27a71be, []int{0}
}
func (m *ReviewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReviewRequest.Unmarshal(m, b)
}
func (m *ReviewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReviewRequest.Marshal(b, m, deterministic)
}
func (dst *ReviewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReviewRequest.Merge(dst, src)
}
func (m *ReviewRequest) XXX_Size() int {
	return xxx_messageInfo_ReviewRequest.Size(m)
}
func (m *ReviewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReviewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReviewRequest proto.InternalMessageInfo

func (m *ReviewRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReviewRequest) GetReview() *ReviewSchema {
	if m != nil {
		return m.Review
	}
	return nil
}

// The review interface schema
type ReviewSchema struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Rating               int32    `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Review               string   `protobuf:"bytes,3,opt,name=review,proto3" json:"review,omitempty"`
	CustomerId           int32    `protobuf:"varint,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ProductId            int32    `protobuf:"varint,5,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReviewSchema) Reset()         { *m = ReviewSchema{} }
func (m *ReviewSchema) String() string { return proto.CompactTextString(m) }
func (*ReviewSchema) ProtoMessage()    {}
func (*ReviewSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_review_f618f0afd27a71be, []int{1}
}
func (m *ReviewSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReviewSchema.Unmarshal(m, b)
}
func (m *ReviewSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReviewSchema.Marshal(b, m, deterministic)
}
func (dst *ReviewSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReviewSchema.Merge(dst, src)
}
func (m *ReviewSchema) XXX_Size() int {
	return xxx_messageInfo_ReviewSchema.Size(m)
}
func (m *ReviewSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_ReviewSchema.DiscardUnknown(m)
}

var xxx_messageInfo_ReviewSchema proto.InternalMessageInfo

func (m *ReviewSchema) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReviewSchema) GetRating() int32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *ReviewSchema) GetReview() string {
	if m != nil {
		return m.Review
	}
	return ""
}

func (m *ReviewSchema) GetCustomerId() int32 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *ReviewSchema) GetProductId() int32 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

// ReviewsRequest describes a request to get count reviews starting at id start
type ReviewsRequest struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Start                int32    `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReviewsRequest) Reset()         { *m = ReviewsRequest{} }
func (m *ReviewsRequest) String() string { return proto.CompactTextString(m) }
func (*ReviewsRequest) ProtoMessage()    {}
func (*ReviewsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_review_f618f0afd27a71be, []int{2}
}
func (m *ReviewsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReviewsRequest.Unmarshal(m, b)
}
func (m *ReviewsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReviewsRequest.Marshal(b, m, deterministic)
}
func (dst *ReviewsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReviewsRequest.Merge(dst, src)
}
func (m *ReviewsRequest) XXX_Size() int {
	return xxx_messageInfo_ReviewsRequest.Size(m)
}
func (m *ReviewsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReviewsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReviewsRequest proto.InternalMessageInfo

func (m *ReviewsRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReviewsRequest) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

// ReviewsResponse is a list of reviews
type ReviewsResponse struct {
	Reviews              []*ReviewSchema `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReviewsResponse) Reset()         { *m = ReviewsResponse{} }
func (m *ReviewsResponse) String() string { return proto.CompactTextString(m) }
func (*ReviewsResponse) ProtoMessage()    {}
func (*ReviewsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_review_f618f0afd27a71be, []int{3}
}
func (m *ReviewsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReviewsResponse.Unmarshal(m, b)
}
func (m *ReviewsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReviewsResponse.Marshal(b, m, deterministic)
}
func (dst *ReviewsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReviewsResponse.Merge(dst, src)
}
func (m *ReviewsResponse) XXX_Size() int {
	return xxx_messageInfo_ReviewsResponse.Size(m)
}
func (m *ReviewsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReviewsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReviewsResponse proto.InternalMessageInfo

func (m *ReviewsResponse) GetReviews() []*ReviewSchema {
	if m != nil {
		return m.Reviews
	}
	return nil
}

func init() {
	proto.RegisterType((*ReviewRequest)(nil), "review.ReviewRequest")
	proto.RegisterType((*ReviewSchema)(nil), "review.ReviewSchema")
	proto.RegisterType((*ReviewsRequest)(nil), "review.ReviewsRequest")
	proto.RegisterType((*ReviewsResponse)(nil), "review.ReviewsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReviewServiceClient is the client API for ReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReviewServiceClient interface {
	// Create a review
	CreateReview(ctx context.Context, in *ReviewSchema, opts ...grpc.CallOption) (*ReviewSchema, error)
	// Get a review by id
	GetReview(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewSchema, error)
	// Update a review by id
	UpdateReview(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewSchema, error)
	// Delete a review by id
	DeleteReview(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*wrappers.StringValue, error)
	// GetReviews gets a list of reviews by start and count
	GetReviews(ctx context.Context, in *ReviewsRequest, opts ...grpc.CallOption) (*ReviewsResponse, error)
}

type reviewServiceClient struct {
	cc *grpc.ClientConn
}

func NewReviewServiceClient(cc *grpc.ClientConn) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) CreateReview(ctx context.Context, in *ReviewSchema, opts ...grpc.CallOption) (*ReviewSchema, error) {
	out := new(ReviewSchema)
	err := c.cc.Invoke(ctx, "/review.ReviewService/CreateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetReview(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewSchema, error) {
	out := new(ReviewSchema)
	err := c.cc.Invoke(ctx, "/review.ReviewService/GetReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) UpdateReview(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewSchema, error) {
	out := new(ReviewSchema)
	err := c.cc.Invoke(ctx, "/review.ReviewService/UpdateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) DeleteReview(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*wrappers.StringValue, error) {
	out := new(wrappers.StringValue)
	err := c.cc.Invoke(ctx, "/review.ReviewService/DeleteReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetReviews(ctx context.Context, in *ReviewsRequest, opts ...grpc.CallOption) (*ReviewsResponse, error) {
	out := new(ReviewsResponse)
	err := c.cc.Invoke(ctx, "/review.ReviewService/GetReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServiceServer is the server API for ReviewService service.
type ReviewServiceServer interface {
	// Create a review
	CreateReview(context.Context, *ReviewSchema) (*ReviewSchema, error)
	// Get a review by id
	GetReview(context.Context, *ReviewRequest) (*ReviewSchema, error)
	// Update a review by id
	UpdateReview(context.Context, *ReviewRequest) (*ReviewSchema, error)
	// Delete a review by id
	DeleteReview(context.Context, *ReviewRequest) (*wrappers.StringValue, error)
	// GetReviews gets a list of reviews by start and count
	GetReviews(context.Context, *ReviewsRequest) (*ReviewsResponse, error)
}

func RegisterReviewServiceServer(s *grpc.Server, srv ReviewServiceServer) {
	s.RegisterService(&_ReviewService_serviceDesc, srv)
}

func _ReviewService_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/CreateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).CreateReview(ctx, req.(*ReviewSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/GetReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetReview(ctx, req.(*ReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_UpdateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).UpdateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/UpdateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).UpdateReview(ctx, req.(*ReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_DeleteReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).DeleteReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/DeleteReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).DeleteReview(ctx, req.(*ReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/GetReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetReviews(ctx, req.(*ReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReviewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "review.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReview",
			Handler:    _ReviewService_CreateReview_Handler,
		},
		{
			MethodName: "GetReview",
			Handler:    _ReviewService_GetReview_Handler,
		},
		{
			MethodName: "UpdateReview",
			Handler:    _ReviewService_UpdateReview_Handler,
		},
		{
			MethodName: "DeleteReview",
			Handler:    _ReviewService_DeleteReview_Handler,
		},
		{
			MethodName: "GetReviews",
			Handler:    _ReviewService_GetReviews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "review/review.proto",
}

func init() { proto.RegisterFile("review/review.proto", fileDescriptor_review_f618f0afd27a71be) }

var fileDescriptor_review_f618f0afd27a71be = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4f, 0xea, 0x40,
	0x14, 0xa5, 0xe5, 0x95, 0x17, 0x2e, 0x7d, 0xbc, 0x64, 0x1e, 0x0f, 0x1b, 0xe2, 0x47, 0xd3, 0x15,
	0x0b, 0x53, 0x12, 0xdc, 0x19, 0x8c, 0x31, 0x6a, 0x0c, 0x0b, 0x37, 0x25, 0xba, 0x35, 0xa5, 0xbd,
	0xd6, 0x26, 0xd0, 0xa9, 0x33, 0x53, 0xf8, 0x17, 0xfe, 0x36, 0x7f, 0x92, 0x61, 0x3e, 0x50, 0x02,
	0xc4, 0xb8, 0x9a, 0xdc, 0x73, 0xcf, 0x39, 0xf7, 0xdc, 0x99, 0x81, 0x7f, 0x0c, 0x17, 0x39, 0x2e,
	0x07, 0xea, 0x08, 0x4b, 0x46, 0x05, 0x25, 0x0d, 0x55, 0xf5, 0x8e, 0x33, 0x4a, 0xb3, 0x19, 0x0e,
	0x24, 0x3a, 0xad, 0x9e, 0x07, 0x4b, 0x16, 0x97, 0x25, 0x32, 0xae, 0x78, 0xc1, 0x3d, 0xfc, 0x89,
	0x24, 0x33, 0xc2, 0xd7, 0x0a, 0xb9, 0x20, 0x6d, 0xb0, 0xf3, 0xd4, 0xb3, 0x7c, 0xab, 0xef, 0x44,
	0x76, 0x9e, 0x92, 0x53, 0xd0, 0x56, 0x9e, 0xed, 0x5b, 0xfd, 0xd6, 0xb0, 0x13, 0xea, 0x39, 0x4a,
	0x36, 0x49, 0x5e, 0x70, 0x1e, 0x47, 0x9a, 0x13, 0xbc, 0x59, 0xe0, 0x7e, 0x6d, 0x6c, 0xd9, 0x75,
	0xa1, 0xc1, 0x62, 0x91, 0x17, 0x99, 0xb4, 0x73, 0x22, 0x5d, 0x49, 0x5c, 0x8d, 0xa9, 0xfb, 0x56,
	0xbf, 0x69, 0x0c, 0xc9, 0x09, 0xb4, 0x92, 0x8a, 0x0b, 0x3a, 0x47, 0xf6, 0x94, 0xa7, 0xde, 0x2f,
	0x29, 0x02, 0x03, 0x8d, 0x53, 0x72, 0x04, 0x50, 0x32, 0x9a, 0x56, 0x89, 0x58, 0xf5, 0x1d, 0xd9,
	0x6f, 0x6a, 0x64, 0x9c, 0x06, 0x23, 0x68, 0xab, 0x3c, 0xdc, 0x2c, 0xd8, 0x01, 0x27, 0xa1, 0x55,
	0x21, 0x74, 0x28, 0x55, 0xac, 0x50, 0x2e, 0x62, 0x26, 0x74, 0x2c, 0x55, 0x04, 0x57, 0xf0, 0x77,
	0xad, 0xe6, 0x25, 0x2d, 0x38, 0x92, 0x10, 0x7e, 0xab, 0x68, 0xdc, 0xb3, 0xfc, 0xfa, 0xde, 0x0b,
	0x31, 0xa4, 0xe1, 0xbb, 0x6d, 0x6e, 0x78, 0x82, 0x6c, 0x91, 0x27, 0x48, 0x46, 0xe0, 0x5e, 0x33,
	0x8c, 0x05, 0x2a, 0x98, 0xec, 0x34, 0xe8, 0xed, 0x44, 0x83, 0x1a, 0x39, 0x87, 0xe6, 0x1d, 0x0a,
	0x2d, 0xfd, 0xbf, 0x49, 0xd2, 0x2b, 0xee, 0xd5, 0x5e, 0x80, 0xfb, 0x50, 0xa6, 0x9f, 0x93, 0x7f,
	0x28, 0xbf, 0x05, 0xf7, 0x06, 0x67, 0xf8, 0x9d, 0xfc, 0x30, 0x54, 0x7f, 0x2e, 0x34, 0x7f, 0x2e,
	0x9c, 0x08, 0x96, 0x17, 0xd9, 0x63, 0x3c, 0xab, 0x30, 0xa8, 0x91, 0x4b, 0x80, 0xf5, 0x06, 0x9c,
	0x74, 0x37, 0x4d, 0xcc, 0x33, 0xf5, 0x0e, 0xb6, 0x70, 0xf5, 0x00, 0x41, 0x6d, 0xda, 0x90, 0xc6,
	0x67, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x82, 0xd4, 0x91, 0xf9, 0x02, 0x00, 0x00,
}
