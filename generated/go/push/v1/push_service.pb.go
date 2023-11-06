// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: push/v1/push_service.proto

package push

import (
	v1 "github.com/code-wallet/code-protobuf-api/generated/go/common/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenType int32

const (
	TokenType_UNKNOWN TokenType = 0
	// FCM registration token for an Android device
	TokenType_FCM_ANDROID TokenType = 1
	// FCM registration token or an iOS device
	TokenType_FCM_APNS TokenType = 2
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0: "UNKNOWN",
		1: "FCM_ANDROID",
		2: "FCM_APNS",
	}
	TokenType_value = map[string]int32{
		"UNKNOWN":     0,
		"FCM_ANDROID": 1,
		"FCM_APNS":    2,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_push_v1_push_service_proto_enumTypes[0].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_push_v1_push_service_proto_enumTypes[0]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{0}
}

type AddTokenResponse_Result int32

const (
	AddTokenResponse_OK AddTokenResponse_Result = 0
	// The push token is invalid and wasn't stored.
	AddTokenResponse_INVALID_PUSH_TOKEN AddTokenResponse_Result = 1
)

// Enum value maps for AddTokenResponse_Result.
var (
	AddTokenResponse_Result_name = map[int32]string{
		0: "OK",
		1: "INVALID_PUSH_TOKEN",
	}
	AddTokenResponse_Result_value = map[string]int32{
		"OK":                 0,
		"INVALID_PUSH_TOKEN": 1,
	}
)

func (x AddTokenResponse_Result) Enum() *AddTokenResponse_Result {
	p := new(AddTokenResponse_Result)
	*p = x
	return p
}

func (x AddTokenResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddTokenResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_push_v1_push_service_proto_enumTypes[1].Descriptor()
}

func (AddTokenResponse_Result) Type() protoreflect.EnumType {
	return &file_push_v1_push_service_proto_enumTypes[1]
}

func (x AddTokenResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddTokenResponse_Result.Descriptor instead.
func (AddTokenResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{1, 0}
}

type AddTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The public key of the owner account that signed this request message.
	OwnerAccountId *v1.SolanaAccountId `protobuf:"bytes,1,opt,name=owner_account_id,json=ownerAccountId,proto3" json:"owner_account_id,omitempty"`
	// The signature is of serialize(AddTokenRequest) without this field set
	// using the private key of owner_account_id. This provides an authentication
	// mechanism to the RPC.
	Signature *v1.Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// The data container where the push token will be stored.
	ContainerId *v1.DataContainerId `protobuf:"bytes,3,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	// The push token to store
	PushToken string `protobuf:"bytes,4,opt,name=push_token,json=pushToken,proto3" json:"push_token,omitempty"`
	// The type of push token
	TokenType TokenType `protobuf:"varint,5,opt,name=token_type,json=tokenType,proto3,enum=code.push.v1.TokenType" json:"token_type,omitempty"`
	// The instance of the app install where the push token was generated. Ideally,
	// the push token is unique to the install.
	AppInstall *v1.AppInstallId `protobuf:"bytes,6,opt,name=app_install,json=appInstall,proto3" json:"app_install,omitempty"`
}

func (x *AddTokenRequest) Reset() {
	*x = AddTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_v1_push_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTokenRequest) ProtoMessage() {}

func (x *AddTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_push_v1_push_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTokenRequest.ProtoReflect.Descriptor instead.
func (*AddTokenRequest) Descriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddTokenRequest) GetOwnerAccountId() *v1.SolanaAccountId {
	if x != nil {
		return x.OwnerAccountId
	}
	return nil
}

func (x *AddTokenRequest) GetSignature() *v1.Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *AddTokenRequest) GetContainerId() *v1.DataContainerId {
	if x != nil {
		return x.ContainerId
	}
	return nil
}

func (x *AddTokenRequest) GetPushToken() string {
	if x != nil {
		return x.PushToken
	}
	return ""
}

func (x *AddTokenRequest) GetTokenType() TokenType {
	if x != nil {
		return x.TokenType
	}
	return TokenType_UNKNOWN
}

func (x *AddTokenRequest) GetAppInstall() *v1.AppInstallId {
	if x != nil {
		return x.AppInstall
	}
	return nil
}

type AddTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result AddTokenResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=code.push.v1.AddTokenResponse_Result" json:"result,omitempty"`
}

func (x *AddTokenResponse) Reset() {
	*x = AddTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_v1_push_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTokenResponse) ProtoMessage() {}

func (x *AddTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_push_v1_push_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTokenResponse.ProtoReflect.Descriptor instead.
func (*AddTokenResponse) Descriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddTokenResponse) GetResult() AddTokenResponse_Result {
	if x != nil {
		return x.Result
	}
	return AddTokenResponse_OK
}

var File_push_v1_push_service_proto protoreflect.FileDescriptor

var file_push_v1_push_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x03, 0x0a, 0x0f, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x55,
	0x0a, 0x10, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x6c, 0x61, 0x6e, 0x61,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x4e, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0a, 0x70, 0x75,
	0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xba, 0xe9, 0xc0, 0x03, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0x80, 0x20, 0x52, 0x09, 0x70, 0x75,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x44, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x0c, 0xba, 0xe9, 0xc0, 0x03, 0x07, 0x82, 0x01, 0x04, 0x18, 0x01,
	0x18, 0x02, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a,
	0x0b, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x49, 0x64,
	0x52, 0x0a, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x22, 0x7b, 0x0a, 0x10,
	0x41, 0x64, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x25, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x28, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x55, 0x53,
	0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x01, 0x2a, 0x37, 0x0a, 0x09, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x43, 0x4d, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f,
	0x49, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x43, 0x4d, 0x5f, 0x41, 0x50, 0x4e, 0x53,
	0x10, 0x02, 0x32, 0x51, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x49, 0x0a, 0x08, 0x41, 0x64,
	0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x75,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x75, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x69, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x69, 0x6e, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x2d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x75, 0x73, 0x68, 0xa2, 0x02, 0x09, 0x41, 0x50, 0x42, 0x50, 0x75, 0x73, 0x68, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_push_v1_push_service_proto_rawDescOnce sync.Once
	file_push_v1_push_service_proto_rawDescData = file_push_v1_push_service_proto_rawDesc
)

