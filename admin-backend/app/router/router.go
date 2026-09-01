package router

import (
	"admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(e *gin.Engine) {
	noAuth(e)
	auth(e)
}

func noAuth(e *gin.Engine) {
	notCheck := e.Group("/")
	{
		swagger(notCheck)
		// 健康检查探针（liveness / readiness），供 K8s 等编排系统使用
		notCheck.GET("/healthz", Healthz)
		notCheck.GET("/readyz", Readyz)
	}
}

func auth(e *gin.Engine) {
	check := e.Group("/", middleware.Auth...)
	{
		// 示例
		demo(check)
		// 管理后台
		admins(check)
	}
}
