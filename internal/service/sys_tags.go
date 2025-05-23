package service

import (
	"strconv"
	"strings"

	"github.com/WuLianN/go-toy/internal/model"
)

type CreateTagRequest struct {
	UserId uint32 `json:"user_id"`
	Name   string `json:"name" binding:"required"`
}

type DeleteTagRequest struct {
	Id     uint32 `json:"id" binding:"required"`
	UserId uint32 `json:"user_id"`
}

type UpdateTagRequest struct {
	Id      uint32 `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Color   string `json:"color"`
	BgColor string `json:"bg_color"`
}

func (svc *Service) GetTagList(userId uint32, ids string) ([]model.Tag, error) {
	idList := strings.Split(ids, ",")

	var idListInt []int
	if ids != "" {
		for _, idStr := range idList {
			idInt, _ := strconv.Atoi(idStr)
			idListInt = append(idListInt, idInt)
		}
	}

	return svc.dao.QueryTagList(userId, idListInt)
}

func (svc *Service) GetMenuTagList(menuId uint32) ([]model.Tag, error) {
	return svc.dao.QueryMenuTags(menuId)
}

func (svc *Service) GetDraftTagList(userId uint32, tagId uint32, draftId uint32) ([]model.Tag, error) {
	if draftId > 0 {
		return svc.dao.QueryDraftTagsT(userId, 0, draftId)
	}

	if tagId > 0 {
		return svc.dao.QueryDraftTagsT(userId, tagId, 0)
	}

	return svc.dao.QueryDraftTagsT(userId, 0, 0)
}

func (svc *Service) CreateTag(req *CreateTagRequest) (uint32, error) {
	return svc.dao.CreateTag(req.UserId, req.Name)
}

func (svc *Service) DeleteTag(req *DeleteTagRequest) error {
	return svc.dao.DeleteTag(req.Id)
}

func (svc *Service) QueryTag(req *CreateTagRequest) ([]model.Tag, error) {
	return svc.dao.QueryTag(req.UserId, req.Name)
}

func (svc *Service) FuzzyQueryTags(req *CreateTagRequest) ([]model.Tag, error) {
	return svc.dao.FuzzyQueryTags(req.UserId, req.Name)
}

func (svc *Service) UpdateTag(req *UpdateTagRequest) error {
	return svc.dao.UpdateTag(req.Id, req.Name, req.Color, req.BgColor)
}

func (svc *Service) BindTag2Menu(menuTags *model.MenuTags, userId uint32) error {
	var tags []model.Tag

	for _, tag := range menuTags.Tags {
		tags = append(tags, model.Tag{
			UserId: userId,
			Name:   tag.Name,
		})
	}

	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}

	exsitTags, _ := svc.dao.QueryTags(userId, tagNames) // tags库中存在该user_id的标签 -> 绑定
	newTags := removeExisting(tags, exsitTags)          // 新标签 -> 创建 + 绑定

	return svc.dao.BindTag2Menu(exsitTags, newTags, menuTags.MenuId, userId)
}

func (svc *Service) UnbindTag2Menu(menuTags *model.MenuTags) error {
	var tagIds []uint32

	for _, tag := range menuTags.Tags {
		tagIds = append(tagIds, tag.Id)
	}

	return svc.dao.UnbindTag2Menu(tagIds, menuTags.MenuId)
}

func (svc *Service) BindTag2Draft(draftTags *model.DraftTags, userId uint32) error {
	var tags []model.Tag

	for _, tag := range draftTags.Tags {
		tags = append(tags, model.Tag{
			UserId: userId,
			Name:   tag.Name,
		})
	}

	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}

	exsitTags, _ := svc.dao.QueryTags(userId, tagNames) // tags库中存在该user_id的标签 -> 绑定
	newTags := removeExisting(tags, exsitTags)          // 新标签 -> 创建 + 绑定

	return svc.dao.BindTag2Draft(exsitTags, newTags, draftTags.DraftId, userId)
}

func (svc *Service) UnbindTag2Draft(draftTags *model.DraftTags) error {
	var tagIds []uint32

	for _, tag := range draftTags.Tags {
		tagIds = append(tagIds, tag.Id)
	}

	return svc.dao.UnbindTag2Draft(tagIds, draftTags.DraftId)
}

func removeExisting(tags []model.Tag, existTags []model.Tag) []model.Tag {
	// 使用 map 来加速查找
	existTagMap := make(map[string]bool)
	for _, exTag := range existTags {
		existTagMap[exTag.Name] = true
	}

	// 创建一个新的切片来存储不重复的 tags
	newTags := []model.Tag{}
	for _, tag := range tags {
		if !existTagMap[tag.Name] {
			newTags = append(newTags, tag)
		}
	}

	return newTags
}
