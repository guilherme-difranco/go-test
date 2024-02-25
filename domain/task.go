package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "tasks"
)

// Task representa a estrutura de uma tarefa dentro do sistema.
type Task struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title          string             `bson:"title" json:"title"`
	Description    string             `bson:"description" json:"description"`
	Priority       Priority           `bson:"priority" json:"priority"`
	Status         Status             `bson:"status" json:"status"`
	CreationDate   time.Time          `bson:"creationDate" json:"creationDate"`
	CompletionDate *time.Time         `bson:"completionDate,omitempty" json:"completionDate,omitempty"`
	UserID         primitive.ObjectID `bson:"userID" json:"userID"`
}

// TaskRepository define as operações de repositório disponíveis para uma Task.
type TaskRepository interface {
	Create(c context.Context, task *Task) error
	CreateBatch(c context.Context, tasks []Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
	FetchTasks(c context.Context, filter bson.M, projection bson.M, limit, skip int64) ([]Task, error)
	// Update(c context.Context, task *Task) error
	// Delete(c context.Context, id primitive.ObjectID) error
}

// TaskUsecase define as operações de caso de uso disponíveis para uma Task.
type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	CreateBatch(c context.Context, tasks []Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
	FetchUserTasks(c context.Context, filter, projection bson.M, limit, skip int64) ([]Task, error)
	// Update(c context.Context, task *Task) error
	// Delete(c context.Context, id primitive.ObjectID) error
}
