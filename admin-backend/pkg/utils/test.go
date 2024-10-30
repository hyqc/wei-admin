package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// TestMainSetupDB 测试Main Setup 毁掉
func TestMainSetupDB(call func(db *gorm.DB)) {
	// 初始化操作，例如设置数据库连接、初始化配置等
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", "root", "root", "localhost", "3306", "wei", "utf8mb4", "True", "Local")
	fmt.Println("dsn: ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Err opening database: %v", err)
	}
	call(db)
}
