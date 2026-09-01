package captcha

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// redisStore 基于 Redis 的验证码存储，适用于集群（多副本）部署
type redisStore struct {
	client *redis.Client
	expire time.Duration
	prefix string
}

// NewRedisStore 创建 Redis 存储。keyPrefix 用于区分不同场景，expire 为验证码有效期
func NewRedisStore(client *redis.Client, expire time.Duration, keyPrefix string) Store {
	if expire <= 0 {
		expire = DefaultExpire
	}
	if keyPrefix == "" {
		keyPrefix = "captcha:"
	}
	return &redisStore{client: client, expire: expire, prefix: keyPrefix}
}

func (s *redisStore) key(id string) string {
	return s.prefix + id
}

func (s *redisStore) Set(id, value string) error {
	if s.client == nil || id == "" {
		return nil
	}
	return s.client.Set(context.Background(), s.key(id), value, s.expire).Err()
}

func (s *redisStore) Get(id string, clear bool) string {
	if s.client == nil || id == "" {
		return ""
	}
	val, err := s.client.Get(context.Background(), s.key(id)).Result()
	if err != nil {
		return ""
	}
	if clear {
		s.client.Del(context.Background(), s.key(id))
	}
	return val
}

func (s *redisStore) Verify(id, answer string, clear bool) bool {
	return s.Get(id, clear) == answer
}
