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
// source: google/ads/googleads/v12/resources/bidding_data_exclusion.proto

package resources

import (
	enums "github.com/adomate-ads/api/pkg/google-ads/pb/v12/enums"
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

// Represents a bidding data exclusion.
//
// See "About data exclusions" at
// https://support.google.com/google-ads/answer/10370710.
type BiddingDataExclusion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the data exclusion.
	// Data exclusion resource names have the form:
	//
	// `customers/{customer_id}/biddingDataExclusions/{data_exclusion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the data exclusion.
	DataExclusionId int64 `protobuf:"varint,2,opt,name=data_exclusion_id,json=dataExclusionId,proto3" json:"data_exclusion_id,omitempty"`
	// The scope of the data exclusion.
	Scope enums.SeasonalityEventScopeEnum_SeasonalityEventScope `protobuf:"varint,3,opt,name=scope,proto3,enum=google.ads.googleads.v12.enums.SeasonalityEventScopeEnum_SeasonalityEventScope" json:"scope,omitempty"`
	// Output only. The status of the data exclusion.
	Status enums.SeasonalityEventStatusEnum_SeasonalityEventStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v12.enums.SeasonalityEventStatusEnum_SeasonalityEventStatus" json:"status,omitempty"`
	// Required. The inclusive start time of the data exclusion in yyyy-MM-dd HH:mm:ss
	// format.
	//
	// A data exclusion is backward looking and should be used for events that
	// start in the past and end either in the past or future.
	StartDateTime string `protobuf:"bytes,5,opt,name=start_date_time,json=startDateTime,proto3" json:"start_date_time,omitempty"`
	// Required. The exclusive end time of the data exclusion in yyyy-MM-dd HH:mm:ss format.
	//
	// The length of [start_date_time, end_date_time) interval must be
	// within (0, 14 days].
	EndDateTime string `protobuf:"bytes,6,opt,name=end_date_time,json=endDateTime,proto3" json:"end_date_time,omitempty"`
	// The name of the data exclusion. The name can be at most 255
	// characters.
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the data exclusion. The description can be at
	// most 2048 characters.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// If not specified, all devices will be included in this exclusion.
	// Otherwise, only the specified targeted devices will be included in this
	// exclusion.
	Devices []enums.DeviceEnum_Device `protobuf:"varint,9,rep,packed,name=devices,proto3,enum=google.ads.googleads.v12.enums.DeviceEnum_Device" json:"devices,omitempty"`
	// The data exclusion will apply to the campaigns listed when the scope of
	// this exclusion is CAMPAIGN. The maximum number of campaigns per event is
	// 2000.
	// Note: a data exclusion with both advertising_channel_types and
	// campaign_ids is not supported.
	Campaigns []string `protobuf:"bytes,10,rep,name=campaigns,proto3" json:"campaigns,omitempty"`
	// The data_exclusion will apply to all the campaigns under the listed
	// channels retroactively as well as going forward when the scope of this
	// exclusion is CHANNEL.
	// The supported advertising channel types are DISPLAY, SEARCH and SHOPPING.
	// Note: a data exclusion with both advertising_channel_types and
	// campaign_ids is not supported.
	AdvertisingChannelTypes []enums.AdvertisingChannelTypeEnum_AdvertisingChannelType `protobuf:"varint,11,rep,packed,name=advertising_channel_types,json=advertisingChannelTypes,proto3,enum=google.ads.googleads.v12.enums.AdvertisingChannelTypeEnum_AdvertisingChannelType" json:"advertising_channel_types,omitempty"`
}

func (x *BiddingDataExclusion) Reset() {
	*x = BiddingDataExclusion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BiddingDataExclusion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BiddingDataExclusion) ProtoMessage() {}

func (x *BiddingDataExclusion) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BiddingDataExclusion.ProtoReflect.Descriptor instead.
func (*BiddingDataExclusion) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_rawDescGZIP(), []int{0}
}

func (x *BiddingDataExclusion) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *BiddingDataExclusion) GetDataExclusionId() int64 {
	if x != nil {
		return x.DataExclusionId
	}
	return 0
}

