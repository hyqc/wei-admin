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

// GetCustomClaims 取当前登录用户 claims
// 免鉴权接口可能没有 claims，此时返回空对象，避免调用方 nil 指针解引用
func GetCustomClaims(ctx *gin.Context) *jwt.CustomClaims {
	val, ok := ctx.Get(ContextClaims)
	if ok {
		if claims, ok := val.(*jwt.CustomClaims); ok && claims != nil {
			return claims
		}
	}
	return &jwt.CustomClaims{}
}
