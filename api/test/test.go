package test

import (
	"github.com/gin-gonic/gin"
)

// @Summary 测试Ping
// @Produce json
// @Success 200 {string} string "成功"
// @Router /api/test/ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}