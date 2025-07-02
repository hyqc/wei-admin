package config

type NacosConfig struct {
	Hosts     []string `json:"hosts"`
	Namespace string   `json:"namespace"`
	User      string   `json:"user"`
	Password  string   `json:"password"` //密码
	DataId    string   `json:"dataId"`
	Group     string   `json:"group"`   //分组
	Cluster   string   `json:"cluster"` //tag 集群名
	Local     bool     `json:"open"`    //开关 true开启 false关闭（使用本地）
}

func (n *NacosConfig) Original() interface{} {
	return n
}

func (n *NacosConfig) Handle() error {
	return nil
}

type Server struct {
	Name    string `json:"name"`    //服务名
	Version string `json:"version"` //服务版本
	Host    string `json:"host"`    //服务地址
	Port    string `json:"port"`    //端口号
	Id      string `json:"id"`      //服务唯一ID,md5(addr)
}

type Broker struct {
	Address []string `json:"address"`
}

type Jwt struct {
	Private string `json:"private"`
	Public  string `json:"public"`
	Expire  string `json:"expire"`
}
