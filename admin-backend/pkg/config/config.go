package config

import (
	json2 "encoding/json"
	"flag"
	"fmt"
	"github.com/micro/plugins/v5/config/encoder/yaml"
	nacosSource "github.com/micro/plugins/v5/config/source/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"go-micro.dev/v5/config"
	"go-micro.dev/v5/config/reader"
	"go-micro.dev/v5/config/reader/json"
	"go-micro.dev/v5/config/source"
	"go-micro.dev/v5/config/source/file"
	"go-micro.dev/v5/logger"
	"reflect"
	"strings"
	"time"
)

type IConfig interface {
	Handle() error
	Original() interface{}
}

const (
	cfsValueLocalFile = "local"
	cfsValueNacosFile = "nacos"
)

var (
	cfp = flag.String("cfp", ".", "配置文件路径\n")
	cfs = flag.String("cfs", cfsValueLocalFile, fmt.Sprintf("配置源名称，可能值: \n%s: 本地配置文件\n%s: 解析本地配置中指定的%s源的配置\n",
		cfsValueLocalFile, cfsValueNacosFile, cfsValueNacosFile))
	env   = flag.String("env", "", "自定义环境变量值\n")
	Nacos = &NacosConfig{}
	coder = yaml.NewEncoder()
)

func init() {
	flag.Parse()
}

func Env() string {
	return *env
}

// Init 初始化解析Nacos
func Init(cf IConfig) error {
	logger.Infof("Server Run Env=%s", Env())
	//解析config.yaml配置文件的nacos配置
	if err := loadFile(cf); err != nil {
		return err
	}
	if *cfs != cfsValueNacosFile {
		//解析本地的配置
		return nil
	}

	if err := loadFile(Nacos, *cfs); err != nil {
		return err
	}
	return loadNacos(Nacos, cf)
}

// LoadByNacosDataId 解析Nacos中的DataID配置项
func LoadByNacosDataId(ncf *NacosConfig, cf IConfig, dataId string, keys ...string) error {
	c, err := getNacosConfig(ncf, dataId)
	if err != nil {
		return err
	}
	return load(cf, c, keys...)
}

// loadFile 加载配置文件配置并解析
func loadFile(cf IConfig, keys ...string) error {
	var (
		filepath string
		mode     = Env()
	)
	if mode == "" {
		filepath = fmt.Sprintf("%s/config.yaml", *cfp)
	} else {
		filepath = fmt.Sprintf("%s/config-%s.yaml", *cfp, mode)
	}

	if len(keys) == 0 {
		logger.Infof("开始加载配置文件,文件路径cfp: %s, 配置源cfs:", filepath)
	} else {
		logger.Infof("开始加载配置文件,文件路径cfp: %s, 配置源cfs: %v", filepath, keys[0])
	}

	opts := []config.Option{
		config.WithSource(file.NewSource(file.WithPath(filepath))),
		config.WithReader(json.NewReader(reader.WithEncoder(coder))),
	}
	// 加载配置
	c, err := config.NewConfig(opts...)
	if err != nil {
		return err
	}
	return load(cf, c, keys...)
}

// loadNacos 加载Nacos中的配置项并解析
func loadNacos(ncf *NacosConfig, cf IConfig, keys ...string) error {
	c, err := getNacosConfig(ncf)
	if err != nil {
		return err
	}
	return load(cf, c, keys...)
}

// load 加载配置
func load(cf IConfig, c config.Config, keys ...string) error {
	//读取配置项的key列表
	val, err := c.Get(keys...)
	if err != nil {
		logger.Errorf("LoadRemote cfg.Get failed,error: %v", err)
		return err
	}
	if err = reload(val, cf); err != nil {
		logger.Errorf("LoadRemote reload failed,error: %v", err)
		return err
	}

	return watch(cf, c, keys...)
}

// 获取nacos配置项
func getNacosConfig(ncf *NacosConfig, dataId ...string) (config.Config, error) {
	id := ncf.DataId
	if len(dataId) != 0 {
		id = dataId[0]
	}

	//加载配置源配置
	sOpts := []source.Option{
		nacosSource.WithClientConfig(newClientConfig(ncf.Namespace, ncf.User, ncf.Password)),
		nacosSource.WithAddress(ncf.Hosts), // Nacos 地址
		nacosSource.WithDataId(id),         // Data ID
		nacosSource.WithGroup(ncf.Group),   // Group
	}

	var opts []config.Option
	if strings.Contains(id, "yaml") {
		opts = append(opts, config.WithReader(json.NewReader(reader.WithEncoder(coder))))
		sOpts = append(sOpts, source.WithEncoder(coder))
	}

	// 加载配置
	c, err := config.NewConfig(opts...)
	if err != nil {
		logger.Errorf("getNacosOptions config.NewConfig failed,error: %v", err)
		return nil, err
	}
	if err = c.Load(nacosSource.NewSource(sOpts...)); err != nil {
		logger.Errorf("getNacosOptions c.Load failed,error: %v", err)
		return nil, err
	}
	return c, nil
}

// 监听nacos变动
func watch(cf IConfig, c config.Config, paths ...string) error {
	w, err := c.Watch(paths...)
	if err != nil {
		logger.Errorf("Watch c.Watch failed,error: %v", err)
		return err
	}
	go func() {
		for {
			next, err := w.Next()
			if err != nil {
				logger.Errorf("Watch watch.Next failed,error: %v", err)
				time.Sleep(time.Second * 5)
				continue
			}
			if err = reload(next, cf); err != nil {
				logger.Errorf("Watch reload failed,error: %v", err)
			} else {
				logger.Infof("Watch reload success")
			}
		}
	}()
	return nil
}

func newClientConfig(namespaceId, user, passwd string) constant.ClientConfig {
	return constant.ClientConfig{
		NamespaceId:         namespaceId,
		TimeoutMs:           50000,
		Username:            user,
		Password:            passwd,
		NotLoadCacheAtStart: true,
		LogDir:              "./logs/nacos/nacos.log",
		CacheDir:            "./logs/nacos/cache",
		LogLevel:            "debug",
	}
}

func reload(read reader.Value, conf IConfig) (err error) {
	data := conf.Original()
	err = resetOrigin(read, data)
	if err != nil {
		return err
	}
	return conf.Handle()
}

// resetOrigin 先清空target原属性值，再从read中重新赋值
func resetOrigin(read reader.Value, target interface{}) error {
	// 实例化一个target新对象
	ntv := reflect.New(reflect.TypeOf(target).Elem())
	nt := ntv.Interface()
	if err := read.Scan(nt); err != nil {
		return err
	}
	reflect.ValueOf(target).Elem().Set(ntv.Elem())
	return nil
}

// resetJsonOrigin 根据raw更新target对象，先清空target原属性值
func resetJsonOrigin(raw []byte, target interface{}) error {
	// 实例化一个target新对象
	ntv := reflect.New(reflect.TypeOf(target).Elem())
	nt := ntv.Interface()
	if err := json2.Unmarshal(raw, nt); err != nil {
		return err
	}
	reflect.ValueOf(target).Elem().Set(ntv.Elem())
	return nil
}
