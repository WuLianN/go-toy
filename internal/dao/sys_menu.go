package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
)

func (d *Dao) GetMenu() []model.Menu {
	var menu []model.Menu
	result := d.engine.Where("is_use = ?", 1).Find(&menu)

	if (result.RowsAffected == 0 || result.Error != nil) {
		return nil
	}

	return menu
}