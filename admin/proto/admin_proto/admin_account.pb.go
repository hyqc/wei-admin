// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: admin_account.proto

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

// 管理员账号详情
type AdminInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId       int32                `protobuf:"varint,1,opt,name=adminId,proto3" json:"adminId,omitempty"`
	Username      string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Nickname      string               `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar        string               `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Email         string               `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	CreateTime    string               `protobuf:"bytes,6,opt,name=createTime,proto3" json:"createTime,omitempty"`
	ModifyTime    string               `protobuf:"bytes,7,opt,name=modifyTime,proto3" json:"modifyTime,omitempty"`
	LastLoginTime string               `protobuf:"bytes,8,opt,name=lastLoginTime,proto3" json:"lastLoginTime,omitempty"`
	LastLoginIp   string               `protobuf:"bytes,9,opt,name=lastLoginIp,proto3" json:"lastLoginIp,omitempty"`
	LoginTotal    int32                `protobuf:"varint,10,opt,name=loginTotal,proto3" json:"loginTotal,omitempty"`
	Enabled       bool                 `protobuf:"varint,11,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Token         string               `protobuf:"bytes,12,opt,name=token,proto3" json:"token,omitempty"`
	Expire        int32                `protobuf:"varint,13,opt,name=expire,proto3" json:"expire,omitempty"`
	Menus         map[string]*MenuItem `protobuf:"bytes,14,rep,name=menus,proto3" json:"menus,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Permissions   map[string]string    `protobuf:"bytes,15,rep,name=permissions,proto3" json:"permissions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AdminInfo) Reset() {
	*x = AdminInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminInfo) ProtoMessage() {}

func (x *AdminInfo) ProtoReflect() protoreflect.Message {
	mi := &file_admin_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminInfo.ProtoReflect.Descriptor instead.
func (*AdminInfo) Descriptor() ([]byte, []int) {
	return file_admin_account_proto_rawDescGZIP(), []int{0}
}

func (x *AdminInfo) GetAdminId() int32 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

func (x *AdminInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AdminInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AdminInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *AdminInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AdminInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *AdminInfo) GetModifyTime() string {
	if x != nil {
		return x.ModifyTime
	}
	return ""
}

func (x *AdminInfo) GetLastLoginTime() string {
	if x != nil {
		return x.LastLoginTime
	}
	return ""
}

func (x *AdminInfo) GetLastLoginIp() string {
	if x != nil {
		return x.LastLoginIp
	}
	return ""
}

func (x *AdminInfo) GetLoginTotal() int32 {
	if x != nil {
		return x.LoginTotal
	}
	return 0
}

func (x *AdminInfo) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *AdminInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AdminInfo) GetExpire() int32 {
	if x != nil {
		return x.Expire
	}
	return 0
}

func (x *AdminInfo) GetMenus() map[string]*MenuItem {
	if x != nil {
		return x.Menus
	}
	return nil
}

func (x *AdminInfo) GetPermissions() map[string]string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

// 登录
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"` // 用户名
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // 密码
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_admin_account_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// 账号详情
type AccountDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken bool `protobuf:"varint,1,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"` //刷新token
}

func (x *AccountDetailReq) Reset() {
	*x = AccountDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountDetailReq) ProtoMessage() {}

func (x *AccountDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountDetailReq.ProtoReflect.Descriptor instead.
func (*AccountDetailReq) Descriptor() ([]byte, []int) {
	return file_admin_account_proto_rawDescGZIP(), []int{2}
}

func (x *AccountDetailReq) GetRefreshToken() bool {
	if x != nil {
		return x.RefreshToken
	}
	return false
}

// 账号编辑
type AccountEditReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"` //名称
	Avatar   string `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`     //头像
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`       //邮箱
}

func (x *AccountEditReq) Reset() {
	*x = AccountEditReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountEditReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountEditReq) ProtoMessage() {}

func (x *AccountEditReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountEditReq.ProtoReflect.Descriptor instead.
func (*AccountEditReq) Descriptor() ([]byte, []int) {
	return file_admin_account_proto_rawDescGZIP(), []int{3}
}

func (x *AccountEditReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AccountEditReq) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *AccountEditReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AccountEditResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AccountEditResp) Reset() {
	*x = AccountEditResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountEditResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountEditResp) ProtoMessage() {}

func (x *AccountEditResp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountEditResp.ProtoReflect.Descriptor instead.
func (*AccountEditResp) Descriptor() ([]byte, []int) {
	return file_admin_account_proto_rawDescGZIP(), []int{4}
}

// 修改密码
type AccountPasswordEditReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldPassword     string `protobuf:"bytes,1,opt,name=oldPassword,proto3" json:"oldPassword,omitempty"`         //旧密码
	Password        string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`               //新密码
	ConfirmPassword string `protobuf:"bytes,3,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty"` //确认密码
}

func (x *AccountPasswordEditReq) Reset() {
	*x = AccountPasswordEditReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountPasswordEditReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountPasswordEditReq) ProtoMessage() {}

