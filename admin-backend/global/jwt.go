package global

import (
	"admin/pkg/utils/jwt"
	"strings"
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
		// 请求方法统一大写，与 ctx.Request.Method（POST/GET）保持一致
		method := strings.ToUpper(strings.TrimSpace(val.Method))
		for _, v := range val.Paths {
			if _, ok := AppConfig.JWT.IgnoresMap[method]; !ok {
				AppConfig.JWT.IgnoresMap[method] = map[string]struct{}{}
			}
			AppConfig.JWT.IgnoresMap[method][v] = struct{}{}
		}
	}

	return nil
}
