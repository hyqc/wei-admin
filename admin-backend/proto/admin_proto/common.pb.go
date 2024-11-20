// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: common.proto

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

type RespCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Reason string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *RespCode) Reset() {
	*x = RespCode{}
	mi := &file_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RespCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespCode) ProtoMessage() {}

func (x *RespCode) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespCode.ProtoReflect.Descriptor instead.
func (*RespCode) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *RespCode) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RespCode) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RespCode) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type ReqListBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize        int32  `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`               // 分页偏移量
	PageNum         int32  `protobuf:"varint,2,opt,name=pageNum,proto3" json:"pageNum,omitempty"`                 // 页码
	SortField       string `protobuf:"bytes,3,opt,name=sortField,proto3" json:"sortField,omitempty"`              // 排序字段
	SortType        string `protobuf:"bytes,4,opt,name=sortType,proto3" json:"sortType,omitempty"`                // 排序值
	Enabled         int32  `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`                 // 0：全部 1：启用 2：禁用
	CreateStartTime int64  `protobuf:"varint,6,opt,name=createStartTime,proto3" json:"createStartTime,omitempty"` // 查询开始时间戳秒
	CreateEndTime   int64  `protobuf:"varint,7,opt,name=createEndTime,proto3" json:"createEndTime,omitempty"`     //查询结束时间戳秒
}

func (x *ReqListBase) Reset() {
	*x = ReqListBase{}
	mi := &file_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReqListBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqListBase) ProtoMessage() {}

func (x *ReqListBase) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqListBase.ProtoReflect.Descriptor instead.
func (*ReqListBase) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *ReqListBase) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ReqListBase) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ReqListBase) GetSortField() string {
	if x != nil {
		return x.SortField
	}
	return ""
}

func (x *ReqListBase) GetSortType() string {
	if x != nil {
		return x.SortType
	}
	return ""
}

func (x *ReqListBase) GetEnabled() int32 {
	if x != nil {
		return x.Enabled
	}
	return 0
}

func (x *ReqListBase) GetCreateStartTime() int64 {
	if x != nil {
		return x.CreateStartTime
	}
	return 0
}

func (x *ReqListBase) GetCreateEndTime() int64 {
	if x != nil {
		return x.CreateEndTime
	}
	return 0
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
	CreatedAt    string `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt    string `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *ApiItem) Reset() {
	*x = ApiItem{}
	mi := &file_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiItem) ProtoMessage() {}

func (x *ApiItem) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if x != nil {
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
	return file_common_proto_rawDescGZIP(), []int{2}
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

func (x *ApiItem) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ApiItem) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x48, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22,
	0xe7, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xcd, 0x01, 0x0a, 0x07, 0x41, 0x70,
	0x69, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_proto_goTypes = []any{
	(*RespCode)(nil),    // 0: admin.RespCode
	(*ReqListBase)(nil), // 1: admin.ReqListBase
	(*ApiItem)(nil),     // 2: admin.ApiItem
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
