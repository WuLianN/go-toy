package routers

import (
	"github.com/gin-gonic/gin"
	api "github.com/WuLianN/go-toy/api"
)

func InitAuthRouter(Router *gin.RouterGroup) {
	authApi := api.ApiGroupApp.AuthApi
	{
		Router.GET("/checkAuth", authApi.CheckAuth)
	}
}
