package storage

import (
	"admin/pkg/config"
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	cos "github.com/tencentyun/cos-go-sdk-v5"
)

// Qcloud 腾讯云对象存储 COS
type Qcloud struct {
	client *cos.Client
	domain string
	base   string
}

func init() {
	Register(DriverQcloud, func(cfg *config.Upload) (Driver, error) {
		c := cfg.Qcloud
		var missing []string
		if c.Region == "" {
			missing = append(missing, "region")
		}
		if c.SecretID == "" {
			missing = append(missing, "secret_id")
		}
		if c.SecretKey == "" {
			missing = append(missing, "secret_key")
		}
		if c.Bucket == "" {
			missing = append(missing, "bucket")
		}
		if len(missing) > 0 {
			return nil, fmt.Errorf("storage(qcloud): missing config: %v", missing)
		}
		bucketURL, err := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", c.Bucket, c.Region))
		if err != nil {
			return nil, fmt.Errorf("storage(qcloud): parse bucket url: %w", err)
		}
		client := cos.NewClient(
			&cos.BaseURL{BucketURL: bucketURL},
			&http.Client{Transport: &cos.AuthorizationTransport{SecretID: c.SecretID, SecretKey: c.SecretKey}},
		)
		return &Qcloud{
			client: client,
			domain: cfg.Domain,
			base:   c.BasePath,
		}, nil
	})
}

func (q *Qcloud) Name() string { return DriverQcloud }

func (q *Qcloud) Put(ctx context.Context, objectKey string, data []byte, contentType string) error {
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentLength: int64(len(data)),
		},
	}
	if contentType != "" {
		opt.ObjectPutHeaderOptions.ContentType = contentType
	}
	_, err := q.client.Object.Put(ctx, q.fullKey(objectKey), bytes.NewReader(data), opt)
	return err
}

func (q *Qcloud) Delete(ctx context.Context, objectKey string) error {
	_, err := q.client.Object.Delete(ctx, q.fullKey(objectKey))
	if err != nil {
		// 对象不存在视为成功
		var resp *cos.ErrorResponse
		if errors.As(err, &resp) && resp.Response != nil && resp.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		return err
	}
	return nil
}

func (q *Qcloud) URL(objectKey string) string {
	if q.domain != "" {
		return joinURL(q.domain, q.fullKey(objectKey))
	}
	// 未配置域名时使用 COS 默认域名
	bucketURL := q.client.BaseURL.BucketURL
	return joinURL(fmt.Sprintf("%s://%s", bucketURL.Scheme, bucketURL.Host), q.fullKey(objectKey))
}

func (q *Qcloud) fullKey(objectKey string) string {
	return joinKey(q.base, objectKey)
}
