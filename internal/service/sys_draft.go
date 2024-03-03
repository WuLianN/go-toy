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
	Id           uint32 `json:"id" binding:"required"`
	UserId       uint32 `json:"user_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	IsPublish    uint8  `json:"is_publish"`
	IsPrivacy    uint8  `json:"is_privacy"`
	OperatedType uint8  `json:"operated_type"` // 操作类型  0:新增[默认] 1:修改
}

type DeleteRequest struct {
	Id uint32 `json:"id" binding:"required"`
}

type PublishRequest struct {
	Id uint32 `json:"id" binding:"required"`
}

type DraftListRequest struct {
	UserId   uint32      `json:"user_id"`
	Page     int         `json:"page" default:"1"`
	PageSize int         `json:"page_size" default:"10"`
	Status   uint32      `json:"status"`
	Tags     []model.Tag `json:"tags"`
}

func (svc *Service) GetDraft(id uint32, userId uint32) (model.Draft, error) {
	return svc.dao.QueryPublishDraft(id, userId)
}

func (svc *Service) GetUserDraft(id uint32, userId uint32) (model.Draft, error) {
	return svc.dao.QueryUserDraft(id, userId)
}

func (svc *Service) CreateDraft(userId uint32) (id uint32) {
	loc, err := time.LoadLocation("Asia/Shanghai")

	if err != nil {
		loc = time.FixedZone("CST", 8*3600) // 替换上海时间
	}

	draft := model.Draft{
		UserId:     userId,
		CreateTime: time.Now().In(loc).Format(time.DateTime),
		UpdateTime: time.Now().In(loc).Format(time.DateTime),
		IsPublish:  0,
		IsDelete:   0,
		IsPrivacy:  0,
	}

	return svc.dao.CreateDraft(&draft)
}

func (svc *Service) UpdateDraft(request SaveRequest) error {
	draft := model.Draft{
		Id:         request.Id,
		UserId:     request.UserId,
		Title:      request.Title,
		Content:    request.Content,
		IsPublish:  request.IsPublish,
		IsPrivacy:  request.IsPrivacy,
		UpdateTime: time.Now().Format(time.DateTime),
	}

	// 编辑 保存草稿
	if request.OperatedType == 1 {
		return svc.dao.EditSaveDraft(&draft)
	}
	// 新增 保存草稿
	return svc.dao.AddSaveDraft(&draft)
}

func (svc *Service) DeleteDraft(request DeleteRequest) error {
	return svc.dao.DeleteDraft(request.Id)
}

func (svc *Service) PublishDraft(request PublishRequest) error {
	return svc.dao.PublishDraft(request.Id)
}

func (svc *Service) GetDraftList(request *DraftListRequest) ([]model.DraftWithTags, error) {
	return svc.dao.QueryDraftList(request.UserId, request.Status, request.Page, request.PageSize)
}

func (svc *Service) SearchDrafts(userId uint32, keyword string, page int, pageSize int, isSelf uint8) ([]model.DraftWithTags, error) {
	return svc.dao.QuerySearchDraftList(userId, keyword, page, pageSize, isSelf)
}
