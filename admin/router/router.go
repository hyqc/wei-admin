package router

import (
	demoCtl "admin/app/demo/controller"
	"github.com/gin-gonic/gin"
)

func Routes(e *gin.Engine) {
	root := e.Group("/")

	demo(root)
}

// demo demo
func demo(g *gin.RouterGroup) {
	api := demoCtl.DemoController{}
	g.GET("/demo", api.Demo)
}
