package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
	"github.com/WuLianN/go-toy/pkg/app"
)

func (d *Dao) QueryRecommendList(userId uint32, page int, pageSize int, tagIds []uint32, isSelf uint8) ([]model.RecommendList, error) {
	offset := app.GetPageOffset(page, pageSize)
	var list []model.RecommendList
	var err error

	if isSelf == 1 {
		// 获取自己文章 包括私密
		err = d.engine.Table("drafts").Order("update_time DESC").Where("drafts.user_id = ? AND is_publish = ? AND is_delete = ?", userId, 1, 0).Limit(pageSize).Offset(offset).Find(&list).Error
	} else {
		// 获取用户文章 非私密
		if len(tagIds) == 0 {
			err = d.engine.Table("drafts").Order("update_time DESC").Where("drafts.user_id = ? AND is_publish = ? AND is_delete = ? AND is_privacy = ?", userId, 1, 0, 0).Limit(pageSize).Offset(offset).Find(&list).Error
		}
	}

	if err != nil {
		return list, err
	}

	tagList, _ := d.QueryDraftTagsDT(userId, tagIds)

	// 获取指定标签的文章
	if len(tagIds) > 0 {
		var draftIds []uint32
		var draftList []model.RecommendList

		if len(tagList) == 0 {
			return draftList, err
		}

		for _, tag := range tagList {
			draftIds = append(draftIds, tag.DraftId)
		}
		var joinList []model.DraftAndTag
		if len(draftIds) > 0 {
			err = d.engine.Table("drafts").Select("drafts.id as id, drafts.title as title, drafts.content as content, drafts.bg_image as bg_image, drafts.create_time as create_time, drafts.update_time as update_time, drafts.user_id as user_id, tags.id as tag_id, tags.name as tag_name, tags.bg_color as tag_bg_color, tags.color as tag_color").Joins("left join draft_tags on drafts.id = draft_tags.draft_id").Joins("left join tags on tags.id = draft_tags.tag_id").Where("drafts.id IN (?)", draftIds).Limit(pageSize).Offset(offset).Find(&joinList).Error
		}

		if len(joinList) > 0 {
			var keyMap = make(map[uint32]bool)
			for _, joinListItem := range joinList {
				if !keyMap[joinListItem.Id] {
					keyMap[joinListItem.Id] = true

					draftList = append(draftList, model.RecommendList{
						Id:         joinListItem.Id,
						Title:      joinListItem.Title,
						Content:    joinListItem.Content,
						BgImage:    joinListItem.BgImage,
						CreateTime: joinListItem.CreateTime,
						UpdateTime: joinListItem.UpdateTime,
						UserId:     joinListItem.UserId,
						Tags: []model.Tag{
							{
								Id:      joinListItem.TagId,
								Name:    joinListItem.TagName,
								BgColor: joinListItem.TagBgColor,
								Color:   joinListItem.TagColor,
							},
						},
					})
				} else {
					for i := range draftList {
						if draftList[i].Id == joinListItem.Id {
							draftList[i].Tags = append(draftList[i].Tags, model.Tag{
								Id:      joinListItem.TagId,
								Name:    joinListItem.TagName,
								BgColor: joinListItem.TagBgColor,
								Color:   joinListItem.TagColor,
							})
						}
					}
				}
			}
		}
		return draftList, err
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
