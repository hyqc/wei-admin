package gormgen

import (
	"fmt"
	"github.com/golang/mock/mockgen/model"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const Utf8mb4 = "utf8mb4"

type GenType struct {
	Table string
	Type  string
}

func Init(name, pwd, ip, database, charset string, tables []GenType) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", name, pwd, ip, database, charset)
	// dsn := "root:123456@tcp(192.168.9.100:3306)/hope_city_chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db)
	for _, item := range tables {
		g.ApplyInterface(func(method model.Method) {}, g.GenerateModelAs(item.Table, item.Type))
	}
	g.Execute()
}

func demo() {
	dsn := "root:123456@tcp(192.168.9.100:3306)/hope_city_chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db)
	g.GenerateModelAs("hc_chat_channel", "ChatChannel")
	g.GenerateModelAs("hc_chat_config", "ChatConfig")
	//g.ApplyInterface(func(method logic.Method) {}, g.GenerateModelAs("hc_chat_channel", "ChatChannel"))
	//g.ApplyInterface(func(method logic.Method) {}, g.GenerateModelAs("hc_chat_config", "ChatConfig"))
	g.Execute()
}
