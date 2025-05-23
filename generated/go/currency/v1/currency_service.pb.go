// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: currency/v1/currency_service.proto

package currency

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAllRatesResponse_Result int32

const (
	GetAllRatesResponse_OK GetAllRatesResponse_Result = 0
	// No currency data is available for the requested timestamp.
	GetAllRatesResponse_MISSING_DATA GetAllRatesResponse_Result = 1
)

// Enum value maps for GetAllRatesResponse_Result.
var (
	GetAllRatesResponse_Result_name = map[int32]string{
		0: "OK",
		1: "MISSING_DATA",
	}
	GetAllRatesResponse_Result_value = map[string]int32{
		"OK":           0,
		"MISSING_DATA": 1,
	}
)

func (x GetAllRatesResponse_Result) Enum() *GetAllRatesResponse_Result {
	p := new(GetAllRatesResponse_Result)
	*p = x
	return p
}

func (x GetAllRatesResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetAllRatesResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_currency_v1_currency_service_proto_enumTypes[0].Descriptor()
}

func (GetAllRatesResponse_Result) Type() protoreflect.EnumType {
	return &file_currency_v1_currency_service_proto_enumTypes[0]
}

func (x GetAllRatesResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetAllRatesResponse_Result.Descriptor instead.
func (GetAllRatesResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_currency_v1_currency_service_proto_rawDescGZIP(), []int{1, 0}
}

type GetAllRatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If timestamp is included, the returned rate will be the most recent available
	// exchange rate prior to the provided timestamp within the same day. Otherwise,
	// the latest rates will be returned.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetAllRatesRequest) Reset() {
	*x = GetAllRatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_v1_currency_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRatesRequest) ProtoMessage() {}

func (x *GetAllRatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_currency_v1_currency_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRatesRequest.ProtoReflect.Descriptor instead.
func (*GetAllRatesRequest) Descriptor() ([]byte, []int) {
	return file_currency_v1_currency_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllRatesRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type GetAllRatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result GetAllRatesResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=code.currency.v1.GetAllRatesResponse_Result" json:"result,omitempty"`
	// The time the exchange rates were observed
	AsOf *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=as_of,json=asOf,proto3" json:"as_of,omitempty"`
	// The price of 1 core mint token in different currencies, keyed on 3- or 4-
	// letter lowercase currency code.
	Rates map[string]float64 `protobuf:"bytes,3,rep,name=rates,proto3" json:"rates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *GetAllRatesResponse) Reset() {
	*x = GetAllRatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_v1_currency_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRatesResponse) ProtoMessage() {}

func (x *GetAllRatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_currency_v1_currency_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRatesResponse.ProtoReflect.Descriptor instead.
func (*GetAllRatesResponse) Descriptor() ([]byte, []int) {
	return file_currency_v1_currency_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllRatesResponse) GetResult() GetAllRatesResponse_Result {
	if x != nil {
		return x.Result
	}
	return GetAllRatesResponse_OK
}

func (x *GetAllRatesResponse) GetAsOf() *timestamppb.Timestamp {
	if x != nil {
		return x.AsOf
	}
	return nil
}

func (x *GetAllRatesResponse) GetRates() map[string]float64 {
	if x != nil {
		return x.Rates
	}
	return nil
}

var File_currency_v1_currency_service_proto protoreflect.FileDescriptor

var file_currency_v1_currency_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0xda, 0x02, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b,
	0x0a, 0x05, 0x61, 0x73, 0x5f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05,
	0xb2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x04, 0x61, 0x73, 0x4f, 0x66, 0x12, 0x62, 0x0a, 0x05, 0x72,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x1a, 0xba, 0xe9,
	0xc0, 0x03, 0x15, 0x9a, 0x01, 0x12, 0x22, 0x10, 0x72, 0x0e, 0x32, 0x0c, 0x5e, 0x5b, 0x61, 0x2d,
	0x7a, 0x5d, 0x7b, 0x33, 0x2c, 0x34, 0x7d, 0x24, 0x52, 0x05, 0x72, 0x61, 0x74, 0x65, 0x73, 0x1a,
	0x38, 0x0a, 0x0a, 0x52, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x22, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x01, 0x32, 0x66, 0x0a,
	0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x5a, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x7b, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x69, 0x6e, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x76, 0x31, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0xa2, 0x02, 0x0d, 0x43, 0x50, 0x42, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_currency_v1_currency_service_proto_rawDescOnce sync.Once
	file_currency_v1_currency_service_proto_rawDescData = file_currency_v1_currency_service_proto_rawDesc
)

func file_currency_v1_currency_service_proto_rawDescGZIP() []byte {
	file_currency_v1_currency_service_proto_rawDescOnce.Do(func() {
		file_currency_v1_currency_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_currency_v1_currency_service_proto_rawDescData)
	})
	return file_currency_v1_currency_service_proto_rawDescData
}

var file_currency_v1_currency_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_currency_v1_currency_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_currency_v1_currency_service_proto_goTypes = []interface{}{
	(GetAllRatesResponse_Result)(0), // 0: code.currency.v1.GetAllRatesResponse.Result
	(*GetAllRatesRequest)(nil),      // 1: code.currency.v1.GetAllRatesRequest
	(*GetAllRatesResponse)(nil),     // 2: code.currency.v1.GetAllRatesResponse
	nil,                             // 3: code.currency.v1.GetAllRatesResponse.RatesEntry
	(*timestamppb.Timestamp)(nil),   // 4: google.protobuf.Timestamp
}
var file_currency_v1_currency_service_proto_depIdxs = []int32{
	4, // 0: code.currency.v1.GetAllRatesRequest.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: code.currency.v1.GetAllRatesResponse.result:type_name -> code.currency.v1.GetAllRatesResponse.Result
	4, // 2: code.currency.v1.GetAllRatesResponse.as_of:type_name -> google.protobuf.Timestamp
	3, // 3: code.currency.v1.GetAllRatesResponse.rates:type_name -> code.currency.v1.GetAllRatesResponse.RatesEntry
	1, // 4: code.currency.v1.Currency.GetAllRates:input_type -> code.currency.v1.GetAllRatesRequest
	2, // 5: code.currency.v1.Currency.GetAllRates:output_type -> code.currency.v1.GetAllRatesResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_currency_v1_currency_service_proto_init() }
func file_currency_v1_currency_service_proto_init() {
	if File_currency_v1_currency_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_currency_v1_currency_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRatesRequest); i {
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
		file_currency_v1_currency_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRatesResponse); i {
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
			RawDescriptor: file_currency_v1_currency_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_currency_v1_currency_service_proto_goTypes,
		DependencyIndexes: file_currency_v1_currency_service_proto_depIdxs,
		EnumInfos:         file_currency_v1_currency_service_proto_enumTypes,
		MessageInfos:      file_currency_v1_currency_service_proto_msgTypes,
	}.Build()
	File_currency_v1_currency_service_proto = out.File
	file_currency_v1_currency_service_proto_rawDesc = nil
	file_currency_v1_currency_service_proto_goTypes = nil
	file_currency_v1_currency_service_proto_depIdxs = nil
}
