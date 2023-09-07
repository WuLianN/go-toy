package system

import (
	"github.com/gin-gonic/gin"
	api "github.com/WuLianN/go-toy/api"
)

type BaseRouter struct {}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := api.ApiGroupApp.SystemApiGroup.BaseApi
	{
		Router.GET("/ping", baseRouter.Ping)
		Router.GET("/login", baseRouter.Login)
	}
}
