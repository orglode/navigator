package service

import "context"

func (s *Service) WxProgram(ctx context.Context) (interface{}, error) {
	return s.dao.GetUserAll(ctx)
}
