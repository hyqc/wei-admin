// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: error_code.proto

package code_proto

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

type ErrorCode int32

const (
	ErrorCode_Success                      ErrorCode = 0      // 成功
	ErrorCode_Error                        ErrorCode = 1      // 失败
	ErrorCode_ReadContextRequestBodyFailed ErrorCode = 100001 // 获取请求体内容失败
	ErrorCode_RecordNotExist               ErrorCode = 100002 // 记录不存在
	ErrorCode_RecordNValidCanNotDeleted    ErrorCode = 100003 // 生效中的记录不能删除
	ErrorCode_AuthTokenFailed              ErrorCode = 200001 // 鉴权失败
	ErrorCode_AuthTokenInvalid             ErrorCode = 200002 // BearerToken无效
	ErrorCode_AuthTokenInspectInvalid      ErrorCode = 200003 // BearerToken无效
	ErrorCode_AuthTokenInfoInvalid         ErrorCode = 200004 // Token信息无效
	ErrorCode_AuthTokenForbidden           ErrorCode = 200005 // 没有访问权限
	ErrorCode_RequestBodyInvalid           ErrorCode = 300001 // 请求体参数无效
	ErrorCode_RequestQueryInvalid          ErrorCode = 300002 // 请求查询参数无效
	ErrorCode_RequestParamsInvalid         ErrorCode = 300003 // 请求参数无效
	ErrorCode_AdminAccountPasswordInvalid  ErrorCode = 400001 // 账号或密码错误
	ErrorCode_AdminAccountNotExist         ErrorCode = 400002 // 账号不存在或已被删除
	ErrorCode_AdminAccountInvalid          ErrorCode = 400003 // 账号已被封禁
	ErrorCode_AdminApiNameExist            ErrorCode = 500001 // 接口名称已存在
	ErrorCode_AdminApiPathExist            ErrorCode = 500002 // 接口路径已存在
	ErrorCode_AdminApiKeyExist             ErrorCode = 500003 // 接口键名已存在
	ErrorCode_AdminApiNotExist             ErrorCode = 500004 // 接口不存在货已失效
	ErrorCode_AdminPermissionKeyExist      ErrorCode = 600001 // 权限键名已存在
	ErrorCode_AdminPermissionExist         ErrorCode = 600002 // 该菜单下权限已存在
	ErrorCode_AdminPermissionKeyNeed       ErrorCode = 600003 // 权限键名不能为空
	ErrorCode_AdminPermissionNameNeed      ErrorCode = 600004 // 权限名称不能为空
	ErrorCode_AdminPermissionTypeInvalid   ErrorCode = 600005 // 权限类型无效
	ErrorCode_AdminMenuNotExist            ErrorCode = 700001 // 菜单不存在
	ErrorCode_AdminRoleNotExist            ErrorCode = 800001 // 角色不存在
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:      "Success",
		1:      "Error",
		100001: "ReadContextRequestBodyFailed",
		100002: "RecordNotExist",
		100003: "RecordNValidCanNotDeleted",
		200001: "AuthTokenFailed",
		200002: "AuthTokenInvalid",
		200003: "AuthTokenInspectInvalid",
		200004: "AuthTokenInfoInvalid",
		200005: "AuthTokenForbidden",
		300001: "RequestBodyInvalid",
		300002: "RequestQueryInvalid",
		300003: "RequestParamsInvalid",
		400001: "AdminAccountPasswordInvalid",
		400002: "AdminAccountNotExist",
		400003: "AdminAccountInvalid",
		500001: "AdminApiNameExist",
		500002: "AdminApiPathExist",
		500003: "AdminApiKeyExist",
		500004: "AdminApiNotExist",
		600001: "AdminPermissionKeyExist",
		600002: "AdminPermissionExist",
		600003: "AdminPermissionKeyNeed",
		600004: "AdminPermissionNameNeed",
		600005: "AdminPermissionTypeInvalid",
		700001: "AdminMenuNotExist",
		800001: "AdminRoleNotExist",
	}
	ErrorCode_value = map[string]int32{
		"Success":                      0,
		"Error":                        1,
		"ReadContextRequestBodyFailed": 100001,
		"RecordNotExist":               100002,
		"RecordNValidCanNotDeleted":    100003,
		"AuthTokenFailed":              200001,
		"AuthTokenInvalid":             200002,
		"AuthTokenInspectInvalid":      200003,
		"AuthTokenInfoInvalid":         200004,
		"AuthTokenForbidden":           200005,
		"RequestBodyInvalid":           300001,
		"RequestQueryInvalid":          300002,
		"RequestParamsInvalid":         300003,
		"AdminAccountPasswordInvalid":  400001,
		"AdminAccountNotExist":         400002,
		"AdminAccountInvalid":          400003,
		"AdminApiNameExist":            500001,
		"AdminApiPathExist":            500002,
		"AdminApiKeyExist":             500003,
		"AdminApiNotExist":             500004,
		"AdminPermissionKeyExist":      600001,
		"AdminPermissionExist":         600002,
		"AdminPermissionKeyNeed":       600003,
		"AdminPermissionNameNeed":      600004,
		"AdminPermissionTypeInvalid":   600005,
		"AdminMenuNotExist":            700001,
		"AdminRoleNotExist":            800001,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_error_code_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_error_code_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_error_code_proto_rawDescGZIP(), []int{0}
}

