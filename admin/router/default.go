package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	Engine = gin.New()
	root   = Engine.Group("/")

	groups = []func(){
		check,
		admin,
	}
)

func init() {
	for _, f := range groups {
		f()
	}
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
