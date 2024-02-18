package routers

import (
	api "github.com/WuLianN/go-toy/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userApi := api.ApiGroupApp.UserApi
	{
		Router.POST("/changePassword", userApi.ChangePassword)
		Router.POST("/updateUserInfo", userApi.UpdateUserInfo)
		Router.POST("/bingUser", userApi.BingUser)
		Router.POST("/unbindUser", userApi.UnbindUser)
		Router.GET("/getBindedUserList", userApi.GetBindedUserList)
	}
}
