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

type Appear interface {
	parseCmd()    // 解下cmd参数
	parseConfig() // 解析配置
	initConfig()  // 初始化配置
	runServe()    // 启动http服务
	signalListen(ctx context.Context)
	shutdown(ctx context.Context, ser *http.Server)
	Run()
}

func init() {

}

var (
	sig = make(chan os.Signal, 1) // 接收停止信号
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		global.LoggerClose()
	}()

	parseCmd()
	parseConfig()
	initConfig()
	runServe()
	signalListen(ctx)
}

func parseConfig() {
	if err := global.ParseConfig(configFullPath); err != nil {
		utils.PrintfLn("parse yaml config: %s , error: %v", configFilePath, err.Error())
		os.Exit(1)
		return
	}
	utils.PrintfLn("parse yaml config success")
}

func initConfig() {

	if err := global.InitLogger(); err != nil {
		utils.PrintfLn("init logger config error: %v", err.Error())
		os.Exit(1)
		return
	}

	if err := global.InitMySQLDB(); err != nil {
		utils.PrintfLn("init database config error: %v", err.Error())
		os.Exit(2)
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
	server.Addr = fmt.Sprintf(":%s", port)
	server.Handler = e
	go func() {
		utils.PrintfLn(fmt.Sprintf("start serve port: %s", port))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			utils.PrintfLn("start serve port: %s, error: %v", port, err)
			return
		}
	}()
}

func signalListen(ctx context.Context) {
	utils.PrintfLn("listen os signal...")
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	run := true
	for run {
		select {
		case <-sig:
			run = false
			shutdown(ctx, server)
		}
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
