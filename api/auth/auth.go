package auth

import (
	"github.com/gin-gonic/gin"
)

func CheckAuth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "有权限",
	})
}