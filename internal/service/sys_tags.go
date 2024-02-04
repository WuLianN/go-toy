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
	Id   uint32 `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (svc *Service) GetTagList(userId uint32, ids string, menuId uint32) ([]model.Tag, error) {
	idList := strings.Split(ids, ",")

	var idListInt []int
	if ids != "" {
		for _, idStr := range idList {
			idInt, _ := strconv.Atoi(idStr)
			idListInt = append(idListInt, idInt)
		}
	}

	return svc.dao.QueryTagList(userId, idListInt, menuId)
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

func (svc *Service) UpdateTag(req *UpdateTagRequest) error {
	return svc.dao.UpdateTag(req.Id, req.Name)
}

func (svc *Service) BindTag2Menu(menuTags *model.MenuTags, userId uint32) error {
	var tags []model.Tag

	for _, tag := range menuTags.Tags {
		tags = append(tags, model.Tag{
			UserId: userId,
			Name:   tag.Name,
		})
	}

	return svc.dao.BindTag2Menu(tags, menuTags.MenuId, userId)
}

func (svc *Service) UnbindTag2Menu(menuTags *model.MenuTags) error {
	var tagIds []uint32

	for _, tag := range menuTags.Tags {
		tagIds = append(tagIds, tag.Id)
	}

	return svc.dao.UnbindTag2Menu(tagIds, menuTags.MenuId)
}
