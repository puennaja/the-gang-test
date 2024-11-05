package auth

import (
	"context"
	"daveslist/internal/core/domain/dto"
)

func (s *Service) Authorize(ctx context.Context, req *dto.AuthorizeRequest) (bool, error) {
	ok, err := s.casbinEnforcer.Enforce(req.Role, req.Object, req.Action)
	if err != nil {
		return false, err
	}
	return ok, nil
}
