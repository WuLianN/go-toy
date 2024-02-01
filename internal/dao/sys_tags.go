package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
)

func (d *Dao) QueryTagList(userId uint32) ([]model.Tag, error) {
	var list []model.Tag
	err := d.engine.Table("tags").Where("user_id = ?", userId).Find(&list).Error

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
