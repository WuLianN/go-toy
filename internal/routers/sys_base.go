package routers

import (
	"github.com/gin-gonic/gin"
	api "github.com/WuLianN/go-toy/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	baseApi := api.ApiGroupApp.BaseApi
	uploadApi := api.NewUpload()
	{
		Router.GET("/ping", baseApi.Ping)
		Router.POST("/login", baseApi.Login)
		Router.POST("/register", baseApi.Register)
		Router.POST("/upload/file", uploadApi.UploadFile)
	}
}
