package service

import (
	"context"

	"github.com/SuperArnold/GO_Blog/global"
	"github.com/SuperArnold/GO_Blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	global.Logger.Info(param.Name)

	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}
