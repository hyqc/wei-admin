package serves

import (
	"admin/app/middleware"
	"admin/app/router"
	"admin/global"
	"admin/pkg/utils"
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {

}

var (
	sig    = make(chan os.Signal, 1) // 接收停止信号
	server = &http.Server{}          // http server
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		global.CloseLogger()
	}()
	//初始化配置
	initConfig()
	//启动服务
	runServe()
	//服务器监听
	signalListen(ctx)
}

func initConfig() {
	if err := global.Init(); err != nil {
		utils.PrintfLn("init config error: %v", err.Error())
		os.Exit(1)
		return
	}
	utils.PrintfLn("init config success")
}

func runServe() {

	if !global.AppConfig.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	e := gin.Default()

	if global.AppConfig.Server.Pprof {
		pprof.Register(e)
	}

	e.Use(middleware.Global...)
	router.Routes(e)

	port := global.AppConfig.Server.Port
	server.Addr = fmt.Sprintf(":%d", port)
	server.Handler = e

	go func() {
		utils.PrintfLn(fmt.Sprintf("start serve port: %v", port))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			utils.PrintfLn("start serve port: %v, error: %v", port, err)
			return
		}
	}()
}

func signalListen(ctx context.Context) {
	utils.PrintfLn("listen os signal...")
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sig:
		shutdown(ctx, server)
	}
	utils.PrintfLn("stop serve")
}

func shutdown(ctx context.Context, ser *http.Server) {
	err := ser.Shutdown(ctx)
	if err != nil {
		utils.PrintfLn("shutdown server error: %v", err)
		return
	}
	utils.PrintfLn("serves shutdown success")
}
