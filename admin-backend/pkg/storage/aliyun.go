package storage

import (
	"admin/pkg/config"
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Aliyun 阿里云对象存储 OSS
type Aliyun struct {
	bucket *oss.Bucket
	domain string
	base   string
}

func init() {
	Register(DriverAliyun, func(cfg *config.Upload) (Driver, error) {
		c := cfg.Aliyun
		var missing []string
		if c.Endpoint == "" {
			missing = append(missing, "endpoint")
		}
		if c.AccessKeyID == "" {
			missing = append(missing, "access_key_id")
		}
		if c.AccessKeySecret == "" {
			missing = append(missing, "access_key_secret")
		}
		if c.Bucket == "" {
			missing = append(missing, "bucket")
		}
		if len(missing) > 0 {
			return nil, fmt.Errorf("storage(aliyun): missing config: %v", missing)
		}
		client, err := oss.New(c.Endpoint, c.AccessKeyID, c.AccessKeySecret)
		if err != nil {
			return nil, fmt.Errorf("storage(aliyun): init client: %w", err)
		}
		bucket, err := client.Bucket(c.Bucket)
		if err != nil {
			return nil, fmt.Errorf("storage(aliyun): init bucket: %w", err)
		}
		return &Aliyun{
			bucket: bucket,
			domain: cfg.Domain,
			base:   c.BasePath,
		}, nil
	})
}

func (a *Aliyun) Name() string { return DriverAliyun }

func (a *Aliyun) Put(_ context.Context, objectKey string, data []byte, contentType string) error {
	opts := make([]oss.Option, 0, 1)
	if contentType != "" {
		opts = append(opts, oss.ContentType(contentType))
	}
	return a.bucket.PutObject(a.fullKey(objectKey), bytes.NewReader(data), opts...)
}

func (a *Aliyun) Delete(_ context.Context, objectKey string) error {
	if err := a.bucket.DeleteObject(a.fullKey(objectKey)); err != nil {
		// 对象已不存在时视为成功，保证删除幂等
		var svcErr oss.ServiceError
		if errors.As(err, &svcErr) && svcErr.StatusCode == 404 {
			return nil
		}
		return err
	}
	return nil
}

func (a *Aliyun) URL(objectKey string) string {
	if a.domain != "" {
		return joinURL(a.domain, a.fullKey(objectKey))
	}
	// 未配置域名时回退到 OSS 默认公网地址
	return joinURL(fmt.Sprintf("https://%s.%s", a.bucket.BucketName, a.bucket.Client.Config.Endpoint), a.fullKey(objectKey))
}

func (a *Aliyun) fullKey(objectKey string) string {
	return joinKey(a.base, objectKey)
}
