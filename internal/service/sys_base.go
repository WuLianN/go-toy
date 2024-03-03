package service

import (
	"github.com/WuLianN/go-toy/internal/model"
)

func (svc *Service) GetRecommendList(userId uint32, page int, pageSize int, tagId uint32, isSelf uint8) ([]model.RecommendList, error) {
	return svc.dao.QueryRecommendList(userId, page, pageSize, tagId, isSelf)
}
