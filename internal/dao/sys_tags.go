package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
	"gorm.io/gorm"
)

func (d *Dao) QueryTagList(userId uint32, idList []int, menuId uint32) ([]model.Tag, error) {
	var list []model.Tag
	var err error

	if menuId == 0 {
		// 查询用户标签
		if len(idList) == 0 {
			// 所有标签
			err = d.engine.Table("tags").Where("user_id = ?", userId).Find(&list).Error
		} else {
			// 指定标签
			err = d.engine.Table("tags").Where("user_id = ?", userId).Where("id IN ?", idList).Find(&list).Error
		}
	} else {
		// 查询菜单关联的标签
		err = d.engine.Table("menu_tags").Select("tags.name as name, tags.id as id, tags.user_id").Where("menu_id = ?", menuId).Joins("left join tags on menu_tags.tag_id = tags.id").Find(&list).Error
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

func (d *Dao) UpdateTag(tagId uint32, name string) error {
	return d.engine.Table("tags").Where("id = ?", tagId).Update("name", name).Error
}

func (d *Dao) BindTag2Menu(tags []model.Tag, menuId uint32, userId uint32) error {
	err := d.engine.Transaction(func(tx *gorm.DB) error {
		var err error
		if err = tx.Table("tags").Create(&tags).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		for _, tag := range tags {
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
