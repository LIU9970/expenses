package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sample struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUserInfo() gin.HandlerFunc {
	fmt.Println(11)
	return func(c *gin.Context) {
		result := sample{1, "sample"}
		c.JSON(http.StatusOK, result)
	}
}
