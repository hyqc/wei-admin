package code

import "admin/proto/code_proto"

var enMsg = map[code_proto.ErrorCode]string{
	code_proto.ErrorCode_Success:                      "success",
	code_proto.ErrorCode_Error:                        "error",
	code_proto.ErrorCode_ReadContextRequestBodyFailed: "read context request body failed",
	code_proto.ErrorCode_RecordNotExist:               "record not exist",
	code_proto.ErrorCode_RecordNValidCanNotDeleted:    "record valid, can not be deleted",

	code_proto.ErrorCode_AuthTokenFailed:         "token expired or not login",
	code_proto.ErrorCode_AuthTokenInvalid:        "token invalid",
	code_proto.ErrorCode_AuthTokenInspectInvalid: "token inspect invalid",
	code_proto.ErrorCode_AuthTokenInfoInvalid:    "token info invalid",

	code_proto.ErrorCode_RequestBodyInvalid:   "request body invalid",
	code_proto.ErrorCode_RequestQueryInvalid:  "request query invalid",
	code_proto.ErrorCode_RequestParamsInvalid: "request params invalid",

	code_proto.ErrorCode_AdminAccountPasswordInvalid: "account or pwd invalid",
	code_proto.ErrorCode_AdminAccountNotExist:        "account not exist or deleted",
	code_proto.ErrorCode_AdminAccountInvalid:         "account invalid",

	code_proto.ErrorCode_AdminApiNameExist: "api name exist",
	code_proto.ErrorCode_AdminApiPathExist: "api path exist",
	code_proto.ErrorCode_AdminApiKeyExist:  "api key exist",
	code_proto.ErrorCode_AdminApiNotExist:  "api not exist or invalid",

	code_proto.ErrorCode_AdminPermissionKeyExist:    "permission key exist",
	code_proto.ErrorCode_AdminPermissionExist:       "permission exist",
	code_proto.ErrorCode_AdminPermissionKeyNeed:     "permission key need",
	code_proto.ErrorCode_AdminPermissionNameNeed:    "permission name need",
	code_proto.ErrorCode_AdminPermissionTypeInvalid: "permission type invalid",

	code_proto.ErrorCode_AdminMenuNotExist: "menu not exist",
}
