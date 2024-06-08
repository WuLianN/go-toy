package routers

import (
	api "github.com/WuLianN/go-toy/api"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) {
	menuApi := api.ApiGroupApp.MenuApi
	{
		Router.GET("/getMenuList", menuApi.GetRoleMenu)
		Router.POST("/addMenuItem", menuApi.AddMenuItem)
		Router.POST("/deleteMenuItem", menuApi.DeleteMenuItem)
		Router.POST("/updateMenuItem", menuApi.UpdateMenuItem)
		Router.POST("/saveMenuSort", menuApi.SaveMenuSort)
	}
}
