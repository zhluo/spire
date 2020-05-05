// Code generated by protoc-gen-go. DO NOT EDIT.
// source: svid.proto

package svid

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	types "github.com/spiffe/spire/proto/spire-next/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MintX509SVIDRequest struct {
	// Required. ASN.1 DER encoded CSR. The CSR is used to convey the public
	// key and the SPIFFE ID (via the URI SAN). Only one URI SAN can be set.
	// Optionally, the subject and any number of DNS SANs can also be set.
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	// The desired TTL of the X509-SVID, in seconds. The server default will be
	// used if unset. The TTL is advisory only. The actual lifetime of the
	// X509-SVID may be lower depending on the remaining lifetime of the active
	// SPIRE Server CA.
	Ttl                  int32    `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MintX509SVIDRequest) Reset()         { *m = MintX509SVIDRequest{} }
func (m *MintX509SVIDRequest) String() string { return proto.CompactTextString(m) }
func (*MintX509SVIDRequest) ProtoMessage()    {}
func (*MintX509SVIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{0}
}

func (m *MintX509SVIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MintX509SVIDRequest.Unmarshal(m, b)
}
func (m *MintX509SVIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MintX509SVIDRequest.Marshal(b, m, deterministic)
}
func (m *MintX509SVIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintX509SVIDRequest.Merge(m, src)
}
func (m *MintX509SVIDRequest) XXX_Size() int {
	return xxx_messageInfo_MintX509SVIDRequest.Size(m)
}
func (m *MintX509SVIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MintX509SVIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MintX509SVIDRequest proto.InternalMessageInfo

func (m *MintX509SVIDRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

func (m *MintX509SVIDRequest) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type MintX509SVIDResponse struct {
	// The newly issued X509-SVID.
	Svid                 *types.X509SVID `protobuf:"bytes,1,opt,name=svid,proto3" json:"svid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MintX509SVIDResponse) Reset()         { *m = MintX509SVIDResponse{} }
func (m *MintX509SVIDResponse) String() string { return proto.CompactTextString(m) }
func (*MintX509SVIDResponse) ProtoMessage()    {}
func (*MintX509SVIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{1}
}