var File_error_code_proto protoreflect.FileDescriptor

var file_error_code_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2a, 0xdb,
	0x05, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1c, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x46, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x10, 0xa1, 0x8d, 0x06, 0x12, 0x14, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xa2, 0x8d, 0x06, 0x12, 0x1f,
	0x0a, 0x19, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x61,
	0x6e, 0x4e, 0x6f, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0xa3, 0x8d, 0x06, 0x12,
	0x15, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x10, 0xc1, 0x9a, 0x0c, 0x12, 0x16, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0xc2, 0x9a, 0x0c, 0x12, 0x1d,
	0x0a, 0x17, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0xc3, 0x9a, 0x0c, 0x12, 0x1a, 0x0a,
	0x14, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0xc4, 0x9a, 0x0c, 0x12, 0x18, 0x0a, 0x12, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x10,
	0xc5, 0x9a, 0x0c, 0x12, 0x18, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f,
	0x64, 0x79, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0xe1, 0xa7, 0x12, 0x12, 0x19, 0x0a,
	0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x10, 0xe2, 0xa7, 0x12, 0x12, 0x1a, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x10, 0xe3, 0xa7, 0x12, 0x12, 0x21, 0x0a, 0x1b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x10, 0x81, 0xb5, 0x18, 0x12, 0x1a, 0x0a, 0x14, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10,
	0x82, 0xb5, 0x18, 0x12, 0x19, 0x0a, 0x13, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x83, 0xb5, 0x18, 0x12, 0x17,
	0x0a, 0x11, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x10, 0xa1, 0xc2, 0x1e, 0x12, 0x17, 0x0a, 0x11, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x41, 0x70, 0x69, 0x50, 0x61, 0x74, 0x68, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xa2, 0xc2, 0x1e,
	0x12, 0x16, 0x0a, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x10, 0xa3, 0xc2, 0x1e, 0x12, 0x16, 0x0a, 0x10, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x41, 0x70, 0x69, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xa4, 0xc2, 0x1e,
	0x12, 0x1d, 0x0a, 0x17, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xc1, 0xcf, 0x24, 0x12,
	0x1a, 0x0a, 0x14, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xc2, 0xcf, 0x24, 0x12, 0x1c, 0x0a, 0x16, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65,
	0x79, 0x4e, 0x65, 0x65, 0x64, 0x10, 0xc3, 0xcf, 0x24, 0x12, 0x1d, 0x0a, 0x17, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x4e, 0x65, 0x65, 0x64, 0x10, 0xc4, 0xcf, 0x24, 0x12, 0x20, 0x0a, 0x1a, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0xc5, 0xcf, 0x24, 0x12, 0x17, 0x0a, 0x11, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x4d, 0x65, 0x6e, 0x75, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10,
	0xe1, 0xdc, 0x2a, 0x12, 0x17, 0x0a, 0x11, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65,
	0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0x81, 0xea, 0x30, 0x42, 0x19, 0x5a, 0x17,
	0x2e, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x6f, 0x64,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_code_proto_rawDescOnce sync.Once
	file_error_code_proto_rawDescData = file_error_code_proto_rawDesc
)

func file_error_code_proto_rawDescGZIP() []byte {
	file_error_code_proto_rawDescOnce.Do(func() {
		file_error_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_code_proto_rawDescData)
	})
	return file_error_code_proto_rawDescData
}

var file_error_code_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_code_proto_goTypes = []any{
	(ErrorCode)(0), // 0: error_code.ErrorCode
}
var file_error_code_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_code_proto_init() }
func file_error_code_proto_init() {
	if File_error_code_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_code_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_code_proto_goTypes,
		DependencyIndexes: file_error_code_proto_depIdxs,
		EnumInfos:         file_error_code_proto_enumTypes,
	}.Build()
	File_error_code_proto = out.File
	file_error_code_proto_rawDesc = nil
	file_error_code_proto_goTypes = nil
	file_error_code_proto_depIdxs = nil
}
