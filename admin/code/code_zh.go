package code

import (
	"admin/proto/code_proto"
)

var zhMsg = map[code_proto.ErrorCode]string{
	code_proto.ErrorCode_Success:                      "成功",
	code_proto.ErrorCode_Error:                        "请求失败",
	code_proto.ErrorCode_ReadContextRequestBodyFailed: "读取请求体数据失败",
	code_proto.ErrorCode_RecordNotExist:               "记录不存在",
	code_proto.ErrorCode_RecordNValidCanNotDeleted:    "生效中的记录不能删除",

	code_proto.ErrorCode_AuthTokenFailed:         "未登录或登录令牌已过期",
	code_proto.ErrorCode_AuthTokenInvalid:        "登录令牌无效",
	code_proto.ErrorCode_AuthTokenInspectInvalid: "登录令牌检查失败",
	code_proto.ErrorCode_AuthTokenInfoInvalid:    "登录令牌信息无效",

	code_proto.ErrorCode_RequestBodyInvalid:   "请求体参数无效",
	code_proto.ErrorCode_RequestQueryInvalid:  "查询参数无效",
	code_proto.ErrorCode_RequestParamsInvalid: "请求参数无效",

	code_proto.ErrorCode_AdminAccountPasswordInvalid: "账号或密码错误",
	code_proto.ErrorCode_AdminAccountNotExist:        "账号不存在或已被删除",
	code_proto.ErrorCode_AdminAccountInvalid:         "账号无效",

	code_proto.ErrorCode_AdminApiNameExist: "接口名称已存在",
	code_proto.ErrorCode_AdminApiPathExist: "接口路径已存在",
	code_proto.ErrorCode_AdminApiKeyExist:  "接口键名已存在",
	code_proto.ErrorCode_AdminApiNotExist:  "接口不存在货已失效",

	code_proto.ErrorCode_AdminPermissionKeyExist:    "权限键名已存在",
	code_proto.ErrorCode_AdminPermissionExist:       "该菜单下权限已存在",
	code_proto.ErrorCode_AdminPermissionKeyNeed:     "权限键名必填",
	code_proto.ErrorCode_AdminPermissionNameNeed:    "权限名称必填",
	code_proto.ErrorCode_AdminPermissionTypeInvalid: "权限类型无效",

	code_proto.ErrorCode_AdminMenuNotExist: "菜单不存在",
}
