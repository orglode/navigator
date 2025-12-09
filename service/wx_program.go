package service

import (
	"context"
	apiErr "navigator/api/error"

	"github.com/orglode/hades/logger"
)

func (s *Service) WxProgram(ctx context.Context) (interface{}, error) {
	logger.Debug(ctx, "service.WxProgram")

	return s.dao.GetUserAll(ctx)

}

func (s *Service) TestError(ctx context.Context) (interface{}, error) {
	return nil, apiErr.ErrInvalidPermission
}
