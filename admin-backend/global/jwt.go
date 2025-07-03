package global

import (
	"admin/pkg/utils/jwt"
	"time"
)

func initJwt() error {
	if AppConfig.JWT.Expire <= 0 {
		AppConfig.JWT.Expire = time.Hour * 24 * 7
	}

	//初始化Auth
	auth, err := jwt.NewAuth(AppConfig.JWT.Public, AppConfig.JWT.Private)
	if err != nil {
		LogSugar.Error("initServer jwt.NewAuth error: %v", err)
		return err
	}
	AppAuth = auth

	return nil
}
