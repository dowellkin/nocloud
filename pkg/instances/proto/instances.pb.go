// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: pkg/instances/proto/instances.proto

package proto

import (
	_ "github.com/slntopp/nocloud/pkg/hasher/hasherpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string                     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title     string                     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Config    map[string]*structpb.Value `protobuf:"bytes,3,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Resources map[string]*structpb.Value `protobuf:"bytes,4,rep,name=resources,proto3" json:"resources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data      map[string]*structpb.Value `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Hash      string                     `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Instance) Reset() {
	*x = Instance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_instances_proto_instances_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_instances_proto_instances_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instance.ProtoReflect.Descriptor instead.
func (*Instance) Descriptor() ([]byte, []int) {
	return file_pkg_instances_proto_instances_proto_rawDescGZIP(), []int{0}
}

func (x *Instance) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Instance) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Instance) GetConfig() map[string]*structpb.Value {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Instance) GetResources() map[string]*structpb.Value {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Instance) GetData() map[string]*structpb.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Instance) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type InstancesGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string                     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Type      string                     `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Config    map[string]*structpb.Value `protobuf:"bytes,3,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Instances []*Instance                `protobuf:"bytes,4,rep,name=instances,proto3" json:"instances,omitempty"`
	Resources map[string]*structpb.Value `protobuf:"bytes,5,rep,name=resources,proto3" json:"resources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data      map[string]*structpb.Value `protobuf:"bytes,6,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Hash      string                     `protobuf:"bytes,7,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *InstancesGroup) Reset() {
	*x = InstancesGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_instances_proto_instances_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstancesGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstancesGroup) ProtoMessage() {}

func (x *InstancesGroup) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_instances_proto_instances_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstancesGroup.ProtoReflect.Descriptor instead.
func (*InstancesGroup) Descriptor() ([]byte, []int) {
	return file_pkg_instances_proto_instances_proto_rawDescGZIP(), []int{1}
}

func (x *InstancesGroup) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *InstancesGroup) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *InstancesGroup) GetConfig() map[string]*structpb.Value {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *InstancesGroup) GetInstances() []*Instance {
	if x != nil {
		return x.Instances
	}
	return nil
}

func (x *InstancesGroup) GetResources() map[string]*structpb.Value {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *InstancesGroup) GetData() map[string]*structpb.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *InstancesGroup) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type TestInstancesGroupConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *InstancesGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *TestInstancesGroupConfigRequest) Reset() {
	*x = TestInstancesGroupConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_instances_proto_instances_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestInstancesGroupConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestInstancesGroupConfigRequest) ProtoMessage() {}

func (x *TestInstancesGroupConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_instances_proto_instances_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestInstancesGroupConfigRequest.ProtoReflect.Descriptor instead.
func (*TestInstancesGroupConfigRequest) Descriptor() ([]byte, []int) {
	return file_pkg_instances_proto_instances_proto_rawDescGZIP(), []int{2}
}

func (x *TestInstancesGroupConfigRequest) GetGroup() *InstancesGroup {
	if x != nil {
		return x.Group
	}
	return nil
}

type TestInstancesGroupConfigError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Instance string `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *TestInstancesGroupConfigError) Reset() {
	*x = TestInstancesGroupConfigError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_instances_proto_instances_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestInstancesGroupConfigError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestInstancesGroupConfigError) ProtoMessage() {}

func (x *TestInstancesGroupConfigError) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_instances_proto_instances_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestInstancesGroupConfigError.ProtoReflect.Descriptor instead.
func (*TestInstancesGroupConfigError) Descriptor() ([]byte, []int) {
	return file_pkg_instances_proto_instances_proto_rawDescGZIP(), []int{3}
}

func (x *TestInstancesGroupConfigError) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *TestInstancesGroupConfigError) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

type TestInstancesGroupConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool                             `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Errors []*TestInstancesGroupConfigError `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *TestInstancesGroupConfigResponse) Reset() {
	*x = TestInstancesGroupConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_instances_proto_instances_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestInstancesGroupConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestInstancesGroupConfigResponse) ProtoMessage() {}

func (x *TestInstancesGroupConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_instances_proto_instances_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestInstancesGroupConfigResponse.ProtoReflect.Descriptor instead.
func (*TestInstancesGroupConfigResponse) Descriptor() ([]byte, []int) {
	return file_pkg_instances_proto_instances_proto_rawDescGZIP(), []int{4}
}

func (x *TestInstancesGroupConfigResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *TestInstancesGroupConfigResponse) GetErrors() []*TestInstancesGroupConfigError {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_pkg_instances_proto_instances_proto protoreflect.FileDescriptor

var file_pkg_instances_proto_instances_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x68, 0x61, 0x73, 0x68,
	0x65, 0x72, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x68, 0x61, 0x73, 0x68,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x04, 0x0a, 0x08, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x45, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4e, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6e, 0x6f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x1a, 0x51, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x54, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4f,
	0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xe5, 0x04, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6e, 0x6f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x6f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x12, 0x54, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x52, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x1a, 0x51, 0x0a,
	0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x54, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4f, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5a, 0x0a, 0x1f, 0x54, 0x65, 0x73, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x6f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x22, 0x51, 0x0a, 0x1d, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x20, 0x54, 0x65, 0x73, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x48, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0xbc, 0x01,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6c, 0x6e, 0x74, 0x6f, 0x70, 0x70, 0x2f, 0x6e, 0x6f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x4e, 0x49, 0x58, 0xaa,
	0x02, 0x11, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0xca, 0x02, 0x11, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0xe2, 0x02, 0x1d, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_instances_proto_instances_proto_rawDescOnce sync.Once
	file_pkg_instances_proto_instances_proto_rawDescData = file_pkg_instances_proto_instances_proto_rawDesc
)

func file_pkg_instances_proto_instances_proto_rawDescGZIP() []byte {
	file_pkg_instances_proto_instances_proto_rawDescOnce.Do(func() {
		file_pkg_instances_proto_instances_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_instances_proto_instances_proto_rawDescData)
	})
	return file_pkg_instances_proto_instances_proto_rawDescData
}

