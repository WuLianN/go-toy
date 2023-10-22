package service

import (
	"github.com/WuLianN/go-toy/global"
	"github.com/WuLianN/go-toy/internal/model"
	"time"
	"fmt"
)

func (svc *Service) VisitCron(spec string) {
	enterId, _ := global.Cron.AddFunc(spec, func() {
		svc.GetVisitCount()
		// TODO 配合消息推送
	})

	fmt.Println("enterId------", enterId)
}

func (svc *Service) GetVisitCount() model.Result {
	now := time.Now()
	year := now.Year()
	month := int(now.Month())

	return svc.dao.VisitStatisticsByOneMonth(year, month)
}

func (svc *Service) CronStart() {
	visitSpec := "*/1 * * * *"
	svc.VisitCron(visitSpec)

	global.Cron.Start()
}