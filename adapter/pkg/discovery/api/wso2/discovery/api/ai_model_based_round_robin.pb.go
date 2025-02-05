//  Copyright (c) 2025, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
//
//  WSO2 LLC. licenses this file to you under the Apache License,
//  Version 2.0 (the "License"); you may not use this file except
//  in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an
//  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
//  KIND, either express or implied.  See the License for the
//  specific language governing permissions and limitations
//  under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/api/ai_model_based_round_robin.proto

package api

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

// Holds AIModelBasedRoundRobin configs
type AIModelBasedRoundRobin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnQuotaExceedSuspendDuration int32          `protobuf:"varint,1,opt,name=onQuotaExceedSuspendDuration,proto3" json:"onQuotaExceedSuspendDuration,omitempty"`
	ProductionModels             []*ModelWeight `protobuf:"bytes,2,rep,name=productionModels,proto3" json:"productionModels,omitempty"`
	SandboxModels                []*ModelWeight `protobuf:"bytes,3,rep,name=sandboxModels,proto3" json:"sandboxModels,omitempty"`
	Enabled                      bool           `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *AIModelBasedRoundRobin) Reset() {
	*x = AIModelBasedRoundRobin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_ai_model_based_round_robin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AIModelBasedRoundRobin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AIModelBasedRoundRobin) ProtoMessage() {}

func (x *AIModelBasedRoundRobin) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_ai_model_based_round_robin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AIModelBasedRoundRobin.ProtoReflect.Descriptor instead.
func (*AIModelBasedRoundRobin) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_ai_model_based_round_robin_proto_rawDescGZIP(), []int{0}
}

func (x *AIModelBasedRoundRobin) GetOnQuotaExceedSuspendDuration() int32 {
	if x != nil {
		return x.OnQuotaExceedSuspendDuration
	}
	return 0
}

func (x *AIModelBasedRoundRobin) GetProductionModels() []*ModelWeight {
	if x != nil {
		return x.ProductionModels
	}
	return nil
}

func (x *AIModelBasedRoundRobin) GetSandboxModels() []*ModelWeight {
	if x != nil {
		return x.SandboxModels
	}
	return nil
}

func (x *AIModelBasedRoundRobin) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// Holds ModelWeight configs
type ModelWeight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model    string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Weight   int32  `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *ModelWeight) Reset() {
	*x = ModelWeight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_ai_model_based_round_robin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelWeight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelWeight) ProtoMessage() {}

func (x *ModelWeight) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_ai_model_based_round_robin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelWeight.ProtoReflect.Descriptor instead.
func (*ModelWeight) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_ai_model_based_round_robin_proto_rawDescGZIP(), []int{1}
}

func (x *ModelWeight) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ModelWeight) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *ModelWeight) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

var File_wso2_discovery_api_ai_model_based_round_robin_proto protoreflect.FileDescriptor

var file_wso2_discovery_api_ai_model_based_round_robin_proto_rawDesc = []byte{
	0x0a, 0x33, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x69, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x62, 0x61,
	0x73, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x8a, 0x02, 0x0a, 0x16, 0x41, 0x49,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x61, 0x73, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52,
	0x6f, 0x62, 0x69, 0x6e, 0x12, 0x42, 0x0a, 0x1c, 0x6f, 0x6e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x45,
	0x78, 0x63, 0x65, 0x65, 0x64, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1c, 0x6f, 0x6e, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x45, 0x78, 0x63, 0x65, 0x65, 0x64, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x45, 0x0a, 0x0d, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x77,
	0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x0d, 0x73,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x57, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42,
	0x83, 0x01, 0x0a, 0x23, 0x6f, 0x72, 0x67, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x61, 0x70, 0x6b,
	0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x1b, 0x41, 0x49, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x42, 0x61, 0x73, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x6f, 0x62, 0x69, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f,
	0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77,
	0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_api_ai_model_based_round_robin_proto_rawDescOnce sync.Once
	file_wso2_discovery_api_ai_model_based_round_robin_proto_rawDescData = file_wso2_discovery_api_ai_model_based_round_robin_proto_rawDesc
)

func file_wso2_discovery_api_ai_model_based_round_robin_proto_rawDescGZIP() []byte {
	file_wso2_discovery_api_ai_model_based_round_robin_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_api_ai_model_based_round_robin_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_api_ai_model_based_round_robin_proto_rawDescData)
	})
	return file_wso2_discovery_api_ai_model_based_round_robin_proto_rawDescData
}

var file_wso2_discovery_api_ai_model_based_round_robin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wso2_discovery_api_ai_model_based_round_robin_proto_goTypes = []interface{}{
	(*AIModelBasedRoundRobin)(nil), // 0: wso2.discovery.api.AIModelBasedRoundRobin
	(*ModelWeight)(nil),            // 1: wso2.discovery.api.ModelWeight
}
var file_wso2_discovery_api_ai_model_based_round_robin_proto_depIdxs = []int32{
	1, // 0: wso2.discovery.api.AIModelBasedRoundRobin.productionModels:type_name -> wso2.discovery.api.ModelWeight
	1, // 1: wso2.discovery.api.AIModelBasedRoundRobin.sandboxModels:type_name -> wso2.discovery.api.ModelWeight
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wso2_discovery_api_ai_model_based_round_robin_proto_init() }
func file_wso2_discovery_api_ai_model_based_round_robin_proto_init() {
	if File_wso2_discovery_api_ai_model_based_round_robin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_api_ai_model_based_round_robin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AIModelBasedRoundRobin); i {
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
		file_wso2_discovery_api_ai_model_based_round_robin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelWeight); i {
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
			RawDescriptor: file_wso2_discovery_api_ai_model_based_round_robin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_api_ai_model_based_round_robin_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_api_ai_model_based_round_robin_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_api_ai_model_based_round_robin_proto_msgTypes,
	}.Build()
	File_wso2_discovery_api_ai_model_based_round_robin_proto = out.File
	file_wso2_discovery_api_ai_model_based_round_robin_proto_rawDesc = nil
	file_wso2_discovery_api_ai_model_based_round_robin_proto_goTypes = nil
	file_wso2_discovery_api_ai_model_based_round_robin_proto_depIdxs = nil
}
