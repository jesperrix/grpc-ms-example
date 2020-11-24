// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: my_model_api.proto

package mymodelapi

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MyModelInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X []int32 `protobuf:"varint,1,rep,packed,name=x,proto3" json:"x,omitempty"`
	Y []int32 `protobuf:"varint,2,rep,packed,name=y,proto3" json:"y,omitempty"`
}

func (x *MyModelInput) Reset() {
	*x = MyModelInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_model_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyModelInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyModelInput) ProtoMessage() {}

func (x *MyModelInput) ProtoReflect() protoreflect.Message {
	mi := &file_my_model_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyModelInput.ProtoReflect.Descriptor instead.
func (*MyModelInput) Descriptor() ([]byte, []int) {
	return file_my_model_api_proto_rawDescGZIP(), []int{0}
}

func (x *MyModelInput) GetX() []int32 {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *MyModelInput) GetY() []int32 {
	if x != nil {
		return x.Y
	}
	return nil
}

type MyModelOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelId string  `protobuf:"bytes,1,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	Y       []int32 `protobuf:"varint,2,rep,packed,name=y,proto3" json:"y,omitempty"`
}

func (x *MyModelOutput) Reset() {
	*x = MyModelOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_model_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyModelOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyModelOutput) ProtoMessage() {}

func (x *MyModelOutput) ProtoReflect() protoreflect.Message {
	mi := &file_my_model_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyModelOutput.ProtoReflect.Descriptor instead.
func (*MyModelOutput) Descriptor() ([]byte, []int) {
	return file_my_model_api_proto_rawDescGZIP(), []int{1}
}

func (x *MyModelOutput) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

func (x *MyModelOutput) GetY() []int32 {
	if x != nil {
		return x.Y
	}
	return nil
}

var File_my_model_api_proto protoreflect.FileDescriptor

var file_my_model_api_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x79, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x61, 0x70, 0x69,
	0x22, 0x2a, 0x0a, 0x0c, 0x4d, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x38, 0x0a, 0x0d,
	0x4d, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x01, 0x79, 0x32, 0x4e, 0x0a, 0x0a, 0x4d, 0x79, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x41, 0x70, 0x69, 0x12, 0x40, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x12,
	0x18, 0x2e, 0x6d, 0x79, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x79, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x79, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_my_model_api_proto_rawDescOnce sync.Once
	file_my_model_api_proto_rawDescData = file_my_model_api_proto_rawDesc
)

func file_my_model_api_proto_rawDescGZIP() []byte {
	file_my_model_api_proto_rawDescOnce.Do(func() {
		file_my_model_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_my_model_api_proto_rawDescData)
	})
	return file_my_model_api_proto_rawDescData
}

var file_my_model_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_my_model_api_proto_goTypes = []interface{}{
	(*MyModelInput)(nil),  // 0: mymodelapi.MyModelInput
	(*MyModelOutput)(nil), // 1: mymodelapi.MyModelOutput
}
var file_my_model_api_proto_depIdxs = []int32{
	0, // 0: mymodelapi.MyModelApi.Predict:input_type -> mymodelapi.MyModelInput
	1, // 1: mymodelapi.MyModelApi.Predict:output_type -> mymodelapi.MyModelOutput
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_my_model_api_proto_init() }
func file_my_model_api_proto_init() {
	if File_my_model_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_my_model_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyModelInput); i {
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
		file_my_model_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyModelOutput); i {
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
			RawDescriptor: file_my_model_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_my_model_api_proto_goTypes,
		DependencyIndexes: file_my_model_api_proto_depIdxs,
		MessageInfos:      file_my_model_api_proto_msgTypes,
	}.Build()
	File_my_model_api_proto = out.File
	file_my_model_api_proto_rawDesc = nil
	file_my_model_api_proto_goTypes = nil
	file_my_model_api_proto_depIdxs = nil
}
