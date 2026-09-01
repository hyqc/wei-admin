// Package mongodb MongoDB 数据源配置与客户端创建
package mongodb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// pingTimeout 建立连接时的连通性校验超时
const pingTimeout = 5 * time.Second

// Config 单个 MongoDB 数据源配置
type Config struct {
	// URI 连接串，如 mongodb://user:password@127.0.0.1:27017/?authSource=admin
	URI string `json:"uri"`
	// Database 默认库名
	Database string `json:"database"`
	// MaxPoolSize 连接池最大连接数，0 表示使用驱动默认值
	MaxPoolSize uint64 `json:"max_pool_size"`
	// MinPoolSize 连接池最小连接数，0 表示使用驱动默认值
	MinPoolSize uint64 `json:"min_pool_size"`
}

// New 创建 MongoDB 客户端并校验连通性
func New(ctx context.Context, conf *Config) (*mongo.Client, error) {
	if conf == nil || conf.URI == "" {
		return nil, errors.New("mongodb: uri 不能为空")
	}
	opts := options.Client().ApplyURI(conf.URI)
	if conf.MaxPoolSize > 0 {
		opts = opts.SetMaxPoolSize(conf.MaxPoolSize)
	}
	if conf.MinPoolSize > 0 {
		opts = opts.SetMinPoolSize(conf.MinPoolSize)
	}
	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("mongodb: 连接失败: %w", err)
	}
	pingCtx, cancel := context.WithTimeout(ctx, pingTimeout)
	defer cancel()
	if err = client.Ping(pingCtx, readpref.Primary()); err != nil {
		// 连通性校验失败时释放资源
		_ = client.Disconnect(context.Background())
		return nil, fmt.Errorf("mongodb: 连通性校验失败: %w", err)
	}
	return client, nil
}
