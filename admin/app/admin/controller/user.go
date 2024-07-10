package controller

import (
	"admin/app/admin/logic"
	"admin/app/common"
	"admin/code"
	"admin/config"
	"admin/constant"
	"admin/pkg/core"
	"admin/pkg/validator"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AdminUserController struct {
	core.Controller
}

// List 管理员列表
func (AdminUserController) List(ctx *gin.Context) {
	msg := "AdminUserController.List"
	params := &admin_proto.AdminUserListParamsReq{Base: common.NewListBaseReq()}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminUser.List(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// All 全部有效管理员
func (AdminUserController) All(ctx *gin.Context) {

}

// Add 添加管理员
func (AdminUserController) Add(ctx *gin.Context) {

}

// Detail 详情
func (AdminUserController) Detail(ctx *gin.Context) {

}

// Edit 编辑管理员
func (AdminUserController) Edit(ctx *gin.Context) {

}

// Enable 启用禁用
func (AdminUserController) Enable(ctx *gin.Context) {

}

// Delete 删除管理员
func (AdminUserController) Delete(ctx *gin.Context) {

}

// BindRoles 管理员绑定角色
func (AdminUserController) BindRoles(ctx *gin.Context) {

}
