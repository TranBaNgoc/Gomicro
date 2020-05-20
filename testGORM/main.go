package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
}

func main() {
	db, err := gorm.Open("mysql", "root:bangocntt@49@/test?charset=utf8&parseTime=True&loc=Local")
	db.DB()
	if err != nil {
		panic("failed to connect database") // Kiểm tra kết nối tới database
	}
	defer db.Close()

	//db.DropTableIfExists(&UserAccount{})
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	user := User{Username: "ngoctb3"}
	db.Create(&user)

	rows, _ := db.Where("username = ?", "ngoctb3").Find(&User{}).Rows()
	for rows.Next() {
		ua := User{}
		err = rows.Scan(&ua.ID, &ua.Username, &ua.CreatedAt, &ua.DeletedAt, &ua.UpdatedAt)
		if err == nil {
			fmt.Println(ua)
		} else {
			fmt.Println(err)
		}
	}
}
