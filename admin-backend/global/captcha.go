package global

import (
	"fmt"
	"strings"
	"time"

	"admin/pkg/captcha"
	"admin/pkg/config"
)

// initCaptcha 初始化验证码管理器。
// 存储方式由 captcha.store 决定：memory（单机，默认）| redis（集群，需同时配置 redis.addr）。
func initCaptcha() error {
	cfg := AppConfig.Captcha
	manager, err := captcha.New(storeFactory(cfg), buildCaptchaScenes(cfg))
	if err != nil {
		return err
	}
	AppCaptcha = manager
	return nil
}

// buildCaptchaScenes 构建各场景配置：未配置 scenes 时用全局默认值注册登录场景
func buildCaptchaScenes(cfg config.Captcha) map[captcha.Scene]captcha.Config {
	base := captcha.Config{
		Length: cfg.Length,
		Expire: time.Duration(cfg.Expire) * time.Second,
		Width:  cfg.Width,
		Height: cfg.Height,
	}
	scenes := make(map[captcha.Scene]captcha.Config, len(cfg.Scenes))
	if len(cfg.Scenes) == 0 {
		scenes[captcha.SceneLogin] = base
		return scenes
	}
	for name, sc := range cfg.Scenes {
		c := base
		if sc.Length > 0 {
			c.Length = sc.Length
		}
		if sc.Expire > 0 {
			c.Expire = time.Duration(sc.Expire) * time.Second
		}
		if sc.Width > 0 {
			c.Width = sc.Width
		}
		if sc.Height > 0 {
			c.Height = sc.Height
		}
		scenes[captcha.Scene(name)] = c
	}
	return scenes
}

// storeFactory 按部署模式选择验证码存储：
// memory（默认，单机部署）| redis（集群部署，必须配置可用的 redis.addr，否则启动失败）
func storeFactory(cfg config.Captcha) func(captcha.Scene, captcha.Config) (captcha.Store, error) {
	return func(scene captcha.Scene, c captcha.Config) (captcha.Store, error) {
		if strings.EqualFold(strings.TrimSpace(cfg.Store), "redis") {
			if AppRedis == nil {
				return nil, fmt.Errorf("captcha.store=redis 需要配置可用的 redis.addr")
			}
			return captcha.NewRedisStore(AppRedis, c.Expire, fmt.Sprintf("captcha:%s:", scene)), nil
		}
		return captcha.NewMemoryStore(c.Expire), nil
	}
}
