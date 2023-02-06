package utils

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// "github.com/joho/godotenv"
)

func OpenDB() *gorm.DB {

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Printf("読み込み出来ませんでした: %v", err)
	// }
	// .envの SAMPLE_MESSAGEを取得して、messageに代入します。
	// dbms := os.Getenv("DBMS")
	// dbUser := os.Getenv("USER")
	// dbPass := os.Getenv("PASS")
	// protocol := os.Getenv("PROTOCOL")
	// dbName := os.Getenv("DBNAME")
	// CONNECT := dbUser + ":" + dbPass + "@" + protocol + "/" + dbName + "?parseTime=true"
	// db, err := gorm.Open(dbms, CONNECT)

	DBMS := "mysql"
	USER := "deegle"
	PASS := "deegle"
	PROTOCOL := "tcp(deegle-db:3306)"
	DBNAME := "deegle"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		fmt.Println("接続失敗")
		panic(err.Error())
	}

	return db
}