func (x *BiddingDataExclusion) GetScope() enums.SeasonalityEventScopeEnum_SeasonalityEventScope {
	if x != nil {
		return x.Scope
	}
	return enums.SeasonalityEventScopeEnum_SeasonalityEventScope(0)
}

func (x *BiddingDataExclusion) GetStatus() enums.SeasonalityEventStatusEnum_SeasonalityEventStatus {
	if x != nil {
		return x.Status
	}
	return enums.SeasonalityEventStatusEnum_SeasonalityEventStatus(0)
}

func (x *BiddingDataExclusion) GetStartDateTime() string {
	if x != nil {
		return x.StartDateTime
	}
	return ""
}

func (x *BiddingDataExclusion) GetEndDateTime() string {
	if x != nil {
		return x.EndDateTime
	}
	return ""
}

func (x *BiddingDataExclusion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BiddingDataExclusion) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BiddingDataExclusion) GetDevices() []enums.DeviceEnum_Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

func (x *BiddingDataExclusion) GetCampaigns() []string {
	if x != nil {
		return x.Campaigns
	}
	return nil
}

func (x *BiddingDataExclusion) GetAdvertisingChannelTypes() []enums.AdvertisingChannelTypeEnum_AdvertisingChannelType {
	if x != nil {
		return x.AdvertisingChannelTypes
	}
	return nil
}

var File_google_ads_googleads_v12_resources_bidding_data_exclusion_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x07, 0x0a, 0x14, 0x42,
	0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x5a, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0xe0, 0x41, 0x05, 0xfa,
	0x41, 0x2f, 0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x64,
	0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2f, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0f, 0x64, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x65, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x4f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x6e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x51, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45,
	0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0b, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x44, 0x0a, 0x09, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x26, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x09, 0x63, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x19, 0x61, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x51, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69,
	0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x17, 0x61,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x78, 0xea, 0x41, 0x75, 0x0a, 0x2d, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d,
	0x42, 0x8b, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x19, 0x42, 0x69, 0x64,
	0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x31, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x32, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_rawDescData = file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_rawDesc
)

func file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_rawDescData
}

var file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_goTypes = []interface{}{
	(*BiddingDataExclusion)(nil),                                 // 0: google.ads.googleads.v12.resources.BiddingDataExclusion
	(enums.SeasonalityEventScopeEnum_SeasonalityEventScope)(0),   // 1: google.ads.googleads.v12.enums.SeasonalityEventScopeEnum.SeasonalityEventScope
	(enums.SeasonalityEventStatusEnum_SeasonalityEventStatus)(0), // 2: google.ads.googleads.v12.enums.SeasonalityEventStatusEnum.SeasonalityEventStatus
	(enums.DeviceEnum_Device)(0),                                 // 3: google.ads.googleads.v12.enums.DeviceEnum.Device
	(enums.AdvertisingChannelTypeEnum_AdvertisingChannelType)(0), // 4: google.ads.googleads.v12.enums.AdvertisingChannelTypeEnum.AdvertisingChannelType
}
var file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v12.resources.BiddingDataExclusion.scope:type_name -> google.ads.googleads.v12.enums.SeasonalityEventScopeEnum.SeasonalityEventScope
	2, // 1: google.ads.googleads.v12.resources.BiddingDataExclusion.status:type_name -> google.ads.googleads.v12.enums.SeasonalityEventStatusEnum.SeasonalityEventStatus
	3, // 2: google.ads.googleads.v12.resources.BiddingDataExclusion.devices:type_name -> google.ads.googleads.v12.enums.DeviceEnum.Device
	4, // 3: google.ads.googleads.v12.resources.BiddingDataExclusion.advertising_channel_types:type_name -> google.ads.googleads.v12.enums.AdvertisingChannelTypeEnum.AdvertisingChannelType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_init() }
func file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_init() {
	if File_google_ads_googleads_v12_resources_bidding_data_exclusion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BiddingDataExclusion); i {
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
			RawDescriptor: file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_resources_bidding_data_exclusion_proto = out.File
	file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_rawDesc = nil
	file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_goTypes = nil
	file_google_ads_googleads_v12_resources_bidding_data_exclusion_proto_depIdxs = nil
}
