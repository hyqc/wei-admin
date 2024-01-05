package service

import (
	"admin/app/gen/query"
	"admin/pkg/utils"
	"admin/pkg/utils/files"
	"admin/proto/admin_account"
	"context"
	"encoding/json"
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
	defer func() {
		_ = db.Close()
	}()
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

func TestAdminUserService_Login(t *testing.T) {
	srv := NewAdminUserService()
	resp, err := srv.Login(context.Background(), &admin_account.LoginReq{
		Username: "admin",
		Password: "123456",
	}, "127.0.0.1")
	assert.Nil(t, err, err)
	body, err := json.Marshal(resp)
	assert.Nil(t, err, err)
	err = files.Override("login.resp.json", body)
	assert.Nil(t, err, err)
}
