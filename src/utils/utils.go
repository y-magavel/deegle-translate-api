package utils

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func OpenDB() *gorm.DB {
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

func LoadEnv() {
	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// .envの SAMPLE_MESSAGEを取得して、messageに代入します。
	message := os.Getenv("SAMPLE_MESSAGE")

	fmt.Println(message)
}
