package admin

import "github.com/gin-gonic/gin"

const (
	AdministerId      int32 = 1 // 超管ID
	TokenAdminIdField       = "admin_id"
)

func IsAdministrator(adminId int32) bool {
	return adminId == AdministerId
}

func TokenParseAdminId(ctx *gin.Context) int32 {
	return int32(ctx.GetInt(TokenAdminIdField))
}
