package main

import (
	// "expenses/handler/user"

	"expenses/handler/user"

	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	r := gin.Default()
	r.GET("/user", user.GetUserInfo())
	r.Run(":3000")
	return r
}
