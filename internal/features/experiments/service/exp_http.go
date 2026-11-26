package exp_service

import (
	"context"
	"fmt"
)

func (s *ExpService) ExpHttp(ctx context.Context) error {

	err := s.expRepository.ExpHttp(ctx)
	if err != nil {
		return fmt.Errorf("exp http service: %w", err)
	}

	return nil
}
