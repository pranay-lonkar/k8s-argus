// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/keyword_plan_negative_keyword_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for
// [KeywordPlanNegativeKeywordService.GetKeywordPlanNegativeKeyword][google.ads.googleads.v3.services.KeywordPlanNegativeKeywordService.GetKeywordPlanNegativeKeyword].
type GetKeywordPlanNegativeKeywordRequest struct {
	// Required. The resource name of the plan to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeywordPlanNegativeKeywordRequest) Reset()         { *m = GetKeywordPlanNegativeKeywordRequest{} }
func (m *GetKeywordPlanNegativeKeywordRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeywordPlanNegativeKeywordRequest) ProtoMessage()    {}
func (*GetKeywordPlanNegativeKeywordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6976b9c3cb471824, []int{0}
}

func (m *GetKeywordPlanNegativeKeywordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.Unmarshal(m, b)
}
func (m *GetKeywordPlanNegativeKeywordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.Marshal(b, m, deterministic)
}
func (m *GetKeywordPlanNegativeKeywordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.Merge(m, src)
}
func (m *GetKeywordPlanNegativeKeywordRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.Size(m)
}
func (m *GetKeywordPlanNegativeKeywordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest proto.InternalMessageInfo

func (m *GetKeywordPlanNegativeKeywordRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [KeywordPlanNegativeKeywordService.MutateKeywordPlanNegativeKeywords][google.ads.googleads.v3.services.KeywordPlanNegativeKeywordService.MutateKeywordPlanNegativeKeywords].
type MutateKeywordPlanNegativeKeywordsRequest struct {
	// Required. The ID of the customer whose negative keywords are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual Keyword Plan negative
	// keywords.
	Operations []*KeywordPlanNegativeKeywordOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateKeywordPlanNegativeKeywordsRequest) Reset() {
	*m = MutateKeywordPlanNegativeKeywordsRequest{}
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanNegativeKeywordsRequest) ProtoMessage()    {}
func (*MutateKeywordPlanNegativeKeywordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6976b9c3cb471824, []int{1}
}

func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.Unmarshal(m, b)
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.Merge(m, src)
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.Size(m)
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest proto.InternalMessageInfo

func (m *MutateKeywordPlanNegativeKeywordsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateKeywordPlanNegativeKeywordsRequest) GetOperations() []*KeywordPlanNegativeKeywordOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateKeywordPlanNegativeKeywordsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateKeywordPlanNegativeKeywordsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a Keyword Plan negative
// keyword.
type KeywordPlanNegativeKeywordOperation struct {
	// The FieldMask that determines which resource fields are modified in an
	// update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*KeywordPlanNegativeKeywordOperation_Create
	//	*KeywordPlanNegativeKeywordOperation_Update
	//	*KeywordPlanNegativeKeywordOperation_Remove
	Operation            isKeywordPlanNegativeKeywordOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *KeywordPlanNegativeKeywordOperation) Reset()         { *m = KeywordPlanNegativeKeywordOperation{} }
func (m *KeywordPlanNegativeKeywordOperation) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanNegativeKeywordOperation) ProtoMessage()    {}
func (*KeywordPlanNegativeKeywordOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6976b9c3cb471824, []int{2}
}

func (m *KeywordPlanNegativeKeywordOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanNegativeKeywordOperation.Unmarshal(m, b)
}
func (m *KeywordPlanNegativeKeywordOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanNegativeKeywordOperation.Marshal(b, m, deterministic)
}
func (m *KeywordPlanNegativeKeywordOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanNegativeKeywordOperation.Merge(m, src)
}
func (m *KeywordPlanNegativeKeywordOperation) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanNegativeKeywordOperation.Size(m)
}
func (m *KeywordPlanNegativeKeywordOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanNegativeKeywordOperation.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanNegativeKeywordOperation proto.InternalMessageInfo

func (m *KeywordPlanNegativeKeywordOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isKeywordPlanNegativeKeywordOperation_Operation interface {
	isKeywordPlanNegativeKeywordOperation_Operation()
}

type KeywordPlanNegativeKeywordOperation_Create struct {
	Create *resources.KeywordPlanNegativeKeyword `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type KeywordPlanNegativeKeywordOperation_Update struct {
	Update *resources.KeywordPlanNegativeKeyword `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type KeywordPlanNegativeKeywordOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*KeywordPlanNegativeKeywordOperation_Create) isKeywordPlanNegativeKeywordOperation_Operation() {}

func (*KeywordPlanNegativeKeywordOperation_Update) isKeywordPlanNegativeKeywordOperation_Operation() {}

func (*KeywordPlanNegativeKeywordOperation_Remove) isKeywordPlanNegativeKeywordOperation_Operation() {}

func (m *KeywordPlanNegativeKeywordOperation) GetOperation() isKeywordPlanNegativeKeywordOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *KeywordPlanNegativeKeywordOperation) GetCreate() *resources.KeywordPlanNegativeKeyword {
	if x, ok := m.GetOperation().(*KeywordPlanNegativeKeywordOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *KeywordPlanNegativeKeywordOperation) GetUpdate() *resources.KeywordPlanNegativeKeyword {
	if x, ok := m.GetOperation().(*KeywordPlanNegativeKeywordOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *KeywordPlanNegativeKeywordOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*KeywordPlanNegativeKeywordOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*KeywordPlanNegativeKeywordOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*KeywordPlanNegativeKeywordOperation_Create)(nil),
		(*KeywordPlanNegativeKeywordOperation_Update)(nil),
		(*KeywordPlanNegativeKeywordOperation_Remove)(nil),
	}
}

// Response message for a Keyword Plan negative keyword mutate.
type MutateKeywordPlanNegativeKeywordsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateKeywordPlanNegativeKeywordResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *MutateKeywordPlanNegativeKeywordsResponse) Reset() {
	*m = MutateKeywordPlanNegativeKeywordsResponse{}
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanNegativeKeywordsResponse) ProtoMessage()    {}
func (*MutateKeywordPlanNegativeKeywordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6976b9c3cb471824, []int{3}
}

func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.Unmarshal(m, b)
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.Merge(m, src)
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.Size(m)
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse proto.InternalMessageInfo

func (m *MutateKeywordPlanNegativeKeywordsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateKeywordPlanNegativeKeywordsResponse) GetResults() []*MutateKeywordPlanNegativeKeywordResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the Keyword Plan negative keyword mutate.
type MutateKeywordPlanNegativeKeywordResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateKeywordPlanNegativeKeywordResult) Reset() {
	*m = MutateKeywordPlanNegativeKeywordResult{}
}
func (m *MutateKeywordPlanNegativeKeywordResult) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanNegativeKeywordResult) ProtoMessage()    {}
func (*MutateKeywordPlanNegativeKeywordResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_6976b9c3cb471824, []int{4}
}

