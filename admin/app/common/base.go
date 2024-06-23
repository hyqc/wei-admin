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

func NewListBaseReq() *admin_proto.ListBaseReq {
	return &admin_proto.ListBaseReq{}
}

// HandleListBaseReq 处理通用列表查询参数
func HandleListBaseReq(params *admin_proto.ApiListReq) (offset, limit int) {
	if params.Base == nil {
		params.Base = &admin_proto.ListBaseReq{}
	}
	if params.Base.PageSize == 0 {
		params.Base.PageSize = 10
	}
	if params.Base.PageNum == 0 {
		params.Base.PageNum = 1
	}
	offset = int((params.Base.PageNum - 1) * params.Base.PageSize)
	limit = int(params.Base.PageSize)
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
