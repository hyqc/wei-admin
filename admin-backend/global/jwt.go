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

	//转ignore
	if AppConfig.JWT.IgnoresMap == nil {
		AppConfig.JWT.IgnoresMap = map[string]map[string]struct{}{}
	}
	for _, val := range AppConfig.JWT.Ignores {
		for _, v := range val.Paths {
			if _, ok := AppConfig.JWT.IgnoresMap[val.Method]; !ok {
				AppConfig.JWT.IgnoresMap[val.Method] = map[string]struct{}{}
			}
			AppConfig.JWT.IgnoresMap[val.Method][v] = struct{}{}
		}
	}

	return nil
}