func (x *AccountPasswordEditReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountPasswordEditReq.ProtoReflect.Descriptor instead.
func (*AccountPasswordEditReq) Descriptor() ([]byte, []int) {
	return file_admin_account_proto_rawDescGZIP(), []int{5}
}

func (x *AccountPasswordEditReq) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *AccountPasswordEditReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AccountPasswordEditReq) GetConfirmPassword() string {
	if x != nil {
		return x.ConfirmPassword
	}
	return ""
}

type AccountPasswordEditResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AccountPasswordEditResp) Reset() {
	*x = AccountPasswordEditResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountPasswordEditResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountPasswordEditResp) ProtoMessage() {}

func (x *AccountPasswordEditResp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountPasswordEditResp.ProtoReflect.Descriptor instead.
func (*AccountPasswordEditResp) Descriptor() ([]byte, []int) {
	return file_admin_account_proto_rawDescGZIP(), []int{6}
}

// 个人权限
type AccountPermissionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuId int32 `protobuf:"varint,1,opt,name=menuId,proto3" json:"menuId,omitempty"` //菜单ID
}

func (x *AccountPermissionReq) Reset() {
	*x = AccountPermissionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountPermissionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountPermissionReq) ProtoMessage() {}

func (x *AccountPermissionReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountPermissionReq.ProtoReflect.Descriptor instead.
func (*AccountPermissionReq) Descriptor() ([]byte, []int) {
	return file_admin_account_proto_rawDescGZIP(), []int{7}
}

func (x *AccountPermissionReq) GetMenuId() int32 {
	if x != nil {
		return x.MenuId
	}
	return 0
}

type AccountPermissionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AccountPermissionResp) Reset() {
	*x = AccountPermissionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountPermissionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountPermissionResp) ProtoMessage() {}

func (x *AccountPermissionResp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountPermissionResp.ProtoReflect.Descriptor instead.
func (*AccountPermissionResp) Descriptor() ([]byte, []int) {
	return file_admin_account_proto_rawDescGZIP(), []int{8}
}

var File_admin_account_proto protoreflect.FileDescriptor

var file_admin_account_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x10, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe,
	0x04, 0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49,
	0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x49, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x31, 0x0a, 0x05,
	0x6d, 0x65, 0x6e, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x65,
	0x6e, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x12,
	0x43, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x49, 0x0a, 0x0a, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x6e, 0x75,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3e, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x42, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x36, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5a, 0x0a, 0x0e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x11, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x80, 0x01, 0x0a, 0x16, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x64,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x19, 0x0a,
	0x17, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2e, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_account_proto_rawDescOnce sync.Once
	file_admin_account_proto_rawDescData = file_admin_account_proto_rawDesc
)

func file_admin_account_proto_rawDescGZIP() []byte {
	file_admin_account_proto_rawDescOnce.Do(func() {
		file_admin_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_account_proto_rawDescData)
	})
	return file_admin_account_proto_rawDescData
}

var file_admin_account_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_admin_account_proto_goTypes = []interface{}{
	(*AdminInfo)(nil),               // 0: admin.AdminInfo
	(*LoginReq)(nil),                // 1: admin.LoginReq
	(*AccountDetailReq)(nil),        // 2: admin.AccountDetailReq
	(*AccountEditReq)(nil),          // 3: admin.AccountEditReq
	(*AccountEditResp)(nil),         // 4: admin.AccountEditResp
	(*AccountPasswordEditReq)(nil),  // 5: admin.AccountPasswordEditReq
	(*AccountPasswordEditResp)(nil), // 6: admin.AccountPasswordEditResp
	(*AccountPermissionReq)(nil),    // 7: admin.AccountPermissionReq
	(*AccountPermissionResp)(nil),   // 8: admin.AccountPermissionResp
	nil,                             // 9: admin.AdminInfo.MenusEntry
	nil,                             // 10: admin.AdminInfo.PermissionsEntry
	(*MenuItem)(nil),                // 11: admin.MenuItem
}
var file_admin_account_proto_depIdxs = []int32{
	9,  // 0: admin.AdminInfo.menus:type_name -> admin.AdminInfo.MenusEntry
	10, // 1: admin.AdminInfo.permissions:type_name -> admin.AdminInfo.PermissionsEntry
	11, // 2: admin.AdminInfo.MenusEntry.value:type_name -> admin.MenuItem
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_admin_account_proto_init() }
func file_admin_account_proto_init() {
	if File_admin_account_proto != nil {
		return
	}
	file_admin_menu_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_admin_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminInfo); i {
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
		file_admin_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_admin_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountDetailReq); i {
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
		file_admin_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountEditReq); i {
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
		file_admin_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountEditResp); i {
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
		file_admin_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountPasswordEditReq); i {
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
		file_admin_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountPasswordEditResp); i {
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
		file_admin_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountPermissionReq); i {
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
		file_admin_account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountPermissionResp); i {
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
			RawDescriptor: file_admin_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_account_proto_goTypes,
		DependencyIndexes: file_admin_account_proto_depIdxs,
		MessageInfos:      file_admin_account_proto_msgTypes,
	}.Build()
	File_admin_account_proto = out.File
	file_admin_account_proto_rawDesc = nil
	file_admin_account_proto_goTypes = nil
	file_admin_account_proto_depIdxs = nil
}
