package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := app()
	application := app()
	application.Gin.GET("/ping", func(c *gin.Context) {
		request := newRequest(c)

		request.NotAuth()
		//request.Context.JSON(http.StatusOK, gin.H{
		//	"message": "pong",
		//})
	})

	application.Gin.Run(":8081")
}
