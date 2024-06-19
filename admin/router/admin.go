package router

import (
	adminCtl "admin/app/admin/controller"
	"github.com/gin-gonic/gin"
)

func admins(g *gin.RouterGroup) {
	admin := g.Group("/admin")
	account := admin.Group("/account")
	{
		accountApi := adminCtl.AccountController{}
		account.POST("/login", accountApi.Login)
		account.GET("/info", accountApi.Info)
		account.POST("/edit", accountApi.Edit)
		account.POST("/password", accountApi.Password)
		account.POST("/menu", accountApi.Menu)
		account.POST("/permission", accountApi.Permission)
	}
}
