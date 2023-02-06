package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmsic/deegle-translate-api/db/dbrepo"
)

// 全てのいいね数を取得
func AllLikes() gin.HandlerFunc {
	return func(c *gin.Context) {
		//実際の処理
		result := dbrepo.AllLikes()
		//json形式で画面に値返却
		c.JSON(http.StatusOK, gin.H{
			"likeCount": result,
		})
	}
}
