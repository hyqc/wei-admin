package middleware

import (
	"admin/app/admin/dao"
	"admin/code"
	"admin/constant"
	"admin/global"
	"admin/proto/code_proto"
	"github.com/gin-gonic/gin"
	"strings"
)

func getAuthorization(ctx *gin.Context) string {
	tokenStr := ctx.GetHeader(constant.Authorization)
	return strings.Trim(strings.Replace(tokenStr, "Bearer", "", -1), " ")
}

func auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if m, ok := global.AppConfig.JWT.IgnoresMap[ctx.Request.Method]; ok {
			if _, ok := m[ctx.Request.URL.Path]; ok {
				ctx.Next()
				return
			}

			for v := range m {
				if len(v) != 0 && strings.HasSuffix(v, "*") && strings.HasPrefix(ctx.Request.URL.Path, v[:len(v)-1]) {
				}
			}
		}

		//需要鉴权的路由
		token := getAuthorization(ctx)
		cla, err := global.AppAuth.Inspect(token)
		if err != nil {
			code.JSON(ctx, code.NewCodeError(code_proto.ErrorCode_AuthTokenInspectInvalid, err))
			return
		}

		// 查询管理员的权限
		pass, err := dao.H.AdminPermission.IsAdminCanAccessPath(ctx, cla.AccountId, ctx.Request.URL.Path)
		if err != nil || !pass {
			code.JSON(ctx, code.NewCodeError(code_proto.ErrorCode_AuthTokenForbidden, err))
			return
		}
		ctx.Set(constant.ContextClaims, cla)

		ctx.Next()
	}
}
