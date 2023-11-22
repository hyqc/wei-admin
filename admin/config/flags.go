package config

import (
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"os"
)

type IFlags interface {
	Init() micro.Option
}

type Flags struct {
	Version   FlagItem
	ConfigDir FlagItem
}

type FlagItem struct {
	Name  string
	Value any
	Usage string
}

var (
	AppFlags = NewFlags()
)

func NewFlags() *Flags {
	return &Flags{
		Version: FlagItem{
			Name:  "v",
			Usage: "the server version",
			Value: AppVersion,
		},
		ConfigDir: FlagItem{
			Name:  "cd",
			Usage: "config dir, cd is config dir acronym",
			Value: configFileFullPath,
		},
	}
}

func (f *Flags) Init() []micro.Option {
	fg := AppFlags
	return []micro.Option{
		micro.Flags(
			&cli.BoolFlag{
				Name:  fg.Version.Name,
				Usage: fg.Version.Usage,
			},
			&cli.StringFlag{
				Name:  fg.ConfigDir.Name,
				Usage: fg.ConfigDir.Usage,
				Value: fg.ConfigDir.Value.(string),
			},
		),
		micro.Action(func(c *cli.Context) error {
			if c.Bool(AppFlags.Version.Name) {
				logger.Infof("The version is: %s", AppFlags.Version.Value)
				os.Exit(0)
			}
			if cd := c.String(AppFlags.ConfigDir.Name); cd != "" {
				AppFlags.ConfigDir.Value = cd
				logger.Infof("The config dir is: %s", cd)
			}
			return nil
		}),
	}
}
