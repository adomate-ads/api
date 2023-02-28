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
// source: google/ads/googleads/v12/enums/user_list_size_range.proto

package enums

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

// Enum containing possible user list size ranges.
type UserListSizeRangeEnum_UserListSizeRange int32

const (
	// Not specified.
	UserListSizeRangeEnum_UNSPECIFIED UserListSizeRangeEnum_UserListSizeRange = 0
	// Used for return value only. Represents value unknown in this version.
	UserListSizeRangeEnum_UNKNOWN UserListSizeRangeEnum_UserListSizeRange = 1
	// User list has less than 500 users.
	UserListSizeRangeEnum_LESS_THAN_FIVE_HUNDRED UserListSizeRangeEnum_UserListSizeRange = 2
	// User list has number of users in range of 500 to 1000.
	UserListSizeRangeEnum_LESS_THAN_ONE_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 3
	// User list has number of users in range of 1000 to 10000.
	UserListSizeRangeEnum_ONE_THOUSAND_TO_TEN_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 4
	// User list has number of users in range of 10000 to 50000.
	UserListSizeRangeEnum_TEN_THOUSAND_TO_FIFTY_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 5
	// User list has number of users in range of 50000 to 100000.
	UserListSizeRangeEnum_FIFTY_THOUSAND_TO_ONE_HUNDRED_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 6
	// User list has number of users in range of 100000 to 300000.
	UserListSizeRangeEnum_ONE_HUNDRED_THOUSAND_TO_THREE_HUNDRED_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 7
	// User list has number of users in range of 300000 to 500000.
	UserListSizeRangeEnum_THREE_HUNDRED_THOUSAND_TO_FIVE_HUNDRED_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 8
	// User list has number of users in range of 500000 to 1 million.
	UserListSizeRangeEnum_FIVE_HUNDRED_THOUSAND_TO_ONE_MILLION UserListSizeRangeEnum_UserListSizeRange = 9
	// User list has number of users in range of 1 to 2 millions.
	UserListSizeRangeEnum_ONE_MILLION_TO_TWO_MILLION UserListSizeRangeEnum_UserListSizeRange = 10
	// User list has number of users in range of 2 to 3 millions.
	UserListSizeRangeEnum_TWO_MILLION_TO_THREE_MILLION UserListSizeRangeEnum_UserListSizeRange = 11
	// User list has number of users in range of 3 to 5 millions.
	UserListSizeRangeEnum_THREE_MILLION_TO_FIVE_MILLION UserListSizeRangeEnum_UserListSizeRange = 12
	// User list has number of users in range of 5 to 10 millions.
	UserListSizeRangeEnum_FIVE_MILLION_TO_TEN_MILLION UserListSizeRangeEnum_UserListSizeRange = 13
	// User list has number of users in range of 10 to 20 millions.
	UserListSizeRangeEnum_TEN_MILLION_TO_TWENTY_MILLION UserListSizeRangeEnum_UserListSizeRange = 14
	// User list has number of users in range of 20 to 30 millions.
	UserListSizeRangeEnum_TWENTY_MILLION_TO_THIRTY_MILLION UserListSizeRangeEnum_UserListSizeRange = 15
	// User list has number of users in range of 30 to 50 millions.
	UserListSizeRangeEnum_THIRTY_MILLION_TO_FIFTY_MILLION UserListSizeRangeEnum_UserListSizeRange = 16
	// User list has over 50 million users.
	UserListSizeRangeEnum_OVER_FIFTY_MILLION UserListSizeRangeEnum_UserListSizeRange = 17
)

