package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
)

type Proto struct {
	FileName string
	Dir      string
	Tables   []Item
	Headers  []string
}

type Item struct {
	TableName    string
	MessageName  string
	IgnoreFields map[string]struct{}
}

var (
	//数据库连接字符串
	dataSourceName = "root:123456@tcp(127.0.0.1:3306)/wei"
	// 指定要转换的表名 "admin_role", "admin_menu", "admin_api", "admin_user", "admin_permission"
	tables = []Item{
		{
			TableName:   "admin_role",
			MessageName: "AdminRoleModel",
		},
		{
			TableName:   "admin_menu",
			MessageName: "AdminMenuModel",
		},
		{
			TableName:   "admin_api",
			MessageName: "AdminApiModel",
		},
		{
			TableName:   "admin_user",
			MessageName: "AdminUserModel",
			IgnoreFields: map[string]struct{}{
				"password": {},
			},
		},
	}

	conf = Proto{
		FileName: "admin_model",
		Dir:      "./proto/",
		Tables:   tables,
		Headers: []string{
			"syntax = \"proto3\";",
			"package admin;",
			"option go_package = \"./admin_proto;admin_proto\";",
			//"import \"google/protobuf/timestamp.proto\";",
		},
	}
)

func main() {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	gen(db)
}

func gen(db *sql.DB) {
	var protoContent strings.Builder
	for _, str := range conf.Headers {
		protoContent.WriteString(fmt.Sprintf("%s;\n", str))
	}
	protoContent.WriteString(fmt.Sprintf("\n\n"))
	for _, table := range tables {
		rows, err := db.Query(fmt.Sprintf("DESCRIBE %s", table.TableName))
		if err != nil {
			panic(err)
		}
		generateProtoFile(&protoContent, table, rows)
	}

	fileName := fmt.Sprintf("%s%s.proto", conf.Dir, conf.FileName)
	fmt.Println("生成文件：", fileName)
	err := writeToFile(fileName, protoContent.String())
	if err != nil {
		log.Fatal(err)
	}
}

func generateProtoFile(protoContent *strings.Builder, table Item, rows *sql.Rows) {
	protoContent.WriteString(fmt.Sprintf("message %s {\n", table.MessageName))
	fieldIndex := 1
	for rows.Next() {
		var (
			null, key, def, extra any
			fieldName, fieldType  string
		)
		if err := rows.Scan(&fieldName, &fieldType, &null, &key, &def, &extra); err != nil {
			log.Fatal(err)
		}
		if table.IgnoreFields != nil {
			if _, ok := table.IgnoreFields[fieldName]; ok {
				continue
			}
		}
		protoType := mapMySQLTypeToProtoType(fieldName, fieldType)
		protoContent.WriteString(fmt.Sprintf("  %s %s = %d;\n", protoType, SnakeToCamel(fieldName), fieldIndex))
		fieldIndex++
	}
	protoContent.WriteString("}\n\n\n")

}

// SnakeToCamel 将蛇形命名转换为小驼峰命名
func SnakeToCamel(snakeStr string) string {
	var camelStrBuilder strings.Builder
	capitalizeNext := false

	for i, r := range snakeStr {
		if r == '_' {
			capitalizeNext = true
		} else {
			if capitalizeNext && i > 0 {
				camelStrBuilder.WriteRune(unicode.ToUpper(r))
				capitalizeNext = false
			} else {
				camelStrBuilder.WriteRune(unicode.ToLower(r))
			}
		}
	}

	return camelStrBuilder.String()
}

func mapMySQLTypeToProtoType(fieldName, mysqlType string) string {
	if strings.Contains(fieldName, "is_") {
		return "bool"
	}
	arr := strings.Split(mysqlType, " ")
	mysqlType = arr[0]
	switch mysqlType {
	case "int", "tinyint", "smallint", "mediumint":
		return "int32"
	case "bigint":
		return "int64"
	case "float", "double":
		return "double"
	case "decimal":
		return "fixed64" // 或者使用 fixed64
	case "char", "varchar", "text", "blob":
		return "string"
	case "date", "datetime", "timestamp":
		//return "google.protobuf.Timestamp" // 或者使用 google.protobuf.Timestamp
		return "string" // 或者使用 google.protobuf.Timestamp
	case "bool", "boolean":
		return "bool"
	case "null":
		return "ni"
	default:
		return "string"
	}
}

func writeToFile(fileName, content string) error {
	return os.WriteFile(fileName, []byte(content), 0644)
}
