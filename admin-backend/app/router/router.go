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
