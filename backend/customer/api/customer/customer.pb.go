// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
<<<<<<< HEAD
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
=======
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
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

<<<<<<< HEAD
=======
type EstimatePriceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin      string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *EstimatePriceReq) Reset() {
	*x = EstimatePriceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_customer_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstimatePriceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimatePriceReq) ProtoMessage() {}

func (x *EstimatePriceReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstimatePriceReq.ProtoReflect.Descriptor instead.
func (*EstimatePriceReq) Descriptor() ([]byte, []int) {
	return file_api_customer_customer_proto_rawDescGZIP(), []int{0}
}

func (x *EstimatePriceReq) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *EstimatePriceReq) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type EstimatePriceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message     string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Origin      string `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination string `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	Price       int64  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *EstimatePriceResp) Reset() {
	*x = EstimatePriceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_customer_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstimatePriceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimatePriceResp) ProtoMessage() {}

func (x *EstimatePriceResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstimatePriceResp.ProtoReflect.Descriptor instead.
func (*EstimatePriceResp) Descriptor() ([]byte, []int) {
	return file_api_customer_customer_proto_rawDescGZIP(), []int{1}
}

func (x *EstimatePriceResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *EstimatePriceResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EstimatePriceResp) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *EstimatePriceResp) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *EstimatePriceResp) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type LogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutReq) Reset() {
	*x = LogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_customer_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReq) ProtoMessage() {}

func (x *LogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReq.ProtoReflect.Descriptor instead.
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return file_api_customer_customer_proto_rawDescGZIP(), []int{2}
}

type LogoutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LogoutResp) Reset() {
	*x = LogoutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_customer_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResp) ProtoMessage() {}

func (x *LogoutResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResp.ProtoReflect.Descriptor instead.
func (*LogoutResp) Descriptor() ([]byte, []int) {
	return file_api_customer_customer_proto_rawDescGZIP(), []int{3}
}

func (x *LogoutResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LogoutResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// 登录的消息
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Telephone  string `protobuf:"bytes,1,opt,name=telephone,proto3" json:"telephone,omitempty"`
	VerifyCode string `protobuf:"bytes,2,opt,name=verify_code,json=verifyCode,proto3" json:"verify_code,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_customer_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_api_customer_customer_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReq) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *LoginReq) GetVerifyCode() string {
	if x != nil {
		return x.VerifyCode
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// token,登录表示，特殊的字符串，JWT 编码格式
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	// 生成时间 unix timestamp
	TokenCreateAt int64 `protobuf:"varint,4,opt,name=token_create_at,json=tokenCreateAt,proto3" json:"token_create_at,omitempty"`
	// 有效期，单位 second
	TokenLife int32 `protobuf:"varint,5,opt,name=token_life,json=tokenLife,proto3" json:"token_life,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_customer_customer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_customer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_api_customer_customer_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResp) GetTokenCreateAt() int64 {
	if x != nil {
		return x.TokenCreateAt
	}
	return 0
}

func (x *LoginResp) GetTokenLife() int32 {
	if x != nil {
		return x.TokenLife
	}
	return 0
}

// 获取验证码的消息
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
type GetVerifyCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Telephone string `protobuf:"bytes,1,opt,name=telephone,proto3" json:"telephone,omitempty"`
}

