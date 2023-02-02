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
// source: google/ads/googleads/v12/services/campaign_bid_modifier_service.proto

package services

import (
	context "context"
	enums "github.com/adomate-ads/api/pkg/google-ads/pb/v12/enums"
	resources "github.com/adomate-ads/api/pkg/google-ads/pb/v12/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for
// [CampaignBidModifierService.MutateCampaignBidModifiers][google.ads.googleads.v12.services.CampaignBidModifierService.MutateCampaignBidModifiers].
type MutateCampaignBidModifiersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the customer whose campaign bid modifiers are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual campaign bid modifiers.
	Operations []*CampaignBidModifierOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// The response content type setting. Determines whether the mutable resource
	// or just the resource name should be returned post mutation.
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,5,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v12.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
}

func (x *MutateCampaignBidModifiersRequest) Reset() {
	*x = MutateCampaignBidModifiersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCampaignBidModifiersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCampaignBidModifiersRequest) ProtoMessage() {}

func (x *MutateCampaignBidModifiersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCampaignBidModifiersRequest.ProtoReflect.Descriptor instead.
func (*MutateCampaignBidModifiersRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateCampaignBidModifiersRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateCampaignBidModifiersRequest) GetOperations() []*CampaignBidModifierOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateCampaignBidModifiersRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateCampaignBidModifiersRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateCampaignBidModifiersRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_ResponseContentType(0)
}

// A single operation (create, remove, update) on a campaign bid modifier.
type CampaignBidModifierOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*CampaignBidModifierOperation_Create
	//	*CampaignBidModifierOperation_Update
	//	*CampaignBidModifierOperation_Remove
	Operation isCampaignBidModifierOperation_Operation `protobuf_oneof:"operation"`
}

func (x *CampaignBidModifierOperation) Reset() {
	*x = CampaignBidModifierOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignBidModifierOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignBidModifierOperation) ProtoMessage() {}

