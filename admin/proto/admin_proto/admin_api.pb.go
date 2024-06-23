// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: admin_api.proto

package admin_proto

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

// 接口列表请求参数
type ApiListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *ListBaseReq `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Key  string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`   //接口键名
	Name string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"` //接口名称
	Path string       `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"` //接口路由
}

func (x *ApiListReq) Reset() {
	*x = ApiListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiListReq) ProtoMessage() {}

func (x *ApiListReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiListReq.ProtoReflect.Descriptor instead.
func (*ApiListReq) Descriptor() ([]byte, []int) {
	return file_admin_api_proto_rawDescGZIP(), []int{0}
}

func (x *ApiListReq) GetBase() *ListBaseReq {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *ApiListReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ApiListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApiListReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// 接口列表返回结构
type ApiItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Path         string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Key          string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Name         string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Enabled      bool   `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	PermissionId int32  `protobuf:"varint,6,opt,name=permissionId,proto3" json:"permissionId,omitempty"`
	CreatedAt    int64  `protobuf:"varint,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt    int64  `protobuf:"varint,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *ApiItem) Reset() {
	*x = ApiItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiItem) ProtoMessage() {}

func (x *ApiItem) ProtoReflect() protoreflect.Message {
	mi := &file_admin_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiItem.ProtoReflect.Descriptor instead.
func (*ApiItem) Descriptor() ([]byte, []int) {
	return file_admin_api_proto_rawDescGZIP(), []int{1}
}

func (x *ApiItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ApiItem) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ApiItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ApiItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApiItem) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *ApiItem) GetPermissionId() int32 {
	if x != nil {
		return x.PermissionId
	}
	return 0
}

func (x *ApiItem) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ApiItem) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

// 接口列表返回数据
type ApiListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64      `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Rows  []*ApiItem `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *ApiListResp) Reset() {
	*x = ApiListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiListResp) ProtoMessage() {}

func (x *ApiListResp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiListResp.ProtoReflect.Descriptor instead.
func (*ApiListResp) Descriptor() ([]byte, []int) {
	return file_admin_api_proto_rawDescGZIP(), []int{2}
}

func (x *ApiListResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ApiListResp) GetRows() []*ApiItem {
	if x != nil {
		return x.Rows
	}
	return nil
}

// 全部接口列表
type ApiAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApiAllReq) Reset() {
	*x = ApiAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiAllReq) ProtoMessage() {}

func (x *ApiAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiAllReq.ProtoReflect.Descriptor instead.
func (*ApiAllReq) Descriptor() ([]byte, []int) {
	return file_admin_api_proto_rawDescGZIP(), []int{3}
}

// 创建接口
type ApiAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`         //接口路由
	Key      string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`           //接口键名
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`         //接口名称
	Describe string `protobuf:"bytes,4,opt,name=describe,proto3" json:"describe,omitempty"` //接口描述
	Enabled  bool   `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`  //启用状态
}

func (x *ApiAddReq) Reset() {
	*x = ApiAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiAddReq) ProtoMessage() {}

func (x *ApiAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiAddReq.ProtoReflect.Descriptor instead.
func (*ApiAddReq) Descriptor() ([]byte, []int) {
	return file_admin_api_proto_rawDescGZIP(), []int{4}
}

func (x *ApiAddReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ApiAddReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ApiAddReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApiAddReq) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *ApiAddReq) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// 接口详情
type ApiInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //接口ID
}

func (x *ApiInfoReq) Reset() {
	*x = ApiInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiInfoReq) ProtoMessage() {}

func (x *ApiInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiInfoReq.ProtoReflect.Descriptor instead.
func (*ApiInfoReq) Descriptor() ([]byte, []int) {
	return file_admin_api_proto_rawDescGZIP(), []int{5}
}

func (x *ApiInfoReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 接口编辑
type ApiEditReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`            //接口ID
	Path     string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`         //接口路由
	Key      string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`           //接口键名
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`         //接口名称
	Describe string `protobuf:"bytes,5,opt,name=describe,proto3" json:"describe,omitempty"` //接口描述
	Enabled  bool   `protobuf:"varint,6,opt,name=enabled,proto3" json:"enabled,omitempty"`  //接口状态
}

func (x *ApiEditReq) Reset() {
	*x = ApiEditReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiEditReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiEditReq) ProtoMessage() {}

func (x *ApiEditReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiEditReq.ProtoReflect.Descriptor instead.
func (*ApiEditReq) Descriptor() ([]byte, []int) {
	return file_admin_api_proto_rawDescGZIP(), []int{6}
}

func (x *ApiEditReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ApiEditReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ApiEditReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ApiEditReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApiEditReq) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *ApiEditReq) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// 接口禁用启用
type ApiEnableReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`           //接口ID
	Enabled bool  `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"` //接口状态
}

