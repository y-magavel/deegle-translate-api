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

	return r
}