// Enum value maps for UserListSizeRangeEnum_UserListSizeRange.
var (
	UserListSizeRangeEnum_UserListSizeRange_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "LESS_THAN_FIVE_HUNDRED",
		3:  "LESS_THAN_ONE_THOUSAND",
		4:  "ONE_THOUSAND_TO_TEN_THOUSAND",
		5:  "TEN_THOUSAND_TO_FIFTY_THOUSAND",
		6:  "FIFTY_THOUSAND_TO_ONE_HUNDRED_THOUSAND",
		7:  "ONE_HUNDRED_THOUSAND_TO_THREE_HUNDRED_THOUSAND",
		8:  "THREE_HUNDRED_THOUSAND_TO_FIVE_HUNDRED_THOUSAND",
		9:  "FIVE_HUNDRED_THOUSAND_TO_ONE_MILLION",
		10: "ONE_MILLION_TO_TWO_MILLION",
		11: "TWO_MILLION_TO_THREE_MILLION",
		12: "THREE_MILLION_TO_FIVE_MILLION",
		13: "FIVE_MILLION_TO_TEN_MILLION",
		14: "TEN_MILLION_TO_TWENTY_MILLION",
		15: "TWENTY_MILLION_TO_THIRTY_MILLION",
		16: "THIRTY_MILLION_TO_FIFTY_MILLION",
		17: "OVER_FIFTY_MILLION",
	}
	UserListSizeRangeEnum_UserListSizeRange_value = map[string]int32{
		"UNSPECIFIED":                                     0,
		"UNKNOWN":                                         1,
		"LESS_THAN_FIVE_HUNDRED":                          2,
		"LESS_THAN_ONE_THOUSAND":                          3,
		"ONE_THOUSAND_TO_TEN_THOUSAND":                    4,
		"TEN_THOUSAND_TO_FIFTY_THOUSAND":                  5,
		"FIFTY_THOUSAND_TO_ONE_HUNDRED_THOUSAND":          6,
		"ONE_HUNDRED_THOUSAND_TO_THREE_HUNDRED_THOUSAND":  7,
		"THREE_HUNDRED_THOUSAND_TO_FIVE_HUNDRED_THOUSAND": 8,
		"FIVE_HUNDRED_THOUSAND_TO_ONE_MILLION":            9,
		"ONE_MILLION_TO_TWO_MILLION":                      10,
		"TWO_MILLION_TO_THREE_MILLION":                    11,
		"THREE_MILLION_TO_FIVE_MILLION":                   12,
		"FIVE_MILLION_TO_TEN_MILLION":                     13,
		"TEN_MILLION_TO_TWENTY_MILLION":                   14,
		"TWENTY_MILLION_TO_THIRTY_MILLION":                15,
		"THIRTY_MILLION_TO_FIFTY_MILLION":                 16,
		"OVER_FIFTY_MILLION":                              17,
	}
)

func (x UserListSizeRangeEnum_UserListSizeRange) Enum() *UserListSizeRangeEnum_UserListSizeRange {
	p := new(UserListSizeRangeEnum_UserListSizeRange)
	*p = x
	return p
}

func (x UserListSizeRangeEnum_UserListSizeRange) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserListSizeRangeEnum_UserListSizeRange) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v12_enums_user_list_size_range_proto_enumTypes[0].Descriptor()
}

func (UserListSizeRangeEnum_UserListSizeRange) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v12_enums_user_list_size_range_proto_enumTypes[0]
}

func (x UserListSizeRangeEnum_UserListSizeRange) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserListSizeRangeEnum_UserListSizeRange.Descriptor instead.
func (UserListSizeRangeEnum_UserListSizeRange) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_enums_user_list_size_range_proto_rawDescGZIP(), []int{0, 0}
}

// Size range in terms of number of users of a UserList.
type UserListSizeRangeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserListSizeRangeEnum) Reset() {
	*x = UserListSizeRangeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_enums_user_list_size_range_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListSizeRangeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListSizeRangeEnum) ProtoMessage() {}

