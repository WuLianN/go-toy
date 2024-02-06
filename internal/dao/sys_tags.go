package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
	"gorm.io/gorm"
)

func (d *Dao) QueryTagList(userId uint32, idList []int) ([]model.Tag, error) {
	var list []model.Tag
	var err error

	// 查询用户标签
	if len(idList) == 0 {
		// 所有标签
		err = d.engine.Table("tags").Where("user_id = ?", userId).Find(&list).Error
	} else {
		// 指定标签
		err = d.engine.Table("tags").Where("user_id = ?", userId).Where("id IN ?", idList).Find(&list).Error
	}

	if err != nil {
		return list, err
	}

	return list, nil
}

func (d *Dao) QueryMenuTags(menuId uint32) ([]model.Tag, error) {
	var list []model.Tag

	err := d.engine.Table("menu_tags").Select("tags.name as name, tags.id as id, tags.user_id").Where("menu_id = ?", menuId).Joins("left join tags on menu_tags.tag_id = tags.id").Find(&list).Error

	if err != nil {
		return list, err
	}

	return list, nil
}

func (d *Dao) QueryDraftTagsDT(userId uint32, tagId uint32) ([]model.DraftTag, error) {
	var list []model.DraftTag
	var err error

	if tagId > 0 {
		err = d.engine.Table("draft_tags").Select("tags.id as id, draft_tags.draft_id, tags.name").Joins("left join tags on draft_tags.tag_id = tags.id").Where("tags.user_id = ? AND draft_tags.tag_id = ?", userId, tagId).Find(&list).Error
	}

	if err != nil {
		return list, err
	}

	return list, nil
}

func (d *Dao) QueryDraftTagsT(userId uint32, tagId uint32, draftId uint32) ([]model.Tag, error) {
	var list []model.Tag
	var err error

	if tagId > 0 {
		err = d.engine.Table("draft_tags").Select("tags.id as id, draft_tags.draft_id, tags.name").Joins("left join tags on draft_tags.tag_id = tags.id").Where("tags.user_id = ? AND draft_tags.tag_id = ?", userId, tagId).Find(&list).Error
	}

	if draftId > 0 {
		err = d.engine.Table("draft_tags").Select("tags.id as id, draft_tags.draft_id, tags.name").Where("draft_id = ?", draftId).Joins("left join tags on draft_tags.tag_id = tags.id").Find(&list).Error
	}

	if err != nil {
		return list, err
	}

	return list, nil
}

func (d *Dao) CreateTag(userId uint32, name string) (uint32, error) {
	tag := &model.Tag{
		UserId: userId,
		Name:   name,
	}

	err := d.engine.Table("tags").Create(&tag).Error

	if err != nil {
		return 0, err
	}

	return tag.Id, nil
}

func (d *Dao) DeleteTag(tagId uint32) error {
	return d.engine.Table("tags").Delete(&model.Tag{}, "id = ?", tagId).Error
}

func (d *Dao) QueryTag(userId uint32, name string) ([]model.Tag, error) {
	var list []model.Tag
	err := d.engine.Table("tags").Where("user_id = ? AND name = ?", userId, name).First(&list).Error

	if err != nil {
		return list, err
	}
	return list, nil
}

func (d *Dao) QueryTags(userId uint32, names []string) ([]model.Tag, error) {
	var list []model.Tag
	err := d.engine.Table("tags").Where("user_id = ? AND name in ?", userId, names).Find(&list).Error

	if err != nil {
		return list, err
	}
	return list, nil
}

func (d *Dao) UpdateTag(tagId uint32, name string) error {
	return d.engine.Table("tags").Where("id = ?", tagId).Update("name", name).Error
}

func (d *Dao) BindTag2Menu(exsitTags []model.Tag, newTags []model.Tag, menuId uint32, userId uint32) error {
	err := d.engine.Transaction(func(tx *gorm.DB) error {
		var err error

		// 创建新标签
		if len(newTags) > 0 {
			if err = tx.Table("tags").Create(&newTags).Error; err != nil {
				// 返回任何错误都会回滚事务
				return err
			}
		}

		var bindTags []model.Tag
		bindTags = append(bindTags, exsitTags...)
		bindTags = append(bindTags, newTags...)

		for _, tag := range bindTags {
			if err = tx.Table("menu_tags").Create(&model.MenuTag{
				TagId:  tag.Id,
				MenuId: menuId,
			}).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) UnbindTag2Menu(tags []uint32, menuId uint32) error {
	err := d.engine.Table("menu_tags").Where("menu_id = ? AND tag_id in ?", menuId, tags).Delete(model.MenuTag{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) BindTag2Draft(exsitTags []model.Tag, newTags []model.Tag, draftId uint32, userId uint32) error {
	err := d.engine.Transaction(func(tx *gorm.DB) error {
		var err error

		// 创建新标签
		if len(newTags) > 0 {
			if err = tx.Table("tags").Create(&newTags).Error; err != nil {
				// 返回任何错误都会回滚事务
				return err
			}
		}

		var bindTags []model.Tag
		bindTags = append(bindTags, exsitTags...)
		bindTags = append(bindTags, newTags...)

		type temp struct {
			TagId   uint32 `json:"tag_id"`
			DraftId uint32 `json:"draft_id"`
		}

		for _, tag := range bindTags {
			draftTags := temp{
				TagId:   tag.Id,
				DraftId: draftId,
			}

			if err = tx.Table("draft_tags").Create(&draftTags).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) UnbindTag2Draft(tags []uint32, draftId uint32) error {
	err := d.engine.Table("draft_tags").Where("draft_id = ? AND tag_id in ?", draftId, tags).Delete(model.DraftTag{}).Error

	if err != nil {
		return err
	}

	return nil
}
