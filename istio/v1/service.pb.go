// Code generated by protoc-gen-go. DO NOT EDIT.
// source: istio/v1/service.proto

package istio

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

func init() { proto.RegisterFile("istio/v1/service.proto", fileDescriptor_b27436ea3ecfaad9) }

var fileDescriptor_b27436ea3ecfaad9 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc7, 0x79, 0x2e, 0x8f, 0x12, 0x05, 0x21, 0x14, 0x0b, 0xbd, 0x78, 0x55, 0x34, 0xbb, 0xd6,
	0x17, 0x10, 0xf5, 0x52, 0x9b, 0x22, 0x12, 0xbd, 0x58, 0xf0, 0xe0, 0x45, 0xb6, 0xe9, 0x34, 0x59,
	0x92, 0x7d, 0x31, 0x3b, 0x29, 0xe4, 0xec, 0x87, 0xf1, 0x6b, 0x4a, 0xf3, 0x52, 0x5b, 0x31, 0x36,
	0xb9, 0x2d, 0xb3, 0xbf, 0xff, 0xfc, 0x06, 0x86, 0xb1, 0xf6, 0xb9, 0x41, 0xae, 0xe8, 0xbc, 0x4f,
	0x0d, 0x24, 0x73, 0xee, 0x03, 0xd1, 0x89, 0x42, 0x65, 0x77, 0x31, 0x64, 0x31, 0x18, 0xe2, 0xeb,
	0x98, 0x44, 0xc2, 0x90, 0x1c, 0x23, 0xf3, 0x7e, 0xef, 0x20, 0x50, 0x2a, 0x88, 0x81, 0xe6, 0xd8,
	0x24, 0x9d, 0x51, 0xe4, 0x02, 0x0c, 0x32, 0xa1, 0x8b, 0x64, 0xaf, 0xbb, 0xec, 0x28, 0xc0, 0x18,
	0x16, 0x80, 0x29, 0x3e, 0xce, 0x3e, 0xb7, 0xac, 0x8e, 0x07, 0xd9, 0x13, 0x93, 0x2c, 0x00, 0x01,
	0x12, 0xc7, 0x85, 0xd1, 0x0e, 0xad, 0x9d, 0x7b, 0x90, 0x90, 0x30, 0x04, 0x6f, 0xe4, 0xd9, 0xc7,
	0xa4, 0xc6, 0x4d, 0x56, 0xa8, 0x67, 0x78, 0x4f, 0xc1, 0x60, 0xef, 0xa4, 0x19, 0x6c, 0xb4, 0x92,
	0x66, 0xcd, 0xe4, 0x36, 0x32, 0xb9, 0x6d, 0x4c, 0xee, 0x8a, 0x29, 0xb2, 0x76, 0xab, 0xf2, 0xd8,
	0x83, 0xcc, 0xde, 0x9c, 0x5e, 0x60, 0x95, 0xcb, 0x69, 0x48, 0x97, 0xb2, 0x37, 0x6b, 0xfb, 0x51,
	0xb1, 0x69, 0x2e, 0x3a, 0xac, 0x8d, 0x56, 0x48, 0x25, 0x39, 0x6a, 0x40, 0x96, 0x82, 0x8f, 0x7f,
	0x56, 0x67, 0x90, 0x62, 0x08, 0x12, 0xb9, 0xcf, 0x10, 0xa6, 0x23, 0xe9, 0x27, 0x99, 0x46, 0xfb,
	0xa2, 0xb6, 0xc7, 0x6f, 0x78, 0x65, 0xbe, 0x6c, 0x99, 0xaa, 0x9b, 0xc2, 0x85, 0x56, 0x53, 0x94,
	0x78, 0xcb, 0x29, 0x96, 0xa9, 0xef, 0xcd, 0x3e, 0x08, 0xad, 0x12, 0x1c, 0x0e, 0x86, 0x90, 0xe0,
	0x1f, 0x9b, 0x5d, 0xc5, 0x36, 0x6f, 0x76, 0x9d, 0x2e, 0x65, 0x68, 0xed, 0xbd, 0x40, 0xc2, 0x67,
	0xd9, 0xa2, 0x3a, 0x0c, 0x19, 0x97, 0x36, 0xad, 0xed, 0xf0, 0x83, 0xac, 0x94, 0xa7, 0xcd, 0x03,
	0x85, 0xf5, 0xee, 0xf6, 0xf5, 0x3a, 0xe0, 0x18, 0xa6, 0x13, 0xe2, 0x2b, 0x41, 0x8b, 0xb4, 0xaf,
	0x63, 0x87, 0x2b, 0x1a, 0x5d, 0x19, 0x27, 0x12, 0xc6, 0xd1, 0x71, 0x1a, 0x70, 0xe9, 0x30, 0xcd,
	0x0d, 0xad, 0x0e, 0xfe, 0x26, 0x7f, 0x4c, 0xfe, 0xe7, 0xe7, 0x7e, 0xfe, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x5f, 0xff, 0x17, 0x47, 0x5b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyManagementServiceClient is the client API for KeyManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyManagementServiceClient interface {
	// GenerateKEK returns the KID of the GeneratedKEK if allowed/successful
	GenerateKEK(ctx context.Context, in *GenerateKEKRequest, opts ...grpc.CallOption) (*GenerateKEKResponse, error)
	// GenerateDEK returns a wrapped (by HSM handled KEK)
	GenerateDEK(ctx context.Context, in *GenerateDEKRequest, opts ...grpc.CallOption) (*GenerateDEKResponse, error)
	// GenerateSKey returns a wrapped (by provided encrypted DEK ), for later use during loading and signing key generation
	GenerateSKey(ctx context.Context, in *GenerateSKeyRequest, opts ...grpc.CallOption) (*GenerateSKeyResponse, error)
	// LoadSKey returns the SKey unwrapped for the controller to use for CA work...
	LoadSKey(ctx context.Context, in *LoadSKeyRequest, opts ...grpc.CallOption) (*LoadSKeyResponse, error)
	// AuthenticatedEncrypt returns a payload protect by AEAD (Authenticated Encryption with Associated Data) AES-GCM with 256-bit
	AuthenticatedEncrypt(ctx context.Context, in *AuthenticatedEncryptRequest, opts ...grpc.CallOption) (*AuthenticatedEncryptResponse, error)
	// AuthenticatedDecrypt returns a payload decrypted from the previously invoke AE
	AuthenticatedDecrypt(ctx context.Context, in *AuthenticatedDecryptRequest, opts ...grpc.CallOption) (*AuthenticatedDecryptResponse, error)
	// ImportCACert imports the supplied CA Cert as the Offline Root CA Cert Chain
	ImportCACert(ctx context.Context, in *ImportCACertRequest, opts ...grpc.CallOption) (*ImportCACertResponse, error)
	// VerifyCertChain returns a boolean indicating successful verification of a certificate chain (currently needs to be self-contained)
	VerifyCertChain(ctx context.Context, in *VerifyCertChainRequest, opts ...grpc.CallOption) (*VerifyCertChainResponse, error)
}

type keyManagementServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeyManagementServiceClient(cc *grpc.ClientConn) KeyManagementServiceClient {
	return &keyManagementServiceClient{cc}
}

func (c *keyManagementServiceClient) GenerateKEK(ctx context.Context, in *GenerateKEKRequest, opts ...grpc.CallOption) (*GenerateKEKResponse, error) {
	out := new(GenerateKEKResponse)
	err := c.cc.Invoke(ctx, "/thales.cpl.kms.istio.v1.KeyManagementService/GenerateKEK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) GenerateDEK(ctx context.Context, in *GenerateDEKRequest, opts ...grpc.CallOption) (*GenerateDEKResponse, error) {
	out := new(GenerateDEKResponse)
	err := c.cc.Invoke(ctx, "/thales.cpl.kms.istio.v1.KeyManagementService/GenerateDEK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) GenerateSKey(ctx context.Context, in *GenerateSKeyRequest, opts ...grpc.CallOption) (*GenerateSKeyResponse, error) {
	out := new(GenerateSKeyResponse)
	err := c.cc.Invoke(ctx, "/thales.cpl.kms.istio.v1.KeyManagementService/GenerateSKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) LoadSKey(ctx context.Context, in *LoadSKeyRequest, opts ...grpc.CallOption) (*LoadSKeyResponse, error) {
	out := new(LoadSKeyResponse)
	err := c.cc.Invoke(ctx, "/thales.cpl.kms.istio.v1.KeyManagementService/LoadSKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) AuthenticatedEncrypt(ctx context.Context, in *AuthenticatedEncryptRequest, opts ...grpc.CallOption) (*AuthenticatedEncryptResponse, error) {
	out := new(AuthenticatedEncryptResponse)
	err := c.cc.Invoke(ctx, "/thales.cpl.kms.istio.v1.KeyManagementService/AuthenticatedEncrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) AuthenticatedDecrypt(ctx context.Context, in *AuthenticatedDecryptRequest, opts ...grpc.CallOption) (*AuthenticatedDecryptResponse, error) {
	out := new(AuthenticatedDecryptResponse)
	err := c.cc.Invoke(ctx, "/thales.cpl.kms.istio.v1.KeyManagementService/AuthenticatedDecrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) ImportCACert(ctx context.Context, in *ImportCACertRequest, opts ...grpc.CallOption) (*ImportCACertResponse, error) {
	out := new(ImportCACertResponse)
	err := c.cc.Invoke(ctx, "/thales.cpl.kms.istio.v1.KeyManagementService/ImportCACert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) VerifyCertChain(ctx context.Context, in *VerifyCertChainRequest, opts ...grpc.CallOption) (*VerifyCertChainResponse, error) {
	out := new(VerifyCertChainResponse)
	err := c.cc.Invoke(ctx, "/thales.cpl.kms.istio.v1.KeyManagementService/VerifyCertChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyManagementServiceServer is the server API for KeyManagementService service.
type KeyManagementServiceServer interface {
	// GenerateKEK returns the KID of the GeneratedKEK if allowed/successful
	GenerateKEK(context.Context, *GenerateKEKRequest) (*GenerateKEKResponse, error)
	// GenerateDEK returns a wrapped (by HSM handled KEK)
	GenerateDEK(context.Context, *GenerateDEKRequest) (*GenerateDEKResponse, error)
	// GenerateSKey returns a wrapped (by provided encrypted DEK ), for later use during loading and signing key generation
	GenerateSKey(context.Context, *GenerateSKeyRequest) (*GenerateSKeyResponse, error)
	// LoadSKey returns the SKey unwrapped for the controller to use for CA work...
	LoadSKey(context.Context, *LoadSKeyRequest) (*LoadSKeyResponse, error)
	// AuthenticatedEncrypt returns a payload protect by AEAD (Authenticated Encryption with Associated Data) AES-GCM with 256-bit
	AuthenticatedEncrypt(context.Context, *AuthenticatedEncryptRequest) (*AuthenticatedEncryptResponse, error)
	// AuthenticatedDecrypt returns a payload decrypted from the previously invoke AE
	AuthenticatedDecrypt(context.Context, *AuthenticatedDecryptRequest) (*AuthenticatedDecryptResponse, error)
	// ImportCACert imports the supplied CA Cert as the Offline Root CA Cert Chain
	ImportCACert(context.Context, *ImportCACertRequest) (*ImportCACertResponse, error)
	// VerifyCertChain returns a boolean indicating successful verification of a certificate chain (currently needs to be self-contained)
	VerifyCertChain(context.Context, *VerifyCertChainRequest) (*VerifyCertChainResponse, error)
}

// UnimplementedKeyManagementServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeyManagementServiceServer struct {
}

func (*UnimplementedKeyManagementServiceServer) GenerateKEK(ctx context.Context, req *GenerateKEKRequest) (*GenerateKEKResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKEK not implemented")
}
func (*UnimplementedKeyManagementServiceServer) GenerateDEK(ctx context.Context, req *GenerateDEKRequest) (*GenerateDEKResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateDEK not implemented")
}
func (*UnimplementedKeyManagementServiceServer) GenerateSKey(ctx context.Context, req *GenerateSKeyRequest) (*GenerateSKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateSKey not implemented")
}
func (*UnimplementedKeyManagementServiceServer) LoadSKey(ctx context.Context, req *LoadSKeyRequest) (*LoadSKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadSKey not implemented")
}
func (*UnimplementedKeyManagementServiceServer) AuthenticatedEncrypt(ctx context.Context, req *AuthenticatedEncryptRequest) (*AuthenticatedEncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticatedEncrypt not implemented")
}
func (*UnimplementedKeyManagementServiceServer) AuthenticatedDecrypt(ctx context.Context, req *AuthenticatedDecryptRequest) (*AuthenticatedDecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticatedDecrypt not implemented")
}
func (*UnimplementedKeyManagementServiceServer) ImportCACert(ctx context.Context, req *ImportCACertRequest) (*ImportCACertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportCACert not implemented")
}
func (*UnimplementedKeyManagementServiceServer) VerifyCertChain(ctx context.Context, req *VerifyCertChainRequest) (*VerifyCertChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCertChain not implemented")
}

func RegisterKeyManagementServiceServer(s *grpc.Server, srv KeyManagementServiceServer) {
	s.RegisterService(&_KeyManagementService_serviceDesc, srv)
}

func _KeyManagementService_GenerateKEK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKEKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).GenerateKEK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thales.cpl.kms.istio.v1.KeyManagementService/GenerateKEK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).GenerateKEK(ctx, req.(*GenerateKEKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_GenerateDEK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateDEKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).GenerateDEK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thales.cpl.kms.istio.v1.KeyManagementService/GenerateDEK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).GenerateDEK(ctx, req.(*GenerateDEKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_GenerateSKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateSKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).GenerateSKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thales.cpl.kms.istio.v1.KeyManagementService/GenerateSKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).GenerateSKey(ctx, req.(*GenerateSKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_LoadSKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadSKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).LoadSKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thales.cpl.kms.istio.v1.KeyManagementService/LoadSKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).LoadSKey(ctx, req.(*LoadSKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_AuthenticatedEncrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticatedEncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).AuthenticatedEncrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thales.cpl.kms.istio.v1.KeyManagementService/AuthenticatedEncrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).AuthenticatedEncrypt(ctx, req.(*AuthenticatedEncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_AuthenticatedDecrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticatedDecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).AuthenticatedDecrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thales.cpl.kms.istio.v1.KeyManagementService/AuthenticatedDecrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).AuthenticatedDecrypt(ctx, req.(*AuthenticatedDecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_ImportCACert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportCACertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).ImportCACert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thales.cpl.kms.istio.v1.KeyManagementService/ImportCACert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).ImportCACert(ctx, req.(*ImportCACertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_VerifyCertChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCertChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).VerifyCertChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thales.cpl.kms.istio.v1.KeyManagementService/VerifyCertChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).VerifyCertChain(ctx, req.(*VerifyCertChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thales.cpl.kms.istio.v1.KeyManagementService",
	HandlerType: (*KeyManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKEK",
			Handler:    _KeyManagementService_GenerateKEK_Handler,
		},
		{
			MethodName: "GenerateDEK",
			Handler:    _KeyManagementService_GenerateDEK_Handler,
		},
		{
			MethodName: "GenerateSKey",
			Handler:    _KeyManagementService_GenerateSKey_Handler,
		},
		{
			MethodName: "LoadSKey",
			Handler:    _KeyManagementService_LoadSKey_Handler,
		},
		{
			MethodName: "AuthenticatedEncrypt",
			Handler:    _KeyManagementService_AuthenticatedEncrypt_Handler,
		},
		{
			MethodName: "AuthenticatedDecrypt",
			Handler:    _KeyManagementService_AuthenticatedDecrypt_Handler,
		},
		{
			MethodName: "ImportCACert",
			Handler:    _KeyManagementService_ImportCACert_Handler,
		},
		{
			MethodName: "VerifyCertChain",
			Handler:    _KeyManagementService_VerifyCertChain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "istio/v1/service.proto",
}
