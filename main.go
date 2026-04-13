package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := app()
	application := app()
	application.Gin.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	application.Gin.Run(":8081")
}
