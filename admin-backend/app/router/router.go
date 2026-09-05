package router

import (
	"admin/app/middleware"
	"admin/global"
	"admin/pkg/storage"
	"github.com/gin-gonic/gin"
	"strings"
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
	// 本地存储时把上传目录注册为静态资源，使文件可通过 URL 直接访问（免鉴权）
	if driver := global.AppConfig.Upload.Driver; driver == "" || strings.EqualFold(driver, storage.DriverLocal) {
		root := global.AppConfig.Upload.Local.Root
		if root == "" {
			root = "./upload"
		}
		prefix := global.AppConfig.Upload.Local.Prefix
		if prefix == "" {
			prefix = "/upload"
		}
		e.Static(prefix, root)
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
