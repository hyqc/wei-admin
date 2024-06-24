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
		account.GET("/menu", accountApi.Menu)
		account.GET("/permission", accountApi.Permission)
	}

	api := admin.Group("/api")
	{
		apiAPI := adminCtl.APIController{}
		api.GET("/list", apiAPI.List)
		api.GET("/all", apiAPI.All)
		api.GET("/info", apiAPI.Info)
		api.POST("/add", apiAPI.Add)
		api.POST("/edit", apiAPI.Edit)
		api.POST("/enable", apiAPI.Enable)
		api.POST("/delete", apiAPI.Delete)
	}

	menu := admin.Group("/menu")
	{
		menuAPI := adminCtl.MenuController{}
		menu.GET("/list", menuAPI.List)
		menu.GET("/tree", menuAPI.Tree)
		menu.GET("/info", menuAPI.Info)
		menu.POST("/add", menuAPI.Add)
		menu.POST("/edit", menuAPI.Edit)
		menu.POST("/enable", menuAPI.Enable)
		menu.POST("/delete", menuAPI.Delete)
	}
}
