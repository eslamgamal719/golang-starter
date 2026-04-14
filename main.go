package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := app()
	application := app()
	application.Gin.GET("/ping", func(c *gin.Context) {
		request := newRequest(c)
		fmt.Println(request.Connection.Ping())
		request.closeConnection()
		fmt.Println(request.Connection.Ping())
		request.Context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	application.Gin.Run(":8081")
}
