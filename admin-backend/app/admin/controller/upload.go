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

type UploadController struct {
	core.Controller
}

// List 上传记录列表
//
//	@Summary		上传记录列表
//	@Description	上传记录列表
//	@Tags			上传接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqFrontAdminUploadList					true	"请求参数"
//	@Success		200		{object}	code.Message{data=admin_proto.RespFrontAdminUploadListData}	"desc"
//	@Router			/admin/upload/list [post]
func (UploadController) List(ctx *gin.Context) {
	msg := "UploadController.List"
	// 前端分页参数为平铺结构，先按前端契约接收，再转换为内部参数
	front := &admin_proto.ReqFrontAdminUploadList{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, front); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	params := &admin_proto.ReqAdminUploadList{
		Base: &admin_proto.ReqListBase{
			PageSize:        front.PageSize,
			PageNum:         front.PageNum,
			CreateStartTime: front.CreateStartTime,
			CreateEndTime:   front.CreateEndTime,
		},
		UploadGroup:  front.UploadGroup,
		OriginalName: front.OriginalName,
		Ext:          front.Ext,
		Driver:       front.Driver,
	}
	data, err := logic.H.AdminUpload.List(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	// 转换为前端结构：list + pageInfo
	frontData := &admin_proto.RespFrontAdminUploadListData{}
	if data != nil {
		frontData.List = data.List
		frontData.PageInfo = &admin_proto.PageInfo{
			Total:    data.Total,
			PageNum:  front.PageNum,
			PageSize: front.PageSize,
		}
	}
	result.SetData(frontData)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Upload 上传文件
//
//	@Summary		上传文件
//	@Description	multipart/form-data 上传；file 为文件字段，uploadGroup 为分组（上传路径前缀，如 /admin/user/）
//	@Tags			上传接口相关
//	@Accept			multipart/form-data
//	@Produce		application/json
//	@Param			file		formData	file	true	"文件"
//	@Param			uploadGroup	formData	string	true	"分组（上传路径前缀）"
//	@Success		200			{object}	code.Message{data=admin_proto.AdminUploadItem}	"desc"
//	@Router			/admin/upload/upload [post]
func (UploadController) Upload(ctx *gin.Context) {
	msg := "UploadController.Upload"
	result := code.NewCode(code_proto.ErrorCode_Success)
	file, err := ctx.FormFile("file")
	if err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_AdminUploadFileEmpty, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminUpload.Upload(ctx, constant.GetCustomClaims(ctx).AccountId, file, ctx.PostForm("uploadGroup"))
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Delete 删除上传记录（同时删除存储文件）
//
//	@Summary		删除上传记录
//	@Description	删除上传记录，并尝试删除存储上的文件
//	@Tags			上传接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminUploadDelete	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}				"desc"
//	@Router			/admin/upload/delete [post]
func (UploadController) Delete(ctx *gin.Context) {
	msg := "UploadController.Delete"
	params := &admin_proto.ReqAdminUploadDelete{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminUpload.Delete(ctx, params.Id); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}
