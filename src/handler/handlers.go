package handler

import (
	"github.com/gin-gonic/gin"
)

func AllGoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var msg struct {
		// 	Message string `json:"message"`
		// }
		// result := db.DBConnect()

		c.JSON(200, gin.H{"message": "届いてるよ！"})

		// c.JSON(http.StatusOK, result)
	}
}
