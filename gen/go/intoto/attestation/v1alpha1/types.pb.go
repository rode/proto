// Copyright 2021 The Rode Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.0
// source: intoto/attestation/v1alpha1/types.proto

package v1alpha1

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

// TypeURI is a uniform resource locator that should resolve to a human-readable description of the type to which it refers.
// See the in-toto attestation description for more information: https://github.com/in-toto/attestation/blob/main/spec/field_types.md#TypeURI
type TypeURI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value is the string representing the type URI.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TypeURI) Reset() {
	*x = TypeURI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intoto_attestation_v1alpha1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeURI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeURI) ProtoMessage() {}

func (x *TypeURI) ProtoReflect() protoreflect.Message {
	mi := &file_intoto_attestation_v1alpha1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeURI.ProtoReflect.Descriptor instead.
func (*TypeURI) Descriptor() ([]byte, []int) {
	return file_intoto_attestation_v1alpha1_types_proto_rawDescGZIP(), []int{0}
}

func (x *TypeURI) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// ResourceURI is a uniform resource locator that should resolve to an artifact. ResourceURIs should use the
// package URL (pkg:) or SPDX download location format.
// More information: https://github.com/in-toto/attestation/blob/main/spec/field_types.md#ResourceURI
type ResourceURI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value is the string representing the resource URI.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ResourceURI) Reset() {
	*x = ResourceURI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intoto_attestation_v1alpha1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceURI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceURI) ProtoMessage() {}

func (x *ResourceURI) ProtoReflect() protoreflect.Message {
	mi := &file_intoto_attestation_v1alpha1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceURI.ProtoReflect.Descriptor instead.
func (*ResourceURI) Descriptor() ([]byte, []int) {
	return file_intoto_attestation_v1alpha1_types_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceURI) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// DigestSet is mapping from hash names to digests of an artifact.
// More information: https://github.com/in-toto/attestation/blob/main/spec/field_types.md#DigestSet
type DigestSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sha256 is the sha256 hash of a file. It's the recommended baseline for maximum compatibility.
	Sha256 string `protobuf:"bytes,1,opt,name=sha256,proto3" json:"sha256,omitempty"`
}

func (x *DigestSet) Reset() {
	*x = DigestSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intoto_attestation_v1alpha1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DigestSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DigestSet) ProtoMessage() {}

func (x *DigestSet) ProtoReflect() protoreflect.Message {
	mi := &file_intoto_attestation_v1alpha1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DigestSet.ProtoReflect.Descriptor instead.
func (*DigestSet) Descriptor() ([]byte, []int) {
	return file_intoto_attestation_v1alpha1_types_proto_rawDescGZIP(), []int{2}
}

func (x *DigestSet) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

var File_intoto_attestation_v1alpha1_types_proto protoreflect.FileDescriptor

var file_intoto_attestation_v1alpha1_types_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x69, 0x6e, 0x74, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x1f, 0x0a, 0x07, 0x54, 0x79, 0x70, 0x65, 0x55, 0x52,
	0x49, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x55, 0x52, 0x49, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x29, 0x0a, 0x09,
	0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x53, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61,
	0x32, 0x35, 0x36, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35,
	0x36, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x33, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_intoto_attestation_v1alpha1_types_proto_rawDescOnce sync.Once
	file_intoto_attestation_v1alpha1_types_proto_rawDescData = file_intoto_attestation_v1alpha1_types_proto_rawDesc
)

func file_intoto_attestation_v1alpha1_types_proto_rawDescGZIP() []byte {
	file_intoto_attestation_v1alpha1_types_proto_rawDescOnce.Do(func() {
		file_intoto_attestation_v1alpha1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_intoto_attestation_v1alpha1_types_proto_rawDescData)
	})
	return file_intoto_attestation_v1alpha1_types_proto_rawDescData
}

var file_intoto_attestation_v1alpha1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_intoto_attestation_v1alpha1_types_proto_goTypes = []interface{}{
	(*TypeURI)(nil),     // 0: intoto.attestation.v1alpha1.TypeURI
	(*ResourceURI)(nil), // 1: intoto.attestation.v1alpha1.ResourceURI
	(*DigestSet)(nil),   // 2: intoto.attestation.v1alpha1.DigestSet
}
var file_intoto_attestation_v1alpha1_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_intoto_attestation_v1alpha1_types_proto_init() }
func file_intoto_attestation_v1alpha1_types_proto_init() {
	if File_intoto_attestation_v1alpha1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_intoto_attestation_v1alpha1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeURI); i {
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
		file_intoto_attestation_v1alpha1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceURI); i {
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
		file_intoto_attestation_v1alpha1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DigestSet); i {
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
			RawDescriptor: file_intoto_attestation_v1alpha1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_intoto_attestation_v1alpha1_types_proto_goTypes,
		DependencyIndexes: file_intoto_attestation_v1alpha1_types_proto_depIdxs,
		MessageInfos:      file_intoto_attestation_v1alpha1_types_proto_msgTypes,
	}.Build()
	File_intoto_attestation_v1alpha1_types_proto = out.File
	file_intoto_attestation_v1alpha1_types_proto_rawDesc = nil
	file_intoto_attestation_v1alpha1_types_proto_goTypes = nil
	file_intoto_attestation_v1alpha1_types_proto_depIdxs = nil
}
