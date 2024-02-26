package usecase

import (
	"context"
	"time"

	"github.com/guilherme-difranco/go-test/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (tu *taskUsecase) Create(c context.Context, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Create(ctx, task)
}

func (tu *taskUsecase) CreateBatch(c context.Context, tasks []domain.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.CreateBatch(ctx, tasks)
}

func (tu *taskUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.FetchByUserID(ctx, userID)
}

func (tu *taskUsecase) FetchUserTasks(c context.Context, filter, projection bson.M, limit, skip int64) ([]domain.Task, error) {
	//filter := bson.M{"userID": userID}
	//projection := bson.M{"title": 1, "status": 1}
	tasks, err := tu.taskRepository.FetchTasks(c, filter, projection, limit, skip)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tu *taskUsecase) Update(ctx context.Context, id string, task domain.Task) error {
	return tu.taskRepository.Update(ctx, id, task)
}

func (tu *taskUsecase) Delete(ctx context.Context, id string) error {
	return tu.taskRepository.Delete(ctx, id)
}

func (tu *taskUsecase) FetchTaskByID(ctx context.Context, id primitive.ObjectID) (*domain.Task, error) {
	return tu.taskRepository.FetchTaskByID(ctx, id)
}
