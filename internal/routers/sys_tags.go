package routers

import (
	"github.com/WuLianN/go-toy/api"
	"github.com/gin-gonic/gin"
)

func InitTagRouter(Router *gin.RouterGroup) {
	tagApi := api.ApiGroupApp.TagApi
	{
		Router.GET("/getTagList", tagApi.GetTagList)
		Router.POST("/createTag", tagApi.CreateTag)
		Router.POST("/deleteTag", tagApi.DeleteTag)
	}
}