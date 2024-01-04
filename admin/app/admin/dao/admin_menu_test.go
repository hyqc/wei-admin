package dao

import (
	"admin/app/gen/query"
	"admin/pkg/utils"
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
)

func mockDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// 设置你期望的 SQL 调用和结果
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "example")
	mock.ExpectQuery("^SELECT (.+) FROM users$").WillReturnRows(rows)
	// 这里写你的测试逻辑，使用 db 作为数据库连接
}

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

func TestFindAll(t *testing.T) {
	dao := NewAdminMenu()
	data, err := dao.FindAll(context.Background())
	assert.Nil(t, err, err)
	for _, item := range data {
		fmt.Println(item)
	}
}
