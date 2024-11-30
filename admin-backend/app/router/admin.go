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
		account.POST("/info", accountApi.Info)
		account.POST("/edit", accountApi.Edit)
		account.POST("/password", accountApi.Password)
		account.POST("/menu", accountApi.Menu)
		account.POST("/permission", accountApi.Permission)
	}

	user := admin.Group("/user")
	{
		userAPI := adminCtl.UserController{}
		user.POST("/list", userAPI.List)
		user.POST("/info", userAPI.Info)
		user.POST("/add", userAPI.Add)
		user.POST("/edit", userAPI.Edit)
		user.POST("/edit_pwd", userAPI.EditPassword)
		user.POST("/enable", userAPI.Enable)
		user.POST("/delete", userAPI.Delete)
		user.POST("/bind_roles", userAPI.BindRoles)
	}

	api := admin.Group("/api")
	{
		apiAPI := adminCtl.APIController{}
		api.POST("/list", apiAPI.List)
		api.POST("/all", apiAPI.All)
		api.POST("/info", apiAPI.Info)
		api.POST("/add", apiAPI.Add)
		api.POST("/edit", apiAPI.Edit)
		api.POST("/enable", apiAPI.Enable)
		api.POST("/delete", apiAPI.Delete)
	}

	menu := admin.Group("/menu")
	{
		menuAPI := adminCtl.MenuController{}
		menu.POST("/list", menuAPI.List)
		menu.POST("/modes", menuAPI.Modes)
		menu.POST("/pages", menuAPI.Pages)
		menu.POST("/tree", menuAPI.Tree)
		menu.POST("/info", menuAPI.Info)
		menu.POST("/add", menuAPI.Add)
		menu.POST("/edit", menuAPI.Edit)
		menu.POST("/enable", menuAPI.Enable)
		menu.POST("/show", menuAPI.Show)
		menu.POST("/delete", menuAPI.Delete)
	}

	permission := admin.Group("/permission")
	{
		permissionAPI := adminCtl.PermissionController{}
		permission.POST("/list", permissionAPI.List)
		permission.POST("/info", permissionAPI.Info)
		permission.POST("/add", permissionAPI.Add)
		permission.POST("/edit", permissionAPI.Edit)
		permission.POST("/enable", permissionAPI.Enable)
		permission.POST("/delete", permissionAPI.Delete)
		permission.POST("/bind_api", permissionAPI.BindAPI)
		permission.POST("/add_menu_permission", permissionAPI.AddMenuPermissions)
	}

	role := admin.Group("/role")
	{
		roleAPI := adminCtl.RoleController{}
		role.POST("/list", roleAPI.List)
		role.POST("/all", roleAPI.All)
		role.POST("/info", roleAPI.Info)
		role.POST("/add", roleAPI.Add)
		role.POST("/edit", roleAPI.Edit)
		role.POST("/enable", roleAPI.Enable)
		role.POST("/delete", roleAPI.Delete)
		role.POST("/bind_permissions", roleAPI.BindPermissions)
	}
}
