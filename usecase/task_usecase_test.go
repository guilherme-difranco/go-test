package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/guilherme-difranco/go-test/domain"
	"github.com/guilherme-difranco/go-test/domain/mocks"
	"github.com/guilherme-difranco/go-test/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

func TestTaskUsecase_CreateBatch(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	taskUsecase := usecase.NewTaskUsecase(mockTaskRepo, 10*time.Second)

	tasks := []domain.Task{{Title: "Task 1"}, {Title: "Task 2"}}

	mockTaskRepo.On("CreateBatch", mock.Anything, tasks).Return(nil)

	err := taskUsecase.CreateBatch(context.Background(), tasks)

	assert.NoError(t, err)
	mockTaskRepo.AssertExpectations(t)
}

func TestFetchTasks(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	taskUsecase := usecase.NewTaskUsecase(mockTaskRepo, 10*time.Second)

	mockTasks := []domain.Task{{Title: "Task 1"}, {Title: "Task 2"}}
	filter := bson.M{}
	projection := bson.M{"title": 1}
	limit := int64(10)
	skip := int64(0)

	mockTaskRepo.On("FetchTasks", mock.Anything, filter, projection, limit, skip).Return(mockTasks, nil)

	tasks, err := taskUsecase.FetchUserTasks(context.TODO(), filter, projection, limit, skip)
	assert.NoError(t, err)
	assert.Equal(t, mockTasks, tasks)

	mockTaskRepo.AssertExpectations(t)
}
