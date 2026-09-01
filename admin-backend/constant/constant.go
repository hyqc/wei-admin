package constant

import (
	"admin/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
)

const (
	Authorization                 = "authorization"
	ContextClaims                 = "ctx_jwt_claims"
	AdministerId            int32 = 1 // 超管账号ID
	AdministerRoleId        int32 = 1 // 超管角色ID
	LogResponseMsgField           = "response"
	AdminTokenExpireSeconds       = 3600 * 24 * 7
)

func IsAdministrator(adminId int32) bool {
	return adminId == AdministerId
}

// IsAdministratorRole 是否超级管理员角色
func IsAdministratorRole(roleId int32) bool {
	return roleId == AdministerRoleId
}

func GetCustomClaims(ctx *gin.Context) *jwt.CustomClaims {
	val, ok := ctx.Get(ContextClaims)
	if ok {
		return val.(*jwt.CustomClaims)
	}
	return nil
}
