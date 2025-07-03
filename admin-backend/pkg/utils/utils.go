package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/spf13/viper"
	"io"
	"net"
	"reflect"
	"sort"
	"strings"
	"time"
	"unicode"
)

// GetOutBoundIP 获取对外IP
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func GetLocalBoundIP() ([]string, error) {
	var ips []string
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range interfaces {
		byIndex, err := net.InterfaceByName(i.Name)
		if err != nil {
			continue
		}
		addrs, err := byIndex.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // 不是IPv4地址
			}
			ips = append(ips, ip.String())
		}
	}
	sort.Slice(ips, func(i, j int) bool {
		return ips[i] > ips[j]
	})
	return ips, nil
}

// BeanCopy 结构体深拷贝
func BeanCopy(dst, src any) error {

	if reflect.TypeOf(src).Kind() != reflect.Pointer {
		return errors.New("src must be pointer")
	}
	if reflect.TypeOf(dst).Kind() != reflect.Pointer {
		return errors.New("dst must be pointer")
	}
	sv := reflect.ValueOf(src).Elem()
	if !sv.IsValid() {
		panic("BeanCopy src is nil")
	}

	dv := reflect.ValueOf(dst).Elem()
	if !dv.IsValid() {
		panic("BeanCopy dst is nil")
	}
	for i := 0; i < sv.NumField(); i++ {
		val := sv.Field(i)
		name := sv.Type().Field(i).Name
		kind := sv.Type().Field(i).Type.Kind()
		f := dv.FieldByName(name)
		if f.IsValid() && f.Type().Kind() == kind {
			f.Set(val)
		}
	}
	return nil
}

// GetConfigEnv 获取环境变量
func GetConfigEnv(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

// HandleTime2Local 将时间转换为本地时间
func HandleTime2Local(data ...*time.Time) {
	if data == nil || len(data) == 0 {
		return
	}
	for _, d := range data {
		if d == nil {
			continue
		}
		*d = d.In(time.Local)
	}
}

// HandleTime2String 将时间转换为字符串
func HandleTime2String(data time.Time) string {
	if data.IsZero() {
		return ""
	}
	return data.Format(time.RFC3339)
}

// HandleTimePointer2String 将指针指向的时间转换为字符串
func HandleTimePointer2String(data *time.Time) string {
	if data == nil || data.IsZero() {
		return ""
	}
	return data.Format(time.RFC3339)
}

// CamelToSnake 将小驼峰命名风格的字符串转换为下划线命名风格。
func CamelToSnake(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func Md5(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}
