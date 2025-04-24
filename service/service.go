package service

import (
	"navigator/conf"
	"navigator/dao"
	"navigator/manager"
)

type Service struct {
	c   *conf.Config
	mgr *manager.Manager
	dao *dao.Dao
}

func NewService(conf *conf.Config) *Service {
	return &Service{
		c:   conf,
		mgr: manager.NewManager(conf),
		dao: dao.NewDao(conf),
	}
}
