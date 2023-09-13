package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
)

func (d *Dao) GetMenu() []model.Menu {
	var menu []model.Menu

	d.engine.Table("menu").Joins("left join menu_meta on menu_meta.id = menu.meta_id").Scan(&menu)

	return menu
}