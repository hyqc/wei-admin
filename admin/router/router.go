package router

import (
	adminCtl "admin/app/admin/controller"
	demoCtl "admin/app/demo/controller"
	"github.com/gin-gonic/gin"
)

func Routes(e *gin.Engine) {
	root := e.Group("/")

	demo(root)

	admins(root)
}

// demo demo
func demo(g *gin.RouterGroup) {
	api := demoCtl.DemoController{}
	g.GET("/demo", api.Demo)
}

func admins(g *gin.RouterGroup) {
	admin := g.Group("/admin")
	account := admin.Group("/account")
	{
		accountApi := adminCtl.AccountController{}
		account.POST("/login", accountApi.Login)
	}
}
