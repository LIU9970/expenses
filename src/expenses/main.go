package main

import (
	"expenses/handler/user"
	sqlCon "expenses/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	sqlCon.DBOpen()

	router.GET("/user", user.GetUserInfo("success"))
	router.Run(":3000")
}
