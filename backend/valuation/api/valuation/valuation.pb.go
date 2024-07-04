// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: api/valuation/valuation.proto

package valuation

import (
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

type GetEstimatePriceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin      string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *GetEstimatePriceReq) Reset() {
	*x = GetEstimatePriceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_valuation_valuation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEstimatePriceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEstimatePriceReq) ProtoMessage() {}

func (x *GetEstimatePriceReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_valuation_valuation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEstimatePriceReq.ProtoReflect.Descriptor instead.
func (*GetEstimatePriceReq) Descriptor() ([]byte, []int) {
	return file_api_valuation_valuation_proto_rawDescGZIP(), []int{0}
}

func (x *GetEstimatePriceReq) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *GetEstimatePriceReq) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type GetEstimatePriceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin      string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Price       int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *GetEstimatePriceReply) Reset() {
	*x = GetEstimatePriceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_valuation_valuation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEstimatePriceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEstimatePriceReply) ProtoMessage() {}

func (x *GetEstimatePriceReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_valuation_valuation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEstimatePriceReply.ProtoReflect.Descriptor instead.
func (*GetEstimatePriceReply) Descriptor() ([]byte, []int) {
	return file_api_valuation_valuation_proto_rawDescGZIP(), []int{1}
}

func (x *GetEstimatePriceReply) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *GetEstimatePriceReply) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *GetEstimatePriceReply) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_api_valuation_valuation_proto protoreflect.FileDescriptor

var file_api_valuation_valuation_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4f,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x67, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0x69, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x73, 0x74, 0x69,
	0x6d, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x73, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x23, 0x5a, 0x21, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_valuation_valuation_proto_rawDescOnce sync.Once
	file_api_valuation_valuation_proto_rawDescData = file_api_valuation_valuation_proto_rawDesc
)

func file_api_valuation_valuation_proto_rawDescGZIP() []byte {
	file_api_valuation_valuation_proto_rawDescOnce.Do(func() {
		file_api_valuation_valuation_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_valuation_valuation_proto_rawDescData)
	})
	return file_api_valuation_valuation_proto_rawDescData
}

var file_api_valuation_valuation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_valuation_valuation_proto_goTypes = []interface{}{
	(*GetEstimatePriceReq)(nil),   // 0: api.valuation.GetEstimatePriceReq
	(*GetEstimatePriceReply)(nil), // 1: api.valuation.GetEstimatePriceReply
}
var file_api_valuation_valuation_proto_depIdxs = []int32{
	0, // 0: api.valuation.Valuation.GetEstimatePrice:input_type -> api.valuation.GetEstimatePriceReq
	1, // 1: api.valuation.Valuation.GetEstimatePrice:output_type -> api.valuation.GetEstimatePriceReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_valuation_valuation_proto_init() }
func file_api_valuation_valuation_proto_init() {
	if File_api_valuation_valuation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_valuation_valuation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEstimatePriceReq); i {
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
		file_api_valuation_valuation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEstimatePriceReply); i {
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
			RawDescriptor: file_api_valuation_valuation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_valuation_valuation_proto_goTypes,
		DependencyIndexes: file_api_valuation_valuation_proto_depIdxs,
		MessageInfos:      file_api_valuation_valuation_proto_msgTypes,
	}.Build()
	File_api_valuation_valuation_proto = out.File
	file_api_valuation_valuation_proto_rawDesc = nil
	file_api_valuation_valuation_proto_goTypes = nil
	file_api_valuation_valuation_proto_depIdxs = nil
}
