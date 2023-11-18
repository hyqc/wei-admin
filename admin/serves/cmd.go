package serves

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

const (
	cmdVersion    = "v"           // 版本
	cmdConfigPath = "config-path" // 配置文件路径
	cmdConfigMode = "mode"        // 运行模式

	defaultConfigPrefixName = "config"
)

var (
	server           = &http.Server{} // http server
	configFilePath   = "./"           // yaml config file path
	configFileSuffix = ".yaml"        // 后缀
	configFullPath   = ""             // 配置完整路径，默认配置的名称，不包含后缀
	runMode          = ""             // 运行模式

)

var (
	BuildVersion = ""
	BuildTag     = ""
	BuildCommit  = ""
	BuildTime    = ""
	GoVersion    = ""
)

func parseCmd() {
	var version bool
	flag.BoolVar(&version, cmdVersion, false, "application version")
	flag.StringVar(&configFilePath, cmdConfigPath, configFilePath, "yaml config file path")
	flag.StringVar(&runMode, cmdConfigMode, ModeDefaultName, "yaml config file path")
	flag.Parse()

	if _, ok := ModeNameMap[runMode]; !ok {
		fmt.Printf("run mode should be in : [\"\", \"dev\", \"local\", \"test\", \"prod\"]\n")
		os.Exit(1)
	}
	connector := ""
	if runMode != ModeDefaultName {
		connector = "-"
	}
	configFullPath = fmt.Sprintf("%s%s%s%s%s", configFilePath, defaultConfigPrefixName, connector, runMode, configFileSuffix)
	fmt.Printf("config path: %s\n", configFullPath)

	if version {
		fmt.Printf(" build version: %s\n build tag: %s\n build commit: %s\n build time: %s\n go version: %s\n",
			BuildVersion, BuildTag, BuildCommit, BuildTime, GoVersion)
		os.Exit(0)
		return
	}
	fmt.Println("parse cmd success")
}
