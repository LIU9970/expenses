package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(res string) gin.HandlerFunc {

	return func(c *gin.Context) {
		result := res
		c.JSON(http.StatusOK, result)
	}
}
