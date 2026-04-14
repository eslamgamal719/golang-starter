package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
}

// handle request closure data
func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.Context = c
		request.DB, request.Connection = connectToDatabase()
		return request
	}
}

// init new request closure
func newRequest(c *gin.Context) Request {
	request := req()
	req := request(c)
	return req
}

// connect to database
func connectToDatabase() (*gorm.DB, *sql.DB) {
	dsn := "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database", err.Error())
	}
	connection, err := db.DB()
	if err != nil {
		fmt.Println("Error when get database connection", err.Error())
	}
	return db, connection
}

// close database connection
func (request Request) closeConnection() {
	request.Connection.Close()
}
