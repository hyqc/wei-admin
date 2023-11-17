package router

import (
	account "admin/app/account/controller"
	"github.com/gin-gonic/gin"
)

var (
	accountController = account.DemoController{}
)

func Routes(e *gin.Engine) {
	r := e.Group("/api")
	{
		r.GET("/demo", accountController.Demo)
	}
}
