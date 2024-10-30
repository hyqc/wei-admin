package config

import (
	json2 "encoding/json"
	"flag"
	"fmt"
	"github.com/go-micro/plugins/v4/config/encoder/yaml"
	"github.com/go-micro/plugins/v4/config/source/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/pkg/errors"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
	"go-micro.dev/v4/config/source/file"
	"hope_city/lib/log"
	"hope_city/lib/util"
	"reflect"
	"strings"
)

type GameConfig interface {
	Reload(c reader.Value) error
}

type Config interface {
	Handle() error
	Original() interface{}
}

type Conf interface {
	FileName() string
}

type NaCos struct {
	Host      string `json:"host"`
	Port      uint64 `json:"port"`
	Namespace string `json:"namespace"`
	User      string `json:"user"`
	Password  string `json:"password"` //密码
	DataId    string `json:"dataId"`
	Group     string `json:"group"`   //分组
	Cluster   string `json:"cluster"` //tag 集群名
	Open      bool   `json:"open"`    //开关 true开启 false关闭（使用本地）
}

func (n *NaCos) Original() interface{} {
	return n
}

func (n *NaCos) Handle() error {
	return nil
}

func Init(c Config) error {
	//先读nacos配置
	flag.Parse()
	filename := *ConfFile
	env := util.GetHopeCityEnv()
	if env != "" {
		filename = "config-" + env + ".yaml"
	}
	log.Infof("加载配置文件名:%s", filename)
	err := Load(nil, filename, Nacos, "nacos")
	if err != nil {
		return err
	}
	log.Debug("===========")
	if Nacos.Open {
		filename = Nacos.DataId
	}
	return Load(Nacos, filename, c)
}

func ClientConfig(namespaceid, user, passwd string) constant.ClientConfig {
	return constant.ClientConfig{
		NamespaceId:         namespaceid,
		TimeoutMs:           50000,
		Username:            user,
		Password:            passwd,
		NotLoadCacheAtStart: true,
		LogDir:              "./logs/nacos.log",
		CacheDir:            "./logs/cache",
		LogLevel:            "debug",
	}
}

func LocalFiles(conf GameConfig, fileNames []string, watchOk bool) (err error) {
	if len(fileNames) == 0 {
		return
	}
	//进行反射
	if reflect.TypeOf(conf).Kind() != reflect.Ptr {
		panic("必须是指针")
	}
	log.Debug(reflect.TypeOf(conf).Elem().Kind())
	opts := []config.Option{}
	for _, fileName := range fileNames {
		opts = append(opts, config.WithSource(file.NewSource(
			file.WithPath(*ConfPath+fileName),
		)))
	}
	c, err := config.NewConfig(opts...)
	if err != nil {
		log.Errorf("config.NewConfig err:%v", err)
		return err
	}
	//	log.Debug(string(c.Bytes()))
	err = c.Scan(conf)
	if err != nil {
		panic(err)
	}
	err = conf.Reload(nil)
	if err != nil {
		panic(err)
	}
	if !watchOk {
		return nil
	}
	watch, err := c.Watch()
	if err != nil {
		log.Errorf("监听失败,err:%v", err)
		return
	}
	go func() {
		for {
			next, err := watch.Next()
			if err != nil {
				log.Errorf("监听配置失败 err:%s", err.Error())
				continue
			}
			log.Debug("监听到配置更改")

			//log.Debug(string(next.Bytes()))
			conf.Reload(next)
		}
	}()
	return

}

func Local(filename string, watchOk bool, val Config, path ...string) error {
	flag.Parse()
	opts := []config.Option{config.WithSource(file.NewSource(file.WithPath(filename)))}
	if strings.Contains(filename, "yaml") {
		enc := yaml.NewEncoder()
		opts = append(opts, config.WithReader(json.NewReader(
			reader.WithEncoder(enc),
		)))
	}
	c, err := config.NewConfig(opts...)
	if err != nil {
		return err
	}
	err = Reload(c.Get(path...), val)
	if err != nil {
		return err
	}

	if !watchOk {
		return nil
	}

	watch, err := c.Watch()
	go func() {
		for {
			next, err := watch.Next()
			if err != nil {
				log.Errorf("监听%s配置失败 err:%s", path, err.Error())
				continue
			}
			err = Reload(next, val)
			if err != nil {
				log.Errorf("监听解析%s配置失败 err:%s", path, err.Error())
			}
		}
	}()

	return nil
}

func NewNacos(path string, n *NaCos) error {
	if path == "" {
		return nil
	}

	return nil
}

