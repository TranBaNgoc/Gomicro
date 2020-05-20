package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func createConn() *sql.DB {
	var err error
	db, err = sql.Open("mysql", "root:bangocntt@49@tcp(localhost:3306)/test")
	if err == nil {
		//return db
	} else {
		fmt.Println(err)
		return nil
	}
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	return db
}
