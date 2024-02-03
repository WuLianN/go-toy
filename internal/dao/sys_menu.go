package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
	"gorm.io/gorm"
)

func (d *Dao) GetMenu(UserId uint32) []model.MenuMeat {
	var menu []model.MenuMeat

	d.engine.Table("menu").Select("menu.id as id, name, parent_id, meta_id, category_id, component, icon").Joins("left join menu_meta on menu_meta.id = menu.meta_id").Where("user_id = ? AND is_use = ?", UserId, 1).Scan(&menu)

	return menu
}

func (d *Dao) AddMenuItem(name string, parentId, categoryId, userId uint32) (model.AddMenuItem, error) {
	var menu model.Menu

	err := d.engine.Transaction(func(tx *gorm.DB) error {
		var err error
		meta := model.Meta{
			CategoryId: 0,
		}

		// 事务处理
		if err = tx.Table("menu_meta").Create(&meta).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		menu = model.Menu{
			MetaId:   meta.Id,
			Name:     name,
			ParentId: parentId,
			UserId:   userId,
			IsUse:    1,
		}

		if err = tx.Table("menu").Create(&menu).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		return nil
	})

	addMenuItem := model.AddMenuItem{
		Id:       menu.Id,
		ParentId: menu.ParentId,
	}

	if err != nil {
		return addMenuItem, err
	}
	return addMenuItem, nil
}
