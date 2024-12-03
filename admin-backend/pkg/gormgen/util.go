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
	Table  string
	Type   string
	Fields []gen.ModelOpt
}

func Init(name, pwd, ip, database, charset string, tables []GenType, outputPath string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", name, pwd, ip, database, charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:       outputPath,
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})
	g.UseDB(db)

	for _, item := range tables {
		g.ApplyInterface(func(method model.Method) {}, g.GenerateModelAs(item.Table, item.Type, item.Fields...))
	}
	g.Execute()
}

func demo() {
	dsn := "root:123456@tcp(192.168.91.2:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db)
	g.GenerateModelAs("channel", "Channel")
	g.GenerateModelAs("config", "Config")
	//g.ApplyInterface(func(method logic.Method) {}, g.GenerateModelAs("hc_chat_channel", "ChatChannel"))
	//g.ApplyInterface(func(method logic.Method) {}, g.GenerateModelAs("hc_chat_config", "ChatConfig"))
	g.Execute()
}
