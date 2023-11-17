package core

type ConfigOption interface {
	ConfigReader
	ConfigDecoder
}

type ConfigReader interface {
	Read(path string) ([]byte, error) // 读取配置
}

type ConfigDecoder interface {
	Decode(data []byte, conf interface{}) interface{}
}

type ConfigReloader interface {
	Reload() error
}

type ConfigWatcher interface {
	Watch() error
}
