package common

import (
	"admin/code"
	"admin/config"
	"admin/constant"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	DESC                           = "descend"
	ASC                            = "ascend"
	EnabledAllQueryValue     int32 = 0
	EnabledValidQueryValue   int32 = 1
	EnabledInvalidQueryValue int32 = 2
)

func NewListBaseReq() *admin_proto.ReqListBase {
	return &admin_proto.ReqListBase{}
}

// HandleListBaseReq 处理通用列表查询参数
func HandleListBaseReq(params *admin_proto.ReqListBase) (offset, limit int, data *admin_proto.ReqListBase) {
	if params == nil {
		params = &admin_proto.ReqListBase{}
	}
	if params.PageSize == 0 {
		params.PageSize = 10
	}
	if params.PageNum == 0 {
		params.PageNum = 1
	}
	offset = int((params.PageNum - 1) * params.PageSize)
	limit = int(params.PageSize)
	data = params
	return
}

// HandleLogicError 处理逻辑层返回的错误
func HandleLogicError(ctx *gin.Context, err error, msg string, result code.IMessage) {
	if errors.As(err, &result) {
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	result.SetCodeError(code_proto.ErrorCode_Error, err)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
	code.JSON(ctx, result)
	return
}
