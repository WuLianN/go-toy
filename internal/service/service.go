package service

import (
	"context"

	otgorm "github.com/WuLianN/go-toy/pkg/opentracing-gorm"

	"github.com/WuLianN/go-toy/global"
	"github.com/WuLianN/go-toy/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}

	// 链路追踪之 SQL 追踪
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))

	return svc
}
