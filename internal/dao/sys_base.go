package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
	"github.com/WuLianN/go-toy/pkg/app"

	"strings"
	"time"

	"github.com/spf13/cast" // 用于类型转换
)

func (d *Dao) QueryRecommendList(userId uint32, page int, pageSize int, tagIds []uint32, isSelf uint8) ([]model.RecommendList, error) {
	offset := app.GetPageOffset(page, pageSize)
	var list []model.RecommendList
	var err error

	if isSelf == 1 {
		// 获取自己文章 包括私密
		err = d.engine.Table("drafts").
			Order("update_time DESC, id DESC").
			Where("drafts.user_id = ? AND is_publish = ? AND is_delete = ?", userId, 1, 0).
			Limit(pageSize).Offset(offset).Find(&list).Error
	} else {
		// 获取用户文章 非私密
		if len(tagIds) == 0 {
			err = d.engine.Table("drafts").
				Order("update_time DESC, id DESC").
				Where("drafts.user_id = ? AND is_publish = ? AND is_delete = ? AND is_privacy = ?", userId, 1, 0, 0).
				Limit(pageSize).Offset(offset).Find(&list).Error
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
			return draftList, nil
		}

		for _, tag := range tagList {
			draftIds = append(draftIds, tag.DraftId)
		}

		type DraftWithTags struct {
			Id         uint32    `gorm:"column:id"`
			Title      string    `gorm:"column:title"`
			Content    string    `gorm:"column:content"`
			BgImage    string    `gorm:"column:bg_image"`
			CreateTime time.Time `gorm:"column:create_time"`
			UpdateTime time.Time `gorm:"column:update_time"`
			UserId     uint32    `gorm:"column:user_id"`

			TagIds      string `gorm:"column:tag_ids"`
			TagNames    string `gorm:"column:tag_names"`
			TagBgColors string `gorm:"column:tag_bg_colors"`
			TagColors   string `gorm:"column:tag_colors"`
		}

		var drafts []DraftWithTags

		baseQuery := d.engine.Table("drafts").
			Select(`drafts.id as id,
				drafts.title as title,
				drafts.content as content,
				drafts.bg_image as bg_image,
				drafts.create_time as create_time,
				drafts.update_time as update_time,
				drafts.user_id as user_id,
				GROUP_CONCAT(tags.id) as tag_ids,
				GROUP_CONCAT(tags.name) as tag_names,
				GROUP_CONCAT(tags.bg_color) as tag_bg_colors,
				GROUP_CONCAT(tags.color) as tag_colors`).
			Joins("LEFT JOIN draft_tags ON drafts.id = draft_tags.draft_id").
			Joins("LEFT JOIN tags ON tags.id = draft_tags.tag_id").
			Where("drafts.id IN (?)", draftIds).
			Group("drafts.id").
			Order("drafts.update_time DESC, drafts.id DESC").
			Limit(pageSize).Offset(offset)

		// ✅ 添加统一的 is_publish 和 is_delete 条件
		baseQuery = baseQuery.Where("is_publish = ? AND is_delete = ?", 1, 0)

		// 如果是非本人查看，再加上 is_privacy = 0
		if isSelf == 0 {
			baseQuery = baseQuery.Where("is_privacy = ?", 0)
		}

		err = baseQuery.Find(&drafts).Error

		if err != nil {
			return nil, err
		}

		// 转换为 RecommendList 结构
		for _, item := range drafts {
			var tags []model.Tag

			ids := strings.Split(item.TagIds, ",")
			names := strings.Split(item.TagNames, ",")
			bgColors := strings.Split(item.TagBgColors, ",")
			colors := strings.Split(item.TagColors, ",")

			for i := 0; i < len(names); i++ {
				if i >= len(ids) {
					break
				}
				tags = append(tags, model.Tag{
					Id:      cast.ToUint32(ids[i]),
					Name:    names[i],
					BgColor: bgColors[i],
					Color:   colors[i],
				})
			}

			const layout = "2006-01-02 15:04:05"

			draftList = append(draftList, model.RecommendList{
				Id:         item.Id,
				Title:      item.Title,
				Content:    item.Content,
				BgImage:    item.BgImage,
				CreateTime: item.CreateTime.Format(layout),
				UpdateTime: item.UpdateTime.Format(layout),
				UserId:     item.UserId,
				Tags:       tags,
			})
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
