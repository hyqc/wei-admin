package global

import (
	"context"
	"strings"

	"github.com/redis/go-redis/v9"
)

// initRedis 初始化 Redis 客户端。
// 仅在配置了 redis.addr 时初始化：单机部署不配置即可，集群部署（验证码共享存储）需配置。
func initRedis() error {
	cfg := AppConfig.Redis
	if strings.TrimSpace(cfg.Addr) == "" {
		return nil
	}
	AppRedis = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := AppRedis.Ping(context.Background()).Err(); err != nil {
		AppRedis = nil
		return err
	}
	return nil
}

// CloseRedis 关闭 Redis 连接
func CloseRedis() {
	if AppRedis != nil {
		_ = AppRedis.Close()
	}
}
