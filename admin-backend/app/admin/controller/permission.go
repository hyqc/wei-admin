package controller

import (
	"admin/app/admin/logic"
	"admin/app/admin/validate"
	"admin/app/common"
	"admin/code"
	"admin/constant"
	"admin/global"
	"admin/pkg/core"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PermissionController struct {
	core.Controller
}

// List 权限列表
//
//	@Summary		权限列表
//	@Description	权限列表
//	@Tags			权限接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminPermissionList							true	"请求参数"
//	@Success		200		{object}	code.Message{data=admin_proto.RespAdminPermissionListData}	"desc"
//	@Router			/admin/permission/list [post]
func (PermissionController) List(ctx *gin.Context) {
	msg := "PermissionController.List"
	params := &admin_proto.ReqAdminPermissionList{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminPermission.List(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Add 添加权限
//
//	@Summary		添加权限
//	@Description	添加权限
//	@Tags			权限接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminPermissionAdd	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}				"desc"
//	@Router			/admin/permission/add [post]
func (PermissionController) Add(ctx *gin.Context) {
	msg := "PermissionController.Add"
	params := &admin_proto.ReqAdminPermissionAdd{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.Add(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
}

// Info 权限详情
//
//	@Summary		权限详情
//	@Description	权限详情
//	@Tags			权限接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminPermissionInfo					true	"请求参数"
//	@Success		200		{object}	code.Message{data=admin_proto.AdminPermissionInfo}	"desc"
//	@Router			/admin/permission/info [post]
func (PermissionController) Info(ctx *gin.Context) {
	msg := "PermissionController.Info"
	params := &admin_proto.ReqAdminPermissionInfo{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	info, err := logic.H.AdminPermission.Info(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(info)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Edit 编辑权限
//
//	@Summary		编辑权限
//	@Description	编辑权限
//	@Tags			权限接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminPermissionEdit	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}				"desc"
//	@Router			/admin/permission/edit [post]
func (PermissionController) Edit(ctx *gin.Context) {
	msg := "PermissionController.Edit"
	params := &admin_proto.ReqAdminPermissionEdit{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.Edit(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Enable 权限启用|禁用
//
//	@Summary		权限启用|禁用
//	@Description	权限启用|禁用
//	@Tags			权限接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminPermissionEnable	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}					"desc"
//	@Router			/admin/permission/enable [post]
func (PermissionController) Enable(ctx *gin.Context) {
	msg := "PermissionController.Enable"
	params := &admin_proto.ReqAdminPermissionEnable{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.Enable(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Delete 删除权限
//
//	@Summary		删除权限
//	@Description	删除权限
//	@Tags			权限接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminPermissionDelete	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}					"desc"
//	@Router			/admin/permission/delete [post]
func (PermissionController) Delete(ctx *gin.Context) {
	msg := "PermissionController.Delete"
	params := &admin_proto.ReqAdminPermissionDelete{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.Delete(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// BindAPI 绑定权限接口
//
//	@Summary		绑定权限接口
//	@Description	绑定权限接口
//	@Tags			权限接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminPermissionBindApis	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}					"desc"
//	@Router			/admin/permission/bind_apis [post]
func (PermissionController) BindAPI(ctx *gin.Context) {
	msg := "PermissionController.BindAPI"
	params := &admin_proto.ReqAdminPermissionBindApis{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.BindAPI(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// UnBindAPI 解绑权限接口
//
//	@Summary		解绑权限接口
//	@Description	解绑权限接口
//	@Tags			权限接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminPermissionUnBindApi	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}					"desc"
//	@Router			/admin/permission/unbind_api [post]
func (PermissionController) UnBindAPI(ctx *gin.Context) {
	msg := "PermissionController.UnBindAPI"
	params := &admin_proto.ReqAdminPermissionUnBindApi{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.UnBindAPI(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// AddMenuPermissions 指定菜单创建查看编辑删除权限
//
//	@Summary		指定菜单创建查看编辑删除权限
//	@Description	指定菜单创建查看编辑删除权限
//	@Tags			权限接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminPermissionBindMenu	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}					"desc"
//	@Router			/admin/permission/add_menu_permissions [post]
func (PermissionController) AddMenuPermissions(ctx *gin.Context) {
	msg := "PermissionController.AddMenuPermissions"
	params := &admin_proto.ReqAdminPermissionBindMenu{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.AddMenuPermissions(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}
