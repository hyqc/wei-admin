package storage

import (
	"admin/pkg/config"
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// S3 亚马逊 S3 及其他 S3 兼容实现（MinIO、Ceph 等）
type S3 struct {
	client *s3.Client
	bucket string
	domain string
	base   string
}

func init() {
	Register(DriverS3, func(cfg *config.Upload) (Driver, error) {
		c := cfg.S3
		var missing []string
		if c.Region == "" {
			missing = append(missing, "region")
		}
		if c.AccessKeyID == "" {
			missing = append(missing, "access_key_id")
		}
		if c.SecretAccessKey == "" {
			missing = append(missing, "secret_access_key")
		}
		if c.Bucket == "" {
			missing = append(missing, "bucket")
		}
		if len(missing) > 0 {
			return nil, fmt.Errorf("storage(s3): missing config: %v", missing)
		}
		client := s3.NewFromConfig(
			aws.Config{
				Region:      c.Region,
				Credentials: credentials.NewStaticCredentialsProvider(c.AccessKeyID, c.SecretAccessKey, ""),
			},
			func(o *s3.Options) {
				// 自定义端点用于 MinIO 等 S3 兼容存储，这类实现通常要求 path-style 寻址
				if c.Endpoint != "" {
					o.BaseEndpoint = aws.String(c.Endpoint)
					o.UsePathStyle = true
				}
			},
		)
		return &S3{
			client: client,
			bucket: c.Bucket,
			domain: cfg.Domain,
			base:   c.BasePath,
		}, nil
	})
}

func (s *S3) Name() string { return DriverS3 }

func (s *S3) Put(ctx context.Context, objectKey string, data []byte, contentType string) error {
	input := &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    aws.String(s.fullKey(objectKey)),
		Body:   bytes.NewReader(data),
	}
	if contentType != "" {
		input.ContentType = aws.String(contentType)
	}
	_, err := s.client.PutObject(ctx, input)
	return err
}

func (s *S3) Delete(ctx context.Context, objectKey string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &s.bucket,
		Key:    aws.String(s.fullKey(objectKey)),
	})
	if err != nil {
		// NoSuchKey 视为删除成功
		var nsk *types.NoSuchKey
		if errors.As(err, &nsk) {
			return nil
		}
		return err
	}
	return nil
}

func (s *S3) URL(objectKey string) string {
	if s.domain != "" {
		return joinURL(s.domain, s.fullKey(objectKey))
	}
	// 未配置域名时使用 AWS S3 默认域名
	return joinURL(fmt.Sprintf("https://%s.s3.amazonaws.com", s.bucket), s.fullKey(objectKey))
}

func (s *S3) fullKey(objectKey string) string {
	return joinKey(s.base, objectKey)
}