func (x *ApiEnableReq) Reset() {
	*x = ApiEnableReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiEnableReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiEnableReq) ProtoMessage() {}

func (x *ApiEnableReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiEnableReq.ProtoReflect.Descriptor instead.
func (*ApiEnableReq) Descriptor() ([]byte, []int) {
	return file_admin_api_proto_rawDescGZIP(), []int{7}
}

func (x *ApiEnableReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ApiEnableReq) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// 删除接口
type ApiDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //接口ID
}

func (x *ApiDeleteReq) Reset() {
	*x = ApiDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDeleteReq) ProtoMessage() {}

func (x *ApiDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDeleteReq.ProtoReflect.Descriptor instead.
func (*ApiDeleteReq) Descriptor() ([]byte, []int) {
	return file_admin_api_proto_rawDescGZIP(), []int{8}
}

func (x *ApiDeleteReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_admin_api_proto protoreflect.FileDescriptor

var file_admin_api_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xcd, 0x01, 0x0a, 0x07, 0x41, 0x70, 0x69, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x47, 0x0a, 0x0b, 0x41, 0x70, 0x69, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x04, 0x72,
	0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x41, 0x70, 0x69, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22,
	0x0b, 0x0a, 0x09, 0x41, 0x70, 0x69, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x22, 0x7b, 0x0a, 0x09,
	0x41, 0x70, 0x69, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x41, 0x70, 0x69,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x45,
	0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x38, 0x0a, 0x0c, 0x41, 0x70, 0x69, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x22, 0x1e, 0x0a, 0x0c, 0x41, 0x70, 0x69, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_api_proto_rawDescOnce sync.Once
	file_admin_api_proto_rawDescData = file_admin_api_proto_rawDesc
)

func file_admin_api_proto_rawDescGZIP() []byte {
	file_admin_api_proto_rawDescOnce.Do(func() {
		file_admin_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_api_proto_rawDescData)
	})
	return file_admin_api_proto_rawDescData
}

var file_admin_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_admin_api_proto_goTypes = []interface{}{
	(*ApiListReq)(nil),   // 0: admin.ApiListReq
	(*ApiItem)(nil),      // 1: admin.ApiItem
	(*ApiListResp)(nil),  // 2: admin.ApiListResp
	(*ApiAllReq)(nil),    // 3: admin.ApiAllReq
	(*ApiAddReq)(nil),    // 4: admin.ApiAddReq
	(*ApiInfoReq)(nil),   // 5: admin.ApiInfoReq
	(*ApiEditReq)(nil),   // 6: admin.ApiEditReq
	(*ApiEnableReq)(nil), // 7: admin.ApiEnableReq
	(*ApiDeleteReq)(nil), // 8: admin.ApiDeleteReq
	(*ListBaseReq)(nil),  // 9: admin.ListBaseReq
}
var file_admin_api_proto_depIdxs = []int32{
	9, // 0: admin.ApiListReq.base:type_name -> admin.ListBaseReq
	1, // 1: admin.ApiListResp.rows:type_name -> admin.ApiItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_admin_api_proto_init() }
func file_admin_api_proto_init() {
	if File_admin_api_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_admin_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiListReq); i {
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
		file_admin_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiItem); i {
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
		file_admin_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiListResp); i {
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
		file_admin_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiAllReq); i {
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
		file_admin_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiAddReq); i {
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
		file_admin_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiInfoReq); i {
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
		file_admin_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiEditReq); i {
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
		file_admin_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiEnableReq); i {
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
		file_admin_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDeleteReq); i {
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
			RawDescriptor: file_admin_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_api_proto_goTypes,
		DependencyIndexes: file_admin_api_proto_depIdxs,
		MessageInfos:      file_admin_api_proto_msgTypes,
	}.Build()
	File_admin_api_proto = out.File
	file_admin_api_proto_rawDesc = nil
	file_admin_api_proto_goTypes = nil
	file_admin_api_proto_depIdxs = nil
}
