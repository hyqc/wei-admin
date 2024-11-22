package utils

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"net"
	"reflect"
	"strings"
	"time"
)

// GetOutBoundIP 获取对外IP
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

// BeanCopy 结构体深拷贝
func BeanCopy(dst, src interface{}) error {

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

func GetConfigEnv(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

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

func HandleTime2String(data time.Time) string {
	if data.IsZero() {
		return ""
	}
	return data.Format(time.RFC3339)
}

func HandleTimePointer2String(data *time.Time) string {
	if data == nil || data.IsZero() {
		return ""
	}
	return data.Format(time.RFC3339)
}
