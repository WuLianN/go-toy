package service

import (
	"time"
	"github.com/WuLianN/go-toy/internal/model"
)

func (svc *Service) Visit(ip string) {
	if ip == "" {
		return
	}
	time := time.Now().Format(time.DateTime)
	visitInfo := model.VisitInfo {
		VisitTime: time,
		IP: ip,
	}
	svc.dao.SaveVisitInfo(visitInfo)
}