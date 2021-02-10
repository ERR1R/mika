// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/role.proto

package rpc

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

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId          uint32    `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	RoleName        string    `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	RemoteId        uint64    `protobuf:"varint,3,opt,name=remote_id,json=remoteId,proto3" json:"remote_id,omitempty"`
	Priority        int32     `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	DownloadEnabled bool      `protobuf:"varint,5,opt,name=download_enabled,json=downloadEnabled,proto3" json:"download_enabled,omitempty"`
	UploadEnabled   bool      `protobuf:"varint,6,opt,name=upload_enabled,json=uploadEnabled,proto3" json:"upload_enabled,omitempty"`
	MultiUp         float64   `protobuf:"fixed64,7,opt,name=multi_up,json=multiUp,proto3" json:"multi_up,omitempty"`
	MultiDown       float64   `protobuf:"fixed64,8,opt,name=multi_down,json=multiDown,proto3" json:"multi_down,omitempty"`
	Time            *TimeMeta `protobuf:"bytes,9,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_proto_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_proto_role_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetRoleId() uint32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *Role) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *Role) GetRemoteId() uint64 {
	if x != nil {
		return x.RemoteId
	}
	return 0
}

func (x *Role) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Role) GetDownloadEnabled() bool {
	if x != nil {
		return x.DownloadEnabled
	}
	return false
}

func (x *Role) GetUploadEnabled() bool {
	if x != nil {
		return x.UploadEnabled
	}
	return false
}

func (x *Role) GetMultiUp() float64 {
	if x != nil {
		return x.MultiUp
	}
	return 0
}

func (x *Role) GetMultiDown() float64 {
	if x != nil {
		return x.MultiDown
	}
	return 0
}

func (x *Role) GetTime() *TimeMeta {
	if x != nil {
		return x.Time
	}
	return nil
}

type RoleID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId   uint32 `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	RoleName string `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
}

func (x *RoleID) Reset() {
	*x = RoleID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleID) ProtoMessage() {}

func (x *RoleID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleID.ProtoReflect.Descriptor instead.
func (*RoleID) Descriptor() ([]byte, []int) {
	return file_proto_role_proto_rawDescGZIP(), []int{1}
}

func (x *RoleID) GetRoleId() uint32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *RoleID) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type RoleAddParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleName        string  `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	RemoteId        uint64  `protobuf:"varint,3,opt,name=remote_id,json=remoteId,proto3" json:"remote_id,omitempty"`
	Priority        int32   `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	DownloadEnabled bool    `protobuf:"varint,5,opt,name=download_enabled,json=downloadEnabled,proto3" json:"download_enabled,omitempty"`
	UploadEnabled   bool    `protobuf:"varint,6,opt,name=upload_enabled,json=uploadEnabled,proto3" json:"upload_enabled,omitempty"`
	MultiUp         float64 `protobuf:"fixed64,7,opt,name=multi_up,json=multiUp,proto3" json:"multi_up,omitempty"`
	MultiDown       float64 `protobuf:"fixed64,8,opt,name=multi_down,json=multiDown,proto3" json:"multi_down,omitempty"`
}

func (x *RoleAddParams) Reset() {
	*x = RoleAddParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleAddParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleAddParams) ProtoMessage() {}

func (x *RoleAddParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleAddParams.ProtoReflect.Descriptor instead.
func (*RoleAddParams) Descriptor() ([]byte, []int) {
	return file_proto_role_proto_rawDescGZIP(), []int{2}
}

func (x *RoleAddParams) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *RoleAddParams) GetRemoteId() uint64 {
	if x != nil {
		return x.RemoteId
	}
	return 0
}

func (x *RoleAddParams) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *RoleAddParams) GetDownloadEnabled() bool {
	if x != nil {
		return x.DownloadEnabled
	}
	return false
}

func (x *RoleAddParams) GetUploadEnabled() bool {
	if x != nil {
		return x.UploadEnabled
	}
	return false
}

func (x *RoleAddParams) GetMultiUp() float64 {
	if x != nil {
		return x.MultiUp
	}
	return 0
}

func (x *RoleAddParams) GetMultiDown() float64 {
	if x != nil {
		return x.MultiDown
	}
	return 0
}

type RoleSetParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdatedKeys     []string `protobuf:"bytes,1,rep,name=updated_keys,json=updatedKeys,proto3" json:"updated_keys,omitempty"`
	RoleName        string   `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	RemoteId        uint64   `protobuf:"varint,3,opt,name=remote_id,json=remoteId,proto3" json:"remote_id,omitempty"`
	Priority        int32    `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	DownloadEnabled bool     `protobuf:"varint,5,opt,name=download_enabled,json=downloadEnabled,proto3" json:"download_enabled,omitempty"`
	UploadEnabled   bool     `protobuf:"varint,6,opt,name=upload_enabled,json=uploadEnabled,proto3" json:"upload_enabled,omitempty"`
	MultiUp         float64  `protobuf:"fixed64,7,opt,name=multi_up,json=multiUp,proto3" json:"multi_up,omitempty"`
	MultiDown       float64  `protobuf:"fixed64,8,opt,name=multi_down,json=multiDown,proto3" json:"multi_down,omitempty"`
}

func (x *RoleSetParams) Reset() {
	*x = RoleSetParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleSetParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleSetParams) ProtoMessage() {}

func (x *RoleSetParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleSetParams.ProtoReflect.Descriptor instead.
func (*RoleSetParams) Descriptor() ([]byte, []int) {
	return file_proto_role_proto_rawDescGZIP(), []int{3}
}

func (x *RoleSetParams) GetUpdatedKeys() []string {
	if x != nil {
		return x.UpdatedKeys
	}
	return nil
}

func (x *RoleSetParams) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *RoleSetParams) GetRemoteId() uint64 {
	if x != nil {
		return x.RemoteId
	}
	return 0
}

func (x *RoleSetParams) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *RoleSetParams) GetDownloadEnabled() bool {
	if x != nil {
		return x.DownloadEnabled
	}
	return false
}

func (x *RoleSetParams) GetUploadEnabled() bool {
	if x != nil {
		return x.UploadEnabled
	}
	return false
}

func (x *RoleSetParams) GetMultiUp() float64 {
	if x != nil {
		return x.MultiUp
	}
	return 0
}

func (x *RoleSetParams) GetMultiDown() float64 {
	if x != nil {
		return x.MultiDown
	}
	return 0
}

var File_proto_role_proto protoreflect.FileDescriptor

var file_proto_role_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6d, 0x69, 0x6b, 0x61, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x02, 0x0a,
	0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f,
	0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x55,
	0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x44, 0x6f, 0x77, 0x6e,
	0x12, 0x22, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6d, 0x69, 0x6b, 0x61, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xf1, 0x01, 0x0a, 0x0d, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x64, 0x64,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x10,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x55, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x44, 0x6f, 0x77, 0x6e, 0x22, 0x94, 0x02, 0x0a, 0x0d, 0x52, 0x6f, 0x6c,
	0x65, 0x53, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x75,
	0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x55, 0x70,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x44, 0x6f, 0x77, 0x6e, 0x42,
	0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65,
	0x69, 0x67, 0x68, 0x6d, 0x61, 0x63, 0x64, 0x6f, 0x6e, 0x61, 0x6c, 0x64, 0x2f, 0x6d, 0x69, 0x6b,
	0x61, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_role_proto_rawDescOnce sync.Once
	file_proto_role_proto_rawDescData = file_proto_role_proto_rawDesc
)

func file_proto_role_proto_rawDescGZIP() []byte {
	file_proto_role_proto_rawDescOnce.Do(func() {
		file_proto_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_role_proto_rawDescData)
	})
	return file_proto_role_proto_rawDescData
}

var file_proto_role_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_role_proto_goTypes = []interface{}{
	(*Role)(nil),          // 0: mika.Role
	(*RoleID)(nil),        // 1: mika.RoleID
	(*RoleAddParams)(nil), // 2: mika.RoleAddParams
	(*RoleSetParams)(nil), // 3: mika.RoleSetParams
	(*TimeMeta)(nil),      // 4: mika.TimeMeta
}
var file_proto_role_proto_depIdxs = []int32{
	4, // 0: mika.Role.time:type_name -> mika.TimeMeta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_role_proto_init() }
func file_proto_role_proto_init() {
	if File_proto_role_proto != nil {
		return
	}
	file_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_proto_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleID); i {
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
		file_proto_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleAddParams); i {
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
		file_proto_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleSetParams); i {
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
			RawDescriptor: file_proto_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_role_proto_goTypes,
		DependencyIndexes: file_proto_role_proto_depIdxs,
		MessageInfos:      file_proto_role_proto_msgTypes,
	}.Build()
	File_proto_role_proto = out.File
	file_proto_role_proto_rawDesc = nil
	file_proto_role_proto_goTypes = nil
	file_proto_role_proto_depIdxs = nil
}
