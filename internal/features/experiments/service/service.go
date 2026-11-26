package exp_service

import "context"

type ExpService struct {
	expRepository ExpRepository
}

type ExpRepository interface {
	ExpHttp(ctx context.Context) error
}

func NewExpService(expRepository ExpRepository) *ExpService {
	return &ExpService{
		expRepository: expRepository,
	}
}
