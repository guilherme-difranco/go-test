// cache/cache.go
package cache

import "github.com/guilherme-difranco/go-test/domain"

// CacheService define a interface para o serviço de cache.
type CacheService interface {
	StoreStatus(status domain.Status)
	GetStatus(name string) (domain.Status, bool)
	StorePriority(priority domain.Priority)
	GetPriority(name string) (domain.Priority, bool)
	StoreAllPriorities(priorities []domain.Priority)
	StoreAllStatuses(statuses []domain.Status)
	InitializeCache()
	GetAllPriorities() ([]domain.Priority, error)
	GetAllStatuses() ([]domain.Status, error)
}
