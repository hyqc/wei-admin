package router

import (
	demoCtl "admin/app/demo/controller"
	"github.com/gin-gonic/gin"
)

// demo demo
func demo(g *gin.RouterGroup) {
	api := demoCtl.DemoController{}
	g.GET("/demo", api.Demo)
}
