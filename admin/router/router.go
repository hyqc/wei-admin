package router

import (
	"github.com/gin-gonic/gin"
)

func Routes(e *gin.Engine) {
	root := e.Group("/")
	// 示例
	demo(root)
	// 管理后台
	admins(root)
}
