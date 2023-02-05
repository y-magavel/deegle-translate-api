package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//db.ConnectDB()
	// ルータ、ミドルウェアを設定
	// r := router.GetRouter()
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "届いてるよ！"})
	})

	r.Run()
}