var file_pkg_instances_proto_instances_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pkg_instances_proto_instances_proto_goTypes = []interface{}{
	(*Instance)(nil),                         // 0: nocloud.instances.Instance
	(*InstancesGroup)(nil),                   // 1: nocloud.instances.InstancesGroup
	(*TestInstancesGroupConfigRequest)(nil),  // 2: nocloud.instances.TestInstancesGroupConfigRequest
	(*TestInstancesGroupConfigError)(nil),    // 3: nocloud.instances.TestInstancesGroupConfigError
	(*TestInstancesGroupConfigResponse)(nil), // 4: nocloud.instances.TestInstancesGroupConfigResponse
	nil,                                      // 5: nocloud.instances.Instance.ConfigEntry
	nil,                                      // 6: nocloud.instances.Instance.ResourcesEntry
	nil,                                      // 7: nocloud.instances.Instance.DataEntry
	nil,                                      // 8: nocloud.instances.InstancesGroup.ConfigEntry
	nil,                                      // 9: nocloud.instances.InstancesGroup.ResourcesEntry
	nil,                                      // 10: nocloud.instances.InstancesGroup.DataEntry
	(*structpb.Value)(nil),                   // 11: google.protobuf.Value
}
var file_pkg_instances_proto_instances_proto_depIdxs = []int32{
	5,  // 0: nocloud.instances.Instance.config:type_name -> nocloud.instances.Instance.ConfigEntry
	6,  // 1: nocloud.instances.Instance.resources:type_name -> nocloud.instances.Instance.ResourcesEntry
	7,  // 2: nocloud.instances.Instance.data:type_name -> nocloud.instances.Instance.DataEntry
	8,  // 3: nocloud.instances.InstancesGroup.config:type_name -> nocloud.instances.InstancesGroup.ConfigEntry
	0,  // 4: nocloud.instances.InstancesGroup.instances:type_name -> nocloud.instances.Instance
	9,  // 5: nocloud.instances.InstancesGroup.resources:type_name -> nocloud.instances.InstancesGroup.ResourcesEntry
	10, // 6: nocloud.instances.InstancesGroup.data:type_name -> nocloud.instances.InstancesGroup.DataEntry
	1,  // 7: nocloud.instances.TestInstancesGroupConfigRequest.group:type_name -> nocloud.instances.InstancesGroup
	3,  // 8: nocloud.instances.TestInstancesGroupConfigResponse.errors:type_name -> nocloud.instances.TestInstancesGroupConfigError
	11, // 9: nocloud.instances.Instance.ConfigEntry.value:type_name -> google.protobuf.Value
	11, // 10: nocloud.instances.Instance.ResourcesEntry.value:type_name -> google.protobuf.Value
	11, // 11: nocloud.instances.Instance.DataEntry.value:type_name -> google.protobuf.Value
	11, // 12: nocloud.instances.InstancesGroup.ConfigEntry.value:type_name -> google.protobuf.Value
	11, // 13: nocloud.instances.InstancesGroup.ResourcesEntry.value:type_name -> google.protobuf.Value
	11, // 14: nocloud.instances.InstancesGroup.DataEntry.value:type_name -> google.protobuf.Value
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_pkg_instances_proto_instances_proto_init() }
func file_pkg_instances_proto_instances_proto_init() {
	if File_pkg_instances_proto_instances_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_instances_proto_instances_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instance); i {
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
		file_pkg_instances_proto_instances_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstancesGroup); i {
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
		file_pkg_instances_proto_instances_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestInstancesGroupConfigRequest); i {
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
		file_pkg_instances_proto_instances_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestInstancesGroupConfigError); i {
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
		file_pkg_instances_proto_instances_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestInstancesGroupConfigResponse); i {
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
			RawDescriptor: file_pkg_instances_proto_instances_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_instances_proto_instances_proto_goTypes,
		DependencyIndexes: file_pkg_instances_proto_instances_proto_depIdxs,
		MessageInfos:      file_pkg_instances_proto_instances_proto_msgTypes,
	}.Build()
	File_pkg_instances_proto_instances_proto = out.File
	file_pkg_instances_proto_instances_proto_rawDesc = nil
	file_pkg_instances_proto_instances_proto_goTypes = nil
	file_pkg_instances_proto_instances_proto_depIdxs = nil
}
