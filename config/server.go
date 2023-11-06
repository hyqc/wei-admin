package config

type Server struct {
	Debug bool   `yaml:"debug"`
	Port  string `yaml:"port"`
	Pprof bool   `yaml:"pprof"`
}
