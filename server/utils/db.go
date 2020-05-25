package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

func GetConnection() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	var err error
	db, err = gorm.Open("mysql", "root:bangocntt@49@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	db.DB().SetMaxIdleConns(250)
	db.DB().SetMaxOpenConns(500)
	return db, nil
}
