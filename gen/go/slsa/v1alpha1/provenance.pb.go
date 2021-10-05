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
// source: slsa/v1alpha1/provenance.proto

package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
	v1alpha1 "github.com/rode/proto/gen/go/intoto/attestation/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Provenance represents a v0.1 SLSA (https://slsa.dev/) build provenance. The schema for a provenance is described by
// the SLSA framework at https://slsa.dev/provenance/v0.1.
type Provenance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Builder represents the identity of the service performing a build.
	Builder *Provenance_Builder `protobuf:"bytes,1,opt,name=builder,proto3" json:"builder,omitempty"`
	// Metadata is a collection of miscellaneous properties about the build.
	Metadata *Provenance_Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Recipe describes the build configuration.
	Recipe *Provenance_Recipe `protobuf:"bytes,3,opt,name=recipe,proto3" json:"recipe,omitempty"`
	// Materials is a collection of artifacts used in the build.
	Materials []*Provenance_Material `protobuf:"bytes,4,rep,name=materials,proto3" json:"materials,omitempty"`
}

func (x *Provenance) Reset() {
	*x = Provenance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slsa_v1alpha1_provenance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provenance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provenance) ProtoMessage() {}

func (x *Provenance) ProtoReflect() protoreflect.Message {
	mi := &file_slsa_v1alpha1_provenance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provenance.ProtoReflect.Descriptor instead.
func (*Provenance) Descriptor() ([]byte, []int) {
	return file_slsa_v1alpha1_provenance_proto_rawDescGZIP(), []int{0}
}

func (x *Provenance) GetBuilder() *Provenance_Builder {
	if x != nil {
		return x.Builder
	}
	return nil
}

func (x *Provenance) GetMetadata() *Provenance_Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Provenance) GetRecipe() *Provenance_Recipe {
	if x != nil {
		return x.Recipe
	}
	return nil
}

func (x *Provenance) GetMaterials() []*Provenance_Material {
	if x != nil {
		return x.Materials
	}
	return nil
}

// Builder represents the identity of the service performing a build.
type Provenance_Builder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id is the unique identifier of the Builder. An example value from the SLSA GitHub Actions demo:
	// https://github.com/slsa-framework/github-actions-demo/Attestations/GitHubHostedActions@v1
	Id *v1alpha1.TypeURI `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Provenance_Builder) Reset() {
	*x = Provenance_Builder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slsa_v1alpha1_provenance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provenance_Builder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provenance_Builder) ProtoMessage() {}

func (x *Provenance_Builder) ProtoReflect() protoreflect.Message {
	mi := &file_slsa_v1alpha1_provenance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provenance_Builder.ProtoReflect.Descriptor instead.
func (*Provenance_Builder) Descriptor() ([]byte, []int) {
	return file_slsa_v1alpha1_provenance_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Provenance_Builder) GetId() *v1alpha1.TypeURI {
	if x != nil {
		return x.Id
	}
	return nil
}

// Metadata is a collection of miscellaneous properties about the build.
type Provenance_Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// BuildInvocationId is a unique identifier of a particular build attempt, with a format that will be unique to each
	// builder. An example: https://github.com/slsa-framework/github-actions-demo/actions/runs/940146003
	BuildInvocationId string `protobuf:"bytes,1,opt,name=build_invocation_id,json=buildInvocationId,proto3" json:"build_invocation_id,omitempty"`
	// Completeness describes the extent to which inputs were captured for the recipe and materials.
	Completeness *Provenance_Metadata_Completeness `protobuf:"bytes,2,opt,name=completeness,proto3" json:"completeness,omitempty"`
	// BuildStartedOn is a timestamp captured when the build first began.
	BuildStartedOn *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=build_started_on,json=buildStartedOn,proto3" json:"build_started_on,omitempty"`
	// BuildFinishedOn is a timestamp representing when the build completed.
	BuildFinishedOn *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=build_finished_on,json=buildFinishedOn,proto3" json:"build_finished_on,omitempty"`
}

func (x *Provenance_Metadata) Reset() {
	*x = Provenance_Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slsa_v1alpha1_provenance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provenance_Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provenance_Metadata) ProtoMessage() {}