func (m *MutateKeywordPlanNegativeKeywordResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.Unmarshal(m, b)
}
func (m *MutateKeywordPlanNegativeKeywordResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanNegativeKeywordResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.Merge(m, src)
}
func (m *MutateKeywordPlanNegativeKeywordResult) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.Size(m)
}
func (m *MutateKeywordPlanNegativeKeywordResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult proto.InternalMessageInfo

func (m *MutateKeywordPlanNegativeKeywordResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetKeywordPlanNegativeKeywordRequest)(nil), "google.ads.googleads.v3.services.GetKeywordPlanNegativeKeywordRequest")
	proto.RegisterType((*MutateKeywordPlanNegativeKeywordsRequest)(nil), "google.ads.googleads.v3.services.MutateKeywordPlanNegativeKeywordsRequest")
	proto.RegisterType((*KeywordPlanNegativeKeywordOperation)(nil), "google.ads.googleads.v3.services.KeywordPlanNegativeKeywordOperation")
	proto.RegisterType((*MutateKeywordPlanNegativeKeywordsResponse)(nil), "google.ads.googleads.v3.services.MutateKeywordPlanNegativeKeywordsResponse")
	proto.RegisterType((*MutateKeywordPlanNegativeKeywordResult)(nil), "google.ads.googleads.v3.services.MutateKeywordPlanNegativeKeywordResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/keyword_plan_negative_keyword_service.proto", fileDescriptor_6976b9c3cb471824)
}

