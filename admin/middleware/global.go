package middleware

import (
	"admin/code"
	"admin/config"
	"admin/pkg/response"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			msg := code.NewCode(code.ReadContextRequestBodyFailed)
			config.AppLoggerSugared.Errorf("[middleware.global.Logger] [code:%v] [error:%v]", msg, err)
			response.JSON(ctx, msg)
			return
		}
		ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
		config.AppLoggerSugared.Debugf("[middleware.global.Logger] [request.body:%v]", string(body))
		config.AppLoggerSugared.Debugf("[middleware.global.Logger] [request.query:%v]", ctx.Request.URL.Query().Encode())
		ctx.Next()
	}
}
