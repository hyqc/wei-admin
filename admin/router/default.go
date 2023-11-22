package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	root   *gin.RouterGroup // 根路由
	groups = []func(){      // 注册路由
		check,
		admin,
	}
)

func Init() *gin.Engine {
	Engine := gin.New()
	root = Engine.Group("/")
	for _, f := range groups {
		f()
	}
	return Engine
}

func check() {
	root.GET("/check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
}

func admin() {
	g := root.Group("/admin")
	{
		g.GET("/demo", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{"code": 0, "msg": "success"})
		})
	}
}
