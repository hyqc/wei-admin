package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Config struct {
	User                   string `json:"user"`
	Password               string `json:"pwd"`
	Host                   string `json:"host"`
	Port                   int    `json:"port"`
	Dbname                 string `json:"dbname"`
	Charset                string `json:"charset"`
	ParseTime              bool   `json:"parse_time"`
	Location               string `json:"location"`
	MaxIdleCons            int    `json:"max_idle_cons"`
	MaxOpenCons            int    `json:"max_open_cons"`
	ConnMaxLifetimeMinutes int    `json:"conn_max_lifetime_minutes"`
}

func New(conf *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Dbname, conf.Charset, conf.ParseTime, conf.Location)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.Debug().DB()
	if err != nil {
		return nil, err
	}
	if conf.MaxIdleCons == 0 {
		conf.MaxOpenCons = 1
	}
	if conf.MaxOpenCons == 0 {
		conf.MaxOpenCons = 10
	}
	if conf.ConnMaxLifetimeMinutes == 0 {
		conf.ConnMaxLifetimeMinutes = 60
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(conf.MaxOpenCons)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(conf.MaxOpenCons)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifetimeMinutes) * time.Minute)

	return db, nil
}
