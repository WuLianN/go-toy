package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
	"github.com/WuLianN/go-toy/pkg/app"
)

func (d *Dao) QueryRecommendList(userId uint32, page int, pageSize int, tagId uint32, isSelf uint8) ([]model.RecommendList, error) {
	offset := app.GetPageOffset(page, pageSize)
	var list []model.RecommendList
	var err error

	if isSelf == 1 {
		// 获取自己文章 包括私密
		err = d.engine.Table("drafts").Order("update_time DESC").Where("drafts.user_id = ? AND is_publish = ? AND is_delete = ?", userId, 1, 0).Limit(pageSize).Offset(offset).Find(&list).Error
	} else {
		// 获取用户文章 非私密
		err = d.engine.Table("drafts").Order("update_time DESC").Where("drafts.user_id = ? AND is_publish = ? AND is_delete = ? AND is_privacy = ?", userId, 1, 0, 0).Limit(pageSize).Offset(offset).Find(&list).Error
	}

	if err != nil {
		return list, err
	}

	tagList, _ := d.QueryDraftTagsDT(userId, tagId)

	// 获取指定标签的文章
	if tagId > 0 {
		var tempList []model.RecommendList
		if len(tagList) == 0 {
			return tempList, err
		}

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

		return tempList, err
	}

	// 无指定tag_id 获取全部文章
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

	return list, nil
}