func (x *Provenance_Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_slsa_v1alpha1_provenance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provenance_Metadata.ProtoReflect.Descriptor instead.
func (*Provenance_Metadata) Descriptor() ([]byte, []int) {
	return file_slsa_v1alpha1_provenance_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Provenance_Metadata) GetBuildInvocationId() string {
	if x != nil {
		return x.BuildInvocationId
	}
	return ""
}

func (x *Provenance_Metadata) GetCompleteness() *Provenance_Metadata_Completeness {
	if x != nil {
		return x.Completeness
	}
	return nil
}

func (x *Provenance_Metadata) GetBuildStartedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.BuildStartedOn
	}
	return nil
}

func (x *Provenance_Metadata) GetBuildFinishedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.BuildFinishedOn
	}
	return nil
}

// Recipe describes the build configuration. Along with materials, it should be sufficient to reproduce the build.
// Some recipe components may not have been fully captured, see Metadata.Completeness.
type Provenance_Recipe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type is a TypeURI that describes what type of recipe was performed. An example from the SLSA GitHub Actions demo:
	// https://github.com/Attestations/GitHubActionsWorkflow@v1
	Type *v1alpha1.TypeURI `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// DefinedInMaterial is the index in Provenance.Materials where the recipe steps are defined.
	DefinedInMaterial int32 `protobuf:"varint,2,opt,name=defined_in_material,json=definedInMaterial,proto3" json:"defined_in_material,omitempty"`
	// EntryPoint describes the initializing step of the build, e.g., a configuration file or step. Should be constrained
	// to prevent errors (e.g., creating the wrong type of build).
	EntryPoint string `protobuf:"bytes,3,opt,name=entry_point,json=entryPoint,proto3" json:"entry_point,omitempty"`
	// Arguments represents any external inputs to the build, with a schema defined by the recipe type.
	Arguments *structpb.Struct `protobuf:"bytes,4,opt,name=arguments,proto3" json:"arguments,omitempty"`
	// Environment represents any other inputs supplied by the builder that influence the build. Useful for reproducing
	// the build, but not expected to be part of any policy.
	Environment *structpb.Struct `protobuf:"bytes,5,opt,name=environment,proto3" json:"environment,omitempty"`
}

func (x *Provenance_Recipe) Reset() {
	*x = Provenance_Recipe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slsa_v1alpha1_provenance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provenance_Recipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provenance_Recipe) ProtoMessage() {}

func (x *Provenance_Recipe) ProtoReflect() protoreflect.Message {
	mi := &file_slsa_v1alpha1_provenance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provenance_Recipe.ProtoReflect.Descriptor instead.
func (*Provenance_Recipe) Descriptor() ([]byte, []int) {
	return file_slsa_v1alpha1_provenance_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Provenance_Recipe) GetType() *v1alpha1.TypeURI {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Provenance_Recipe) GetDefinedInMaterial() int32 {
	if x != nil {
		return x.DefinedInMaterial
	}
	return 0
}

func (x *Provenance_Recipe) GetEntryPoint() string {
	if x != nil {
		return x.EntryPoint
	}
	return ""
}

func (x *Provenance_Recipe) GetArguments() *structpb.Struct {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *Provenance_Recipe) GetEnvironment() *structpb.Struct {
	if x != nil {
		return x.Environment
	}
	return nil
}

// Material describes an artifact that was used in the build, including source code and build tools.
type Provenance_Material struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URI represents a resolvable location for the artifact.
	Uri *v1alpha1.ResourceURI `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Digest is a set of hashes for the material that can be used to verify the artifact fetched from the URI.
	Digest *v1alpha1.DigestSet `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *Provenance_Material) Reset() {
	*x = Provenance_Material{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slsa_v1alpha1_provenance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provenance_Material) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provenance_Material) ProtoMessage() {}

func (x *Provenance_Material) ProtoReflect() protoreflect.Message {
	mi := &file_slsa_v1alpha1_provenance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provenance_Material.ProtoReflect.Descriptor instead.
func (*Provenance_Material) Descriptor() ([]byte, []int) {
	return file_slsa_v1alpha1_provenance_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Provenance_Material) GetUri() *v1alpha1.ResourceURI {
	if x != nil {
		return x.Uri
	}
	return nil
}

func (x *Provenance_Material) GetDigest() *v1alpha1.DigestSet {
	if x != nil {
		return x.Digest
	}
	return nil
}

// Completeness describes the extent to which inputs were captured for the recipe and materials.
type Provenance_Metadata_Completeness struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Arguments indicates that the recipe arguments were completely captured.
	Arguments bool `protobuf:"varint,1,opt,name=arguments,proto3" json:"arguments,omitempty"`
	// Environment indicates that the recipe environment field is complete.
	Environment bool `protobuf:"varint,2,opt,name=environment,proto3" json:"environment,omitempty"`
	// Materials indicates that the described materials are complete.
	Materials bool `protobuf:"varint,3,opt,name=materials,proto3" json:"materials,omitempty"`
}

func (x *Provenance_Metadata_Completeness) Reset() {
	*x = Provenance_Metadata_Completeness{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slsa_v1alpha1_provenance_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provenance_Metadata_Completeness) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provenance_Metadata_Completeness) ProtoMessage() {}

func (x *Provenance_Metadata_Completeness) ProtoReflect() protoreflect.Message {
	mi := &file_slsa_v1alpha1_provenance_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provenance_Metadata_Completeness.ProtoReflect.Descriptor instead.
func (*Provenance_Metadata_Completeness) Descriptor() ([]byte, []int) {
	return file_slsa_v1alpha1_provenance_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *Provenance_Metadata_Completeness) GetArguments() bool {
	if x != nil {
		return x.Arguments
	}
	return false
}

func (x *Provenance_Metadata_Completeness) GetEnvironment() bool {
	if x != nil {
		return x.Environment
	}
	return false
}

func (x *Provenance_Metadata_Completeness) GetMaterials() bool {
	if x != nil {
		return x.Materials
	}
	return false
}

var File_slsa_v1alpha1_provenance_proto protoreflect.FileDescriptor

var file_slsa_v1alpha1_provenance_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x6c, 0x73, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x73, 0x6c, 0x73, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27,
	0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x08, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x76,
	0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6c, 0x73, 0x61, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6c, 0x73, 0x61, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6c, 0x73, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x40, 0x0a,
	0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x73, 0x6c, 0x73, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x1a,
	0x3f, 0x0a, 0x07, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x55, 0x52, 0x49, 0x52, 0x02, 0x69, 0x64,
	0x1a, 0x8b, 0x03, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a,
	0x13, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x53, 0x0a,
	0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x6c, 0x73, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65,
	0x73, 0x73, 0x12, 0x44, 0x0a, 0x10, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x46, 0x0a, 0x11, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x4f, 0x6e,
	0x1a, 0x6c, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x85,
	0x02, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x55, 0x52, 0x49, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x69,
	0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0b, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x86, 0x01, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x52, 0x49, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12,
	0x3e, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x53, 0x65, 0x74, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x42,
	0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f,
	0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x6c, 0x73, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_slsa_v1alpha1_provenance_proto_rawDescOnce sync.Once
	file_slsa_v1alpha1_provenance_proto_rawDescData = file_slsa_v1alpha1_provenance_proto_rawDesc
)

func file_slsa_v1alpha1_provenance_proto_rawDescGZIP() []byte {
	file_slsa_v1alpha1_provenance_proto_rawDescOnce.Do(func() {
		file_slsa_v1alpha1_provenance_proto_rawDescData = protoimpl.X.CompressGZIP(file_slsa_v1alpha1_provenance_proto_rawDescData)
	})
	return file_slsa_v1alpha1_provenance_proto_rawDescData
}

var file_slsa_v1alpha1_provenance_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_slsa_v1alpha1_provenance_proto_goTypes = []interface{}{
	(*Provenance)(nil),                       // 0: slsa.v1alpha1.Provenance
	(*Provenance_Builder)(nil),               // 1: slsa.v1alpha1.Provenance.Builder
	(*Provenance_Metadata)(nil),              // 2: slsa.v1alpha1.Provenance.Metadata
	(*Provenance_Recipe)(nil),                // 3: slsa.v1alpha1.Provenance.Recipe
	(*Provenance_Material)(nil),              // 4: slsa.v1alpha1.Provenance.Material
	(*Provenance_Metadata_Completeness)(nil), // 5: slsa.v1alpha1.Provenance.Metadata.Completeness
	(*v1alpha1.TypeURI)(nil),                 // 6: intoto.attestation.v1alpha1.TypeURI
	(*timestamppb.Timestamp)(nil),            // 7: google.protobuf.Timestamp
	(*structpb.Struct)(nil),                  // 8: google.protobuf.Struct
	(*v1alpha1.ResourceURI)(nil),             // 9: intoto.attestation.v1alpha1.ResourceURI
	(*v1alpha1.DigestSet)(nil),               // 10: intoto.attestation.v1alpha1.DigestSet
}
var file_slsa_v1alpha1_provenance_proto_depIdxs = []int32{
	1,  // 0: slsa.v1alpha1.Provenance.builder:type_name -> slsa.v1alpha1.Provenance.Builder
	2,  // 1: slsa.v1alpha1.Provenance.metadata:type_name -> slsa.v1alpha1.Provenance.Metadata
	3,  // 2: slsa.v1alpha1.Provenance.recipe:type_name -> slsa.v1alpha1.Provenance.Recipe
	4,  // 3: slsa.v1alpha1.Provenance.materials:type_name -> slsa.v1alpha1.Provenance.Material
	6,  // 4: slsa.v1alpha1.Provenance.Builder.id:type_name -> intoto.attestation.v1alpha1.TypeURI
	5,  // 5: slsa.v1alpha1.Provenance.Metadata.completeness:type_name -> slsa.v1alpha1.Provenance.Metadata.Completeness
	7,  // 6: slsa.v1alpha1.Provenance.Metadata.build_started_on:type_name -> google.protobuf.Timestamp
	7,  // 7: slsa.v1alpha1.Provenance.Metadata.build_finished_on:type_name -> google.protobuf.Timestamp
	6,  // 8: slsa.v1alpha1.Provenance.Recipe.type:type_name -> intoto.attestation.v1alpha1.TypeURI
	8,  // 9: slsa.v1alpha1.Provenance.Recipe.arguments:type_name -> google.protobuf.Struct
	8,  // 10: slsa.v1alpha1.Provenance.Recipe.environment:type_name -> google.protobuf.Struct
	9,  // 11: slsa.v1alpha1.Provenance.Material.uri:type_name -> intoto.attestation.v1alpha1.ResourceURI
	10, // 12: slsa.v1alpha1.Provenance.Material.digest:type_name -> intoto.attestation.v1alpha1.DigestSet
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_slsa_v1alpha1_provenance_proto_init() }
func file_slsa_v1alpha1_provenance_proto_init() {
	if File_slsa_v1alpha1_provenance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_slsa_v1alpha1_provenance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provenance); i {
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
		file_slsa_v1alpha1_provenance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provenance_Builder); i {
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
		file_slsa_v1alpha1_provenance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provenance_Metadata); i {
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
		file_slsa_v1alpha1_provenance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provenance_Recipe); i {
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
		file_slsa_v1alpha1_provenance_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provenance_Material); i {
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
		file_slsa_v1alpha1_provenance_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provenance_Metadata_Completeness); i {
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
			RawDescriptor: file_slsa_v1alpha1_provenance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_slsa_v1alpha1_provenance_proto_goTypes,
		DependencyIndexes: file_slsa_v1alpha1_provenance_proto_depIdxs,
		MessageInfos:      file_slsa_v1alpha1_provenance_proto_msgTypes,
	}.Build()
	File_slsa_v1alpha1_provenance_proto = out.File
	file_slsa_v1alpha1_provenance_proto_rawDesc = nil
	file_slsa_v1alpha1_provenance_proto_goTypes = nil
	file_slsa_v1alpha1_provenance_proto_depIdxs = nil
}
