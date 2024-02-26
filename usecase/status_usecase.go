package usecase

import (
	"context"

	"github.com/guilherme-difranco/go-test/cache"
	"github.com/guilherme-difranco/go-test/domain"
)

type StatusUseCase interface {
	FetchAll(ctx context.Context) ([]domain.Status, error)
}

type statusUseCase struct {
	statusRepository domain.StatusRepository
	cacheService     cache.CacheService
}

func NewStatusUseCase(statusRepository domain.StatusRepository, cacheService cache.CacheService) StatusUseCase {
	return &statusUseCase{statusRepository: statusRepository, cacheService: cacheService}
}

func (su *statusUseCase) FetchAll(ctx context.Context) ([]domain.Status, error) {
	return su.cacheService.GetAllStatuses()
}
