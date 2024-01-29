package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
)

func (d *Dao) GetMenu(UserId uint32) []model.Menu {
	var menu []model.Menu

	d.engine.Table("menu").Select("menu.id as id, title, hide_children_in_menu, name, parent_id, meta_id, hide_menu, category, component, icon, path, redirect").Joins("left join menu_meta on menu_meta.id = menu.meta_id").Where("user_id = ? AND is_use = ?", UserId, 1).Scan(&menu)

	return menu
}