func (x *CampaignBidModifierOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignBidModifierOperation.ProtoReflect.Descriptor instead.
func (*CampaignBidModifierOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDescGZIP(), []int{1}
}

func (x *CampaignBidModifierOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *CampaignBidModifierOperation) GetOperation() isCampaignBidModifierOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *CampaignBidModifierOperation) GetCreate() *resources.CampaignBidModifier {
	if x, ok := x.GetOperation().(*CampaignBidModifierOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *CampaignBidModifierOperation) GetUpdate() *resources.CampaignBidModifier {
	if x, ok := x.GetOperation().(*CampaignBidModifierOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *CampaignBidModifierOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*CampaignBidModifierOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isCampaignBidModifierOperation_Operation interface {
	isCampaignBidModifierOperation_Operation()
}

type CampaignBidModifierOperation_Create struct {
	// Create operation: No resource name is expected for the new campaign bid
	// modifier.
	Create *resources.CampaignBidModifier `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignBidModifierOperation_Update struct {
	// Update operation: The campaign bid modifier is expected to have a valid
	// resource name.
	Update *resources.CampaignBidModifier `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignBidModifierOperation_Remove struct {
	// Remove operation: A resource name for the removed campaign bid modifier
	// is expected, in this format:
	//
	// `customers/{customer_id}/CampaignBidModifiers/{campaign_id}~{criterion_id}`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignBidModifierOperation_Create) isCampaignBidModifierOperation_Operation() {}

func (*CampaignBidModifierOperation_Update) isCampaignBidModifierOperation_Operation() {}

func (*CampaignBidModifierOperation_Remove) isCampaignBidModifierOperation_Operation() {}

// Response message for campaign bid modifiers mutate.
type MutateCampaignBidModifiersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (for example, auth
	// errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateCampaignBidModifierResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateCampaignBidModifiersResponse) Reset() {
	*x = MutateCampaignBidModifiersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCampaignBidModifiersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCampaignBidModifiersResponse) ProtoMessage() {}

func (x *MutateCampaignBidModifiersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCampaignBidModifiersResponse.ProtoReflect.Descriptor instead.
func (*MutateCampaignBidModifiersResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateCampaignBidModifiersResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateCampaignBidModifiersResponse) GetResults() []*MutateCampaignBidModifierResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateCampaignBidModifierResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated campaign bid modifier with only mutable fields after mutate.
	// The field will only be returned when response_content_type is set to
	// "MUTABLE_RESOURCE".
	CampaignBidModifier *resources.CampaignBidModifier `protobuf:"bytes,2,opt,name=campaign_bid_modifier,json=campaignBidModifier,proto3" json:"campaign_bid_modifier,omitempty"`
}

func (x *MutateCampaignBidModifierResult) Reset() {
	*x = MutateCampaignBidModifierResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCampaignBidModifierResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCampaignBidModifierResult) ProtoMessage() {}

func (x *MutateCampaignBidModifierResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCampaignBidModifierResult.ProtoReflect.Descriptor instead.
func (*MutateCampaignBidModifierResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateCampaignBidModifierResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateCampaignBidModifierResult) GetCampaignBidModifier() *resources.CampaignBidModifier {
	if x != nil {
		return x.CampaignBidModifier
	}
	return nil
}

var File_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x62, 0x69, 0x64,
	0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x21, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x64, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x7f, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xdb, 0x02, 0x0a, 0x1c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x73, 0x6b, 0x12, 0x51, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xfa, 0x41, 0x2e, 0x0a, 0x2c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xca, 0x01, 0x0a, 0x22, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x5c, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x22, 0xe6, 0x01, 0x0a, 0x1f, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x56, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xfa, 0x41, 0x2e,
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x6b, 0x0a, 0x15,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x13, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x69,
	0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x32, 0xef, 0x02, 0x0a, 0x1a, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x02, 0x0a, 0x1a, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3f, 0x22, 0x3a, 0x2f, 0x76,
	0x31, 0x32, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x63, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x16, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x8b, 0x02, 0x0a, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x1f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42,
	0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x32, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a,
	0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDescData = file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDesc
)

func file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDescData
}

var file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_goTypes = []interface{}{
	(*MutateCampaignBidModifiersRequest)(nil),              // 0: google.ads.googleads.v12.services.MutateCampaignBidModifiersRequest
	(*CampaignBidModifierOperation)(nil),                   // 1: google.ads.googleads.v12.services.CampaignBidModifierOperation
	(*MutateCampaignBidModifiersResponse)(nil),             // 2: google.ads.googleads.v12.services.MutateCampaignBidModifiersResponse
	(*MutateCampaignBidModifierResult)(nil),                // 3: google.ads.googleads.v12.services.MutateCampaignBidModifierResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 4: google.ads.googleads.v12.enums.ResponseContentTypeEnum.ResponseContentType
	(*fieldmaskpb.FieldMask)(nil),                          // 5: google.protobuf.FieldMask
	(*resources.CampaignBidModifier)(nil),                  // 6: google.ads.googleads.v12.resources.CampaignBidModifier
	(*status.Status)(nil),                                  // 7: google.rpc.Status
}
var file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v12.services.MutateCampaignBidModifiersRequest.operations:type_name -> google.ads.googleads.v12.services.CampaignBidModifierOperation
	4, // 1: google.ads.googleads.v12.services.MutateCampaignBidModifiersRequest.response_content_type:type_name -> google.ads.googleads.v12.enums.ResponseContentTypeEnum.ResponseContentType
	5, // 2: google.ads.googleads.v12.services.CampaignBidModifierOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 3: google.ads.googleads.v12.services.CampaignBidModifierOperation.create:type_name -> google.ads.googleads.v12.resources.CampaignBidModifier
	6, // 4: google.ads.googleads.v12.services.CampaignBidModifierOperation.update:type_name -> google.ads.googleads.v12.resources.CampaignBidModifier
	7, // 5: google.ads.googleads.v12.services.MutateCampaignBidModifiersResponse.partial_failure_error:type_name -> google.rpc.Status
	3, // 6: google.ads.googleads.v12.services.MutateCampaignBidModifiersResponse.results:type_name -> google.ads.googleads.v12.services.MutateCampaignBidModifierResult
	6, // 7: google.ads.googleads.v12.services.MutateCampaignBidModifierResult.campaign_bid_modifier:type_name -> google.ads.googleads.v12.resources.CampaignBidModifier
	0, // 8: google.ads.googleads.v12.services.CampaignBidModifierService.MutateCampaignBidModifiers:input_type -> google.ads.googleads.v12.services.MutateCampaignBidModifiersRequest
	2, // 9: google.ads.googleads.v12.services.CampaignBidModifierService.MutateCampaignBidModifiers:output_type -> google.ads.googleads.v12.services.MutateCampaignBidModifiersResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_init() }
func file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_init() {
	if File_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCampaignBidModifiersRequest); i {
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
		file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignBidModifierOperation); i {
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
		file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCampaignBidModifiersResponse); i {
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
		file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCampaignBidModifierResult); i {
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
	file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CampaignBidModifierOperation_Create)(nil),
		(*CampaignBidModifierOperation_Update)(nil),
		(*CampaignBidModifierOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto = out.File
	file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_rawDesc = nil
	file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_goTypes = nil
	file_google_ads_googleads_v12_services_campaign_bid_modifier_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CampaignBidModifierServiceClient is the client API for CampaignBidModifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignBidModifierServiceClient interface {
	// Creates, updates, or removes campaign bid modifiers.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ContextError]()
	//   [CriterionError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateCampaignBidModifiers(ctx context.Context, in *MutateCampaignBidModifiersRequest, opts ...grpc.CallOption) (*MutateCampaignBidModifiersResponse, error)
}

type campaignBidModifierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignBidModifierServiceClient(cc grpc.ClientConnInterface) CampaignBidModifierServiceClient {
	return &campaignBidModifierServiceClient{cc}
}

func (c *campaignBidModifierServiceClient) MutateCampaignBidModifiers(ctx context.Context, in *MutateCampaignBidModifiersRequest, opts ...grpc.CallOption) (*MutateCampaignBidModifiersResponse, error) {
	out := new(MutateCampaignBidModifiersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v12.services.CampaignBidModifierService/MutateCampaignBidModifiers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignBidModifierServiceServer is the server API for CampaignBidModifierService service.
type CampaignBidModifierServiceServer interface {
	// Creates, updates, or removes campaign bid modifiers.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ContextError]()
	//   [CriterionError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateCampaignBidModifiers(context.Context, *MutateCampaignBidModifiersRequest) (*MutateCampaignBidModifiersResponse, error)
}

// UnimplementedCampaignBidModifierServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignBidModifierServiceServer struct {
}

func (*UnimplementedCampaignBidModifierServiceServer) MutateCampaignBidModifiers(context.Context, *MutateCampaignBidModifiersRequest) (*MutateCampaignBidModifiersResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCampaignBidModifiers not implemented")
}

func RegisterCampaignBidModifierServiceServer(s *grpc.Server, srv CampaignBidModifierServiceServer) {
	s.RegisterService(&_CampaignBidModifierService_serviceDesc, srv)
}

func _CampaignBidModifierService_MutateCampaignBidModifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignBidModifiersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignBidModifierServiceServer).MutateCampaignBidModifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v12.services.CampaignBidModifierService/MutateCampaignBidModifiers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignBidModifierServiceServer).MutateCampaignBidModifiers(ctx, req.(*MutateCampaignBidModifiersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignBidModifierService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v12.services.CampaignBidModifierService",
	HandlerType: (*CampaignBidModifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignBidModifiers",
			Handler:    _CampaignBidModifierService_MutateCampaignBidModifiers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v12/services/campaign_bid_modifier_service.proto",
}
