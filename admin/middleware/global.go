package middleware

import (
	"admin/code"
	"admin/config"
	"admin/pkg/utils"
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var Global = []gin.HandlerFunc{
	logger(),
	cors(),
	recovery(),
	auth(),
}

func cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}

func logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			msg := code.NewCode(code.ReadContextRequestBodyFailed)
			config.AppLoggerSugared.Errorw("read context request body error", msg, err)
			code.JSON(ctx, msg)
			return
		}
		ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
		config.AppLoggerSugared.Debugw("body", zap.ByteString("body", body))
		config.AppLoggerSugared.Debugw("query", zap.String("raw", ctx.Request.URL.RawQuery))
		cost := time.Since(start)
		config.AppLoggerSugared.Infow("request",
			zap.Duration("cost", cost),
			zap.String("method", ctx.Request.Method),
			zap.Int("status", ctx.Writer.Status()),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
		ctx.Next()
	}
}

func recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := config.AppLoggerSugared
		defer func() {
			if err := recover(); err != nil {
				errInfo := string(debug.Stack())
				utils.PrintfLn("panic", err)
				utils.PrintfLn("panic", errInfo)
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				if brokenPipe {
					logger.Error(ctx.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = ctx.Error(err.(error)) // nolint: errcheck
					ctx.Abort()
					return
				}

				if config.AppConfig.Server.Debug {
					logger.Error("[Recovery from panic]",
						zap.String("path", ctx.Request.URL.Path),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", errInfo),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.String("path", ctx.Request.URL.Path),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}

				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}
