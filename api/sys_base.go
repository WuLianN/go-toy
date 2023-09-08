package api

import (
	"github.com/gin-gonic/gin"
)

type BaseApi struct {}

// @Summary 测试Ping
// @Produce json
// @Success 200 {string} string "成功"
// @Router /api/test/ping [get]
func (b *BaseApi) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (b *BaseApi) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}