func file_push_v1_push_service_proto_rawDescGZIP() []byte {
	file_push_v1_push_service_proto_rawDescOnce.Do(func() {
		file_push_v1_push_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_push_v1_push_service_proto_rawDescData)
	})
	return file_push_v1_push_service_proto_rawDescData
}

var file_push_v1_push_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_push_v1_push_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_push_v1_push_service_proto_goTypes = []interface{}{
	(TokenType)(0),               // 0: code.push.v1.TokenType
	(AddTokenResponse_Result)(0), // 1: code.push.v1.AddTokenResponse.Result
	(*AddTokenRequest)(nil),      // 2: code.push.v1.AddTokenRequest
	(*AddTokenResponse)(nil),     // 3: code.push.v1.AddTokenResponse
	(*v1.SolanaAccountId)(nil),   // 4: code.common.v1.SolanaAccountId
	(*v1.Signature)(nil),         // 5: code.common.v1.Signature
	(*v1.DataContainerId)(nil),   // 6: code.common.v1.DataContainerId
	(*v1.AppInstallId)(nil),      // 7: code.common.v1.AppInstallId
}
var file_push_v1_push_service_proto_depIdxs = []int32{
	4, // 0: code.push.v1.AddTokenRequest.owner_account_id:type_name -> code.common.v1.SolanaAccountId
	5, // 1: code.push.v1.AddTokenRequest.signature:type_name -> code.common.v1.Signature
	6, // 2: code.push.v1.AddTokenRequest.container_id:type_name -> code.common.v1.DataContainerId
	0, // 3: code.push.v1.AddTokenRequest.token_type:type_name -> code.push.v1.TokenType
	7, // 4: code.push.v1.AddTokenRequest.app_install:type_name -> code.common.v1.AppInstallId
	1, // 5: code.push.v1.AddTokenResponse.result:type_name -> code.push.v1.AddTokenResponse.Result
	2, // 6: code.push.v1.Push.AddToken:input_type -> code.push.v1.AddTokenRequest
	3, // 7: code.push.v1.Push.AddToken:output_type -> code.push.v1.AddTokenResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_push_v1_push_service_proto_init() }
func file_push_v1_push_service_proto_init() {
	if File_push_v1_push_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_push_v1_push_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_push_v1_push_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTokenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_push_v1_push_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_push_v1_push_service_proto_goTypes,
		DependencyIndexes: file_push_v1_push_service_proto_depIdxs,
		EnumInfos:         file_push_v1_push_service_proto_enumTypes,
		MessageInfos:      file_push_v1_push_service_proto_msgTypes,
	}.Build()
	File_push_v1_push_service_proto = out.File
	file_push_v1_push_service_proto_rawDesc = nil
	file_push_v1_push_service_proto_goTypes = nil
	file_push_v1_push_service_proto_depIdxs = nil
}
