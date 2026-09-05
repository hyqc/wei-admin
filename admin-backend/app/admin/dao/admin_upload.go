package dao

import (
	"admin/app/admin/gen/model"
	"admin/app/admin/gen/query"
	"admin/app/common"
	"admin/proto/admin_proto"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type IAdminUpload interface {
	Create(ctx *gin.Context, data *model.AdminUpload) error
	Info(ctx *gin.Context, id int32) (*model.AdminUpload, error)
	Delete(ctx *gin.Context, id int32) error
	List(ctx *gin.Context, params *admin_proto.ReqAdminUploadList) (total int64, list []*model.AdminUpload, err error)
}

type AdminUpload struct {
}

func newAdminUpload() *AdminUpload {
	return &AdminUpload{}
}

func (a *AdminUpload) Create(ctx *gin.Context, data *model.AdminUpload) error {
	return query.AdminUpload.WithContext(ctx).Create(data)
}

func (a *AdminUpload) Info(ctx *gin.Context, id int32) (*model.AdminUpload, error) {
	return query.AdminUpload.WithContext(ctx).Where(query.AdminUpload.ID.Eq(id)).First()
}

func (a *AdminUpload) Delete(ctx *gin.Context, id int32) error {
	_, err := query.AdminUpload.WithContext(ctx).Where(query.AdminUpload.ID.Eq(id)).Delete()
	return err
}

func (a *AdminUpload) List(ctx *gin.Context, params *admin_proto.ReqAdminUploadList) (total int64, list []*model.AdminUpload, err error) {
	offset, limit, base := common.HandleListBaseReq(params.Base)
	params.Base = base
	q := a.handleListReq(ctx, params, base)
	total, err = q.Count()
	if err != nil {
		return total, list, err
	}
	// 最新上传排在最前
	list, err = q.Order(query.AdminUpload.ID.Desc()).Limit(limit).Offset(offset).Find()
	return total, list, nil
}

func (a *AdminUpload) handleListReq(ctx context.Context, params *admin_proto.ReqAdminUploadList, base *admin_proto.ReqListBase) (q query.IAdminUploadDo) {
	DB := query.AdminUpload
	q = DB.WithContext(ctx)

	if base.CreateStartTime > 0 {
		q = q.Where(DB.CreatedAt.Gte(time.Unix(base.CreateStartTime, 0)))
	}
	if base.CreateEndTime > 0 {
		q = q.Where(DB.CreatedAt.Lte(time.Unix(base.CreateEndTime, 0)))
	}
	if params.UploadGroup != "" {
		// 分组按路径前缀匹配，便于按模块筛选
		q = q.Where(DB.UploadGroup.Like(params.UploadGroup + "%"))
	}
	if params.OriginalName != "" {
		q = q.Where(DB.OriginalName.Like("%" + params.OriginalName + "%"))
	}
	if params.Ext != "" {
		q = q.Where(DB.Ext.Eq(params.Ext))
	}
	if params.Driver != "" {
		q = q.Where(DB.Driver.Eq(params.Driver))
	}
	return q
}
