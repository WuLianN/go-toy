package service

import "github.com/WuLianN/go-toy/internal/model"

type CreateTagRequest struct {
	UserId uint32 `json:"user_id"`
	Name   string `json:"name" binding:"required"`
}

type DeleteTagRequest struct {
	Id     uint32 `json:"id" binding:"required"`
	UserId uint32 `json:"user_id"`
}

func (svc *Service) GetTagList(userId uint32) ([]model.Tag, error) {
	return svc.dao.QueryTagList(userId)
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
