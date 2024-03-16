package routers

import (
	api "github.com/WuLianN/go-toy/api"
	"github.com/gin-gonic/gin"
)

func InitUploadRouter(Router *gin.RouterGroup) {
	uploadApi := api.ApiGroupApp.UploadApi
	{
		Router.GET("/getUploadFileList", uploadApi.GetUploadFileList)
		Router.POST("/deleteUploadFiles", uploadApi.DeleteUploadFile)
	}
}
