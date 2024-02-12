package routers

import (
	api "github.com/WuLianN/go-toy/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	baseApi := api.ApiGroupApp.BaseApi
	draftApi := api.ApiGroupApp.DraftApi
	uploadApi := api.NewUpload()
	{
		Router.POST("/login", baseApi.Login)
		Router.POST("/register", baseApi.Register)
		Router.POST("/upload/file", uploadApi.UploadFile)
		Router.GET("/visit", baseApi.Visit)
		Router.GET("/getRecommendList", baseApi.GetRecommendList)
		Router.GET("/getDraft", draftApi.GetDraft)
	}
}
