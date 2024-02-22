package service

import (
	"fmt"

	"github.com/WuLianN/go-toy/global"
)

func (svc *Service) SystemCron(spec string) {
	enterId, _ := global.Cron.AddFunc(spec, func() {
		// TODO 配合消息推送
	})

	fmt.Println("enterId------", enterId)
}

func (svc *Service) CronStart() {
	visitSpec := "0 0 */1 * *"
	svc.SystemCron(visitSpec)

	global.Cron.Start()
}

// # ┌───────────── minute (0–59)
// # │ ┌───────────── hour (0–23)
// # │ │ ┌───────────── day of the month (1–31)
// # │ │ │ ┌───────────── month (1–12)
// # │ │ │ │ ┌───────────── day of the week (0–6) (Sunday to Saturday;
// # │ │ │ │ │                                   7 is also Sunday on some systems)
// # │ │ │ │ │
// # │ │ │ │ │
// # * * * * * <command to execute>
