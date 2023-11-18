package utils

import (
	"github.com/spf13/viper"
)

func GetConfigEnv(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}
