// Code generated by protoc-gen-go. DO NOT EDIT.
// source: k8s/v1beta1/messages.proto

package k8s

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

type DecryptRequest struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The data to be decrypted.
	Cipher []byte `protobuf:"bytes,2,opt,name=cipher,proto3" json:"cipher,omitempty"`
	// Required. The Keyring that is holding the key to use
	KeyringId string `protobuf:"bytes,3,opt,name=keyring_id,json=keyringId,proto3" json:"keyring_id,omitempty"`
	// Required. The Key to use
	KeyId                string   `protobuf:"bytes,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptRequest) Reset()         { *m = DecryptRequest{} }
func (m *DecryptRequest) String() string { return proto.CompactTextString(m) }
func (*DecryptRequest) ProtoMessage()    {}
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_758226cc6df752dd, []int{0}
}

func (m *DecryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptRequest.Unmarshal(m, b)
}
func (m *DecryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptRequest.Marshal(b, m, deterministic)
}
func (m *DecryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptRequest.Merge(m, src)
}
func (m *DecryptRequest) XXX_Size() int {
	return xxx_messageInfo_DecryptRequest.Size(m)
}
func (m *DecryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptRequest proto.InternalMessageInfo

func (m *DecryptRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DecryptRequest) GetCipher() []byte {
	if m != nil {
		return m.Cipher
	}
	return nil
}

func (m *DecryptRequest) GetKeyringId() string {
	if m != nil {
		return m.KeyringId
	}
	return ""
}

func (m *DecryptRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

type DecryptResponse struct {
	// The decrypted data.
	Plain                []byte   `protobuf:"bytes,1,opt,name=plain,proto3" json:"plain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptResponse) Reset()         { *m = DecryptResponse{} }
func (m *DecryptResponse) String() string { return proto.CompactTextString(m) }
func (*DecryptResponse) ProtoMessage()    {}
func (*DecryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_758226cc6df752dd, []int{1}
}

