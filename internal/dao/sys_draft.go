package dao

import (
	"errors"

	"github.com/WuLianN/go-toy/internal/model"
	"github.com/WuLianN/go-toy/pkg/app"
	"gorm.io/gorm"
)

// 查询已发布草稿
func (d *Dao) QueryPublishDraft(id uint32, userId uint32) (model.Draft, error) {
	var draft model.Draft

	err := d.engine.Table("drafts").Where("id = ? AND is_publish = ? AND is_delete = ?", id, 1, 0).First(&draft).Error
	if err != nil {
		return draft, err
	}

	// 是用户自己的草稿，直接返回结果
	if userId == draft.UserId {
		return draft, nil
	}

	// 非用户自己的草稿，并且是私密文章，抛错
	if draft.IsPrivacy == 1 {
		return draft, errors.New("私密文章")
	}

	bool, user := d.IsSystemUser("", draft.UserId)

	// 私密账号 抛错
	if bool && user.IsPrivacy == 1 {
		return draft, errors.New("私密账号")
	}

	return draft, nil
}

// 查询指定用户的草稿
func (d *Dao) QueryUserDraft(id uint32, userId uint32) (model.Draft, error) {
	var draft model.Draft
	err := d.engine.Table("drafts").Where("id = ? AND user_id = ? AND is_delete = ?", id, userId, 0).First(&draft).Error
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
	err := d.engine.Table("drafts").Where("id = ? AND user_id = ?", draft.Id, draft.UserId).Updates(map[string]interface{}{
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
	err := d.engine.Table("drafts").Where("id = ? AND user_id = ?", draft.Id, draft.UserId).Updates(map[string]interface{}{
		"title":       draft.Title,
		"update_time": draft.UpdateTime,
		"is_publish":  draft.IsPublish,
		"is_privacy":  draft.IsPrivacy,
		"bg_image":    draft.BgImage,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) DeleteDraft(id uint32) error {
	// 删除记录
	// err := d.engine.Table("draft").Where("id = ?", id).Delete(&model.Draft{}).Error

	d.engine.Transaction(func(tx *gorm.DB) error {
		var err error
		if err = tx.Table("drafts").Where("id = ?", id).Update("is_delete", 1).Error; err != nil {
			return err
		}

		draft := model.DraftTag{
			DraftId: id,
		}
		if err = tx.Table("draft_tags").Where("draft_id = ?", id).Delete(&draft).Error; err != nil {
			return err
		}

		return nil
	})

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
		err = d.engine.Table("draft_tags").Select("tags.name AS name, tags.Id AS id, tags.user_id, tags.bg_color, tags.color").Where("draft_tags.draft_id = ?", item.Id).Joins("left join tags on draft_tags.tag_id = tags.id").Find(&tags).Error
		list[index].Tags = append(list[index].Tags, tags...)
	}

	if err != nil {
		return list, err
	}

	return list, nil
}

func (d *Dao) QuerySearchDraftList(userId uint32, keyword string, page int, pageSize int, isSelf uint8) ([]model.DraftWithTags, error) {
	offset := app.GetPageOffset(page, pageSize)
	var list []model.DraftWithTags
	var err error

	if isSelf == 1 {
		err = d.engine.Table("drafts").Where("user_id = ? AND title LIKE ? AND is_publish = 1 AND is_delete = 0", userId, "%"+keyword+"%").Limit(pageSize).Offset(offset).Find(&list).Error
	} else {
		err = d.engine.Table("drafts").Where("user_id = ? AND title LIKE ? AND is_publish = 1 AND is_delete = 0 AND is_privacy = ?", userId, "%"+keyword+"%", 0).Limit(pageSize).Offset(offset).Find(&list).Error
	}

	for index, item := range list {
		var tags []model.Tag
		err = d.engine.Table("draft_tags").Select("tags.name AS name, tags.Id AS id, tags.user_id, tags.bg_color, tags.color").Where("draft_tags.draft_id = ?", item.Id).Joins("left join tags on draft_tags.tag_id = tags.id").Find(&tags).Error
		list[index].Tags = append(list[index].Tags, tags...)
	}

	if err != nil {
		return list, err
	}

	return list, nil
}
