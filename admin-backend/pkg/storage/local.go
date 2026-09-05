package storage

import (
	"admin/pkg/config"
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// 本地存储的默认配置
const (
	defaultLocalRoot   = "./upload"
	defaultLocalPrefix = "/upload"
)

// Local 本地磁盘存储，适用于单机部署
type Local struct {
	root   string // 磁盘根目录
	prefix string // 静态访问前缀
	domain string // 访问域名前缀
}

func init() {
	Register(DriverLocal, func(cfg *config.Upload) (Driver, error) {
		root := strings.TrimSpace(cfg.Local.Root)
		if root == "" {
			root = defaultLocalRoot
		}
		prefix := strings.TrimSpace(cfg.Local.Prefix)
		if prefix == "" {
			prefix = defaultLocalPrefix
		}
		if !strings.HasPrefix(prefix, "/") {
			prefix = "/" + prefix
		}
		return &Local{
			root:   filepath.Clean(root),
			prefix: strings.TrimRight(prefix, "/"),
			domain: strings.TrimSpace(cfg.Domain),
		}, nil
	})
}

func (l *Local) Name() string { return DriverLocal }

// Put 写入文件；自动创建目录并做路径穿越校验
func (l *Local) Put(_ context.Context, objectKey string, data []byte, _ string) error {
	path, err := l.absPath(objectKey)
	if err != nil {
		return err
	}
	if err = os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// Delete 删除文件；文件不存在视为成功
func (l *Local) Delete(_ context.Context, objectKey string) error {
	path, err := l.absPath(objectKey)
	if err != nil {
		return err
	}
	if err = os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

// URL 返回访问链接：配置了 domain 时返回绝对地址，否则返回站点相对路径
func (l *Local) URL(objectKey string) string {
	return joinURL(l.domain, l.prefix+"/"+strings.TrimLeft(objectKey, "/"))
}

// absPath 把对象键转换为磁盘绝对路径，并拦截 ../ 穿越
func (l *Local) absPath(objectKey string) (string, error) {
	key := filepath.FromSlash(strings.TrimLeft(objectKey, "/"))
	if strings.Contains(key, "..") {
		return "", errors.New("storage(local): illegal object key")
	}
	abs, err := filepath.Abs(l.root)
	if err != nil {
		return "", err
	}
	path := filepath.Join(abs, key)
	if !strings.HasPrefix(path, abs+string(os.PathSeparator)) {
		return "", errors.New("storage(local): illegal object key")
	}
	return path, nil
}
