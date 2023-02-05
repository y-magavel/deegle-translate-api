package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmsic/deegle-translate-api/db"
)

func AllGoods() gin.HandlerFunc {
	return func(c *gin.Context) {

		result := db.DBConnect()
		fmt.Println(result)

		c.JSON(http.StatusOK, gin.H{
			"likeCount": result,
		})
	}
}
