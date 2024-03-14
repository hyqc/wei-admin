package middleware

import (
	"admin/code"
	"admin/config"
	"admin/pkg/core"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	Authorization = "authorization"
	ContextClaims = "ctx_jwt_claims"
)

func getAuthorization(ctx *gin.Context) string {
	tokenStr := ctx.GetHeader(Authorization)
	return strings.Trim(tokenStr, "Bearer ")
}

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
		ctx.Set(ContextClaims, cla)
		ctx.Next()
	}
}