var fileDescriptor_6976b9c3cb471824 = []byte{
	// 784 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x4f, 0x6b, 0xe3, 0x46,
	0x14, 0xc0, 0x6b, 0x39, 0xa4, 0xcd, 0x38, 0x69, 0x61, 0x4a, 0x5b, 0xe3, 0xb6, 0xd4, 0x51, 0x4c,
	0xea, 0x9a, 0x22, 0x81, 0x7d, 0x53, 0x30, 0xad, 0x0c, 0x76, 0xd2, 0xa6, 0x49, 0x8c, 0x42, 0x53,
	0x28, 0x06, 0x31, 0x96, 0x26, 0xae, 0x1a, 0x49, 0xa3, 0xce, 0x8c, 0x5c, 0x42, 0xc8, 0x65, 0x8f,
	0x0b, 0x7b, 0xda, 0x6f, 0x90, 0xe3, 0x7e, 0x85, 0xfd, 0x06, 0xb9, 0xee, 0x61, 0x21, 0xa7, 0x3d,
	0xec, 0x5e, 0xf6, 0x3b, 0x2c, 0x2c, 0xd2, 0x68, 0xfc, 0x27, 0xc4, 0x91, 0x21, 0xb9, 0x3d, 0xbd,
	0xf7, 0xfc, 0x7b, 0x7f, 0xe7, 0x19, 0xfc, 0x31, 0x22, 0x64, 0xe4, 0x63, 0x1d, 0xb9, 0x4c, 0x17,
	0x62, 0x22, 0x8d, 0x5b, 0x3a, 0xc3, 0x74, 0xec, 0x39, 0x98, 0xe9, 0x67, 0xf8, 0xfc, 0x7f, 0x42,
	0x5d, 0x3b, 0xf2, 0x51, 0x68, 0x87, 0x78, 0x84, 0xb8, 0x37, 0xc6, 0xb6, 0xd4, 0x66, 0x6e, 0x5a,
	0x44, 0x09, 0x27, 0xb0, 0x2a, 0x10, 0x1a, 0x72, 0x99, 0x36, 0xa1, 0x69, 0xe3, 0x96, 0x26, 0x69,
	0x95, 0xee, 0xa2, 0x78, 0x14, 0x33, 0x12, 0xd3, 0xdc, 0x80, 0x22, 0x50, 0xe5, 0x3b, 0x89, 0x89,
	0x3c, 0x1d, 0x85, 0x21, 0xe1, 0x88, 0x7b, 0x24, 0x64, 0x99, 0xf5, 0x9b, 0x19, 0xab, 0xe3, 0x7b,
	0x38, 0xe4, 0x99, 0xe1, 0x87, 0x19, 0xc3, 0xa9, 0x87, 0x7d, 0xd7, 0x1e, 0xe2, 0x7f, 0xd0, 0xd8,
	0x23, 0x34, 0x73, 0xc8, 0x0a, 0xd0, 0xd3, 0xaf, 0x61, 0x7c, 0x9a, 0x79, 0x05, 0x88, 0x9d, 0xdd,
	0x62, 0xd3, 0xc8, 0xd1, 0x19, 0x47, 0x3c, 0xce, 0x82, 0xaa, 0x7d, 0x50, 0xdb, 0xc5, 0x7c, 0x5f,
	0xa4, 0xd9, 0xf7, 0x51, 0x78, 0x98, 0xa5, 0x9e, 0xa9, 0x2c, 0xfc, 0x5f, 0x8c, 0x19, 0x87, 0x75,
	0xb0, 0x21, 0x6b, 0xb5, 0x43, 0x14, 0xe0, 0x72, 0xa1, 0x5a, 0xa8, 0xaf, 0x75, 0x8a, 0x6f, 0x4c,
	0xc5, 0x5a, 0x97, 0x96, 0x43, 0x14, 0x60, 0xf5, 0xa9, 0x02, 0xea, 0x07, 0x31, 0x47, 0x1c, 0x2f,
	0xa6, 0x32, 0x89, 0xad, 0x81, 0x92, 0x13, 0x33, 0x4e, 0x02, 0x4c, 0x6d, 0xcf, 0x9d, 0x85, 0x02,
	0xa9, 0xff, 0xcd, 0x85, 0xff, 0x02, 0x40, 0x22, 0x4c, 0x45, 0xb7, 0xca, 0x4a, 0xb5, 0x58, 0x2f,
	0x35, 0xbb, 0x5a, 0xde, 0xd4, 0xb4, 0xc5, 0xf1, 0x8f, 0x24, 0x2d, 0x8b, 0x35, 0xa5, 0xc3, 0x1f,
	0xc1, 0x17, 0x11, 0xa2, 0xdc, 0x43, 0xbe, 0x7d, 0x8a, 0x3c, 0x3f, 0xa6, 0xb8, 0x5c, 0xac, 0x16,
	0xea, 0x9f, 0x59, 0x9f, 0x67, 0xea, 0x9e, 0xd0, 0xc2, 0x2d, 0xb0, 0x31, 0x46, 0xbe, 0xe7, 0x22,
	0x8e, 0x6d, 0x12, 0xfa, 0xe7, 0xe5, 0x95, 0xd4, 0x6d, 0x5d, 0x2a, 0x8f, 0x42, 0xff, 0x5c, 0x7d,
	0xa9, 0x80, 0xad, 0x25, 0xd2, 0x80, 0x3b, 0xa0, 0x14, 0x47, 0x29, 0x2a, 0x19, 0x5a, 0x8a, 0x2a,
	0x35, 0x2b, 0xb2, 0x44, 0x39, 0x57, 0xad, 0x97, 0xcc, 0xf5, 0x00, 0xb1, 0x33, 0x0b, 0x08, 0xf7,
	0x44, 0x86, 0x7f, 0x81, 0x55, 0x87, 0x62, 0xc4, 0xc5, 0x50, 0x4a, 0xcd, 0xf6, 0xc2, 0xd6, 0x4c,
	0xd6, 0xf5, 0x9e, 0xde, 0xec, 0x7d, 0x62, 0x65, 0xb8, 0x04, 0x2c, 0xc2, 0x94, 0x95, 0x47, 0x02,
	0x0b, 0x1c, 0x2c, 0x83, 0x55, 0x8a, 0x03, 0x32, 0x16, 0xbd, 0x5d, 0x4b, 0x2c, 0xe2, 0xbb, 0x53,
	0x02, 0x6b, 0x93, 0x61, 0xa8, 0xaf, 0x0b, 0xe0, 0xa7, 0x25, 0x56, 0x89, 0x45, 0x24, 0x64, 0x18,
	0xf6, 0xc0, 0x57, 0xb7, 0x26, 0x67, 0x63, 0x4a, 0x09, 0x4d, 0x63, 0x94, 0x9a, 0x50, 0x26, 0x4f,
	0x23, 0x47, 0x3b, 0x4e, 0xdf, 0x80, 0xf5, 0xe5, 0xfc, 0x4c, 0xbb, 0x89, 0x3b, 0x1c, 0x82, 0x4f,
	0x29, 0x66, 0xb1, 0xcf, 0xe5, 0xaa, 0xed, 0xe5, 0xaf, 0x5a, 0x5e, 0x96, 0x56, 0x0a, 0xb4, 0x24,
	0x58, 0x3d, 0x00, 0xdb, 0xcb, 0xfd, 0x24, 0x59, 0xb3, 0x3b, 0x1e, 0xde, 0xfc, 0x9b, 0x6b, 0xbe,
	0x5b, 0x01, 0x9b, 0x8b, 0x49, 0xc7, 0x22, 0x4b, 0xf8, 0xa1, 0x00, 0xbe, 0xbf, 0xf7, 0xb1, 0xc3,
	0x5e, 0x7e, 0xa5, 0xcb, 0x5c, 0x8b, 0xca, 0xc3, 0x16, 0x45, 0xfd, 0xf3, 0xc6, 0x9c, 0x2f, 0xfa,
	0xc9, 0xab, 0xb7, 0xcf, 0x95, 0x5f, 0x60, 0x3b, 0x39, 0xb9, 0x17, 0x73, 0x96, 0xb6, 0x3c, 0x13,
	0x4c, 0x6f, 0xc8, 0x1b, 0x7c, 0xd7, 0x96, 0xe8, 0x8d, 0x4b, 0x78, 0xa5, 0x80, 0xcd, 0xdc, 0x75,
	0x82, 0xbf, 0x3f, 0x7c, 0xda, 0xf2, 0xbc, 0x55, 0xf6, 0x1f, 0x85, 0x25, 0xf6, 0x5b, 0x75, 0x6f,
	0xcc, 0xaf, 0x67, 0x8e, 0xe5, 0xcf, 0xd3, 0xa3, 0x95, 0xb6, 0xa7, 0xab, 0xfe, 0x9a, 0xb4, 0x67,
	0xda, 0x8f, 0x8b, 0x19, 0xe7, 0x76, 0xe3, 0xf2, 0xbe, 0xee, 0x18, 0x41, 0x9a, 0x85, 0x51, 0x68,
	0x54, 0xbe, 0xbd, 0x36, 0xcb, 0xd3, 0x4c, 0x33, 0x29, 0xf2, 0x98, 0xe6, 0x90, 0xa0, 0xf3, 0x4c,
	0x01, 0x35, 0x87, 0x04, 0xb9, 0x55, 0x75, 0xb6, 0x73, 0xb7, 0xb1, 0x9f, 0xdc, 0xb4, 0x7e, 0xe1,
	0xef, 0xbd, 0x8c, 0x35, 0x22, 0x3e, 0x0a, 0x47, 0x1a, 0xa1, 0x23, 0x7d, 0x84, 0xc3, 0xf4, 0xe2,
	0xe9, 0xd3, 0xe8, 0x8b, 0xff, 0xe9, 0x77, 0xa4, 0x70, 0xa5, 0x14, 0x77, 0x4d, 0xf3, 0x85, 0x52,
	0xdd, 0x15, 0x40, 0xd3, 0x65, 0x9a, 0x10, 0x13, 0xe9, 0xa4, 0xa5, 0x65, 0x81, 0xd9, 0xb5, 0x74,
	0x19, 0x98, 0x2e, 0x1b, 0x4c, 0x5c, 0x06, 0x27, 0xad, 0x81, 0x74, 0x79, 0xaf, 0xd4, 0x84, 0xde,
	0x30, 0x4c, 0x97, 0x19, 0xc6, 0xc4, 0xc9, 0x30, 0x4e, 0x5a, 0x86, 0x21, 0xdd, 0x86, 0xab, 0x69,
	0x9e, 0xad, 0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x32, 0xf9, 0xc4, 0x90, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeywordPlanNegativeKeywordServiceClient is the client API for KeywordPlanNegativeKeywordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanNegativeKeywordServiceClient interface {
	// Returns the requested plan in full detail.
	GetKeywordPlanNegativeKeyword(ctx context.Context, in *GetKeywordPlanNegativeKeywordRequest, opts ...grpc.CallOption) (*resources.KeywordPlanNegativeKeyword, error)
	// Creates, updates, or removes Keyword Plan negative keywords. Operation
	// statuses are returned.
	MutateKeywordPlanNegativeKeywords(ctx context.Context, in *MutateKeywordPlanNegativeKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanNegativeKeywordsResponse, error)
}

type keywordPlanNegativeKeywordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanNegativeKeywordServiceClient(cc grpc.ClientConnInterface) KeywordPlanNegativeKeywordServiceClient {
	return &keywordPlanNegativeKeywordServiceClient{cc}
}

func (c *keywordPlanNegativeKeywordServiceClient) GetKeywordPlanNegativeKeyword(ctx context.Context, in *GetKeywordPlanNegativeKeywordRequest, opts ...grpc.CallOption) (*resources.KeywordPlanNegativeKeyword, error) {
	out := new(resources.KeywordPlanNegativeKeyword)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.KeywordPlanNegativeKeywordService/GetKeywordPlanNegativeKeyword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanNegativeKeywordServiceClient) MutateKeywordPlanNegativeKeywords(ctx context.Context, in *MutateKeywordPlanNegativeKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanNegativeKeywordsResponse, error) {
	out := new(MutateKeywordPlanNegativeKeywordsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.KeywordPlanNegativeKeywordService/MutateKeywordPlanNegativeKeywords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanNegativeKeywordServiceServer is the server API for KeywordPlanNegativeKeywordService service.
type KeywordPlanNegativeKeywordServiceServer interface {
	// Returns the requested plan in full detail.
	GetKeywordPlanNegativeKeyword(context.Context, *GetKeywordPlanNegativeKeywordRequest) (*resources.KeywordPlanNegativeKeyword, error)
	// Creates, updates, or removes Keyword Plan negative keywords. Operation
	// statuses are returned.
	MutateKeywordPlanNegativeKeywords(context.Context, *MutateKeywordPlanNegativeKeywordsRequest) (*MutateKeywordPlanNegativeKeywordsResponse, error)
}

// UnimplementedKeywordPlanNegativeKeywordServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanNegativeKeywordServiceServer struct {
}

func (*UnimplementedKeywordPlanNegativeKeywordServiceServer) GetKeywordPlanNegativeKeyword(ctx context.Context, req *GetKeywordPlanNegativeKeywordRequest) (*resources.KeywordPlanNegativeKeyword, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetKeywordPlanNegativeKeyword not implemented")
}
func (*UnimplementedKeywordPlanNegativeKeywordServiceServer) MutateKeywordPlanNegativeKeywords(ctx context.Context, req *MutateKeywordPlanNegativeKeywordsRequest) (*MutateKeywordPlanNegativeKeywordsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateKeywordPlanNegativeKeywords not implemented")
}

func RegisterKeywordPlanNegativeKeywordServiceServer(s *grpc.Server, srv KeywordPlanNegativeKeywordServiceServer) {
	s.RegisterService(&_KeywordPlanNegativeKeywordService_serviceDesc, srv)
}

func _KeywordPlanNegativeKeywordService_GetKeywordPlanNegativeKeyword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordPlanNegativeKeywordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanNegativeKeywordServiceServer).GetKeywordPlanNegativeKeyword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.KeywordPlanNegativeKeywordService/GetKeywordPlanNegativeKeyword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanNegativeKeywordServiceServer).GetKeywordPlanNegativeKeyword(ctx, req.(*GetKeywordPlanNegativeKeywordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanNegativeKeywordService_MutateKeywordPlanNegativeKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanNegativeKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanNegativeKeywordServiceServer).MutateKeywordPlanNegativeKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.KeywordPlanNegativeKeywordService/MutateKeywordPlanNegativeKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanNegativeKeywordServiceServer).MutateKeywordPlanNegativeKeywords(ctx, req.(*MutateKeywordPlanNegativeKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanNegativeKeywordService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.KeywordPlanNegativeKeywordService",
	HandlerType: (*KeywordPlanNegativeKeywordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordPlanNegativeKeyword",
			Handler:    _KeywordPlanNegativeKeywordService_GetKeywordPlanNegativeKeyword_Handler,
		},
		{
			MethodName: "MutateKeywordPlanNegativeKeywords",
			Handler:    _KeywordPlanNegativeKeywordService_MutateKeywordPlanNegativeKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/keyword_plan_negative_keyword_service.proto",
}
