package routers

import (
	"github.com/WuLianN/go-toy/api"
	"github.com/gin-gonic/gin"
)

func InitTagRouter(Router *gin.RouterGroup) {
	tagApi := api.ApiGroupApp.TagApi
	{
		Router.GET("/getTagList", tagApi.GetTagList)
		Router.GET("/getDraftTagList", tagApi.GetDraftTagList)
		Router.GET("/getMenuTagList", tagApi.GetMenuTagList)
		Router.POST("/createTag", tagApi.CreateTag)
		Router.POST("/deleteTag", tagApi.DeleteTag)
		Router.POST("/updateTag", tagApi.UpdateTag)
		Router.POST("/bindTag2Menu", tagApi.BindTag2Menu)
		Router.POST("/unbindTag2Menu", tagApi.UnbindTag2Menu)
		Router.POST("/bindTag2Draft", tagApi.BindTag2Draft)
		Router.POST("/unbindTag2Draft", tagApi.UnbindTag2Draft)
	}
}

func InitBaseTagRouter(Router *gin.RouterGroup) {
	tagApi := api.ApiGroupApp.TagApi
	{
		Router.POST("/searchTags", tagApi.SearchTags)
	}
}
