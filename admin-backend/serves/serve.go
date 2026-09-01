package serves

import (
	"admin/app/middleware"
	"admin/app/router"
	"admin/global"
	"admin/pkg/config"
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
	"time"
)

func init() {

}

const (
	// shutdownTimeout 优雅关闭的最长等待时间，超时后强制关闭（配合 K8s terminationGracePeriodSeconds）
	shutdownTimeout = 15 * time.Second
	readTimeout     = 30 * time.Second
	writeTimeout    = 30 * time.Second
	idleTimeout     = 60 * time.Second
)

var (
	sig    = make(chan os.Signal, 1) // 接收停止信号
	server = &http.Server{}          // http server
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		global.CloseLogger()
		global.CloseRedis()
		global.CloseDatabase()
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
	// 显式设置超时，避免慢连接耗尽资源
	server.ReadTimeout = readTimeout
	server.WriteTimeout = writeTimeout
	server.IdleTimeout = idleTimeout

	go func() {
		utils.PrintfLn(fmt.Sprintf("start serve env: %v, port: %v", config.Env(), port))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			utils.PrintfLn("start serve port: %v, error: %v", port, err)
			// 监听失败（端口占用等）必须退出进程，否则编排系统无法感知故障
			os.Exit(1)
		}
	}()
}

func signalListen(ctx context.Context) {
	utils.PrintfLn("listen os signal")
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sig:
		shutdown(ctx, server)
	}
	utils.PrintfLn("stop serve")
}

func shutdown(ctx context.Context, ser *http.Server) {
	// 给在途请求留出处理时间，超时后强制关闭，避免 Pod 停在 Terminating
	shutdownCtx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()
	err := ser.Shutdown(shutdownCtx)
	if err != nil {
		utils.PrintfLn("shutdown server error: %v", err)
		return
	}
	utils.PrintfLn("serves shutdown success")
}
