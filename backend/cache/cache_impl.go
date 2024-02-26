package cache

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/guilherme-difranco/go-test/domain"
	"github.com/guilherme-difranco/go-test/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type cacheServiceImpl struct {
	db            *mongo.Database
	statusCache   map[string]domain.Status
	priorityCache map[string]domain.Priority
	cacheMutex    sync.RWMutex
}

func NewCacheService(db *mongo.Database) CacheService {
	service := &cacheServiceImpl{
		db:            db,
		statusCache:   make(map[string]domain.Status),
		priorityCache: make(map[string]domain.Priority),
	}
	service.InitializeCache() // Carrega os dados iniciais no cache
	return service
}

func (c *cacheServiceImpl) StoreStatus(status domain.Status) {
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()
	c.statusCache[status.Name] = status
}

func (c *cacheServiceImpl) GetStatus(name string) (domain.Status, bool) {
	c.cacheMutex.RLock()
	status, found := c.statusCache[name]
	c.cacheMutex.RUnlock()

	if !found {
		c.InitializeCacheStatus()
		c.cacheMutex.RLock()
		status, found = c.statusCache[name]
		c.cacheMutex.RUnlock()
	}
	return status, found
}

func (c *cacheServiceImpl) GetAllStatuses() ([]domain.Status, error) {
	c.cacheMutex.RLock()
	defer c.cacheMutex.RUnlock()

	// Se a cache estiver vazia, inicializa a cache de prioridades.
	if len(c.statusCache) == 0 {
		c.InitializeCachePriorities()
	}

	var statuses []domain.Status
	for _, status := range c.statusCache {
		statuses = append(statuses, status)
	}

	if len(statuses) == 0 {
		return nil, fmt.Errorf("no statuses found")
	}

	return statuses, nil
}

func (c *cacheServiceImpl) StorePriority(priority domain.Priority) {
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()
	c.priorityCache[priority.Name] = priority
}

func (c *cacheServiceImpl) GetPriority(name string) (domain.Priority, bool) {
	c.cacheMutex.RLock()
	priority, found := c.priorityCache[name]
	c.cacheMutex.RUnlock()

	if !found {
		c.InitializeCachePriorities()
		c.cacheMutex.RLock()
		priority, found = c.priorityCache[name]
		c.cacheMutex.RUnlock()
	}
	return priority, found
}

func (c *cacheServiceImpl) GetAllPriorities() ([]domain.Priority, error) {
	c.cacheMutex.RLock()
	defer c.cacheMutex.RUnlock()

	// Se a cache estiver vazia, inicializa a cache de prioridades.
	if len(c.priorityCache) == 0 {
		c.InitializeCachePriorities()
	}

	var priorities []domain.Priority
	for _, priority := range c.priorityCache {
		priorities = append(priorities, priority)
	}

	if len(priorities) == 0 {
		return nil, fmt.Errorf("no priorities found")
	}

	return priorities, nil
}

func (c *cacheServiceImpl) StoreAllPriorities(priorities []domain.Priority) {
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()
	for _, priority := range priorities {
		c.priorityCache[priority.Name] = priority
	}
}

func (c *cacheServiceImpl) StoreAllStatuses(statuses []domain.Status) {
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()
	for _, status := range statuses {
		c.statusCache[status.Name] = status
	}
}

func (c *cacheServiceImpl) InitializeCache() {
	c.InitializeCachePriorities()
	c.InitializeCacheStatus()
}

func (c *cacheServiceImpl) InitializeCachePriorities() {
	priorityRepo := repository.NewPriorityRepository(c.db, domain.CollectionPriorities)
	priorities, err := priorityRepo.FetchAll(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch priorities: %v", err)
	}
	c.StoreAllPriorities(priorities)
}

func (c *cacheServiceImpl) InitializeCacheStatus() {
	statusRepo := repository.NewStatusRepository(c.db, domain.CollectionStatus)
	statuses, err := statusRepo.FetchAll(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch statuses: %v", err)
	}
	c.StoreAllStatuses(statuses)
}