package dao

import (
	"encoding/json"
	"strings"
)

type AdminDao struct {
	AdminUser       IAdminUser
	AdminPermission IAdminPermission
	AdminMenu       IAdminMenu
	AdminAPI        IAdminAPI
}

var H = newAdminDao()

func newAdminDao() *AdminDao {
	return &AdminDao{
		AdminUser:       newAdminUser(),
		AdminPermission: newAdminPermission(),
		AdminMenu:       newAdminMenu(),
		AdminAPI:        newAdminAPI(),
	}
}

// HandleAdminUserLastLoginIp 处理AdminUser.LastLoginIp，最多保存上一次登录的IP和本次登录的IP
func HandleAdminUserLastLoginIp(ips string) ([]string, error) {
	arr := make([]string, 0, 3)
	if ips == "[]" {
		return arr, nil
	}
	if err := json.Unmarshal([]byte(ips), &arr); err != nil {
		return nil, err
	}
	return arr, nil
}

// SetAdminUserLastLoginIp 处理AdminUser.LastLoginIp 设置最后登录IP值
func SetAdminUserLastLoginIp(clientIp, ips string) (string, error) {
	clientIp = strings.Trim(clientIp, " ")
	if clientIp == "" {
		return ips, nil
	}
	arr, err := HandleAdminUserLastLoginIp(ips)
	if err != nil {
		return "", err
	}
	arr = append(arr, clientIp)
	if len(arr) > 2 {
		arr = arr[1:3]
	}
	b, err := json.Marshal(arr)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
