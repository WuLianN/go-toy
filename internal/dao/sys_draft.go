package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
	"github.com/WuLianN/go-toy/pkg/app"
)

func (d *Dao) QueryDraft(id uint32) (model.Draft, error) {
	var draft model.Draft
	err := d.engine.Table("drafts").Where("id = ?", id).First(&draft).Error
	if err != nil {
		return draft, err
	}
	return draft, nil
}

func (d *Dao) CreateDraft(draft *model.Draft) (id uint32) {
	err := d.engine.Table("drafts").Create(draft).Error

	if err != nil {
		return 0
	}

	return draft.Id
}

// 新增 保存草稿
func (d *Dao) AddSaveDraft(draft *model.Draft) error {
	err := d.engine.Table("drafts").Where("id = ?", draft.Id).Updates(map[string]interface{}{
		"title":       draft.Title,
		"update_time": draft.UpdateTime,
		"content":     draft.Content,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

// 编辑 保存草稿 不需要更新content
func (d *Dao) EditSaveDraft(draft *model.Draft) error {
	err := d.engine.Table("drafts").Where("id = ?", draft.Id).Updates(map[string]interface{}{
		"title":       draft.Title,
		"update_time": draft.UpdateTime,
		"is_publish":  draft.IsPublish,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) DeleteDraft(id uint32) error {
	// 删除记录
	// err := d.engine.Table("draft").Where("id = ?", id).Delete(&model.Draft{}).Error

	// 假删除 is_delete=1
	err := d.engine.Table("drafts").Where("id = ?", id).Update("is_delete", 1).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) PublishDraft(id uint32) error {
	err := d.engine.Table("drafts").Where("id = ?", id).Update("is_publish", 1).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) QueryDraftList(userId uint32, status uint32, page int, pageSize int) ([]model.DraftWithTags, error) {
	var err error
	var list []model.DraftWithTags
	var draftStatus uint32 // 0全部 1已发布 2草稿
	offset := app.GetPageOffset(page, pageSize)

	if status == 1 {
		draftStatus = 1
	} else if status == 2 {
		draftStatus = 0
	}

	if status > 0 {
		err = d.engine.Table("drafts").Order("update_time DESC").Where("user_id = ? AND is_publish = ? AND is_delete = ?", userId, draftStatus, 0).Limit(pageSize).Offset(offset).Find(&list).Error
	} else {
		err = d.engine.Table("drafts").Order("update_time DESC").Where("user_id = ? AND is_delete = ?", userId, 0).Limit(pageSize).Offset(offset).Find(&list).Error
	}

	for index, item := range list {
		var tags []model.Tag
		err = d.engine.Table("draft_tags").Select("tags.name AS name, tags.Id AS id, tags.user_id").Where("draft_tags.draft_id = ?", item.Id).Joins("left join tags on draft_tags.tag_id = tags.Id").Find(&tags).Error
		list[index].Tags = append(list[index].Tags, tags...)
	}

	if err != nil {
		return list, err
	}

	return list, nil
}
