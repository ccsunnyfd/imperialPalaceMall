// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: mall/v1/error_reason.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_CATEGORY_UNSPECIFIED  ErrorReason = 0
	ErrorReason_CATEGORY_NOT_FOUND    ErrorReason = 1
	ErrorReason_GOODS_UNSPECIFIED     ErrorReason = 2
	ErrorReason_GOODS_NOT_FOUND       ErrorReason = 3
	ErrorReason_GOODS_INFO_NOT_FOUND  ErrorReason = 4
	ErrorReason_SKU_NOT_FOUND         ErrorReason = 5
	ErrorReason_ATTRS_NOT_FOUND       ErrorReason = 6
	ErrorReason_ATTR_VALUES_NOT_FOUND ErrorReason = 7
	ErrorReason_LIST_GOODS_ERROR      ErrorReason = 8
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "CATEGORY_UNSPECIFIED",
		1: "CATEGORY_NOT_FOUND",
		2: "GOODS_UNSPECIFIED",
		3: "GOODS_NOT_FOUND",
		4: "GOODS_INFO_NOT_FOUND",
		5: "SKU_NOT_FOUND",
		6: "ATTRS_NOT_FOUND",
		7: "ATTR_VALUES_NOT_FOUND",
		8: "LIST_GOODS_ERROR",
	}
	ErrorReason_value = map[string]int32{
		"CATEGORY_UNSPECIFIED":  0,
		"CATEGORY_NOT_FOUND":    1,
		"GOODS_UNSPECIFIED":     2,
		"GOODS_NOT_FOUND":       3,
		"GOODS_INFO_NOT_FOUND":  4,
		"SKU_NOT_FOUND":         5,
		"ATTRS_NOT_FOUND":       6,
		"ATTR_VALUES_NOT_FOUND": 7,
		"LIST_GOODS_ERROR":      8,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_mall_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_mall_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_mall_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_mall_v1_error_reason_proto protoreflect.FileDescriptor

var file_mall_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x9a,
	0x02, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x14, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1c,
	0x0a, 0x12, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1b, 0x0a, 0x11,
	0x47, 0x4f, 0x4f, 0x44, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x47, 0x4f, 0x4f,
	0x44, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x1a, 0x04,
	0xa8, 0x45, 0x94, 0x03, 0x12, 0x1e, 0x0a, 0x14, 0x47, 0x4f, 0x4f, 0x44, 0x53, 0x5f, 0x49, 0x4e,
	0x46, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x04, 0x1a, 0x04,
	0xa8, 0x45, 0x94, 0x03, 0x12, 0x17, 0x0a, 0x0d, 0x53, 0x4b, 0x55, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x19, 0x0a,
	0x0f, 0x41, 0x54, 0x54, 0x52, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1f, 0x0a, 0x15, 0x41, 0x54, 0x54, 0x52,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1a, 0x0a, 0x10, 0x4c, 0x49, 0x53,
	0x54, 0x5f, 0x47, 0x4f, 0x4f, 0x44, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x08, 0x1a,
	0x04, 0xa8, 0x45, 0xf4, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x24, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x13, 0x6d, 0x61,
	0x6c, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mall_v1_error_reason_proto_rawDescOnce sync.Once
	file_mall_v1_error_reason_proto_rawDescData = file_mall_v1_error_reason_proto_rawDesc
)

func file_mall_v1_error_reason_proto_rawDescGZIP() []byte {
	file_mall_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_mall_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_mall_v1_error_reason_proto_rawDescData)
	})
	return file_mall_v1_error_reason_proto_rawDescData
}

var file_mall_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mall_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: api.mall.v1.ErrorReason
}
var file_mall_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mall_v1_error_reason_proto_init() }
func file_mall_v1_error_reason_proto_init() {
	if File_mall_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mall_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mall_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_mall_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_mall_v1_error_reason_proto_enumTypes,
	}.Build()
	File_mall_v1_error_reason_proto = out.File
	file_mall_v1_error_reason_proto_rawDesc = nil
	file_mall_v1_error_reason_proto_goTypes = nil
	file_mall_v1_error_reason_proto_depIdxs = nil
}
