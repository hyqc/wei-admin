package constant

import (
	"admin/pkg/core"
	"github.com/gin-gonic/gin"
)

const (
	Authorization       = "authorization"
	ContextClaims       = "ctx_jwt_claims"
	AdministerId  int32 = 1 // 超管ID
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
