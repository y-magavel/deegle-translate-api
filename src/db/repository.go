package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	User_id   int       `json:"user_id"`
	User_name string    `json:"user_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	DeleteFlg string    `json:"delete_flg"`
}

func DBConnect() int {

	DBMS := "mysql"
	USER := "deegle"
	PASS := "deegle"
	//tcp(コンテナ名:port)
	PROTOCOL := "tcp(deegle-db:3306)"
	DBNAME := "deegle"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		fmt.Println("接続失敗")
		panic(err.Error())
	}
	defer db.Close()

	// user := User{
	// 	User_name: "test",
	// 	Email:     "test@example.com",
	// 	CreatedAt: time.Now(),
	// 	UpdateAt:  time.Now(),
	// 	DeleteFlg: "1",
	// }

	// if err := db.Select("user_name", "email", "created_at", "update_at").Create(&user).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// if err := db.Where("user_id = ?", 2).First(&user).Error; err != nil {
	// 	fmt.Println(err)
	// }
	// var users []User
	// if err := db.First(&users).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// if err := db.Find(&users).Error; err != nil {
	// 	fmt.Println(err)
	// }

	var count int
	if err := db.Model(User{}).Count(&count).Error; err != nil {
		fmt.Println(err)
	}

	fmt.Println(count)

	return count
}

func AllGoods() {
}
