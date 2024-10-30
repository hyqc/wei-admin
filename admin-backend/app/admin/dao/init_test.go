package dao

import (
	"admin/app/gen/query"
	"admin/pkg/utils"
	"fmt"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 在这里设置数据库连接
	utils.TestMainSetupDB(func(db *gorm.DB) {
		query.SetDefault(db)
	})
	// 运行测试
	code := m.Run()
	teardown()
	os.Exit(code)
}

func teardown() {
	// 清理操作，例如关闭数据库连接、清除测试数据等
	fmt.Println("执行清理...")
}