func (x *UserListSizeRangeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_enums_user_list_size_range_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListSizeRangeEnum.ProtoReflect.Descriptor instead.
func (*UserListSizeRangeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_enums_user_list_size_range_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v12_enums_user_list_size_range_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_enums_user_list_size_range_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x94, 0x05, 0x0a, 0x15,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xfa, 0x04, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x45, 0x53,
	0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x46, 0x49, 0x56, 0x45, 0x5f, 0x48, 0x55, 0x4e, 0x44,
	0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48,
	0x41, 0x4e, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x48, 0x4f, 0x55, 0x53, 0x41, 0x4e, 0x44, 0x10,
	0x03, 0x12, 0x20, 0x0a, 0x1c, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x48, 0x4f, 0x55, 0x53, 0x41, 0x4e,
	0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x45, 0x4e, 0x5f, 0x54, 0x48, 0x4f, 0x55, 0x53, 0x41, 0x4e,
	0x44, 0x10, 0x04, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x45, 0x4e, 0x5f, 0x54, 0x48, 0x4f, 0x55, 0x53,
	0x41, 0x4e, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x46, 0x49, 0x46, 0x54, 0x59, 0x5f, 0x54, 0x48, 0x4f,
	0x55, 0x53, 0x41, 0x4e, 0x44, 0x10, 0x05, 0x12, 0x2a, 0x0a, 0x26, 0x46, 0x49, 0x46, 0x54, 0x59,
	0x5f, 0x54, 0x48, 0x4f, 0x55, 0x53, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x4f, 0x4e, 0x45,
	0x5f, 0x48, 0x55, 0x4e, 0x44, 0x52, 0x45, 0x44, 0x5f, 0x54, 0x48, 0x4f, 0x55, 0x53, 0x41, 0x4e,
	0x44, 0x10, 0x06, 0x12, 0x32, 0x0a, 0x2e, 0x4f, 0x4e, 0x45, 0x5f, 0x48, 0x55, 0x4e, 0x44, 0x52,
	0x45, 0x44, 0x5f, 0x54, 0x48, 0x4f, 0x55, 0x53, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x54,
	0x48, 0x52, 0x45, 0x45, 0x5f, 0x48, 0x55, 0x4e, 0x44, 0x52, 0x45, 0x44, 0x5f, 0x54, 0x48, 0x4f,
	0x55, 0x53, 0x41, 0x4e, 0x44, 0x10, 0x07, 0x12, 0x33, 0x0a, 0x2f, 0x54, 0x48, 0x52, 0x45, 0x45,
	0x5f, 0x48, 0x55, 0x4e, 0x44, 0x52, 0x45, 0x44, 0x5f, 0x54, 0x48, 0x4f, 0x55, 0x53, 0x41, 0x4e,
	0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x46, 0x49, 0x56, 0x45, 0x5f, 0x48, 0x55, 0x4e, 0x44, 0x52, 0x45,
	0x44, 0x5f, 0x54, 0x48, 0x4f, 0x55, 0x53, 0x41, 0x4e, 0x44, 0x10, 0x08, 0x12, 0x28, 0x0a, 0x24,
	0x46, 0x49, 0x56, 0x45, 0x5f, 0x48, 0x55, 0x4e, 0x44, 0x52, 0x45, 0x44, 0x5f, 0x54, 0x48, 0x4f,
	0x55, 0x53, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x4d, 0x49, 0x4c,
	0x4c, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x4e, 0x45, 0x5f, 0x4d, 0x49,
	0x4c, 0x4c, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x57, 0x4f, 0x5f, 0x4d, 0x49, 0x4c,
	0x4c, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x57, 0x4f, 0x5f, 0x4d, 0x49,
	0x4c, 0x4c, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x4d,
	0x49, 0x4c, 0x4c, 0x49, 0x4f, 0x4e, 0x10, 0x0b, 0x12, 0x21, 0x0a, 0x1d, 0x54, 0x48, 0x52, 0x45,
	0x45, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x46, 0x49, 0x56,
	0x45, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4f, 0x4e, 0x10, 0x0c, 0x12, 0x1f, 0x0a, 0x1b, 0x46,
	0x49, 0x56, 0x45, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x54,
	0x45, 0x4e, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x21, 0x0a, 0x1d,
	0x54, 0x45, 0x4e, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x54,
	0x57, 0x45, 0x4e, 0x54, 0x59, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4f, 0x4e, 0x10, 0x0e, 0x12,
	0x24, 0x0a, 0x20, 0x54, 0x57, 0x45, 0x4e, 0x54, 0x59, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x54, 0x59, 0x5f, 0x4d, 0x49, 0x4c, 0x4c,
	0x49, 0x4f, 0x4e, 0x10, 0x0f, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x48, 0x49, 0x52, 0x54, 0x59, 0x5f,
	0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x46, 0x49, 0x46, 0x54, 0x59,
	0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4f, 0x4e, 0x10, 0x10, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x56,
	0x45, 0x52, 0x5f, 0x46, 0x49, 0x46, 0x54, 0x59, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4f, 0x4e,
	0x10, 0x11, 0x42, 0xf0, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x16, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x32, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a, 0x3a,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v12_enums_user_list_size_range_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_enums_user_list_size_range_proto_rawDescData = file_google_ads_googleads_v12_enums_user_list_size_range_proto_rawDesc
)

func file_google_ads_googleads_v12_enums_user_list_size_range_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_enums_user_list_size_range_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_enums_user_list_size_range_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_enums_user_list_size_range_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_enums_user_list_size_range_proto_rawDescData
}

var file_google_ads_googleads_v12_enums_user_list_size_range_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v12_enums_user_list_size_range_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v12_enums_user_list_size_range_proto_goTypes = []interface{}{
	(UserListSizeRangeEnum_UserListSizeRange)(0), // 0: google.ads.googleads.v12.enums.UserListSizeRangeEnum.UserListSizeRange
	(*UserListSizeRangeEnum)(nil),                // 1: google.ads.googleads.v12.enums.UserListSizeRangeEnum
}
var file_google_ads_googleads_v12_enums_user_list_size_range_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_enums_user_list_size_range_proto_init() }
func file_google_ads_googleads_v12_enums_user_list_size_range_proto_init() {
	if File_google_ads_googleads_v12_enums_user_list_size_range_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_enums_user_list_size_range_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListSizeRangeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v12_enums_user_list_size_range_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v12_enums_user_list_size_range_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_enums_user_list_size_range_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v12_enums_user_list_size_range_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v12_enums_user_list_size_range_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_enums_user_list_size_range_proto = out.File
	file_google_ads_googleads_v12_enums_user_list_size_range_proto_rawDesc = nil
	file_google_ads_googleads_v12_enums_user_list_size_range_proto_goTypes = nil
	file_google_ads_googleads_v12_enums_user_list_size_range_proto_depIdxs = nil
}
