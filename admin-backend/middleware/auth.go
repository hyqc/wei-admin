package middleware

import (
	"admin/code"
	"admin/config"
	"admin/constant"
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
		if !funk.ContainsString(config.AppConfig.Server.JWT.IgnoreUrls, ctx.Request.URL.Path) {
			token := getAuthorization(ctx)
			cla, err := core.JWTCheck(token, config.AppConfig.Server.JWT.Secret)
			if err != nil {
				code.JSON(ctx, code.NewCodeError(code_proto.ErrorCode_AuthTokenInspectInvalid, err))
				return
			}
			ctx.Set(constant.ContextClaims, cla)
		}

		ctx.Next()
	}
}
