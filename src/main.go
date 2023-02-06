package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tmsic/deegle-translate-api/router"
)

func main() {

	// ルータ、ミドルウェアを設定
	r := router.GetRouter()
		r.Run()
}