func (n *NaCos) Load1(dataId string, watchOk bool, val Config) error {
	c, err := config.NewConfig()
	if err != nil {
		return err
	}

	cc := ClientConfig(n.Namespace, n.User, n.Password)
	host := fmt.Sprintf("%s:%d", n.Host, n.Port)

	err = c.Load(nacos.NewSource(nacos.WithClientConfig(cc), nacos.WithDataId(dataId),
		nacos.WithGroup(n.Group), nacos.WithAddress([]string{host})))
	if err != nil {
		return err
	}

	err = Reload(c.Get(), val)
	if err != nil {
		return errors.Wrapf(err, "scan %s 配置失败", dataId)
	}

	if !watchOk {
		return nil
	}

	watch, err := c.Watch()
	go func() {
		for {
			next, err := watch.Next()
			if err != nil {
				log.Errorf("监听%s配置失败 err:%s", dataId, err.Error())
				continue
			}
			err = Reload(next, val)
			if err != nil {
				log.Errorf("监听解析%s配置失败 err:%s", dataId, err.Error())
			}
		}
	}()
	return nil
}

var (
	ConfFile      = flag.String("nf", "config.yaml", "配置文件路径")
	ConfPath      = flag.String("p", "./config/", "配置文件路径 默认./config/")
	CheckRegistry = flag.Bool("ckr", false, "设置是否待K8s检测后再注册服务")
	SstName       = flag.String("sstname", "", "状态服务名")
	ServerAddr    = flag.String("saddr", "", "指定服务启动地址")
	ProtocolType  = flag.String("ptt", "", "通信协议启动类型")
	RpcAddr       = flag.String("rpcaddr", "", "指定rpc服务启动地址")
	HttpAddr      = flag.String("httpaddr", "", "指定http服务启动地址")
	TcpAddr       = flag.String("tcpaddr", "", "指定tcp服务启动地址")
	PublishIp     = flag.String("publiship", "", "指定本机外网ip")
	Nacos         = &NaCos{}
)

func GetPublishIp() (publishIp string) {
	publishIp = *PublishIp
	if len(publishIp) == 0 {
		publishIp = util.GetOutboundIP()
	}
	return
}

func GetSrvAddr(name string, values ...string) (addr string) {
	log.Debug("addr:", *ServerAddr)
	if len(*ServerAddr) > 0 {
		addr = *ServerAddr
		return
	}
	addrKey := fmt.Sprintf("HOPE_CITY_%s_ADDR", strings.ToUpper(name))
	envAddr := util.GetEnvInfo(addrKey)
	log.Debug(addrKey, ":", addrKey)
	if len(envAddr) != 0 {
		addr = envAddr
		return
	}
	if len(values) > 0 {
		addr = values[0]
	}
	return
}

func GetSFN(values ...string) (sfn string) {
	flag.Parse()
	log.Debug("-sstname:", *SstName)
	if len(*SstName) > 0 {
		sfn = *SstName
	}
	if len(sfn) != 0 {
		return
	}
	envSfn := util.GetEnvInfo("HOPE_CITY_SFN")
	log.Debug("HOPE_CITY_SFN:", envSfn)
	if len(envSfn) != 0 {
		sfn = envSfn
		return
	}
	if len(values) > 0 && len(values[0]) > 0 {
		sfn = values[0]
		return
	}
	return sfn
}

func GetProtocolType(values ...string) (ptt string) {
	log.Debug("-ptt:", *ProtocolType)
	if len(*ProtocolType) > 0 {
		ptt = *ProtocolType
	}
	if len(ptt) != 0 {
		return
	}
	envPtt := util.GetEnvInfo("HOPE_CITY_PTT")
	log.Debug("HOPE_CITY_PTT:", envPtt)
	if len(envPtt) != 0 {
		ptt = envPtt
		return
	}
	if len(values) > 0 && len(values[0]) > 0 {
		ptt = values[0]
		return
	}
	return ptt
}

func GetTcpAddr(name string, values ...string) (addr string) {
	log.Debug("addr:", *TcpAddr)
	if len(*TcpAddr) > 0 {
		addr = *TcpAddr
		return
	}
	addrKey := fmt.Sprintf("HOPE_CITY_%s_TCPADDR", strings.ToUpper(name))
	envAddr := util.GetEnvInfo(addrKey)
	log.Debug(addrKey, ":", addrKey)
	if len(envAddr) != 0 {
		addr = envAddr
		return
	}
	if len(values) > 0 {
		addr = values[0]
	}
	return
}

