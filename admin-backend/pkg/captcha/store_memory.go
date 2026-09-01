package captcha

import (
	"sync"
	"time"
)

// 内存存储默认垃圾回收间隔
const defaultCollectInterval = 10 * time.Minute

// memoryStore 进程内存储，适用于单机部署（多副本部署请使用 Redis 存储）
type memoryStore struct {
	mu        sync.Mutex
	items     map[string]memoryItem
	expire    time.Duration
	lastClean time.Time
}

type memoryItem struct {
	value    string
	expireAt time.Time
}

// NewMemoryStore 创建内存存储，expire 为验证码有效期
func NewMemoryStore(expire time.Duration) Store {
	if expire <= 0 {
		expire = DefaultExpire
	}
	return &memoryStore{
		items:     make(map[string]memoryItem),
		expire:    expire,
		lastClean: time.Now(),
	}
}

func (s *memoryStore) Set(id, value string) error {
	if id == "" {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clean()
	s.items[id] = memoryItem{value: value, expireAt: time.Now().Add(s.expire)}
	return nil
}

func (s *memoryStore) Get(id string, clear bool) string {
	if id == "" {
		return ""
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clean()
	item, ok := s.items[id]
	if !ok {
		return ""
	}
	if clear {
		delete(s.items, id)
	}
	if time.Now().After(item.expireAt) {
		delete(s.items, id)
		return ""
	}
	return item.value
}

func (s *memoryStore) Verify(id, answer string, clear bool) bool {
	return s.Get(id, clear) == answer
}

// clean 定期清理过期数据，需在持有锁时调用
func (s *memoryStore) clean() {
	now := time.Now()
	if now.Sub(s.lastClean) < defaultCollectInterval {
		return
	}
	s.lastClean = now
	for k, v := range s.items {
		if now.After(v.expireAt) {
			delete(s.items, k)
		}
	}
}
