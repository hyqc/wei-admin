package middleware

import (
	"admin/code"
	"admin/config"
	"admin/constant"
	"admin/pkg/core"
	"github.com/gin-gonic/gin"
	"strings"
)

func auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := getAuthorization(ctx)
		cla, err := core.JWTCheck(token, config.AppConfig.Server.JWT.Secret)
		if err != nil {
			// TODO 返回
			core.ResponseOk(ctx, core.ResponseData{
				Code:    code.AuthTokenFailed,
				Message: code.Msg(code.AuthTokenFailed),
			})
			return
		}
		ctx.Set(constant.ContextClaims, cla)
		ctx.Next()
	}
}

func getAuthorization(ctx *gin.Context) string {
	tokenStr := ctx.GetHeader(constant.Authorization)
	return strings.Trim(tokenStr, "Bearer ")
}
