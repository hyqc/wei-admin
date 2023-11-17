package router

import (
	demo "admin/app/demo/controller"
	"github.com/gin-gonic/gin"
)

var (
	demoCtl = demo.DemoController{}
)

func Routes(e *gin.Engine) {
	r := e.Group("/api")
	{
		r.GET("/demo", demoCtl.Demo)
	}
}
