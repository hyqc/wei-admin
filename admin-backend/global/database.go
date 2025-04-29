package global

import (
	"admin/app/admin/gen/query"
	"admin/pkg/components/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Wei mysql.Config `yaml:"wei"`
}

type DBClient struct {
	Wei *gorm.DB
}

func InitMySQLDB() error {
	AppDB = &DBClient{}

	db, err := mysql.New(&AppConfig.Database.Wei)
	if err != nil {
		return err
	}

	query.SetDefault(db)

	AppDB.Wei = db

	return nil
}