func (x *GetVerifyCodeReq) Reset() {
	*x = GetVerifyCodeReq{}
<<<<<<< HEAD
	mi := &file_api_customer_customer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
=======
	if protoimpl.UnsafeEnabled {
		mi := &file_api_customer_customer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
}

func (x *GetVerifyCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVerifyCodeReq) ProtoMessage() {}

func (x *GetVerifyCodeReq) ProtoReflect() protoreflect.Message {
<<<<<<< HEAD
	mi := &file_api_customer_customer_proto_msgTypes[0]
	if x != nil {
=======
	mi := &file_api_customer_customer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
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
<<<<<<< HEAD
	return file_api_customer_customer_proto_rawDescGZIP(), []int{0}
=======
	return file_api_customer_customer_proto_rawDescGZIP(), []int{6}
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
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

<<<<<<< HEAD
	Code           int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message        string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	VerifyCode     string `protobuf:"bytes,3,opt,name=verify_code,json=verifyCode,proto3" json:"verify_code,omitempty"`
	VerifyCodeTime int32  `protobuf:"varint,4,opt,name=verify_code_time,json=verifyCodeTime,proto3" json:"verify_code_time,omitempty"`
=======
	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// 验证码
	VerifyCode string `protobuf:"bytes,3,opt,name=verify_code,json=verifyCode,proto3" json:"verify_code,omitempty"`
	// 生成时间 unix timestamp
	VerifyCodeTime int64 `protobuf:"varint,4,opt,name=verify_code_time,json=verifyCodeTime,proto3" json:"verify_code_time,omitempty"`
	// 有效期，单位 second
	VerifyCodeLife int32 `protobuf:"varint,5,opt,name=verify_code_life,json=verifyCodeLife,proto3" json:"verify_code_life,omitempty"`
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
}

func (x *GetVerifyCodeResp) Reset() {
	*x = GetVerifyCodeResp{}
<<<<<<< HEAD
	mi := &file_api_customer_customer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
=======
	if protoimpl.UnsafeEnabled {
		mi := &file_api_customer_customer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
}

func (x *GetVerifyCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVerifyCodeResp) ProtoMessage() {}

func (x *GetVerifyCodeResp) ProtoReflect() protoreflect.Message {
<<<<<<< HEAD
	mi := &file_api_customer_customer_proto_msgTypes[1]
	if x != nil {
=======
	mi := &file_api_customer_customer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
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
<<<<<<< HEAD
	return file_api_customer_customer_proto_rawDescGZIP(), []int{1}
}

func (x *GetVerifyCodeResp) GetCode() int32 {
=======
	return file_api_customer_customer_proto_rawDescGZIP(), []int{7}
}

func (x *GetVerifyCodeResp) GetCode() int64 {
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
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

<<<<<<< HEAD
func (x *GetVerifyCodeResp) GetVerifyCodeTime() int32 {
=======
func (x *GetVerifyCodeResp) GetVerifyCodeTime() int64 {
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
	if x != nil {
		return x.VerifyCodeTime
	}
	return 0
}

<<<<<<< HEAD
=======
func (x *GetVerifyCodeResp) GetVerifyCodeLife() int32 {
	if x != nil {
		return x.VerifyCodeLife
	}
	return 0
}

>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
var File_api_customer_customer_proto protoreflect.FileDescriptor

var file_api_customer_customer_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x63,
<<<<<<< HEAD
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
=======
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x10, 0x45, 0x73, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x91, 0x01, 0x0a, 0x11, 0x45, 0x73, 0x74, 0x69,
	0x6d, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x0b, 0x0a, 0x09, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x22, 0x3a, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x49, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x96, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x28, 0x0a, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x4c,
	0x69, 0x66, 0x65, 0x32, 0xc4, 0x03, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x12, 0x7f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x2d, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x7b, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x7d, 0x12, 0x54, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x55, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x89,
	0x01, 0x0a, 0x0d, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x12, 0x2f, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x2f, 0x7b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x7d, 0x2f, 0x7b, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x42, 0x20, 0x5a, 0x1e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
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

<<<<<<< HEAD
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
=======
var file_api_customer_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_customer_customer_proto_goTypes = []interface{}{
	(*EstimatePriceReq)(nil),  // 0: api.customer.EstimatePriceReq
	(*EstimatePriceResp)(nil), // 1: api.customer.EstimatePriceResp
	(*LogoutReq)(nil),         // 2: api.customer.LogoutReq
	(*LogoutResp)(nil),        // 3: api.customer.LogoutResp
	(*LoginReq)(nil),          // 4: api.customer.LoginReq
	(*LoginResp)(nil),         // 5: api.customer.LoginResp
	(*GetVerifyCodeReq)(nil),  // 6: api.customer.GetVerifyCodeReq
	(*GetVerifyCodeResp)(nil), // 7: api.customer.GetVerifyCodeResp
}
var file_api_customer_customer_proto_depIdxs = []int32{
	6, // 0: api.customer.Customer.GetVerifyCode:input_type -> api.customer.GetVerifyCodeReq
	4, // 1: api.customer.Customer.Login:input_type -> api.customer.LoginReq
	2, // 2: api.customer.Customer.Logout:input_type -> api.customer.LogoutReq
	0, // 3: api.customer.Customer.EstimatePrice:input_type -> api.customer.EstimatePriceReq
	7, // 4: api.customer.Customer.GetVerifyCode:output_type -> api.customer.GetVerifyCodeResp
	5, // 5: api.customer.Customer.Login:output_type -> api.customer.LoginResp
	3, // 6: api.customer.Customer.Logout:output_type -> api.customer.LogoutResp
	1, // 7: api.customer.Customer.EstimatePrice:output_type -> api.customer.EstimatePriceResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_customer_customer_proto_init() }
func file_api_customer_customer_proto_init() {
	if File_api_customer_customer_proto != nil {
		return
	}
<<<<<<< HEAD
=======
	if !protoimpl.UnsafeEnabled {
		file_api_customer_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstimatePriceReq); i {
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
		file_api_customer_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstimatePriceResp); i {
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
		file_api_customer_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReq); i {
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
		file_api_customer_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResp); i {
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
		file_api_customer_customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_api_customer_customer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_api_customer_customer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVerifyCodeReq); i {
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
		file_api_customer_customer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVerifyCodeResp); i {
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
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_customer_customer_proto_rawDesc,
			NumEnums:      0,
<<<<<<< HEAD
			NumMessages:   2,
=======
			NumMessages:   8,
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
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
