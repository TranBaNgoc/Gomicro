package model

import (
	"gomicro/server/utils"
)

type User struct {
	ID   uint64 `gorm:"AUTO_INCREMENT"`
	Name string
	Age  int64
}

func init() {
	//db, err := utils.GetConnection()
	//if err != nil {
	//	log.Println("User not created", err)
	//}
	//if !db.HasTable(&User{}) {
	//	db.CreateTable(&User{})
	//}
}

func AddUser(user *User) error {
	db, err := utils.GetConnection()
	if err != nil {
		panic(err)
	}
	db.Create(&user)
	return nil
}

func GetUsers() ([]*User, error) {
	db, err := utils.GetConnection()
	if err != nil {
		return nil, err
	}
	users, err := db.Model(&User{}).Rows()
	if err != nil {
		return nil, err
	}
	defer users.Close()

	var res []*User
	for users.Next() {
		newUser := User{}
		err = db.ScanRows(users, &newUser)
		if err == nil {
			res = append(res, &newUser)
		} else {
			return nil, err
		}
	}
	return res, nil
}
