package common

import (
	adminproto "admin/proto"
)

var (
	DESC                           = "descend"
	ASC                            = "ascend"
	EnabledAllQueryValue     int32 = 0
	EnabledValidQueryValue   int32 = 1
	EnabledInvalidQueryValue int32 = 2
)

func ListBaseReqHandle(params *adminproto.ApiListReq) (offset, limit int) {
	if params.Base == nil {
		params.Base = &adminproto.ListBaseReq{}
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
