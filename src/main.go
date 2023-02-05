package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	r.GET("/test", func(c *gin.Context) {

		var msg struct {
			Message string `json:"message"`
		}
		msg.Message = "Goからのメッセージ"
		c.JSON(http.StatusOK, msg)
	})
	r.Run()
}
