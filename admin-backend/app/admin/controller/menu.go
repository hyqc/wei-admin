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

type MenuController struct {
	core.Controller
}

// List 菜单列表
//
//	@Summary		菜单列表
//	@Description	菜单列表
//	@Tags			菜单接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminMenuList							true	"请求参数"
//	@Success		200		{object}	code.Message{data=admin_proto.RespAdminMenuListData}	"desc"
//	@Router			/admin/menu/list [post]
func (MenuController) List(ctx *gin.Context) {
	msg := "MenuController.List"
	params := &admin_proto.ReqAdminMenuList{Base: common.NewListBaseReq()}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminMenu.List(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Tree 有效菜单树
//
//	@Summary		有效菜单树
//	@Description	有效菜单树
//	@Tags			菜单接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	code.Message{data=[]admin_proto.MenuTreeItem}	"desc"
//	@Router			/admin/menu/true [post]
func (MenuController) Tree(ctx *gin.Context) {
	msg := "MenuController.Tree"
	result := code.NewCode(code_proto.ErrorCode_Success)
	data, err := logic.H.AdminMenu.Tree(ctx)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Add 添加菜单
//
//	@Summary		添加菜单
//	@Description	添加菜单
//	@Tags			菜单接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminMenuAdd	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}		"desc"
//	@Router			/admin/menu/add [post]
func (MenuController) Add(ctx *gin.Context) {
	msg := "MenuController.Add"
	params := &admin_proto.ReqAdminMenuAdd{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminMenu.Add(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Info 菜单详情
//
//	@Summary		菜单详情
//	@Description	菜单详情
//	@Tags			菜单接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminMenuInfo							true	"请求参数"
//	@Success		200		{object}	code.Message{data=admin_proto.RespAdminMenuInfoData}	"desc"
//	@Router			/admin/menu/info [post]
func (MenuController) Info(ctx *gin.Context) {
	msg := "MenuController.Info"
	params := &admin_proto.ReqAdminMenuInfo{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	info, err := logic.H.AdminMenu.Info(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(info)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Edit 编辑菜单
//
//	@Summary		编辑菜单
//	@Description	编辑菜单
//	@Tags			菜单接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminMenuEdit	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}			"desc"
//	@Router			/admin/menu/edit [post]
func (MenuController) Edit(ctx *gin.Context) {
	msg := "MenuController.Edit"
	params := &admin_proto.ReqAdminMenuEdit{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminMenu.Edit(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Enable 启用禁用菜单
//
//	@Summary		启用禁用菜单
//	@Description	启用禁用菜单
//	@Tags			菜单接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminMenuEnable	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}			"desc"
//	@Router			/admin/menu/enable [post]
func (MenuController) Enable(ctx *gin.Context) {
	msg := "MenuController.Enable"
	params := &admin_proto.ReqAdminMenuEnable{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminMenu.Enable(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Show 显示|隐藏菜单
//
//	@Summary		显示|隐藏菜单
//	@Description	显示|隐藏菜单
//	@Tags			菜单接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminMenuShow	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}			"desc"
//	@Router			/admin/menu/show [post]
func (MenuController) Show(ctx *gin.Context) {
	msg := "MenuController.Show"
	params := &admin_proto.ReqAdminMenuShow{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminMenu.Show(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Delete 删除菜单
//
//	@Summary		删除菜单
//	@Description	删除菜单
//	@Tags			菜单接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminMenuDelete	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}			"desc"
//	@Router			/admin/menu/delete [post]
func (MenuController) Delete(ctx *gin.Context) {
	msg := "MenuController.Delete"
	params := &admin_proto.ReqAdminMenuDelete{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminMenu.Delete(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Permissions 菜单权限列表
//
//	@Summary		菜单权限列表
//	@Description	菜单权限列表
//	@Tags			菜单接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminMenuPermissions							true	"请求参数"
//	@Success		200		{object}	code.Message{data=admin_proto.RespAdminMenuPermissionsData}	"desc"
//	@Router			/admin/menu/permissions [post]
func (MenuController) Permissions(ctx *gin.Context) {
	msg := "MenuController.Permissions"
	params := &admin_proto.ReqAdminMenuPermissions{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminMenu.Permissions(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Pages 页面菜单列表
//
//	@Summary		页面菜单列表
//	@Description	页面菜单列表
//	@Tags			菜单接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminMenuPages						true	"请求参数"
//	@Success		200		{object}	code.Message{data=admin_proto.RespAdminMenuPages}	"desc"
//	@Router			/admin/menu/pages [post]
func (MenuController) Pages(ctx *gin.Context) {
	msg := "MenuController.Permissions"
	params := &admin_proto.ReqAdminMenuPages{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminMenu.Pages(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Modes 页面模块权限列表
//
//	@Summary		页面模块权限列表
//	@Description	页面模块权限列表
//	@Tags			菜单接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	code.Message{data=admin_proto.RespAdminMenuModeData}	"desc"
//	@Router			/admin/menu/modes [post]
func (MenuController) Modes(ctx *gin.Context) {
	msg := "MenuController.Mode"
	result := code.NewCode(code_proto.ErrorCode_Success)
	data, err := logic.H.AdminMenu.AllMode(ctx)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}
