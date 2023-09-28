package api

import (
	"github.com/gin-gonic/gin"
	"github.com/WuLianN/go-toy/internal/service"
	"github.com/WuLianN/go-toy/pkg/app"
	"github.com/WuLianN/go-toy/pkg/errcode"
	"time"
	"strconv"
)

type StatisticsApi struct {}

func (s *StatisticsApi) GetVisitStatistics(c *gin.Context) {
	year := strconv.Itoa(time.Now().Year())
	if s, exist := c.GetQuery("year"); exist {
		year = s
	}
	svc := service.New(c.Request.Context())
	dateList, valueList := svc.GetVisitStatistics(year)

	response := app.NewResponse(c)

	response.ToResponse(gin.H{
		"code": errcode.Success.Code(),
		"message": errcode.Success.Msg(),
		"type": "success",
		"result": gin.H {
			"date": dateList,
			"data": valueList,
		},
	})

}