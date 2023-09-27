package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
)

func (d *Dao) SaveVisitInfo(visitInfo model.VisitInfo) {
	info := model.VisitInfo{
		VisitTime: visitInfo.VisitTime,
		IP: visitInfo.IP,
	}
	d.engine.Table("statistics_visit").Create(&info)
}