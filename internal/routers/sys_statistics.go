package routers

import (
	"github.com/gin-gonic/gin"
	api "github.com/WuLianN/go-toy/api"
)

func InitStatisticsRouter(Router *gin.RouterGroup) {
	statisticsApi := api.ApiGroupApp.StatisticsApi
	{
		Router.GET("/getVisit", statisticsApi.GetVisitStatistics)
	}
}