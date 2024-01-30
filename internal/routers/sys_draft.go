package routers

import (
	api "github.com/WuLianN/go-toy/api"
	"github.com/gin-gonic/gin"
)

func InitDraftRouter(Router *gin.RouterGroup) {
	draftApi := api.ApiGroupApp.DraftApi
	{
		Router.GET("/getDraft", draftApi.GetDraft)
		Router.GET("/createDraft", draftApi.CreateDraft)
		Router.POST("/saveDraft", draftApi.SaveDraft)
		Router.POST("/deleteDraft", draftApi.DeleteDraft)
		Router.POST("/publishDarft", draftApi.PublishDarft)
		Router.GET("/getDraftList", draftApi.GetDraftList)
	}
}
