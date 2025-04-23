package service

import (
	"navigator/api"
	"navigator/conf"
	"navigator/dao"
	"navigator/manager"
	"navigator/model"
)

type Service struct {
	c        *conf.Config
	mgr      *manager.Manager
	dao      *dao.Dao
	Response *model.BaseResponse
}

func NewService(conf *conf.Config) *Service {
	return &Service{
		c:   conf,
		mgr: manager.NewManager(conf),
		dao: dao.NewDao(conf),
		Response: &model.BaseResponse{
			Code: api.Success,
		},
	}
}
