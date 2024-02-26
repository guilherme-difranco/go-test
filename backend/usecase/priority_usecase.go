// Ensure to import the cache package where CacheService is defined
package usecase

import (
	"context"

	"github.com/guilherme-difranco/go-test/cache"
	"github.com/guilherme-difranco/go-test/domain"
	// Adjust this import path
)

type PriorityUseCase interface {
	FetchAll(ctx context.Context) ([]domain.Priority, error)
}

type priorityUseCase struct {
	repo         domain.PriorityRepository
	cacheService cache.CacheService // Ensure CacheService is correctly imported
}

// Ensure the constructor function matches the interface and implementation names
func NewPriorityUseCase(repo domain.PriorityRepository, cacheService cache.CacheService) PriorityUseCase { // Adjusted return type to match interface
	return &priorityUseCase{
		repo:         repo,
		cacheService: cacheService,
	}
}

func (uc *priorityUseCase) FetchAll(ctx context.Context) ([]domain.Priority, error) {
	return uc.cacheService.GetAllPriorities() // Ensure this method exists and is implemented in your cache service
}
