// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.7
// source: google/ads/googleads/v12/errors/date_range_error.proto

package errors

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

// Enum describing possible date range errors.
type DateRangeErrorEnum_DateRangeError int32

const (
	// Enum unspecified.
	DateRangeErrorEnum_UNSPECIFIED DateRangeErrorEnum_DateRangeError = 0
	// The received error code is not known in this version.
	DateRangeErrorEnum_UNKNOWN DateRangeErrorEnum_DateRangeError = 1
	// Invalid date.
	DateRangeErrorEnum_INVALID_DATE DateRangeErrorEnum_DateRangeError = 2
	// The start date was after the end date.
	DateRangeErrorEnum_START_DATE_AFTER_END_DATE DateRangeErrorEnum_DateRangeError = 3
	// Cannot set date to past time
	DateRangeErrorEnum_CANNOT_SET_DATE_TO_PAST DateRangeErrorEnum_DateRangeError = 4
	// A date was used that is past the system "last" date.
	DateRangeErrorEnum_AFTER_MAXIMUM_ALLOWABLE_DATE DateRangeErrorEnum_DateRangeError = 5
	// Trying to change start date on a resource that has started.
	DateRangeErrorEnum_CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED DateRangeErrorEnum_DateRangeError = 6
)

// Enum value maps for DateRangeErrorEnum_DateRangeError.
var (
	DateRangeErrorEnum_DateRangeError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "INVALID_DATE",
		3: "START_DATE_AFTER_END_DATE",
		4: "CANNOT_SET_DATE_TO_PAST",
		5: "AFTER_MAXIMUM_ALLOWABLE_DATE",
		6: "CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED",
	}
	DateRangeErrorEnum_DateRangeError_value = map[string]int32{
		"UNSPECIFIED":                  0,
		"UNKNOWN":                      1,
		"INVALID_DATE":                 2,
		"START_DATE_AFTER_END_DATE":    3,
		"CANNOT_SET_DATE_TO_PAST":      4,
		"AFTER_MAXIMUM_ALLOWABLE_DATE": 5,
		"CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED": 6,
	}
)

func (x DateRangeErrorEnum_DateRangeError) Enum() *DateRangeErrorEnum_DateRangeError {
	p := new(DateRangeErrorEnum_DateRangeError)
	*p = x
	return p
}

func (x DateRangeErrorEnum_DateRangeError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DateRangeErrorEnum_DateRangeError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v12_errors_date_range_error_proto_enumTypes[0].Descriptor()
}

func (DateRangeErrorEnum_DateRangeError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v12_errors_date_range_error_proto_enumTypes[0]
}

func (x DateRangeErrorEnum_DateRangeError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DateRangeErrorEnum_DateRangeError.Descriptor instead.
func (DateRangeErrorEnum_DateRangeError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_errors_date_range_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible date range errors.
type DateRangeErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DateRangeErrorEnum) Reset() {
	*x = DateRangeErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_errors_date_range_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateRangeErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRangeErrorEnum) ProtoMessage() {}

func (x *DateRangeErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_errors_date_range_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRangeErrorEnum.ProtoReflect.Descriptor instead.
func (*DateRangeErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_errors_date_range_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v12_errors_date_range_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_errors_date_range_error_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x32, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xe6, 0x01, 0x0a, 0x12, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d,
	0x22, 0xcf, 0x01, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x44, 0x41, 0x54,
	0x45, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x41, 0x46, 0x54, 0x45, 0x52, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x45,
	0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54,
	0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x50, 0x41, 0x53, 0x54, 0x10, 0x04, 0x12,
	0x20, 0x0a, 0x1c, 0x41, 0x46, 0x54, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x58, 0x49, 0x4d, 0x55, 0x4d,
	0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10,
	0x05, 0x12, 0x2f, 0x0a, 0x2b, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49,
	0x46, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x46,
	0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44,
	0x10, 0x06, 0x42, 0xf3, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x32, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x13, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x32, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x32,
	0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v12_errors_date_range_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_errors_date_range_error_proto_rawDescData = file_google_ads_googleads_v12_errors_date_range_error_proto_rawDesc
)

func file_google_ads_googleads_v12_errors_date_range_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_errors_date_range_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_errors_date_range_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_errors_date_range_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_errors_date_range_error_proto_rawDescData
}

var file_google_ads_googleads_v12_errors_date_range_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v12_errors_date_range_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v12_errors_date_range_error_proto_goTypes = []interface{}{
	(DateRangeErrorEnum_DateRangeError)(0), // 0: google.ads.googleads.v12.errors.DateRangeErrorEnum.DateRangeError
	(*DateRangeErrorEnum)(nil),             // 1: google.ads.googleads.v12.errors.DateRangeErrorEnum
}
var file_google_ads_googleads_v12_errors_date_range_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_errors_date_range_error_proto_init() }
func file_google_ads_googleads_v12_errors_date_range_error_proto_init() {
	if File_google_ads_googleads_v12_errors_date_range_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_errors_date_range_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateRangeErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v12_errors_date_range_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v12_errors_date_range_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_errors_date_range_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v12_errors_date_range_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v12_errors_date_range_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_errors_date_range_error_proto = out.File
	file_google_ads_googleads_v12_errors_date_range_error_proto_rawDesc = nil
	file_google_ads_googleads_v12_errors_date_range_error_proto_goTypes = nil
	file_google_ads_googleads_v12_errors_date_range_error_proto_depIdxs = nil
}
