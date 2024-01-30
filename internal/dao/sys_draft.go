package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
	"github.com/WuLianN/go-toy/pkg/app"
)

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

func (d *Dao) PublishDraft(id uint32) error {
	err := d.engine.Table("draft").Where("id = ?", id).Update("is_publish", 1).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) QueryDraftList(userId uint32, page int, pageSize int) ([]model.Draft, error) {
	var list []model.Draft
	offset := app.GetPageOffset(page, pageSize)

	err := d.engine.Table("draft").Where("user_id = ? AND is_publish = ?", userId, 0).Limit(pageSize).Offset(offset).Find(&list).Error
	if err != nil {
		return list, err
	}

	return list, nil
}
