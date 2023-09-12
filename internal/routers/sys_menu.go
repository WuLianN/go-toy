package routers

import (
	"github.com/gin-gonic/gin"
	api "github.com/WuLianN/go-toy/api"
)

func InitMenuRouter(Router *gin.RouterGroup) {
	menuApi := api.ApiGroupApp.MenuApi
	{
		Router.GET("/getRoleMenu", menuApi.GetRoleMenu)
	}
}
