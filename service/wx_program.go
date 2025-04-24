package service

import (
	"context"
	"github.com/orglode/hades/logger"
)

func (s *Service) WxProgram(ctx context.Context) (interface{}, error) {
	logger.Debug(ctx, "service.WxProgram")

	return s.dao.GetUserAll(ctx)

}
