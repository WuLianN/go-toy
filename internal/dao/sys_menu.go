package dao

import (
	"strconv"

	"github.com/WuLianN/go-toy/internal/model"
	"gorm.io/gorm"
)

func (d *Dao) GetMenu(UserId uint32) []model.MenuMeat {
	var menu []model.MenuMeat

	d.engine.Table("menu").Select("menu.id as id, title, hide_children_in_menu, name, parent_id, meta_id, hide_menu, category, component, icon, path, redirect, tag_id").Joins("left join menu_meta on menu_meta.id = menu.meta_id").Where("user_id = ? AND is_use = ?", UserId, 1).Scan(&menu)

	return menu
}

func (d *Dao) AddMenuItem(menu *model.Menu, userId uint32) (uint32, error) {
	err := d.engine.Transaction(func(tx *gorm.DB) error {
		var err error
		meta := model.Meta{}

		// 事务处理
		if err = tx.Table("menu_meta").Create(&meta).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		menu := model.Menu{
			MetaId:   meta.Id,
			Name:     menu.Name,
			ParentId: menu.ParentId,
			UserId:   userId,
			Category: strconv.Itoa(int(userId)) + "_" + strconv.Itoa(int(meta.Id)),
			IsUse:    1,
		}

		if err = tx.Table("menu").Create(&menu).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		return nil
	})

	if err != nil {
		return menu.Id, err
	}
	return menu.Id, nil
}
