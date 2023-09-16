package routers

import (
	"github.com/gin-gonic/gin"
	api "github.com/WuLianN/go-toy/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userApi := api.ApiGroupApp.UserApi
	{
		Router.GET("/getUserInfo", userApi.GetUserInfo)
		Router.POST("/changePassword", userApi.ChangePassword)
	}
}
