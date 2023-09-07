package system

import (
	"github.com/gin-gonic/gin"
	api "github.com/WuLianN/go-toy/api"
)

type AuthRouter struct {}

func (a *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	authRouter := api.ApiGroupApp.SystemApiGroup.AuthApi
	{
		Router.GET("/checkAuth", authRouter.CheckAuth)
	}
}
