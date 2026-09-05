// Package types 管理后台的传输结构定义（proto 之外的补充常量）
package types

// 存储驱动展示名（与 pkg/storage 的驱动标识对应）
const (
	DriverTextLocal  = "本地"
	DriverTextAliyun = "阿里云OSS"
	DriverTextQcloud = "腾讯云COS"
	DriverTextS3     = "亚马逊S3"
)
