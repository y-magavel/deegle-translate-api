package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tmsic/deegle-translate-api/handler"
	"github.com/tmsic/deegle-translate-api/middleware"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	middleware.CORS(r)

	r.GET("/test", handler.AllLikes())
	//Google翻訳・DeepL APIハンドラ
	r.GET("/google_translate", handler.GoogleTranslate())
	r.GET("/deepL_translate", handler.DeepLTranslate())

	return r
}
