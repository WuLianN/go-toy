package service

import (
	"time"

	"github.com/WuLianN/go-toy/internal/model"
)

type CreateRequest struct {
	UserId     uint32 `json:"user_id"`
	CreateTime string `json:"create_time"`
}

type SaveRequest struct {
	Id      uint32 `json:"id" binding:"required"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type DeleteRequest struct {
	Id uint32 `json:"id" binding:"required"`
}

type PublishRequest struct {
	Id uint32 `json:"id" binding:"required"`
}

type DraftListRequest struct {
	UserId   uint32 `json:"user_id"`
	Page     int    `json:"page" default:"1"`
	PageSize int    `json:"page_size" default:"10"`
	Status   uint32 `json:"status"`
}

func (svc *Service) GetDraft(id uint32) (model.Draft, error) {
	return svc.dao.QueryDraft(id)
}

func (svc *Service) CreateDraft(userId uint32) (id uint32) {
	draft := model.Draft{
		UserId:     userId,
		CreateTime: time.Now().Format(time.DateTime),
		UpdateTime: time.Now().Format(time.DateTime),
		IsPublish:  0,
		IsDelete:   0,
	}

	return svc.dao.CreateDraft(&draft)
}

func (svc *Service) UpdateDraft(request SaveRequest) error {
	draft := model.Draft{
		Id:         request.Id,
		Title:      request.Title,
		Content:    request.Content,
		UpdateTime: time.Now().Format(time.DateTime),
	}
	return svc.dao.UpdateDraft(&draft)
}

func (svc *Service) DeleteDraft(request DeleteRequest) error {
	return svc.dao.DeleteDraft(request.Id)
}

func (svc *Service) PublishDraft(request PublishRequest) error {
	return svc.dao.PublishDraft(request.Id)
}

func (svc *Service) GetDraftList(request *DraftListRequest) ([]model.Draft, error) {
	return svc.dao.QueryDraftList(request.UserId, request.Status, request.Page, request.PageSize)
}
