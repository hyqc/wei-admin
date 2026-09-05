// Package storage 文件存储抽象层
//
// 设计目标：上传逻辑只依赖 Driver 接口，新增一种存储（如七牛云、MinIO、FTP）
// 只需实现 Driver 并在自己的 init / 注册函数中调用 Register，无需改动业务代码。
package storage

import (
	"admin/pkg/config"
	"context"
	"fmt"
	"strings"
	"sync"
)

// 内置驱动名，与配置 upload.driver 对应
const (
	DriverLocal  = "local"
	DriverAliyun = "aliyun"
	DriverQcloud = "qcloud"
	DriverS3     = "s3"
)

// Driver 存储驱动接口
type Driver interface {
	// Name 驱动名，与配置的 driver 一致
	Name() string
	// Put 写入文件；objectKey 为不含域名的相对路径，如 admin/user/2026/09/xxx.png
	Put(ctx context.Context, objectKey string, data []byte, contentType string) error
	// Delete 删除文件；对象不存在时应返回 nil（幂等），便于记录与存储不一致时清理
	Delete(ctx context.Context, objectKey string) error
	// URL 返回对象的可访问链接
	URL(objectKey string) string
}

// Factory 驱动构造函数，配置校验与客户端初始化在此完成
type Factory func(cfg *config.Upload) (Driver, error)

var (
	mu      sync.RWMutex
	factory = map[string]Factory{}
)

// Register 注册存储驱动；通常在驱动的 init 中调用，重复注册会覆盖同名驱动
func Register(name string, f Factory) {
	mu.Lock()
	defer mu.Unlock()
	factory[strings.ToLower(name)] = f
}

// Drivers 返回已注册的驱动名列表
func Drivers() []string {
	mu.RLock()
	defer mu.RUnlock()
	names := make([]string, 0, len(factory))
	for name := range factory {
		names = append(names, name)
	}
	return names
}

// New 按配置创建存储驱动；driver 为空时默认使用本地存储
func New(cfg *config.Upload) (Driver, error) {
	if cfg == nil {
		return nil, fmt.Errorf("storage: upload config is nil")
	}
	name := strings.ToLower(strings.TrimSpace(cfg.Driver))
	if name == "" {
		name = DriverLocal
	}
	mu.RLock()
	f, ok := factory[name]
	mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("storage: unsupported driver %q, registered: %v", name, Drivers())
	}
	return f(cfg)
}

// joinKey 拼接对象键：忽略空片段，统一使用 / 分隔
func joinKey(parts ...string) string {
	segs := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.Trim(p, "/")
		if p != "" {
			segs = append(segs, p)
		}
	}
	return strings.Join(segs, "/")
}

// joinURL 拼接访问链接，避免多余或缺失的斜杠
func joinURL(domain, key string) string {
	key = strings.TrimLeft(key, "/")
	if domain == "" {
		return "/" + key
	}
	return strings.TrimRight(domain, "/") + "/" + key
}
