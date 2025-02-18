// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: api/customer/customer.proto

package customer

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetVerifyCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Telephone string `protobuf:"bytes,1,opt,name=telephone,proto3" json:"telephone,omitempty"`
}

func (x *GetVerifyCodeReq) Reset() {
	*x = GetVerifyCodeReq{}
	mi := &file_api_customer_customer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVerifyCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVerifyCodeReq) ProtoMessage() {}

func (x *GetVerifyCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_customer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVerifyCodeReq.ProtoReflect.Descriptor instead.
func (*GetVerifyCodeReq) Descriptor() ([]byte, []int) {
	return file_api_customer_customer_proto_rawDescGZIP(), []int{0}
}

func (x *GetVerifyCodeReq) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

type GetVerifyCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code           int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message        string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	VerifyCode     string `protobuf:"bytes,3,opt,name=verify_code,json=verifyCode,proto3" json:"verify_code,omitempty"`
	VerifyCodeTime int32  `protobuf:"varint,4,opt,name=verify_code_time,json=verifyCodeTime,proto3" json:"verify_code_time,omitempty"`
}

func (x *GetVerifyCodeResp) Reset() {
	*x = GetVerifyCodeResp{}
	mi := &file_api_customer_customer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVerifyCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVerifyCodeResp) ProtoMessage() {}

func (x *GetVerifyCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_customer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVerifyCodeResp.ProtoReflect.Descriptor instead.
func (*GetVerifyCodeResp) Descriptor() ([]byte, []int) {
	return file_api_customer_customer_proto_rawDescGZIP(), []int{1}
}

func (x *GetVerifyCodeResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetVerifyCodeResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetVerifyCodeResp) GetVerifyCode() string {
	if x != nil {
		return x.VerifyCode
	}
	return ""
}

func (x *GetVerifyCodeResp) GetVerifyCodeTime() int32 {
	if x != nil {
		return x.VerifyCodeTime
	}
	return 0
}

var File_api_customer_customer_proto protoreflect.FileDescriptor

var file_api_customer_customer_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x32, 0x9e, 0x01, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x12, 0x91, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x27, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x7b, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x7d, 0x42, 0x3a, 0x0a, 0x15, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x01, 0x5a,
	0x1f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_customer_customer_proto_rawDescOnce sync.Once
	file_api_customer_customer_proto_rawDescData = file_api_customer_customer_proto_rawDesc
)

func file_api_customer_customer_proto_rawDescGZIP() []byte {
	file_api_customer_customer_proto_rawDescOnce.Do(func() {
		file_api_customer_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_customer_customer_proto_rawDescData)
	})
	return file_api_customer_customer_proto_rawDescData
}

var file_api_customer_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_customer_customer_proto_goTypes = []any{
	(*GetVerifyCodeReq)(nil),  // 0: customer.api.customer.GetVerifyCodeReq
	(*GetVerifyCodeResp)(nil), // 1: customer.api.customer.GetVerifyCodeResp
}
var file_api_customer_customer_proto_depIdxs = []int32{
	0, // 0: customer.api.customer.Customer.GetVerifyCode:input_type -> customer.api.customer.GetVerifyCodeReq
	1, // 1: customer.api.customer.Customer.GetVerifyCode:output_type -> customer.api.customer.GetVerifyCodeResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_customer_customer_proto_init() }
func file_api_customer_customer_proto_init() {
	if File_api_customer_customer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_customer_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_customer_customer_proto_goTypes,
		DependencyIndexes: file_api_customer_customer_proto_depIdxs,
		MessageInfos:      file_api_customer_customer_proto_msgTypes,
	}.Build()
	File_api_customer_customer_proto = out.File
	file_api_customer_customer_proto_rawDesc = nil
	file_api_customer_customer_proto_goTypes = nil
	file_api_customer_customer_proto_depIdxs = nil
}
