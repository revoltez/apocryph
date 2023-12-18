// SPDX-License-Identifier: GPL-3.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: provisioning-capacity.proto

package proto

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

type CapacityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*Resource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *CapacityRequest) Reset() {
	*x = CapacityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provisioning_capacity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapacityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapacityRequest) ProtoMessage() {}

func (x *CapacityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provisioning_capacity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapacityRequest.ProtoReflect.Descriptor instead.
func (*CapacityRequest) Descriptor() ([]byte, []int) {
	return file_provisioning_capacity_proto_rawDescGZIP(), []int{0}
}

func (x *CapacityRequest) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

type CapacityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailableResources []*Resource   `protobuf:"bytes,1,rep,name=availableResources,proto3" json:"availableResources,omitempty"`
	Pricing            *PricingTable `protobuf:"bytes,2,opt,name=pricing,proto3" json:"pricing,omitempty"`
	ProviderAddress    []byte        `protobuf:"bytes,3,opt,name=providerAddress,proto3" json:"providerAddress,omitempty"`
}

func (x *CapacityResponse) Reset() {
	*x = CapacityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provisioning_capacity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapacityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapacityResponse) ProtoMessage() {}

func (x *CapacityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provisioning_capacity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapacityResponse.ProtoReflect.Descriptor instead.
func (*CapacityResponse) Descriptor() ([]byte, []int) {
	return file_provisioning_capacity_proto_rawDescGZIP(), []int{1}
}

func (x *CapacityResponse) GetAvailableResources() []*Resource {
	if x != nil {
		return x.AvailableResources
	}
	return nil
}

func (x *CapacityResponse) GetPricing() *PricingTable {
	if x != nil {
		return x.Pricing
	}
	return nil
}

func (x *CapacityResponse) GetProviderAddress() []byte {
	if x != nil {
		return x.ProviderAddress
	}
	return nil
}

var File_provisioning_capacity_proto protoreflect.FileDescriptor

var file_provisioning_capacity_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x61,
	0x70, 0x6f, 0x63, 0x72, 0x79, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x30,
	0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x1a, 0x09, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x50, 0x0a, 0x0f, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x6f, 0x63, 0x72, 0x79, 0x70, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x30, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x22, 0xd0, 0x01, 0x0a, 0x10, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x6f, 0x63, 0x72, 0x79, 0x70, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x30, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x70, 0x6f, 0x63, 0x72,
	0x79, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x30, 0x2e, 0x70, 0x72, 0x69,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x72, 0x61, 0x64, 0x65, 0x2d, 0x63, 0x6f, 0x6f, 0x70, 0x2f,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x2d, 0x70, 0x6f, 0x64, 0x73, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_provisioning_capacity_proto_rawDescOnce sync.Once
	file_provisioning_capacity_proto_rawDescData = file_provisioning_capacity_proto_rawDesc
)

func file_provisioning_capacity_proto_rawDescGZIP() []byte {
	file_provisioning_capacity_proto_rawDescOnce.Do(func() {
		file_provisioning_capacity_proto_rawDescData = protoimpl.X.CompressGZIP(file_provisioning_capacity_proto_rawDescData)
	})
	return file_provisioning_capacity_proto_rawDescData
}

var file_provisioning_capacity_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_provisioning_capacity_proto_goTypes = []interface{}{
	(*CapacityRequest)(nil),  // 0: apocryph.proto.v0.provisioningCapacity.CapacityRequest
	(*CapacityResponse)(nil), // 1: apocryph.proto.v0.provisioningCapacity.CapacityResponse
	(*Resource)(nil),         // 2: apocryph.proto.v0.pod.Resource
	(*PricingTable)(nil),     // 3: apocryph.proto.v0.pricing.PricingTable
}
var file_provisioning_capacity_proto_depIdxs = []int32{
	2, // 0: apocryph.proto.v0.provisioningCapacity.CapacityRequest.resources:type_name -> apocryph.proto.v0.pod.Resource
	2, // 1: apocryph.proto.v0.provisioningCapacity.CapacityResponse.availableResources:type_name -> apocryph.proto.v0.pod.Resource
	3, // 2: apocryph.proto.v0.provisioningCapacity.CapacityResponse.pricing:type_name -> apocryph.proto.v0.pricing.PricingTable
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_provisioning_capacity_proto_init() }
func file_provisioning_capacity_proto_init() {
	if File_provisioning_capacity_proto != nil {
		return
	}
	file_pod_proto_init()
	file_pricing_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_provisioning_capacity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapacityRequest); i {
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
		file_provisioning_capacity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapacityResponse); i {
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
			RawDescriptor: file_provisioning_capacity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_provisioning_capacity_proto_goTypes,
		DependencyIndexes: file_provisioning_capacity_proto_depIdxs,
		MessageInfos:      file_provisioning_capacity_proto_msgTypes,
	}.Build()
	File_provisioning_capacity_proto = out.File
	file_provisioning_capacity_proto_rawDesc = nil
	file_provisioning_capacity_proto_goTypes = nil
	file_provisioning_capacity_proto_depIdxs = nil
}
