package middleware

import (
	"admin/app/admin/dao"
	"admin/code"
	"admin/constant"
	"admin/global"
	"admin/pkg/core"
	"admin/proto/code_proto"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"strings"
)

func getAuthorization(ctx *gin.Context) string {
	tokenStr := ctx.GetHeader(constant.Authorization)
	return strings.Trim(strings.Replace(tokenStr, "Bearer", "", -1), " ")
}

func auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !funk.ContainsString(global.AppConfig.Server.JWT.IgnoreUrls, ctx.Request.URL.Path) {
			token := getAuthorization(ctx)
			cla, err := core.JWTCheck(token, global.AppConfig.Server.JWT.Secret)
			if err != nil {
				code.JSON(ctx, code.NewCodeError(code_proto.ErrorCode_AuthTokenInspectInvalid, err))
				return
			}

			// 查询管理员的权限
			pass, err := dao.H.AdminPermission.IsAdminCanAccessPath(ctx, cla.AdminID, ctx.Request.URL.Path)
			if err != nil || !pass {
				code.JSON(ctx, code.NewCodeError(code_proto.ErrorCode_AuthTokenForbidden, err))
				return
			}
			ctx.Set(constant.ContextClaims, cla)
		}

		ctx.Next()
	}
}
