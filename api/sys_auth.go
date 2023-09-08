package api

import (
	"github.com/gin-gonic/gin"
)

type AuthApi struct {}

func (a *AuthApi) CheckAuth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "有权限",
	})
}