func (m *MintX509SVIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MintX509SVIDResponse.Unmarshal(m, b)
}
func (m *MintX509SVIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MintX509SVIDResponse.Marshal(b, m, deterministic)
}
func (m *MintX509SVIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintX509SVIDResponse.Merge(m, src)
}
func (m *MintX509SVIDResponse) XXX_Size() int {
	return xxx_messageInfo_MintX509SVIDResponse.Size(m)
}
func (m *MintX509SVIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MintX509SVIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MintX509SVIDResponse proto.InternalMessageInfo

func (m *MintX509SVIDResponse) GetSvid() *types.X509SVID {
	if m != nil {
		return m.Svid
	}
	return nil
}

type MintJWTSVIDRequest struct {
	// Required. SPIFFE ID of the JWT-SVID.
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// Required. List of audience claims to include in the JWT-SVID. At least one must
	// be set.
	Audience []string `protobuf:"bytes,2,rep,name=audience,proto3" json:"audience,omitempty"`
	// Desired TTL of the JWT-SVID, in seconds. The server default will be used
	// if unset. The TTL is advisory only. The actual lifetime of the JWT-SVID
	// may be lower depending on the remaining lifetime of the active SPIRE
	// Server CA.
	Ttl                  int32    `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MintJWTSVIDRequest) Reset()         { *m = MintJWTSVIDRequest{} }
func (m *MintJWTSVIDRequest) String() string { return proto.CompactTextString(m) }
func (*MintJWTSVIDRequest) ProtoMessage()    {}
func (*MintJWTSVIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{2}
}

func (m *MintJWTSVIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MintJWTSVIDRequest.Unmarshal(m, b)
}
func (m *MintJWTSVIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MintJWTSVIDRequest.Marshal(b, m, deterministic)
}
func (m *MintJWTSVIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintJWTSVIDRequest.Merge(m, src)
}
func (m *MintJWTSVIDRequest) XXX_Size() int {
	return xxx_messageInfo_MintJWTSVIDRequest.Size(m)
}
func (m *MintJWTSVIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MintJWTSVIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MintJWTSVIDRequest proto.InternalMessageInfo

func (m *MintJWTSVIDRequest) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *MintJWTSVIDRequest) GetAudience() []string {
	if m != nil {
		return m.Audience
	}
	return nil
}

func (m *MintJWTSVIDRequest) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type MintJWTSVIDResponse struct {
	// The newly issued JWT-SVID.
	Svid                 *types.JWTSVID `protobuf:"bytes,1,opt,name=svid,proto3" json:"svid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MintJWTSVIDResponse) Reset()         { *m = MintJWTSVIDResponse{} }
func (m *MintJWTSVIDResponse) String() string { return proto.CompactTextString(m) }
func (*MintJWTSVIDResponse) ProtoMessage()    {}
func (*MintJWTSVIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{3}
}

func (m *MintJWTSVIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MintJWTSVIDResponse.Unmarshal(m, b)
}
func (m *MintJWTSVIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MintJWTSVIDResponse.Marshal(b, m, deterministic)
}
func (m *MintJWTSVIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintJWTSVIDResponse.Merge(m, src)
}
func (m *MintJWTSVIDResponse) XXX_Size() int {
	return xxx_messageInfo_MintJWTSVIDResponse.Size(m)
}
func (m *MintJWTSVIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MintJWTSVIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MintJWTSVIDResponse proto.InternalMessageInfo

func (m *MintJWTSVIDResponse) GetSvid() *types.JWTSVID {
	if m != nil {
		return m.Svid
	}
	return nil
}

type BatchNewX509SVIDRequest struct {
	// Required. One or more X509-SVID parameters for X509-SVID entries to
	// be signed.
	Params               []*NewX509SVIDParams `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BatchNewX509SVIDRequest) Reset()         { *m = BatchNewX509SVIDRequest{} }
func (m *BatchNewX509SVIDRequest) String() string { return proto.CompactTextString(m) }
func (*BatchNewX509SVIDRequest) ProtoMessage()    {}
func (*BatchNewX509SVIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{4}
}

func (m *BatchNewX509SVIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchNewX509SVIDRequest.Unmarshal(m, b)
}
func (m *BatchNewX509SVIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchNewX509SVIDRequest.Marshal(b, m, deterministic)
}
func (m *BatchNewX509SVIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchNewX509SVIDRequest.Merge(m, src)
}
func (m *BatchNewX509SVIDRequest) XXX_Size() int {
	return xxx_messageInfo_BatchNewX509SVIDRequest.Size(m)
}
func (m *BatchNewX509SVIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchNewX509SVIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchNewX509SVIDRequest proto.InternalMessageInfo

func (m *BatchNewX509SVIDRequest) GetParams() []*NewX509SVIDParams {
	if m != nil {
		return m.Params
	}
	return nil
}

type BatchNewX509SVIDResponse struct {
	// Result for each X509-SVID requested.
	Results              []*BatchNewX509SVIDResponse_Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *BatchNewX509SVIDResponse) Reset()         { *m = BatchNewX509SVIDResponse{} }
func (m *BatchNewX509SVIDResponse) String() string { return proto.CompactTextString(m) }
func (*BatchNewX509SVIDResponse) ProtoMessage()    {}
func (*BatchNewX509SVIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{5}
}

func (m *BatchNewX509SVIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchNewX509SVIDResponse.Unmarshal(m, b)
}
func (m *BatchNewX509SVIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchNewX509SVIDResponse.Marshal(b, m, deterministic)
}
func (m *BatchNewX509SVIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchNewX509SVIDResponse.Merge(m, src)
}
func (m *BatchNewX509SVIDResponse) XXX_Size() int {
	return xxx_messageInfo_BatchNewX509SVIDResponse.Size(m)
}
func (m *BatchNewX509SVIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchNewX509SVIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchNewX509SVIDResponse proto.InternalMessageInfo

func (m *BatchNewX509SVIDResponse) GetResults() []*BatchNewX509SVIDResponse_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type BatchNewX509SVIDResponse_Result struct {
	// The status of creating the X509-SVID.
	Status *types.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// The newly created X509-SVID. This will be set if the status is OK.
	Bundle               *types.X509SVID `protobuf:"bytes,2,opt,name=bundle,proto3" json:"bundle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BatchNewX509SVIDResponse_Result) Reset()         { *m = BatchNewX509SVIDResponse_Result{} }
func (m *BatchNewX509SVIDResponse_Result) String() string { return proto.CompactTextString(m) }
func (*BatchNewX509SVIDResponse_Result) ProtoMessage()    {}
func (*BatchNewX509SVIDResponse_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{5, 0}
}

func (m *BatchNewX509SVIDResponse_Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchNewX509SVIDResponse_Result.Unmarshal(m, b)
}
func (m *BatchNewX509SVIDResponse_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchNewX509SVIDResponse_Result.Marshal(b, m, deterministic)
}
func (m *BatchNewX509SVIDResponse_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchNewX509SVIDResponse_Result.Merge(m, src)
}
func (m *BatchNewX509SVIDResponse_Result) XXX_Size() int {
	return xxx_messageInfo_BatchNewX509SVIDResponse_Result.Size(m)
}
func (m *BatchNewX509SVIDResponse_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchNewX509SVIDResponse_Result.DiscardUnknown(m)
}

var xxx_messageInfo_BatchNewX509SVIDResponse_Result proto.InternalMessageInfo

func (m *BatchNewX509SVIDResponse_Result) GetStatus() *types.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BatchNewX509SVIDResponse_Result) GetBundle() *types.X509SVID {
	if m != nil {
		return m.Bundle
	}
	return nil
}

type NewJWTSVIDRequest struct {
	// Required. The entry ID of the identity being requested.
	EntryId string `protobuf:"bytes,1,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	// Required. List of audience claims to include in the JWT-SVID. At least
	// one must be set.
	Audience             []string `protobuf:"bytes,2,rep,name=audience,proto3" json:"audience,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewJWTSVIDRequest) Reset()         { *m = NewJWTSVIDRequest{} }
func (m *NewJWTSVIDRequest) String() string { return proto.CompactTextString(m) }
func (*NewJWTSVIDRequest) ProtoMessage()    {}
func (*NewJWTSVIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{6}
}

func (m *NewJWTSVIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewJWTSVIDRequest.Unmarshal(m, b)
}
func (m *NewJWTSVIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewJWTSVIDRequest.Marshal(b, m, deterministic)
}
func (m *NewJWTSVIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewJWTSVIDRequest.Merge(m, src)
}
func (m *NewJWTSVIDRequest) XXX_Size() int {
	return xxx_messageInfo_NewJWTSVIDRequest.Size(m)
}
func (m *NewJWTSVIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewJWTSVIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewJWTSVIDRequest proto.InternalMessageInfo

func (m *NewJWTSVIDRequest) GetEntryId() string {
	if m != nil {
		return m.EntryId
	}
	return ""
}

func (m *NewJWTSVIDRequest) GetAudience() []string {
	if m != nil {
		return m.Audience
	}
	return nil
}

type NewJWTSVIDResponse struct {
	// The newly issued JWT-SVID
	Svid                 *types.JWTSVID `protobuf:"bytes,1,opt,name=svid,proto3" json:"svid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NewJWTSVIDResponse) Reset()         { *m = NewJWTSVIDResponse{} }
func (m *NewJWTSVIDResponse) String() string { return proto.CompactTextString(m) }
func (*NewJWTSVIDResponse) ProtoMessage()    {}
func (*NewJWTSVIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{7}
}

func (m *NewJWTSVIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewJWTSVIDResponse.Unmarshal(m, b)
}
func (m *NewJWTSVIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewJWTSVIDResponse.Marshal(b, m, deterministic)
}
func (m *NewJWTSVIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewJWTSVIDResponse.Merge(m, src)
}
func (m *NewJWTSVIDResponse) XXX_Size() int {
	return xxx_messageInfo_NewJWTSVIDResponse.Size(m)
}
func (m *NewJWTSVIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewJWTSVIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewJWTSVIDResponse proto.InternalMessageInfo

func (m *NewJWTSVIDResponse) GetSvid() *types.JWTSVID {
	if m != nil {
		return m.Svid
	}
	return nil
}

type NewDownstreamX509CARequest struct {
	// Required. The entry ID for the identity being requested.
	EntryId string `protobuf:"bytes,1,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	// Required. The ASN.1 DER encoded Certificate Signing Request (CSR). The
	// CSR is only used to convey the public key; other fields in the CSR are
	// ignored. The X509-SVID attributes are determined by the entry.
	Csr                  []byte   `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewDownstreamX509CARequest) Reset()         { *m = NewDownstreamX509CARequest{} }
func (m *NewDownstreamX509CARequest) String() string { return proto.CompactTextString(m) }
func (*NewDownstreamX509CARequest) ProtoMessage()    {}
func (*NewDownstreamX509CARequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{8}
}

func (m *NewDownstreamX509CARequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewDownstreamX509CARequest.Unmarshal(m, b)
}
func (m *NewDownstreamX509CARequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewDownstreamX509CARequest.Marshal(b, m, deterministic)
}
func (m *NewDownstreamX509CARequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewDownstreamX509CARequest.Merge(m, src)
}
func (m *NewDownstreamX509CARequest) XXX_Size() int {
	return xxx_messageInfo_NewDownstreamX509CARequest.Size(m)
}
func (m *NewDownstreamX509CARequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewDownstreamX509CARequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewDownstreamX509CARequest proto.InternalMessageInfo

func (m *NewDownstreamX509CARequest) GetEntryId() string {
	if m != nil {
		return m.EntryId
	}
	return ""
}

func (m *NewDownstreamX509CARequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

type NewDownstreamX509CAResponse struct {
	// CA certificate and any intermediates required to form a chain of trust
	// back to the X.509 authorities (DER encoded). The CA certificate is the
	// first.
	CaCertChain [][]byte `protobuf:"bytes,1,rep,name=ca_cert_chain,json=caCertChain,proto3" json:"ca_cert_chain,omitempty"`
	// X.509 authorities (DER encoded).
	X509Authorities      [][]byte `protobuf:"bytes,2,rep,name=x509_authorities,json=x509Authorities,proto3" json:"x509_authorities,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewDownstreamX509CAResponse) Reset()         { *m = NewDownstreamX509CAResponse{} }
func (m *NewDownstreamX509CAResponse) String() string { return proto.CompactTextString(m) }
func (*NewDownstreamX509CAResponse) ProtoMessage()    {}
func (*NewDownstreamX509CAResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{9}
}

func (m *NewDownstreamX509CAResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewDownstreamX509CAResponse.Unmarshal(m, b)
}
func (m *NewDownstreamX509CAResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewDownstreamX509CAResponse.Marshal(b, m, deterministic)
}
func (m *NewDownstreamX509CAResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewDownstreamX509CAResponse.Merge(m, src)
}
func (m *NewDownstreamX509CAResponse) XXX_Size() int {
	return xxx_messageInfo_NewDownstreamX509CAResponse.Size(m)
}
func (m *NewDownstreamX509CAResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewDownstreamX509CAResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewDownstreamX509CAResponse proto.InternalMessageInfo

func (m *NewDownstreamX509CAResponse) GetCaCertChain() [][]byte {
	if m != nil {
		return m.CaCertChain
	}
	return nil
}

func (m *NewDownstreamX509CAResponse) GetX509Authorities() [][]byte {
	if m != nil {
		return m.X509Authorities
	}
	return nil
}

type NewX509SVIDParams struct {
	// Required. The entry ID for the identity being requested.
	EntryId string `protobuf:"bytes,1,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	// Required. The ASN.1 DER encoded Certificate Signing Request (CSR). The
	// CSR is only used to convey the public key; other fields in the CSR are
	// ignored. The X509-SVID attributes are determined by the entry.
	Csr                  []byte   `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewX509SVIDParams) Reset()         { *m = NewX509SVIDParams{} }
func (m *NewX509SVIDParams) String() string { return proto.CompactTextString(m) }
func (*NewX509SVIDParams) ProtoMessage()    {}
func (*NewX509SVIDParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{10}
}

func (m *NewX509SVIDParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewX509SVIDParams.Unmarshal(m, b)
}
func (m *NewX509SVIDParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewX509SVIDParams.Marshal(b, m, deterministic)
}
func (m *NewX509SVIDParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewX509SVIDParams.Merge(m, src)
}
func (m *NewX509SVIDParams) XXX_Size() int {
	return xxx_messageInfo_NewX509SVIDParams.Size(m)
}
func (m *NewX509SVIDParams) XXX_DiscardUnknown() {
	xxx_messageInfo_NewX509SVIDParams.DiscardUnknown(m)
}

var xxx_messageInfo_NewX509SVIDParams proto.InternalMessageInfo

func (m *NewX509SVIDParams) GetEntryId() string {
	if m != nil {
		return m.EntryId
	}
	return ""
}

func (m *NewX509SVIDParams) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

func init() {
	proto.RegisterType((*MintX509SVIDRequest)(nil), "spire.api.server.svid.v1.MintX509SVIDRequest")
	proto.RegisterType((*MintX509SVIDResponse)(nil), "spire.api.server.svid.v1.MintX509SVIDResponse")
	proto.RegisterType((*MintJWTSVIDRequest)(nil), "spire.api.server.svid.v1.MintJWTSVIDRequest")
	proto.RegisterType((*MintJWTSVIDResponse)(nil), "spire.api.server.svid.v1.MintJWTSVIDResponse")
	proto.RegisterType((*BatchNewX509SVIDRequest)(nil), "spire.api.server.svid.v1.BatchNewX509SVIDRequest")
	proto.RegisterType((*BatchNewX509SVIDResponse)(nil), "spire.api.server.svid.v1.BatchNewX509SVIDResponse")
	proto.RegisterType((*BatchNewX509SVIDResponse_Result)(nil), "spire.api.server.svid.v1.BatchNewX509SVIDResponse.Result")
	proto.RegisterType((*NewJWTSVIDRequest)(nil), "spire.api.server.svid.v1.NewJWTSVIDRequest")
	proto.RegisterType((*NewJWTSVIDResponse)(nil), "spire.api.server.svid.v1.NewJWTSVIDResponse")
	proto.RegisterType((*NewDownstreamX509CARequest)(nil), "spire.api.server.svid.v1.NewDownstreamX509CARequest")
	proto.RegisterType((*NewDownstreamX509CAResponse)(nil), "spire.api.server.svid.v1.NewDownstreamX509CAResponse")
	proto.RegisterType((*NewX509SVIDParams)(nil), "spire.api.server.svid.v1.NewX509SVIDParams")
}

func init() { proto.RegisterFile("svid.proto", fileDescriptor_a3513929f2a405cf) }

var fileDescriptor_a3513929f2a405cf = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4d, 0x6f, 0xda, 0x30,
	0x18, 0x56, 0x4a, 0x47, 0xdb, 0x17, 0xa6, 0x75, 0x6e, 0xa7, 0x65, 0xa9, 0xb6, 0xa1, 0x9c, 0xa8,
	0x5a, 0x92, 0x96, 0xad, 0x07, 0x34, 0x69, 0x2b, 0xd0, 0x0b, 0x95, 0x86, 0xa6, 0x30, 0x6d, 0xd3,
	0x0e, 0x43, 0x26, 0xb8, 0x25, 0x15, 0x24, 0x99, 0xed, 0x40, 0xb9, 0xee, 0xe7, 0xed, 0xbe, 0xff,
	0x33, 0xd9, 0x31, 0x90, 0xf0, 0x25, 0x38, 0x25, 0x79, 0x3f, 0x9e, 0xe7, 0xf1, 0xfb, 0xe1, 0x00,
	0xb0, 0xa1, 0xd7, 0xb5, 0x42, 0x1a, 0xf0, 0x00, 0xe9, 0x2c, 0xf4, 0x28, 0xb1, 0x70, 0xe8, 0x59,
	0x8c, 0xd0, 0x21, 0xa1, 0x96, 0x74, 0x0e, 0x2f, 0x8d, 0x37, 0xd2, 0x53, 0xf2, 0xc9, 0x23, 0xb7,
	0xf9, 0x38, 0x24, 0xcc, 0x7e, 0x18, 0xf1, 0x59, 0xa6, 0xf1, 0x7a, 0xc1, 0xcf, 0x38, 0xe6, 0x11,
	0x53, 0xee, 0xb7, 0x0b, 0xee, 0xc7, 0xab, 0x8b, 0xca, 0x2c, 0xdf, 0xac, 0xc0, 0xd1, 0x67, 0xcf,
	0xe7, 0x3f, 0xae, 0x2e, 0x2a, 0xad, 0x6f, 0x8d, 0x1b, 0x87, 0xfc, 0x8e, 0x08, 0xe3, 0xe8, 0x10,
	0x32, 0x2e, 0xa3, 0xba, 0x56, 0xd0, 0x8a, 0x79, 0x47, 0xbc, 0x0a, 0x0b, 0xe7, 0x7d, 0x3d, 0x53,
	0xd0, 0x8a, 0x4f, 0x1c, 0xf1, 0x6a, 0x56, 0xe1, 0x38, 0x9d, 0xca, 0xc2, 0xc0, 0x67, 0x04, 0x9d,
	0xc2, 0xae, 0x20, 0x90, 0xc9, 0xb9, 0xf2, 0x0b, 0x2b, 0x3e, 0x9b, 0x64, 0xb7, 0xa6, 0xc1, 0x32,
	0xc4, 0x6c, 0x03, 0x12, 0x10, 0xb7, 0xdf, 0xbf, 0x26, 0xc9, 0x4f, 0xe0, 0x80, 0x85, 0xde, 0xdd,
	0x1d, 0x69, 0x2b, 0x94, 0x03, 0x67, 0x3f, 0x36, 0x34, 0xba, 0xc8, 0x80, 0x7d, 0x1c, 0x75, 0x3d,
	0xe2, 0xbb, 0x44, 0xdf, 0x29, 0x64, 0x84, 0x6f, 0xf2, 0xbd, 0x44, 0xe3, 0xa7, 0xf8, 0x78, 0x53,
	0x02, 0x25, 0xb1, 0x98, 0x92, 0x78, 0x9c, 0x92, 0x38, 0x89, 0x8d, 0x15, 0xfe, 0x82, 0x97, 0x35,
	0xcc, 0xdd, 0x5e, 0x93, 0x8c, 0xe6, 0x6b, 0x54, 0x87, 0x6c, 0x88, 0x29, 0x1e, 0x30, 0x5d, 0x2b,
	0x64, 0x8a, 0xb9, 0xf2, 0x99, 0xb5, 0xaa, 0x8b, 0x56, 0x22, 0xfb, 0x8b, 0x4c, 0x71, 0x54, 0xaa,
	0xf9, 0x4f, 0x03, 0x7d, 0x91, 0x40, 0xc9, 0x6c, 0xc1, 0x1e, 0x25, 0x2c, 0xea, 0xf3, 0x09, 0x45,
	0x65, 0x35, 0xc5, 0x2a, 0x10, 0xcb, 0x91, 0x08, 0xce, 0x04, 0xc9, 0xe8, 0x42, 0x36, 0x36, 0xa1,
	0x33, 0xc8, 0xc6, 0xc3, 0xa2, 0xea, 0x70, 0x94, 0xaa, 0x43, 0x4b, 0xba, 0x1c, 0x15, 0x82, 0x4a,
	0x90, 0xed, 0x44, 0x7e, 0xb7, 0x2f, 0xaa, 0xbe, 0xa6, 0xaf, 0x2a, 0xc8, 0xbc, 0x85, 0xe7, 0x4d,
	0x32, 0x9a, 0x6b, 0xec, 0x2b, 0xd8, 0x27, 0x3e, 0xa7, 0xe3, 0x59, 0x5f, 0xf7, 0xe4, 0xf7, 0xfa,
	0xb6, 0x9a, 0x1f, 0x01, 0x25, 0xb1, 0xb6, 0xee, 0x61, 0x03, 0x8c, 0x26, 0x19, 0xdd, 0x04, 0x23,
	0x9f, 0x71, 0x4a, 0xf0, 0x40, 0x88, 0xad, 0x57, 0x37, 0x10, 0xa5, 0xb6, 0x60, 0x67, 0xba, 0x05,
	0x66, 0x1f, 0x4e, 0x96, 0x42, 0x29, 0x4d, 0x26, 0x3c, 0x75, 0x71, 0xdb, 0x25, 0x94, 0xb7, 0xdd,
	0x1e, 0xf6, 0x7c, 0xd9, 0xb6, 0xbc, 0x93, 0x73, 0x71, 0x9d, 0x50, 0x5e, 0x17, 0x26, 0x74, 0x0a,
	0x87, 0x62, 0x07, 0xdb, 0x38, 0xe2, 0xbd, 0x80, 0x7a, 0xdc, 0x23, 0x4c, 0x9e, 0x38, 0xef, 0x3c,
	0x13, 0xf6, 0xea, 0xcc, 0x6c, 0x5e, 0xcb, 0x22, 0xa6, 0x27, 0x67, 0x2b, 0xbd, 0xe5, 0xbf, 0xbb,
	0xb0, 0x2b, 0x72, 0xd1, 0x00, 0xf2, 0xc9, 0x65, 0x45, 0xa5, 0xd5, 0x93, 0xb4, 0xe4, 0x3e, 0x30,
	0xac, 0x4d, 0xc3, 0x55, 0x21, 0x1e, 0x20, 0x97, 0xd8, 0x3b, 0x74, 0xbe, 0x3e, 0x3d, 0x3d, 0x26,
	0x46, 0x69, 0xc3, 0x68, 0xc5, 0x35, 0x86, 0xc3, 0xf9, 0xe1, 0x47, 0x97, 0xdb, 0x2c, 0x4a, 0xcc,
	0x5a, 0xde, 0x7e, 0xb7, 0xd0, 0x3d, 0xc0, 0x6c, 0x32, 0xd1, 0xfa, 0x0b, 0x60, 0xee, 0x90, 0xe7,
	0x9b, 0x05, 0x2b, 0xa2, 0x3f, 0x1a, 0x1c, 0x2d, 0x19, 0x3c, 0xf4, 0x7e, 0x2d, 0xca, 0x8a, 0x91,
	0x37, 0xae, 0xb6, 0xcc, 0x8a, 0x45, 0xd4, 0x6a, 0x3f, 0xaf, 0xef, 0x3d, 0xde, 0x8b, 0x3a, 0x96,
	0x1b, 0x0c, 0xec, 0xf8, 0x46, 0xb6, 0x25, 0x92, 0x2d, 0x7f, 0x26, 0x76, 0xe2, 0x67, 0x83, 0x43,
	0xcf, 0x8e, 0xa1, 0x6d, 0x01, 0x6d, 0x0f, 0x2f, 0x3f, 0x88, 0x67, 0x27, 0x2b, 0x23, 0xdf, 0xfd,
	0x0f, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xf1, 0x93, 0x57, 0xfe, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SVIDClient is the client API for SVID service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SVIDClient interface {
	// Mints a one-off X509-SVID outside of the normal node/workload
	// registration process.
	//
	// The caller must be local or present an admin X509-SVID.
	MintX509SVID(ctx context.Context, in *MintX509SVIDRequest, opts ...grpc.CallOption) (*MintX509SVIDResponse, error)
	// Mints a one-off JWT-SVID outside of the normal node/workload
	// registration process.
	//
	// The caller must be local or present an admin X509-SVID.
	MintJWTSVID(ctx context.Context, in *MintJWTSVIDRequest, opts ...grpc.CallOption) (*MintJWTSVIDResponse, error)
	// Creates one or more X509-SVIDs from registration entries.
	//
	// The caller must present an active agent X509-SVID that is authorized
	// to mint the requested entries. See the Entry GetAuthorizedEntries RPC.
	BatchNewX509SVID(ctx context.Context, in *BatchNewX509SVIDRequest, opts ...grpc.CallOption) (*BatchNewX509SVIDResponse, error)
	// Creates an JWT-SVID from a registration entry.
	//
	// The caller must present an active agent X509-SVID that is authorized
	// to mint the requested entry. See the Entry GetAuthorizedEntries RPC.
	NewJWTSVID(ctx context.Context, in *NewJWTSVIDRequest, opts ...grpc.CallOption) (*NewJWTSVIDResponse, error)
	// Creates an X509 CA certificate appropriate for use by a downstream
	// entity to mint X509-SVIDs.
	//
	// The caller must present a downstream X509-SVID.
	NewDownstreamX509CA(ctx context.Context, in *NewDownstreamX509CARequest, opts ...grpc.CallOption) (*NewDownstreamX509CAResponse, error)
}

type sVIDClient struct {
	cc *grpc.ClientConn
}

func NewSVIDClient(cc *grpc.ClientConn) SVIDClient {
	return &sVIDClient{cc}
}

func (c *sVIDClient) MintX509SVID(ctx context.Context, in *MintX509SVIDRequest, opts ...grpc.CallOption) (*MintX509SVIDResponse, error) {
	out := new(MintX509SVIDResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.svid.v1.SVID/MintX509SVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sVIDClient) MintJWTSVID(ctx context.Context, in *MintJWTSVIDRequest, opts ...grpc.CallOption) (*MintJWTSVIDResponse, error) {
	out := new(MintJWTSVIDResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.svid.v1.SVID/MintJWTSVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sVIDClient) BatchNewX509SVID(ctx context.Context, in *BatchNewX509SVIDRequest, opts ...grpc.CallOption) (*BatchNewX509SVIDResponse, error) {
	out := new(BatchNewX509SVIDResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.svid.v1.SVID/BatchNewX509SVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sVIDClient) NewJWTSVID(ctx context.Context, in *NewJWTSVIDRequest, opts ...grpc.CallOption) (*NewJWTSVIDResponse, error) {
	out := new(NewJWTSVIDResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.svid.v1.SVID/NewJWTSVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sVIDClient) NewDownstreamX509CA(ctx context.Context, in *NewDownstreamX509CARequest, opts ...grpc.CallOption) (*NewDownstreamX509CAResponse, error) {
	out := new(NewDownstreamX509CAResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.svid.v1.SVID/NewDownstreamX509CA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SVIDServer is the server API for SVID service.
type SVIDServer interface {
	// Mints a one-off X509-SVID outside of the normal node/workload
	// registration process.
	//
	// The caller must be local or present an admin X509-SVID.
	MintX509SVID(context.Context, *MintX509SVIDRequest) (*MintX509SVIDResponse, error)
	// Mints a one-off JWT-SVID outside of the normal node/workload
	// registration process.
	//
	// The caller must be local or present an admin X509-SVID.
	MintJWTSVID(context.Context, *MintJWTSVIDRequest) (*MintJWTSVIDResponse, error)
	// Creates one or more X509-SVIDs from registration entries.
	//
	// The caller must present an active agent X509-SVID that is authorized
	// to mint the requested entries. See the Entry GetAuthorizedEntries RPC.
	BatchNewX509SVID(context.Context, *BatchNewX509SVIDRequest) (*BatchNewX509SVIDResponse, error)
	// Creates an JWT-SVID from a registration entry.
	//
	// The caller must present an active agent X509-SVID that is authorized
	// to mint the requested entry. See the Entry GetAuthorizedEntries RPC.
	NewJWTSVID(context.Context, *NewJWTSVIDRequest) (*NewJWTSVIDResponse, error)
	// Creates an X509 CA certificate appropriate for use by a downstream
	// entity to mint X509-SVIDs.
	//
	// The caller must present a downstream X509-SVID.
	NewDownstreamX509CA(context.Context, *NewDownstreamX509CARequest) (*NewDownstreamX509CAResponse, error)
}

// UnimplementedSVIDServer can be embedded to have forward compatible implementations.
type UnimplementedSVIDServer struct {
}

func (*UnimplementedSVIDServer) MintX509SVID(ctx context.Context, req *MintX509SVIDRequest) (*MintX509SVIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintX509SVID not implemented")
}
func (*UnimplementedSVIDServer) MintJWTSVID(ctx context.Context, req *MintJWTSVIDRequest) (*MintJWTSVIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintJWTSVID not implemented")
}
func (*UnimplementedSVIDServer) BatchNewX509SVID(ctx context.Context, req *BatchNewX509SVIDRequest) (*BatchNewX509SVIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchNewX509SVID not implemented")
}
func (*UnimplementedSVIDServer) NewJWTSVID(ctx context.Context, req *NewJWTSVIDRequest) (*NewJWTSVIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewJWTSVID not implemented")
}
func (*UnimplementedSVIDServer) NewDownstreamX509CA(ctx context.Context, req *NewDownstreamX509CARequest) (*NewDownstreamX509CAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewDownstreamX509CA not implemented")
}

func RegisterSVIDServer(s *grpc.Server, srv SVIDServer) {
	s.RegisterService(&_SVID_serviceDesc, srv)
}

func _SVID_MintX509SVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintX509SVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SVIDServer).MintX509SVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.svid.v1.SVID/MintX509SVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SVIDServer).MintX509SVID(ctx, req.(*MintX509SVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SVID_MintJWTSVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintJWTSVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SVIDServer).MintJWTSVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.svid.v1.SVID/MintJWTSVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SVIDServer).MintJWTSVID(ctx, req.(*MintJWTSVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SVID_BatchNewX509SVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchNewX509SVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SVIDServer).BatchNewX509SVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.svid.v1.SVID/BatchNewX509SVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SVIDServer).BatchNewX509SVID(ctx, req.(*BatchNewX509SVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SVID_NewJWTSVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewJWTSVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SVIDServer).NewJWTSVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.svid.v1.SVID/NewJWTSVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SVIDServer).NewJWTSVID(ctx, req.(*NewJWTSVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SVID_NewDownstreamX509CA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewDownstreamX509CARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SVIDServer).NewDownstreamX509CA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.svid.v1.SVID/NewDownstreamX509CA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SVIDServer).NewDownstreamX509CA(ctx, req.(*NewDownstreamX509CARequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SVID_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.api.server.svid.v1.SVID",
	HandlerType: (*SVIDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MintX509SVID",
			Handler:    _SVID_MintX509SVID_Handler,
		},
		{
			MethodName: "MintJWTSVID",
			Handler:    _SVID_MintJWTSVID_Handler,
		},
		{
			MethodName: "BatchNewX509SVID",
			Handler:    _SVID_BatchNewX509SVID_Handler,
		},
		{
			MethodName: "NewJWTSVID",
			Handler:    _SVID_NewJWTSVID_Handler,
		},
		{
			MethodName: "NewDownstreamX509CA",
			Handler:    _SVID_NewDownstreamX509CA_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svid.proto",
}
