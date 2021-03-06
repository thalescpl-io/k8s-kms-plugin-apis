// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simplekms/v1/messages.proto

package simplekms

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type KeyKind int32

const (
	KeyKind_UNKNOWN KeyKind = 0
	KeyKind_AES     KeyKind = 1
	KeyKind_RSA     KeyKind = 2
	KeyKind_ECC     KeyKind = 3
)

var KeyKind_name = map[int32]string{
	0: "UNKNOWN",
	1: "AES",
	2: "RSA",
	3: "ECC",
}

var KeyKind_value = map[string]int32{
	"UNKNOWN": 0,
	"AES":     1,
	"RSA":     2,
	"ECC":     3,
}

func (x KeyKind) String() string {
	return proto.EnumName(KeyKind_name, int32(x))
}

func (KeyKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{0}
}

type VersionRequest struct {
	// Version of the KMS plugin API.
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{0}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

func (m *VersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type VersionResponse struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Name of the KMS provider.
	RuntimeName string `protobuf:"bytes,2,opt,name=runtime_name,json=runtimeName,proto3" json:"runtime_name,omitempty"`
	// Version of the KMS provider. The string must be semver-compatible.
	RuntimeVersion       string   `protobuf:"bytes,3,opt,name=runtime_version,json=runtimeVersion,proto3" json:"runtime_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{1}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionResponse) GetRuntimeName() string {
	if m != nil {
		return m.RuntimeName
	}
	return ""
}

func (m *VersionResponse) GetRuntimeVersion() string {
	if m != nil {
		return m.RuntimeVersion
	}
	return ""
}

type GenerateKEKRequest struct {
	// optional kid, otherwise will be autogenerated as a UUID.v4 in the response
	KekKid               []byte   `protobuf:"bytes,1,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateKEKRequest) Reset()         { *m = GenerateKEKRequest{} }
func (m *GenerateKEKRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateKEKRequest) ProtoMessage()    {}
func (*GenerateKEKRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{2}
}

func (m *GenerateKEKRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKEKRequest.Unmarshal(m, b)
}
func (m *GenerateKEKRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKEKRequest.Marshal(b, m, deterministic)
}
func (m *GenerateKEKRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKEKRequest.Merge(m, src)
}
func (m *GenerateKEKRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateKEKRequest.Size(m)
}
func (m *GenerateKEKRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKEKRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKEKRequest proto.InternalMessageInfo

func (m *GenerateKEKRequest) GetKekKid() []byte {
	if m != nil {
		return m.KekKid
	}
	return nil
}

type GenerateKEKResponse struct {
	// KEK kid
	KekKid               []byte   `protobuf:"bytes,1,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateKEKResponse) Reset()         { *m = GenerateKEKResponse{} }
func (m *GenerateKEKResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateKEKResponse) ProtoMessage()    {}
func (*GenerateKEKResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{3}
}

func (m *GenerateKEKResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKEKResponse.Unmarshal(m, b)
}
func (m *GenerateKEKResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKEKResponse.Marshal(b, m, deterministic)
}
func (m *GenerateKEKResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKEKResponse.Merge(m, src)
}
func (m *GenerateKEKResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateKEKResponse.Size(m)
}
func (m *GenerateKEKResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKEKResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKEKResponse proto.InternalMessageInfo

func (m *GenerateKEKResponse) GetKekKid() []byte {
	if m != nil {
		return m.KekKid
	}
	return nil
}

type DestroyKEKRequest struct {
	// Required kid of KEK to find and delete
	KekKid               []byte   `protobuf:"bytes,1,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyKEKRequest) Reset()         { *m = DestroyKEKRequest{} }
func (m *DestroyKEKRequest) String() string { return proto.CompactTextString(m) }
func (*DestroyKEKRequest) ProtoMessage()    {}
func (*DestroyKEKRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{4}
}

func (m *DestroyKEKRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyKEKRequest.Unmarshal(m, b)
}
func (m *DestroyKEKRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyKEKRequest.Marshal(b, m, deterministic)
}
func (m *DestroyKEKRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyKEKRequest.Merge(m, src)
}
func (m *DestroyKEKRequest) XXX_Size() int {
	return xxx_messageInfo_DestroyKEKRequest.Size(m)
}
func (m *DestroyKEKRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyKEKRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyKEKRequest proto.InternalMessageInfo

func (m *DestroyKEKRequest) GetKekKid() []byte {
	if m != nil {
		return m.KekKid
	}
	return nil
}

type DestroyKEKResponse struct {
	// successful destroy?
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyKEKResponse) Reset()         { *m = DestroyKEKResponse{} }
func (m *DestroyKEKResponse) String() string { return proto.CompactTextString(m) }
func (*DestroyKEKResponse) ProtoMessage()    {}
func (*DestroyKEKResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{5}
}

func (m *DestroyKEKResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyKEKResponse.Unmarshal(m, b)
}
func (m *DestroyKEKResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyKEKResponse.Marshal(b, m, deterministic)
}
func (m *DestroyKEKResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyKEKResponse.Merge(m, src)
}
func (m *DestroyKEKResponse) XXX_Size() int {
	return xxx_messageInfo_DestroyKEKResponse.Size(m)
}
func (m *DestroyKEKResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyKEKResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyKEKResponse proto.InternalMessageInfo

func (m *DestroyKEKResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GenerateCAKRequest struct {
	// key size in bits
	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// What kind of key is it... only Asymmetric kinds will be supported
	Kind KeyKind `protobuf:"varint,2,opt,name=kind,proto3,enum=thales.cpl.kms.simplekms.v1.KeyKind" json:"kind,omitempty"`
	// Root CA Key ID
	RootCaKid            []byte   `protobuf:"bytes,3,opt,name=root_ca_kid,json=rootCaKid,proto3" json:"root_ca_kid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateCAKRequest) Reset()         { *m = GenerateCAKRequest{} }
func (m *GenerateCAKRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateCAKRequest) ProtoMessage()    {}
func (*GenerateCAKRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{6}
}

func (m *GenerateCAKRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateCAKRequest.Unmarshal(m, b)
}
func (m *GenerateCAKRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateCAKRequest.Marshal(b, m, deterministic)
}
func (m *GenerateCAKRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateCAKRequest.Merge(m, src)
}
func (m *GenerateCAKRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateCAKRequest.Size(m)
}
func (m *GenerateCAKRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateCAKRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateCAKRequest proto.InternalMessageInfo

func (m *GenerateCAKRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GenerateCAKRequest) GetKind() KeyKind {
	if m != nil {
		return m.Kind
	}
	return KeyKind_UNKNOWN
}

func (m *GenerateCAKRequest) GetRootCaKid() []byte {
	if m != nil {
		return m.RootCaKid
	}
	return nil
}

type GenerateCAKResponse struct {
	// Root CA Key ID
	RootCaKid            []byte   `protobuf:"bytes,1,opt,name=root_ca_kid,json=rootCaKid,proto3" json:"root_ca_kid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateCAKResponse) Reset()         { *m = GenerateCAKResponse{} }
func (m *GenerateCAKResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateCAKResponse) ProtoMessage()    {}
func (*GenerateCAKResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{7}
}

func (m *GenerateCAKResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateCAKResponse.Unmarshal(m, b)
}
func (m *GenerateCAKResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateCAKResponse.Marshal(b, m, deterministic)
}
func (m *GenerateCAKResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateCAKResponse.Merge(m, src)
}
func (m *GenerateCAKResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateCAKResponse.Size(m)
}
func (m *GenerateCAKResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateCAKResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateCAKResponse proto.InternalMessageInfo

func (m *GenerateCAKResponse) GetRootCaKid() []byte {
	if m != nil {
		return m.RootCaKid
	}
	return nil
}

type GenerateCARequest struct {
	// Root CA Key ID
	RootCaKid            []byte   `protobuf:"bytes,1,opt,name=root_ca_kid,json=rootCaKid,proto3" json:"root_ca_kid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateCARequest) Reset()         { *m = GenerateCARequest{} }
func (m *GenerateCARequest) String() string { return proto.CompactTextString(m) }
func (*GenerateCARequest) ProtoMessage()    {}
func (*GenerateCARequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{8}
}

func (m *GenerateCARequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateCARequest.Unmarshal(m, b)
}
func (m *GenerateCARequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateCARequest.Marshal(b, m, deterministic)
}
func (m *GenerateCARequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateCARequest.Merge(m, src)
}
func (m *GenerateCARequest) XXX_Size() int {
	return xxx_messageInfo_GenerateCARequest.Size(m)
}
func (m *GenerateCARequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateCARequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateCARequest proto.InternalMessageInfo

func (m *GenerateCARequest) GetRootCaKid() []byte {
	if m != nil {
		return m.RootCaKid
	}
	return nil
}

type GenerateCAResponse struct {
	// Root CA cert in PEM format
	Cert                 []byte   `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateCAResponse) Reset()         { *m = GenerateCAResponse{} }
func (m *GenerateCAResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateCAResponse) ProtoMessage()    {}
func (*GenerateCAResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{9}
}

func (m *GenerateCAResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateCAResponse.Unmarshal(m, b)
}
func (m *GenerateCAResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateCAResponse.Marshal(b, m, deterministic)
}
func (m *GenerateCAResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateCAResponse.Merge(m, src)
}
func (m *GenerateCAResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateCAResponse.Size(m)
}
func (m *GenerateCAResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateCAResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateCAResponse proto.InternalMessageInfo

func (m *GenerateCAResponse) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

type DestroyCAKRequest struct {
	// Required kid of KEK to find and delete
	KekKid               []byte   `protobuf:"bytes,1,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyCAKRequest) Reset()         { *m = DestroyCAKRequest{} }
func (m *DestroyCAKRequest) String() string { return proto.CompactTextString(m) }
func (*DestroyCAKRequest) ProtoMessage()    {}
func (*DestroyCAKRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{10}
}

func (m *DestroyCAKRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyCAKRequest.Unmarshal(m, b)
}
func (m *DestroyCAKRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyCAKRequest.Marshal(b, m, deterministic)
}
func (m *DestroyCAKRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyCAKRequest.Merge(m, src)
}
func (m *DestroyCAKRequest) XXX_Size() int {
	return xxx_messageInfo_DestroyCAKRequest.Size(m)
}
func (m *DestroyCAKRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyCAKRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyCAKRequest proto.InternalMessageInfo

func (m *DestroyCAKRequest) GetKekKid() []byte {
	if m != nil {
		return m.KekKid
	}
	return nil
}

type DestroyCAKResponse struct {
	// successful destroy?
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyCAKResponse) Reset()         { *m = DestroyCAKResponse{} }
func (m *DestroyCAKResponse) String() string { return proto.CompactTextString(m) }
func (*DestroyCAKResponse) ProtoMessage()    {}
func (*DestroyCAKResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{11}
}

func (m *DestroyCAKResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyCAKResponse.Unmarshal(m, b)
}
func (m *DestroyCAKResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyCAKResponse.Marshal(b, m, deterministic)
}
func (m *DestroyCAKResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyCAKResponse.Merge(m, src)
}
func (m *DestroyCAKResponse) XXX_Size() int {
	return xxx_messageInfo_DestroyCAKResponse.Size(m)
}
func (m *DestroyCAKResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyCAKResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyCAKResponse proto.InternalMessageInfo

func (m *DestroyCAKResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type DestroyCARequest struct {
	// Required kid of KEK to find and delete
	KekKid               []byte   `protobuf:"bytes,1,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyCARequest) Reset()         { *m = DestroyCARequest{} }
func (m *DestroyCARequest) String() string { return proto.CompactTextString(m) }
func (*DestroyCARequest) ProtoMessage()    {}
func (*DestroyCARequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{12}
}

func (m *DestroyCARequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyCARequest.Unmarshal(m, b)
}
func (m *DestroyCARequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyCARequest.Marshal(b, m, deterministic)
}
func (m *DestroyCARequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyCARequest.Merge(m, src)
}
func (m *DestroyCARequest) XXX_Size() int {
	return xxx_messageInfo_DestroyCARequest.Size(m)
}
func (m *DestroyCARequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyCARequest.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyCARequest proto.InternalMessageInfo

func (m *DestroyCARequest) GetKekKid() []byte {
	if m != nil {
		return m.KekKid
	}
	return nil
}

type DestroyCAResponse struct {
	// successful destroy?
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyCAResponse) Reset()         { *m = DestroyCAResponse{} }
func (m *DestroyCAResponse) String() string { return proto.CompactTextString(m) }
func (*DestroyCAResponse) ProtoMessage()    {}
func (*DestroyCAResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{13}
}

func (m *DestroyCAResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyCAResponse.Unmarshal(m, b)
}
func (m *DestroyCAResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyCAResponse.Marshal(b, m, deterministic)
}
func (m *DestroyCAResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyCAResponse.Merge(m, src)
}
func (m *DestroyCAResponse) XXX_Size() int {
	return xxx_messageInfo_DestroyCAResponse.Size(m)
}
func (m *DestroyCAResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyCAResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyCAResponse proto.InternalMessageInfo

func (m *DestroyCAResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type SignCSRRequest struct {
	// Encrypted blob of DEK
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	// Root CA Key ID
	RootCaKid            []byte   `protobuf:"bytes,2,opt,name=root_ca_kid,json=rootCaKid,proto3" json:"root_ca_kid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignCSRRequest) Reset()         { *m = SignCSRRequest{} }
func (m *SignCSRRequest) String() string { return proto.CompactTextString(m) }
func (*SignCSRRequest) ProtoMessage()    {}
func (*SignCSRRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{14}
}

func (m *SignCSRRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignCSRRequest.Unmarshal(m, b)
}
func (m *SignCSRRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignCSRRequest.Marshal(b, m, deterministic)
}
func (m *SignCSRRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignCSRRequest.Merge(m, src)
}
func (m *SignCSRRequest) XXX_Size() int {
	return xxx_messageInfo_SignCSRRequest.Size(m)
}
func (m *SignCSRRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignCSRRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignCSRRequest proto.InternalMessageInfo

func (m *SignCSRRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

func (m *SignCSRRequest) GetRootCaKid() []byte {
	if m != nil {
		return m.RootCaKid
	}
	return nil
}

type SignCSRResponse struct {
	// Certificate in PEM Form
	Cert                 []byte   `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignCSRResponse) Reset()         { *m = SignCSRResponse{} }
func (m *SignCSRResponse) String() string { return proto.CompactTextString(m) }
func (*SignCSRResponse) ProtoMessage()    {}
func (*SignCSRResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_245adf7c9b02e35d, []int{15}
}

func (m *SignCSRResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignCSRResponse.Unmarshal(m, b)
}
func (m *SignCSRResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignCSRResponse.Marshal(b, m, deterministic)
}
func (m *SignCSRResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignCSRResponse.Merge(m, src)
}
func (m *SignCSRResponse) XXX_Size() int {
	return xxx_messageInfo_SignCSRResponse.Size(m)
}
func (m *SignCSRResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignCSRResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignCSRResponse proto.InternalMessageInfo

func (m *SignCSRResponse) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func init() {
	proto.RegisterEnum("thales.cpl.kms.simplekms.v1.KeyKind", KeyKind_name, KeyKind_value)
	proto.RegisterType((*VersionRequest)(nil), "thales.cpl.kms.simplekms.v1.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "thales.cpl.kms.simplekms.v1.VersionResponse")
	proto.RegisterType((*GenerateKEKRequest)(nil), "thales.cpl.kms.simplekms.v1.GenerateKEKRequest")
	proto.RegisterType((*GenerateKEKResponse)(nil), "thales.cpl.kms.simplekms.v1.GenerateKEKResponse")
	proto.RegisterType((*DestroyKEKRequest)(nil), "thales.cpl.kms.simplekms.v1.DestroyKEKRequest")
	proto.RegisterType((*DestroyKEKResponse)(nil), "thales.cpl.kms.simplekms.v1.DestroyKEKResponse")
	proto.RegisterType((*GenerateCAKRequest)(nil), "thales.cpl.kms.simplekms.v1.GenerateCAKRequest")
	proto.RegisterType((*GenerateCAKResponse)(nil), "thales.cpl.kms.simplekms.v1.GenerateCAKResponse")
	proto.RegisterType((*GenerateCARequest)(nil), "thales.cpl.kms.simplekms.v1.GenerateCARequest")
	proto.RegisterType((*GenerateCAResponse)(nil), "thales.cpl.kms.simplekms.v1.GenerateCAResponse")
	proto.RegisterType((*DestroyCAKRequest)(nil), "thales.cpl.kms.simplekms.v1.DestroyCAKRequest")
	proto.RegisterType((*DestroyCAKResponse)(nil), "thales.cpl.kms.simplekms.v1.DestroyCAKResponse")
	proto.RegisterType((*DestroyCARequest)(nil), "thales.cpl.kms.simplekms.v1.DestroyCARequest")
	proto.RegisterType((*DestroyCAResponse)(nil), "thales.cpl.kms.simplekms.v1.DestroyCAResponse")
	proto.RegisterType((*SignCSRRequest)(nil), "thales.cpl.kms.simplekms.v1.SignCSRRequest")
	proto.RegisterType((*SignCSRResponse)(nil), "thales.cpl.kms.simplekms.v1.SignCSRResponse")
}

func init() { proto.RegisterFile("simplekms/v1/messages.proto", fileDescriptor_245adf7c9b02e35d) }

var fileDescriptor_245adf7c9b02e35d = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x5f, 0x6b, 0xdb, 0x30,
	0x14, 0xc5, 0x97, 0x38, 0x34, 0xab, 0x52, 0x12, 0x57, 0x7b, 0x58, 0x68, 0x61, 0x7f, 0xc2, 0xc6,
	0x4a, 0x57, 0xdb, 0x64, 0x65, 0x50, 0xd8, 0x53, 0x92, 0x86, 0x3d, 0x18, 0x32, 0x70, 0xd8, 0x06,
	0x7b, 0x09, 0x8a, 0x73, 0xe7, 0x0a, 0xdb, 0x92, 0xe7, 0x2b, 0xa7, 0x74, 0x8f, 0xfb, 0xe4, 0xc3,
	0x8a, 0xed, 0x78, 0x81, 0xac, 0x79, 0xca, 0xd5, 0xcd, 0xef, 0x1e, 0x1d, 0x1d, 0x0b, 0x91, 0x73,
	0xe4, 0x71, 0x12, 0x41, 0x18, 0xa3, 0xb3, 0x1e, 0x3a, 0x31, 0x20, 0xb2, 0x00, 0xd0, 0x4e, 0x52,
	0xa9, 0x24, 0x3d, 0x57, 0x77, 0x2c, 0x02, 0xb4, 0xfd, 0x24, 0xb2, 0xc3, 0x18, 0xed, 0x8a, 0xb5,
	0xd7, 0xc3, 0xb3, 0x97, 0x81, 0x94, 0x41, 0x04, 0x8e, 0x46, 0x97, 0xd9, 0x4f, 0x47, 0xf1, 0x18,
	0x50, 0xb1, 0x38, 0xd9, 0x4c, 0x9f, 0x5d, 0xe9, 0x1f, 0xdf, 0x0a, 0x40, 0x58, 0x78, 0xcf, 0x82,
	0x00, 0x52, 0x47, 0x26, 0x8a, 0x4b, 0x81, 0x0e, 0x13, 0x42, 0x2a, 0xa6, 0xeb, 0x0d, 0x3d, 0xb8,
	0x24, 0xdd, 0x6f, 0x90, 0x22, 0x97, 0xc2, 0x83, 0x5f, 0x19, 0xa0, 0xa2, 0x7d, 0xd2, 0x5e, 0x6f,
	0x3a, 0xfd, 0xc6, 0xab, 0xc6, 0xc5, 0xb1, 0x57, 0x2e, 0x07, 0xf7, 0xa4, 0x57, 0xb1, 0x98, 0x48,
	0x81, 0xb0, 0x1f, 0xa6, 0xaf, 0xc9, 0x49, 0x9a, 0x89, 0xdc, 0xdc, 0x42, 0xb0, 0x18, 0xfa, 0x4d,
	0xfd, 0x77, 0xa7, 0xe8, 0xcd, 0x58, 0x0c, 0xf4, 0x1d, 0xe9, 0x95, 0x48, 0x29, 0x62, 0x68, 0xaa,
	0x5b, 0xb4, 0x8b, 0xdd, 0x06, 0x16, 0xa1, 0x9f, 0x41, 0x40, 0xca, 0x14, 0xb8, 0x53, 0xb7, 0x34,
	0xfa, 0x9c, 0xb4, 0x43, 0x08, 0x17, 0x21, 0x5f, 0xe9, 0xbd, 0x4f, 0xbc, 0xa3, 0x10, 0x42, 0x97,
	0xaf, 0x06, 0x36, 0x79, 0xf6, 0x0f, 0x5e, 0x78, 0xdd, 0xcb, 0x5f, 0x91, 0xd3, 0x5b, 0x40, 0x95,
	0xca, 0x87, 0xc3, 0xd4, 0x69, 0x9d, 0xde, 0x06, 0x81, 0x99, 0xef, 0x03, 0xa2, 0xc6, 0x9f, 0x7a,
	0xe5, 0x72, 0xf0, 0xa7, 0xb1, 0x75, 0x3f, 0x19, 0x55, 0xfa, 0x94, 0xb4, 0x90, 0xff, 0x06, 0x4d,
	0x1b, 0x9e, 0xae, 0xe9, 0x0d, 0x69, 0x85, 0x5c, 0xac, 0x74, 0x56, 0xdd, 0x0f, 0x6f, 0xec, 0xff,
	0xdc, 0x03, 0xdb, 0x85, 0x07, 0x97, 0x8b, 0x95, 0xa7, 0x27, 0xe8, 0x0b, 0xd2, 0x49, 0xa5, 0x54,
	0x0b, 0x9f, 0x69, 0xc7, 0x86, 0x76, 0x7c, 0x9c, 0xb7, 0x26, 0x2c, 0x37, 0xfd, 0x71, 0x1b, 0x89,
	0xf6, 0x50, 0xb8, 0xde, 0x19, 0x6b, 0xec, 0x8e, 0x5d, 0x93, 0xd3, 0xed, 0x58, 0xe9, 0xfc, 0xb1,
	0xa1, 0x8b, 0xfa, 0x79, 0xab, 0xad, 0x28, 0x69, 0xf9, 0x90, 0xaa, 0x02, 0xd7, 0x75, 0x2d, 0xf8,
	0x5a, 0x30, 0x07, 0x04, 0x5f, 0x3f, 0xc2, 0xfe, 0xe0, 0xdf, 0x13, 0xb3, 0xe2, 0x1f, 0x15, 0xb7,
	0x6a, 0x56, 0x0e, 0xd0, 0x1e, 0x93, 0xee, 0x9c, 0x07, 0x62, 0x32, 0xf7, 0x4a, 0x65, 0x93, 0x18,
	0x3e, 0xa6, 0x85, 0x6a, 0x5e, 0xee, 0xe6, 0xd4, 0xdc, 0xcd, 0xe9, 0x2d, 0xe9, 0x55, 0x1a, 0xfb,
	0x43, 0xba, 0x1c, 0x92, 0x76, 0xf1, 0xad, 0x69, 0x87, 0xb4, 0xbf, 0xce, 0xdc, 0xd9, 0x97, 0xef,
	0x33, 0xf3, 0x09, 0x6d, 0x13, 0x63, 0x34, 0x9d, 0x9b, 0x8d, 0xbc, 0xf0, 0xe6, 0x23, 0xb3, 0x99,
	0x17, 0xd3, 0xc9, 0xc4, 0x34, 0xc6, 0xb7, 0x3f, 0xc6, 0x01, 0x57, 0x77, 0xd9, 0xd2, 0xf6, 0x65,
	0xec, 0x6c, 0x6e, 0x91, 0x9f, 0x44, 0x16, 0x97, 0x4e, 0x78, 0x83, 0x56, 0x18, 0xa3, 0x95, 0x44,
	0x59, 0xc0, 0x85, 0xc5, 0x12, 0x8e, 0x4e, 0xfd, 0x2d, 0xfa, 0x54, 0x2d, 0x96, 0x47, 0xfa, 0x85,
	0xb8, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x99, 0xd7, 0x38, 0xac, 0x04, 0x00, 0x00,
}
