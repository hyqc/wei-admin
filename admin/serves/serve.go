package serves

import (
	"admin/config"
	"admin/middleware"
	"admin/router"
	"context"
	"errors"
	"flag"
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
	sig      = make(chan os.Signal, 1) // 接收停止信号
	server   = &http.Server{}          // http server
	confPath = "./env/conf.yaml"       // yaml config file path
)

var (
	BuildVersion = ""
	BuildTag     = ""
	BuildCommit  = ""
	BuildTime    = ""
	GoVersion    = ""
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		config.LoggerClose()
	}()

	parseCmd()
	parseConfig()
	initConfig()
	runServe()
	signalListen(ctx)
}

func parseCmd() {
	var version bool
	flag.BoolVar(&version, "v", false, "version")
	flag.StringVar(&confPath, "conf", confPath, "yaml config file path")
	flag.Parse()
	if version {
		fmt.Printf(" build version: %s\n build tag: %s\n build commit: %s\n build time: %s\n go version: %s\n",
			BuildVersion, BuildTag, BuildCommit, BuildTime, GoVersion)
		os.Exit(0)
		return
	}
	fmt.Println("parse cmd success")
}

func parseConfig() {
	if err := config.ParseConfig(confPath); err != nil {
		fmt.Printf("parse yaml config: %s , error: %s\n", confPath, err.Error())
		os.Exit(1)
		return
	}
	fmt.Println("parse yaml config success")
}

func initConfig() {

	if err := config.InitLogger(); err != nil {
		fmt.Printf("init logger config error: %s\n", err.Error())
		os.Exit(1)
		return
	}

	fmt.Println("init config success")
}

func runServe() {

	if config.AppConfig.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	e := gin.Default()

	if config.AppConfig.Server.Pprof {
		pprof.Register(e)
	}

	e.Use(middleware.Cors())
	router.Routes(e)

	port := config.AppConfig.Server.Port
	server.Addr = fmt.Sprintf(":%s", port)
	server.Handler = e
	go func() {
		fmt.Println(fmt.Sprintf("start serve port: %s", port))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("start serve port: %s, error: %s\n", port, err)
			return
		}
	}()
}

func signalListen(ctx context.Context) {
	fmt.Println("listen os signal...")
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	run := true
	for run {
		select {
		case <-sig:
			run = false
			shutdown(ctx, server)
		}
	}
	fmt.Println("stop serve")
}

func shutdown(ctx context.Context, ser *http.Server) {
	err := ser.Shutdown(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("serves shutdown success")
}
