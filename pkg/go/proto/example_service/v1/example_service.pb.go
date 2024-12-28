// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: example_service/v1/example_service.proto

package example_servicev1

import (
	_ "github.com/bufbuild/buf-tour/gen/google/api"
	_type "github.com/bufbuild/buf-tour/gen/google/type"
	_ "github.com/bufbuild/buf-tour/gen/validate"
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

type TypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TypeRequest) Reset() {
	*x = TypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_service_v1_example_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeRequest) ProtoMessage() {}

func (x *TypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_service_v1_example_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeRequest.ProtoReflect.Descriptor instead.
func (*TypeRequest) Descriptor() ([]byte, []int) {
	return file_example_service_v1_example_service_proto_rawDescGZIP(), []int{0}
}

func (x *TypeRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *Type `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *TypeResponse) Reset() {
	*x = TypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_service_v1_example_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeResponse) ProtoMessage() {}

func (x *TypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_service_v1_example_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeResponse.ProtoReflect.Descriptor instead.
func (*TypeResponse) Descriptor() ([]byte, []int) {
	return file_example_service_v1_example_service_proto_rawDescGZIP(), []int{1}
}

func (x *TypeResponse) GetType() *Type {
	if x != nil {
		return x.Type
	}
	return nil
}

type Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *_type.DateTime `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Type) Reset() {
	*x = Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_service_v1_example_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_example_service_v1_example_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_example_service_v1_example_service_proto_rawDescGZIP(), []int{2}
}

func (x *Type) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Type) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Type) GetCreatedAt() *_type.DateTime {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_example_service_v1_example_service_proto protoreflect.FileDescriptor

var file_example_service_v1_example_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x27, 0x0a, 0x0b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x32, 0x03, 0x20, 0xe7, 0x07, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0c, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x72, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x73, 0x0a, 0x0e,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x42, 0xd9, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2d, 0x74, 0x6f, 0x75,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58,
	0xaa, 0x02, 0x11, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_service_v1_example_service_proto_rawDescOnce sync.Once
	file_example_service_v1_example_service_proto_rawDescData = file_example_service_v1_example_service_proto_rawDesc
)

func file_example_service_v1_example_service_proto_rawDescGZIP() []byte {
	file_example_service_v1_example_service_proto_rawDescOnce.Do(func() {
		file_example_service_v1_example_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_service_v1_example_service_proto_rawDescData)
	})
	return file_example_service_v1_example_service_proto_rawDescData
}

var file_example_service_v1_example_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_example_service_v1_example_service_proto_goTypes = []any{
	(*TypeRequest)(nil),    // 0: example_service.v1.TypeRequest
	(*TypeResponse)(nil),   // 1: example_service.v1.TypeResponse
	(*Type)(nil),           // 2: example_service.v1.Type
	(*_type.DateTime)(nil), // 3: google.type.DateTime
}
var file_example_service_v1_example_service_proto_depIdxs = []int32{
	2, // 0: example_service.v1.TypeResponse.type:type_name -> example_service.v1.Type
	3, // 1: example_service.v1.Type.created_at:type_name -> google.type.DateTime
	0, // 2: example_service.v1.ExampleService.Type:input_type -> example_service.v1.TypeRequest
	1, // 3: example_service.v1.ExampleService.Type:output_type -> example_service.v1.TypeResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_example_service_v1_example_service_proto_init() }
func file_example_service_v1_example_service_proto_init() {
	if File_example_service_v1_example_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_service_v1_example_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TypeRequest); i {
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
		file_example_service_v1_example_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TypeResponse); i {
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
		file_example_service_v1_example_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Type); i {
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
			RawDescriptor: file_example_service_v1_example_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_service_v1_example_service_proto_goTypes,
		DependencyIndexes: file_example_service_v1_example_service_proto_depIdxs,
		MessageInfos:      file_example_service_v1_example_service_proto_msgTypes,
	}.Build()
	File_example_service_v1_example_service_proto = out.File
	file_example_service_v1_example_service_proto_rawDesc = nil
	file_example_service_v1_example_service_proto_goTypes = nil
	file_example_service_v1_example_service_proto_depIdxs = nil
}
