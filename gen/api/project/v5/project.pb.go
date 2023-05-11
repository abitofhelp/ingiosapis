////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 1999-2023 - A Bit of Help, Inc.  All Rights Reserved.
// Use of this source code is governed by the content in the LICENSE file in
// the root of this repository.
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: api/project/v5/project.proto

package projectv5

import (
	_ "github.com/abitofhelp/ingiosapis/gen/validate"
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

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AccessTier AccessTier `protobuf:"varint,2,opt,name=access_tier,json=accessTier,proto3,enum=api.project.v5.AccessTier" json:"access_tier,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_v5_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_v5_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_api_project_v5_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetAccessTier() AccessTier {
	if x != nil {
		return x.AccessTier
	}
	return AccessTier_ACCESS_TIER_UNSPECIFIED
}

var File_api_project_v5_project_proto protoreflect.FileDescriptor

var file_api_project_v5_project_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x35,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x35, 0x1a, 0x20,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x35, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2d, 0xfa, 0x42, 0x2a, 0x72, 0x28, 0x28, 0x50, 0x32, 0x24, 0x5e, 0x5b,
	0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x5d, 0x2b, 0x28, 0x20,
	0x5b, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x5d, 0x2b, 0x29,
	0x2a, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x35, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x69, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x69, 0x65, 0x72, 0x42,
	0xbb, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x35, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x62, 0x69, 0x74, 0x6f, 0x66, 0x68, 0x65, 0x6c, 0x70, 0x2f, 0x69, 0x6e,
	0x67, 0x69, 0x6f, 0x73, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x35, 0x3b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x76, 0x35, 0xa2, 0x02, 0x03, 0x41, 0x50, 0x58, 0xaa, 0x02, 0x0e, 0x41, 0x70,
	0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x56, 0x35, 0xca, 0x02, 0x0e, 0x41,
	0x70, 0x69, 0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x56, 0x35, 0xe2, 0x02, 0x1a,
	0x41, 0x70, 0x69, 0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x56, 0x35, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x56, 0x35, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_project_v5_project_proto_rawDescOnce sync.Once
	file_api_project_v5_project_proto_rawDescData = file_api_project_v5_project_proto_rawDesc
)

func file_api_project_v5_project_proto_rawDescGZIP() []byte {
	file_api_project_v5_project_proto_rawDescOnce.Do(func() {
		file_api_project_v5_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_project_v5_project_proto_rawDescData)
	})
	return file_api_project_v5_project_proto_rawDescData
}

var file_api_project_v5_project_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_project_v5_project_proto_goTypes = []interface{}{
	(*Project)(nil), // 0: api.project.v5.Project
	(AccessTier)(0), // 1: api.project.v5.AccessTier
}
var file_api_project_v5_project_proto_depIdxs = []int32{
	1, // 0: api.project.v5.Project.access_tier:type_name -> api.project.v5.AccessTier
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_project_v5_project_proto_init() }
func file_api_project_v5_project_proto_init() {
	if File_api_project_v5_project_proto != nil {
		return
	}
	file_api_project_v5_access_tier_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_project_v5_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
			RawDescriptor: file_api_project_v5_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_project_v5_project_proto_goTypes,
		DependencyIndexes: file_api_project_v5_project_proto_depIdxs,
		MessageInfos:      file_api_project_v5_project_proto_msgTypes,
	}.Build()
	File_api_project_v5_project_proto = out.File
	file_api_project_v5_project_proto_rawDesc = nil
	file_api_project_v5_project_proto_goTypes = nil
	file_api_project_v5_project_proto_depIdxs = nil
}
