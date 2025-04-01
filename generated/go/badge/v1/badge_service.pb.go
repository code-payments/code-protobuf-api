// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: badge/v1/badge_service.proto

package badge

import (
	v1 "github.com/code-payments/code-protobuf-api/generated/go/common/v1"
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

type ResetBadgeCountResponse_Result int32

const (
	ResetBadgeCountResponse_OK ResetBadgeCountResponse_Result = 0
)

// Enum value maps for ResetBadgeCountResponse_Result.
var (
	ResetBadgeCountResponse_Result_name = map[int32]string{
		0: "OK",
	}
	ResetBadgeCountResponse_Result_value = map[string]int32{
		"OK": 0,
	}
)

func (x ResetBadgeCountResponse_Result) Enum() *ResetBadgeCountResponse_Result {
	p := new(ResetBadgeCountResponse_Result)
	*p = x
	return p
}

func (x ResetBadgeCountResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResetBadgeCountResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_badge_v1_badge_service_proto_enumTypes[0].Descriptor()
}

func (ResetBadgeCountResponse_Result) Type() protoreflect.EnumType {
	return &file_badge_v1_badge_service_proto_enumTypes[0]
}

func (x ResetBadgeCountResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResetBadgeCountResponse_Result.Descriptor instead.
func (ResetBadgeCountResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_badge_v1_badge_service_proto_rawDescGZIP(), []int{1, 0}
}

type ResetBadgeCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The owner account to clear badge count
	Owner *v1.SolanaAccountId `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// The signature is of serialize(ResetBadgeCountRequest) without this field set
	// using the private key of the owner account. This provides an authentication
	// mechanism to the RPC.
	Signature *v1.Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ResetBadgeCountRequest) Reset() {
	*x = ResetBadgeCountRequest{}
	mi := &file_badge_v1_badge_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResetBadgeCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetBadgeCountRequest) ProtoMessage() {}

func (x *ResetBadgeCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_badge_v1_badge_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetBadgeCountRequest.ProtoReflect.Descriptor instead.
func (*ResetBadgeCountRequest) Descriptor() ([]byte, []int) {
	return file_badge_v1_badge_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResetBadgeCountRequest) GetOwner() *v1.SolanaAccountId {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *ResetBadgeCountRequest) GetSignature() *v1.Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ResetBadgeCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result ResetBadgeCountResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=code.badge.v1.ResetBadgeCountResponse_Result" json:"result,omitempty"`
}

func (x *ResetBadgeCountResponse) Reset() {
	*x = ResetBadgeCountResponse{}
	mi := &file_badge_v1_badge_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResetBadgeCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetBadgeCountResponse) ProtoMessage() {}

func (x *ResetBadgeCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_badge_v1_badge_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetBadgeCountResponse.ProtoReflect.Descriptor instead.
func (*ResetBadgeCountResponse) Descriptor() ([]byte, []int) {
	return file_badge_v1_badge_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResetBadgeCountResponse) GetResult() ResetBadgeCountResponse_Result {
	if x != nil {
		return x.Result
	}
	return ResetBadgeCountResponse_OK
}

var File_badge_v1_badge_service_proto protoreflect.FileDescriptor

var file_badge_v1_badge_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x61, 0x64, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x64, 0x67, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01,
	0x0a, 0x16, 0x52, 0x65, 0x73, 0x65, 0x74, 0x42, 0x61, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x72, 0x0a, 0x17,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x42, 0x61, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x62,
	0x61, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x42, 0x61, 0x64,
	0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x10,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00,
	0x32, 0x69, 0x0a, 0x05, 0x42, 0x61, 0x64, 0x67, 0x65, 0x12, 0x60, 0x0a, 0x0f, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x42, 0x61, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x42, 0x61, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x42, 0x61, 0x64, 0x67, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x0a, 0x18, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x62,
	0x61, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f,
	0x2f, 0x62, 0x61, 0x64, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x61, 0x64, 0x67, 0x65, 0xa2,
	0x02, 0x0a, 0x43, 0x50, 0x42, 0x42, 0x61, 0x64, 0x67, 0x65, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_badge_v1_badge_service_proto_rawDescOnce sync.Once
	file_badge_v1_badge_service_proto_rawDescData = file_badge_v1_badge_service_proto_rawDesc
)

func file_badge_v1_badge_service_proto_rawDescGZIP() []byte {
	file_badge_v1_badge_service_proto_rawDescOnce.Do(func() {
		file_badge_v1_badge_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_badge_v1_badge_service_proto_rawDescData)
	})
	return file_badge_v1_badge_service_proto_rawDescData
}

var file_badge_v1_badge_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_badge_v1_badge_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_badge_v1_badge_service_proto_goTypes = []any{
	(ResetBadgeCountResponse_Result)(0), // 0: code.badge.v1.ResetBadgeCountResponse.Result
	(*ResetBadgeCountRequest)(nil),      // 1: code.badge.v1.ResetBadgeCountRequest
	(*ResetBadgeCountResponse)(nil),     // 2: code.badge.v1.ResetBadgeCountResponse
	(*v1.SolanaAccountId)(nil),          // 3: code.common.v1.SolanaAccountId
	(*v1.Signature)(nil),                // 4: code.common.v1.Signature
}
var file_badge_v1_badge_service_proto_depIdxs = []int32{
	3, // 0: code.badge.v1.ResetBadgeCountRequest.owner:type_name -> code.common.v1.SolanaAccountId
	4, // 1: code.badge.v1.ResetBadgeCountRequest.signature:type_name -> code.common.v1.Signature
	0, // 2: code.badge.v1.ResetBadgeCountResponse.result:type_name -> code.badge.v1.ResetBadgeCountResponse.Result
	1, // 3: code.badge.v1.Badge.ResetBadgeCount:input_type -> code.badge.v1.ResetBadgeCountRequest
	2, // 4: code.badge.v1.Badge.ResetBadgeCount:output_type -> code.badge.v1.ResetBadgeCountResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_badge_v1_badge_service_proto_init() }
func file_badge_v1_badge_service_proto_init() {
	if File_badge_v1_badge_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_badge_v1_badge_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_badge_v1_badge_service_proto_goTypes,
		DependencyIndexes: file_badge_v1_badge_service_proto_depIdxs,
		EnumInfos:         file_badge_v1_badge_service_proto_enumTypes,
		MessageInfos:      file_badge_v1_badge_service_proto_msgTypes,
	}.Build()
	File_badge_v1_badge_service_proto = out.File
	file_badge_v1_badge_service_proto_rawDesc = nil
	file_badge_v1_badge_service_proto_goTypes = nil
	file_badge_v1_badge_service_proto_depIdxs = nil
}