func GetRpcAddr(name string, values ...string) (addr string) {
	log.Debug("rpcaddr:", *RpcAddr)
	if len(*RpcAddr) > 0 {
		addr = *RpcAddr
		return
	}
	addrKey := fmt.Sprintf("HOPE_CITY_%s_RPCADDR", strings.ToUpper(name))
	envAddr := util.GetEnvInfo(addrKey)
	log.Debug(addrKey, ":", addrKey)
	if len(envAddr) != 0 {
		addr = envAddr
		return
	}
	if len(values) > 0 {
		addr = values[0]
	}
	return
}

func GetHttpAddr(name string, values ...string) (addr string) {
	log.Debug("httpaddr:", *HttpAddr)
	if len(*HttpAddr) > 0 {
		addr = *HttpAddr
		return
	}
	addrKey := fmt.Sprintf("HOPE_CITY_%s_HTTPADDR", strings.ToUpper(name))
	envAddr := util.GetEnvInfo(addrKey)
	log.Debug(addrKey, ":", addrKey)
	if len(envAddr) != 0 {
		addr = envAddr
		return
	}
	if len(values) > 0 {
		addr = values[0]
	}
	return
}

func InitGameConf(c GameConfig, dataIds []string) (err error) {
	flag.Parse()
	err = NewNacos(*ConfFile, Nacos)
	if err != nil {
		return err
	}
	json.NewReader()
	if !Nacos.Open {
		err = LocalFiles(c, dataIds, true)
	} else {

	}
	return nil
}

func Reload(read reader.Value, conf Config) (err error) {
	data := conf.Original()
	err = ResetOrigin(read, data)
	if err != nil {
		return err
	}
	return conf.Handle()
}

func Load(n *NaCos, filename string, val Config, path ...string) (err error) {
	var c config.Config
	enc := yaml.NewEncoder()
	if n != nil && n.Open == true {
		cc := ClientConfig(n.Namespace, n.User, n.Password)
		host := fmt.Sprintf("%s:%d", n.Host, n.Port)
		sOpt := []source.Option{nacos.WithClientConfig(cc), nacos.WithDataId(filename),
			nacos.WithGroup(n.Group), nacos.WithAddress([]string{host})}
		opts := []config.Option{}
		if strings.Contains(filename, "yaml") {
			opts = append(opts, config.WithReader(json.NewReader(
				reader.WithEncoder(enc),
			)))
			sOpt = append(sOpt, source.WithEncoder(enc))
		}
		c, err = config.NewConfig(opts...)
		if err != nil {
			log.Errorf("config.NewConfig fail, err:%v", err)
			return err
		}
		err = c.Load(nacos.NewSource(sOpt...))
		if err != nil {
			log.Errorf("c.Load(nacos.NewSource(sOpt...)) fail, err:%v", err)
			return err
		}
	} else {
		filename = *ConfPath + filename
		opts := []config.Option{config.WithSource(file.NewSource(file.WithPath(filename)))}
		if strings.Contains(filename, "yaml") {
			opts = append(opts, config.WithReader(json.NewReader(
				reader.WithEncoder(enc),
			)))
		}
		c, err = config.NewConfig(opts...)
		if err != nil {
			log.Errorf("config.NewConfig fail, err:%v", err)
			return err
		}
	}
	err = Reload(c.Get(path...), val)
	if err != nil {
		return errors.Wrapf(err, "scan %s 配置失败", filename)
	}
	watch, err := c.Watch(path...)
	go func() {
		for {
			next, err := watch.Next()
			if err != nil {
				log.Errorf("监听%s配置失败 err:%s", filename, err.Error())
				continue
			}
			err = Reload(next, val)
			if err != nil {
				log.Errorf("监听解析%s配置失败 err:%s", filename, err.Error())
			} else {
				log.Errorf("%s-配置修改成功", filename)
			}
		}
	}()
	return nil
}

// ResetJsonOrigin 根据raw更新target对象，先清空target原属性值
func ResetJsonOrigin(raw []byte, target interface{}) error {
	// 实例化一个target新对象
	ntv := reflect.New(reflect.TypeOf(target).Elem())
	nt := ntv.Interface()
	if err := json2.Unmarshal(raw, nt); err != nil {
		return err
	}
	reflect.ValueOf(target).Elem().Set(ntv.Elem())
	return nil
}

// ResetOrigin 先清空target原属性值，再从read中重新赋值
func ResetOrigin(read reader.Value, target interface{}) error {
	// 实例化一个target新对象
	ntv := reflect.New(reflect.TypeOf(target).Elem())
	nt := ntv.Interface()
	if err := read.Scan(nt); err != nil {
		return err
	}
	reflect.ValueOf(target).Elem().Set(ntv.Elem())
	return nil
}
