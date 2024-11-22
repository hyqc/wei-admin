package constant

import (
	"admin/pkg/core"
	"github.com/gin-gonic/gin"
)

const (
	Authorization                 = "authorization"
	ContextClaims                 = "ctx_jwt_claims"
	AdministerId            int32 = 1 // 超管ID
	LogResponseMsgField           = "response"
	AdminTokenExpireSeconds       = 3600 * 24 * 7
)

const (
	DBModelIsEnabledTrue  int32 = 1
	DBModelIsEnabledFalse int32 = 0
)

func IsAdministrator(adminId int32) bool {
	return adminId == AdministerId
}

func GetCustomClaims(ctx *gin.Context) *core.CustomClaims {
	val, ok := ctx.Get(ContextClaims)
	if ok {
		return val.(*core.CustomClaims)
	}
	return nil
}
