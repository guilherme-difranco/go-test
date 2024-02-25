package mocks

import (
	"context"

	"github.com/guilherme-difranco/go-test/domain"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(c context.Context, task *domain.Task) error {
	args := m.Called(c, task)
	return args.Error(0)
}

func (m *MockTaskRepository) CreateBatch(c context.Context, tasks []domain.Task) error {
	args := m.Called(c, tasks)
	return args.Error(0)
}

func (m *MockTaskRepository) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	args := m.Called(c, userID)
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskRepository) FetchTasks(c context.Context, filter bson.M, projection bson.M, limit, skip int64) ([]domain.Task, error) {
	args := m.Called(c, filter, projection, limit, skip)
	return args.Get(0).([]domain.Task), args.Error(1)
}
