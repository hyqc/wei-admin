package global

import (
	"admin/app/admin/gen/query"
	"admin/pkg/config/store/mysql"
	"gorm.io/gorm"
)

type Store struct {
	Wei mysql.Config `json:"wei"`
}

type DBClient struct {
	Wei *gorm.DB
}

func initMySQLDB() error {
	AppDB = &DBClient{}

	db, err := mysql.New(&AppConfig.Store.Wei)
	if err != nil {
		return err
	}

	query.SetDefault(db)

	AppDB.Wei = db

	return nil
}
