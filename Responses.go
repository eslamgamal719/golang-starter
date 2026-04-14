package main

import "github.com/gin-gonic/gin"

func (request Request) Ok(body interface{}) {
	request.Response(200, body)
}

func (request Request) Created(body interface{}) {
	request.Response(201, body)
}

func (request Request) NotAuth() {
	request.Response(401, gin.H{
		"message": "Not Authenticated",
	})
}
