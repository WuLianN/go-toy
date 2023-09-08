package routers

import (
	"github.com/gin-gonic/gin"
	api "github.com/WuLianN/go-toy/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	baseApi := api.ApiGroupApp.BaseApi
	{
		Router.GET("/ping", baseApi.Ping)
		Router.GET("/login", baseApi.Login)
	}
}
