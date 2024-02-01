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
	err := d.engine.Table("drafts").Where("drafts.user_id = ? AND is_publish = ? AND is_delete = ?", userId, 1, 0).Limit(pageSize).Offset(offset).Find(&list).Error

	if err != nil {
		return list, err
	}

	tagList, _ := d.QueryDraftTags(userId, tagId)

	if tagId > 0 && len(tagList) > 0 {
		// 获取指定标签的文章
		var tempList []model.RecommendList
		for _, tag := range tagList {
			for i := range list {
				if list[i].Id == tag.DraftId {
					list[i].Tags = append(list[i].Tags, model.Tag{
						Id:   tag.Id,
						Name: tag.Name,
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
						Id:   tag.Id,
						Name: tag.Name,
					})
				}
			}
		}
	}

	return list, nil
}

func (d *Dao) QueryDraftTags(userId uint32, tagId uint32) ([]model.DraftTag, error) {
	var list []model.DraftTag
	var err error

	if tagId > 0 {
		err = d.engine.Table("draft_tags").Select("tags.id as id, draft_tags.draft_id, tags.name").Joins("left join tags on draft_tags.tag_id = tags.id").Where("draft_tags.user_id = ? AND draft_tags.tag_id = ?", userId, tagId).Find(&list).Error
	} else {
		err = d.engine.Table("draft_tags").Select("tags.id as id, draft_tags.draft_id, tags.name").Joins("left join tags on draft_tags.tag_id = tags.id").Where("draft_tags.user_id = ?", userId).Find(&list).Error
	}

	if err != nil {
		return list, err
	}

	return list, nil
}
