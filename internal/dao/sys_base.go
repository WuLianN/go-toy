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

func (d *Dao) QueryRecommendList(userId uint32, page int, pageSize int, tagId uint32) ([]model.RecommendList, error) {
	offset := app.GetPageOffset(page, pageSize)
	var list []model.RecommendList
	err := d.engine.Table("drafts").Order("update_time DESC").Where("drafts.user_id = ? AND is_publish = ? AND is_delete = ?", userId, 1, 0).Limit(pageSize).Offset(offset).Find(&list).Error

	if err != nil {
		return list, err
	}

	tagList, _ := d.QueryDraftTagsDT(userId, tagId)

	if tagId > 0 && len(tagList) > 0 {
		// 获取指定标签的文章
		var tempList []model.RecommendList
		for _, tag := range tagList {
			for i := range list {
				if list[i].Id == tag.DraftId {
					list[i].Tags = append(list[i].Tags, model.Tag{
						Id:      tag.TagId,
						Name:    tag.Name,
						BgColor: tag.BgColor,
						Color:   tag.Color,
					})
					tempList = append(tempList, list[i])
				}
			}
		}
		list = tempList
	} else {
		// 获取全部文章
		for i := range list {
			for _, tag := range tagList {
				if tag.DraftId == list[i].Id {
					list[i].Tags = append(list[i].Tags, model.Tag{
						Id:      tag.TagId,
						Name:    tag.Name,
						BgColor: tag.BgColor,
						Color:   tag.Color,
					})
				}
			}
		}
	}

	return list, nil
}
