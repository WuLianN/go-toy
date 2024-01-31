package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
	"github.com/WuLianN/go-toy/pkg/app"
)

func (d *Dao) SaveVisitInfo(visitInfo model.VisitInfo) {
	info := model.VisitInfo{
		VisitTime: visitInfo.VisitTime,
		IP:        visitInfo.IP,
	}
	d.engine.Table("statistics_visit").Create(&info)
}

func (d *Dao) QueryRecommendList(userId uint32, page int, pageSize int) ([]model.RecommendList, error) {
	offset := app.GetPageOffset(page, pageSize)
	var list []model.RecommendList
	err := d.engine.Table("draft").Where("user_id = ? AND is_publish = ? AND is_delete = ?", userId, 1, 0).Limit(pageSize).Offset(offset).Find(&list).Error

	if err != nil {
		return list, err
	}

	return list, nil
}
