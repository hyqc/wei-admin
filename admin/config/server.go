package config

type Server struct {
	Debug bool   `yaml:"debug"`
	Port  string `yaml:"port"`
	Pprof bool   `yaml:"pprof"`
	JWT   JWT    `yaml:"jwt"`
}

type JWT struct {
	Secret string `yaml:"secret"`
}
