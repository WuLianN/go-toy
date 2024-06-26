package dao

import (
	"fmt"

	"github.com/WuLianN/go-toy/internal/model"
	"gorm.io/gorm"
)

func (d *Dao) GetMenu(UserId uint32, IsUse uint8, isSelf uint8) []model.MenuMeta {
	var menu []model.MenuMeta

	query := d.engine.Table("menu").
		Select("menu.id as id, name, parent_id, meta_id, component, icon, is_use, is_privacy").
		Joins("LEFT JOIN menu_meta ON menu_meta.id = menu.meta_id").
		Where("user_id = ?", UserId).
		Order("menu.sort ASC")

	// 动态添加条件
	if IsUse == 1 {
		query = query.Where("is_use = ?", IsUse)
	}
	if isSelf == 0 {
		query = query.Where("is_privacy = ?", 0)
	}

	query.Scan(&menu)

	return menu
}

func (d *Dao) AddMenuItem(name string, parentId, userId uint32) (model.AddMenuItem, error) {
	var menu model.Menu

	err := d.engine.Transaction(func(tx *gorm.DB) error {
		var err error
		meta := model.Meta{}

		// 事务处理
		if err = tx.Table("menu_meta").Create(&meta).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		menu = model.Menu{
			MetaId:    meta.Id,
			Name:      name,
			ParentId:  parentId,
			UserId:    userId,
			IsUse:     1,
			Sort:      1,
			IsPrivacy: 0,
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

func (d *Dao) DeleteMenuItem(menuId uint32, userId uint32) error {
	menu, err := d.QueryMenuById(menuId, userId)

	if err != nil {
		return err
	}

	menuFamily := d.GetMenuIdFamily(menuId, userId)
	menuFamily = append(menuFamily, menu) // 添加自身

	err = d.engine.Transaction(func(tx *gorm.DB) error {
		var err error
		var menuIds []uint32
		var metaIds []uint32

		for _, v := range menuFamily {
			menuIds = append(menuIds, v.Id)
			metaIds = append(metaIds, v.MetaId)
		}

		if err = tx.Table("menu").Where("id IN ? AND user_id = ?", menuIds, userId).Delete(&model.Menu{}).Error; err != nil {
			return err
		}

		if err = tx.Table("menu_tags").Where("menu_id IN ?", menuIds).Delete(&model.MenuTag{}).Error; err != nil {
			return err
		}

		if err = tx.Table("menu_meta").Where("id IN ?", metaIds).Delete(&model.MenuMeta{}).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) UpdateMenuItem(menuId uint32, name string, icon string, isUse, isPrivacy uint8) error {
	err := d.engine.Transaction(func(tx *gorm.DB) error {
		var err error
		var menu model.Menu
		if err = tx.Table("menu").Where("id = ?", menuId).Updates(map[string]any{"name": name, "is_use": isUse, "is_privacy": isPrivacy}).Error; err != nil {
			return err
		}

		if err = tx.Table("menu").Where("id = ?", menuId).Find(&menu).Limit(1).Error; err != nil {
			return err
		}

		if err = tx.Table("menu_meta").Where("id = ?", menu.MetaId).Update("icon", icon).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) QueryChildMenu(menuId uint32, userId uint32) []model.Menu {
	var menu []model.Menu

	d.engine.Table("menu").Select("id, meta_id, parent_id").Where("parent_id = ? AND user_id = ?", menuId, userId).Find(&menu)

	return menu
}

func (d *Dao) GetMenuIdFamily(menuId uint32, userId uint32) []model.Menu {
	menu := d.QueryChildMenu(menuId, userId)

	if len(menu) > 0 {
		for _, v := range menu {
			list := d.GetMenuIdFamily(v.Id, userId)
			menu = append(menu, list...)
		}
	}

	return menu
}

func (d *Dao) QueryMenuById(menuId uint32, userId uint32) (model.Menu, error) {
	var menu model.Menu

	err := d.engine.Table("menu").Select("id, meta_id, parent_id").Where("id = ? AND user_id = ?", menuId, userId).Find(&menu).Limit(1).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (d *Dao) SaveMenuSort(sortList []model.SaveMenuSort) error {
	var errors []error
	err := d.engine.Transaction(func(tx *gorm.DB) error {
		for _, v := range sortList {
			if err := tx.Table("menu").Where("id = ?", v.Id).Updates(map[string]interface{}{"parent_id": v.ParentId, "sort": v.Sort}).Error; err != nil {
				errors = append(errors, err)
			}
		}
		if len(errors) > 0 {
			return fmt.Errorf("部分更新失败: %v", errors)
		}
		return nil
	})
	return err
}
