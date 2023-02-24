package service

import (
	"github.com/orglode/go-wake/api"
	"github.com/orglode/go-wake/conf"
	"github.com/orglode/go-wake/dao"
	"github.com/orglode/go-wake/manager"
	"go.uber.org/zap"
)

type Service struct {
	c      *conf.Config
	mgr    *manager.Manager
	dao    *dao.Dao
	logger *zap.Logger
}

func NewService(conf *conf.Config) *Service {
	return &Service{
		c:      conf,
		mgr:    manager.NewManager(conf),
		dao:    dao.NewDao(conf),
		logger: api.InitLogger(),
	}
}
