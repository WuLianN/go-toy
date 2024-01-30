package dao

import "github.com/WuLianN/go-toy/internal/model"

func (d *Dao) QueryDraft(id uint32) (model.Draft, error) {
	var draft model.Draft
	err := d.engine.Table("draft").Where("id = ?", id).First(&draft).Error
	if err != nil {
		return draft, err
	}
	return draft, nil
}

func (d *Dao) CreateDraft(draft *model.Draft) (id uint32) {
	err := d.engine.Table("draft").Create(draft).Error

	if err != nil {
		return 0
	}

	return draft.Id
}

func (d *Dao) UpdateDraft(draft *model.Draft) error {
	err := d.engine.Table("draft").Where("id = ?", draft.Id).Updates(&model.Draft{Title: draft.Title, Content: draft.Content, UpdateTime: draft.UpdateTime}).Error

	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) DeleteDraft(id uint32) error {
	// 删除记录
	// err := d.engine.Table("draft").Where("id = ?", id).Delete(&model.Draft{}).Error

	// 假删除 is_delete=1
	err := d.engine.Table("draft").Where("id = ?", id).Update("is_delete", 1).Error
	if err != nil {
		return err
	}

	return nil
}
