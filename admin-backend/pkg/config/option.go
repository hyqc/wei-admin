package config

type Options struct {
	configFilePath     string //配置文件路径
	configEnvSeparator string //配置文件与环境变量组合分隔符默认.
}

const (
	defaultConfigFilePath = "config"
)

type Option func(options Options)

func NewOptions(opts ...Option) Options {
	c := Options{}
	if len(opts) == 0 {
		opts = append(opts,
			WithConfigFilePath("config"),
			WithConfigEnvSeparator("."),
		)
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func WithConfigFilePath(p string) Option {
	return func(options Options) {
		if p == "" {
			p = "config"
		}
		options.configFilePath = p
	}
}

func WithConfigEnvSeparator(s string) Option {
	return func(options Options) {
		if s == "" {
			s = "."
		}
		options.configEnvSeparator = s
	}
}
