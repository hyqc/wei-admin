package dao

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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

func TestFindAll(t *testing.T) {
	dao := newAdminMenu()
	data, err := dao.FindAllValid(context.Background())
	assert.Nil(t, err, err)
	for _, item := range data {
		fmt.Println(item)
	}
}
