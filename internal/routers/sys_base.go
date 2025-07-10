package routers

import (
	api "github.com/WuLianN/go-toy/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	baseApi := api.ApiGroupApp.BaseApi
	draftApi := api.ApiGroupApp.DraftApi
	userApi := api.ApiGroupApp.UserApi
	uploadApi := api.ApiGroupApp.UploadApi
	{
		Router.POST("/login", baseApi.Login)
		Router.POST("/register", baseApi.Register)
		Router.POST("/upload/file", uploadApi.UploadFile)
		Router.GET("/getRecommendList", baseApi.GetRecommendList)
		Router.GET("/getDraft", draftApi.GetDraft)
		Router.GET("/getUserInfo", userApi.GetUserInfo)
		Router.GET("/searchDrafts", draftApi.SearchDrafts)
		Router.GET("/getUserSetting", userApi.GetUserSetting)
	}
}