func (m *DecryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptResponse.Unmarshal(m, b)
}
func (m *DecryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptResponse.Marshal(b, m, deterministic)
}
func (m *DecryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptResponse.Merge(m, src)
}
func (m *DecryptResponse) XXX_Size() int {
	return xxx_messageInfo_DecryptResponse.Size(m)
}
func (m *DecryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptResponse proto.InternalMessageInfo

func (m *DecryptResponse) GetPlain() []byte {
	if m != nil {
		return m.Plain
	}
	return nil
}

type EncryptRequest struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The data to be encrypted.
	Plain []byte `protobuf:"bytes,2,opt,name=plain,proto3" json:"plain,omitempty"`
	// Required. The Keyring that is holding the key to use
	KeyringId string `protobuf:"bytes,3,opt,name=keyring_id,json=keyringId,proto3" json:"keyring_id,omitempty"`
	// Required. The Key to use
	KeyId                string   `protobuf:"bytes,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptRequest) Reset()         { *m = EncryptRequest{} }
func (m *EncryptRequest) String() string { return proto.CompactTextString(m) }
func (*EncryptRequest) ProtoMessage()    {}
func (*EncryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_758226cc6df752dd, []int{2}
}

func (m *EncryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptRequest.Unmarshal(m, b)
}
func (m *EncryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptRequest.Marshal(b, m, deterministic)
}
func (m *EncryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptRequest.Merge(m, src)
}
func (m *EncryptRequest) XXX_Size() int {
	return xxx_messageInfo_EncryptRequest.Size(m)
}
func (m *EncryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptRequest proto.InternalMessageInfo

func (m *EncryptRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *EncryptRequest) GetPlain() []byte {
	if m != nil {
		return m.Plain
	}
	return nil
}

func (m *EncryptRequest) GetKeyringId() string {
	if m != nil {
		return m.KeyringId
	}
	return ""
}

func (m *EncryptRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

type EncryptResponse struct {
	// The encrypted data.
	Cipher               []byte   `protobuf:"bytes,1,opt,name=cipher,proto3" json:"cipher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptResponse) Reset()         { *m = EncryptResponse{} }
func (m *EncryptResponse) String() string { return proto.CompactTextString(m) }
func (*EncryptResponse) ProtoMessage()    {}
func (*EncryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_758226cc6df752dd, []int{3}
}

func (m *EncryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptResponse.Unmarshal(m, b)
}
func (m *EncryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptResponse.Marshal(b, m, deterministic)
}
func (m *EncryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptResponse.Merge(m, src)
}
func (m *EncryptResponse) XXX_Size() int {
	return xxx_messageInfo_EncryptResponse.Size(m)
}
func (m *EncryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptResponse proto.InternalMessageInfo

func (m *EncryptResponse) GetCipher() []byte {
	if m != nil {
		return m.Cipher
	}
	return nil
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
	return fileDescriptor_758226cc6df752dd, []int{4}
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
	return fileDescriptor_758226cc6df752dd, []int{5}
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

func init() {
	proto.RegisterType((*DecryptRequest)(nil), "thales.cpl.kms.k8s.v1beta1.DecryptRequest")
	proto.RegisterType((*DecryptResponse)(nil), "thales.cpl.kms.k8s.v1beta1.DecryptResponse")
	proto.RegisterType((*EncryptRequest)(nil), "thales.cpl.kms.k8s.v1beta1.EncryptRequest")
	proto.RegisterType((*EncryptResponse)(nil), "thales.cpl.kms.k8s.v1beta1.EncryptResponse")
	proto.RegisterType((*VersionRequest)(nil), "thales.cpl.kms.k8s.v1beta1.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "thales.cpl.kms.k8s.v1beta1.VersionResponse")
}

func init() { proto.RegisterFile("k8s/v1beta1/messages.proto", fileDescriptor_758226cc6df752dd) }

var fileDescriptor_758226cc6df752dd = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcf, 0xab, 0xd3, 0x40,
	0x10, 0x26, 0x4f, 0x5f, 0xa5, 0x6b, 0x49, 0x20, 0xa8, 0x84, 0x80, 0xf8, 0xcc, 0xe5, 0x3d, 0xc5,
	0x64, 0x29, 0x5e, 0x02, 0xc5, 0x8b, 0xe8, 0xa1, 0x17, 0x0f, 0x39, 0x78, 0xf0, 0x52, 0x36, 0xe9,
	0xb8, 0x5d, 0x36, 0xfb, 0xc3, 0xcc, 0xa6, 0x35, 0xff, 0xbd, 0x34, 0xd9, 0xc6, 0x5e, 0x44, 0xf1,
	0x94, 0xcc, 0xcc, 0x37, 0xdf, 0xf7, 0xcd, 0xec, 0x90, 0x54, 0x96, 0x48, 0x8f, 0xeb, 0x1a, 0x1c,
	0x5b, 0x53, 0x05, 0x88, 0x8c, 0x03, 0x16, 0xb6, 0x33, 0xce, 0xc4, 0xa9, 0x3b, 0xb0, 0x16, 0xb0,
	0x68, 0x6c, 0x5b, 0x48, 0x85, 0x85, 0x2c, 0xb1, 0xf0, 0xd0, 0xf4, 0x15, 0x37, 0x86, 0xb7, 0x40,
	0x47, 0x64, 0xdd, 0x7f, 0xa7, 0x4e, 0x28, 0x40, 0xc7, 0x94, 0x9d, 0x9a, 0xd3, 0x77, 0xe3, 0xa7,
	0xc9, 0x39, 0xe8, 0x1c, 0x4f, 0x8c, 0x73, 0xe8, 0xa8, 0xb1, 0x4e, 0x18, 0x8d, 0x94, 0x69, 0x6d,
	0x1c, 0x1b, 0xff, 0x27, 0x74, 0xf6, 0x93, 0x84, 0x9f, 0xa0, 0xe9, 0x06, 0xeb, 0x2a, 0xf8, 0xd1,
	0x03, 0xba, 0x38, 0x21, 0x4f, 0x8e, 0xd0, 0xa1, 0x30, 0x3a, 0x09, 0xee, 0x82, 0x87, 0x65, 0x75,
	0x09, 0xe3, 0x17, 0x64, 0xd1, 0x08, 0x7b, 0x80, 0x2e, 0xb9, 0xb9, 0x0b, 0x1e, 0x56, 0x95, 0x8f,
	0xe2, 0x97, 0x84, 0x48, 0x18, 0x3a, 0xa1, 0xf9, 0x4e, 0xec, 0x93, 0x47, 0x63, 0xd3, 0xd2, 0x67,
	0xb6, 0xfb, 0xf8, 0x39, 0x59, 0x48, 0x18, 0xce, 0xa5, 0xc7, 0x63, 0xe9, 0x56, 0xc2, 0xb0, 0xdd,
	0x67, 0xf7, 0x24, 0x9a, 0x95, 0xd1, 0x1a, 0x8d, 0x10, 0x3f, 0x23, 0xb7, 0xb6, 0x65, 0x62, 0x12,
	0x5e, 0x55, 0x53, 0x90, 0x1d, 0x49, 0xf8, 0x59, 0xff, 0xa3, 0xc5, 0x99, 0xe1, 0xe6, 0x8a, 0xe1,
	0x3f, 0x0d, 0xbe, 0x21, 0xd1, 0xac, 0xeb, 0x0d, 0xfe, 0xde, 0x40, 0x70, 0xbd, 0x81, 0xec, 0x2d,
	0x09, 0xbf, 0x4e, 0x0e, 0xfe, 0x6a, 0x31, 0x3b, 0x91, 0x68, 0xc6, 0x7a, 0xda, 0x3f, 0xcf, 0xf3,
	0x9a, 0xac, 0xba, 0x5e, 0x9f, 0x9f, 0x78, 0xa7, 0x99, 0x82, 0x71, 0xac, 0x65, 0xf5, 0xd4, 0xe7,
	0xbe, 0x30, 0x05, 0xf1, 0x3d, 0x89, 0x2e, 0x90, 0x0b, 0xc9, 0x34, 0x61, 0xe8, 0xd3, 0x5e, 0xed,
	0xe3, 0x87, 0x6f, 0x1b, 0x2e, 0xdc, 0xa1, 0xaf, 0x8b, 0xc6, 0x28, 0x3a, 0x9d, 0x58, 0x63, 0xdb,
	0x5c, 0x18, 0x2a, 0x4b, 0xcc, 0xa5, 0xc2, 0xdc, 0xb6, 0x3d, 0x17, 0x3a, 0x67, 0x56, 0x20, 0xbd,
	0xba, 0xcf, 0x8d, 0x2c, 0xb1, 0x5e, 0x8c, 0x07, 0xf3, 0xfe, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x2c, 0x4b, 0x5c, 0x49, 0xb9, 0x02, 0x00, 0x00,
}
