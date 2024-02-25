package repository

import (
	"context"

	"github.com/guilherme-difranco/go-test/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(db *mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   *db,
		collection: collection,
	}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	collection := tr.database.Collection(tr.collection)
	_, err := collection.InsertOne(c, task)
	return err
}

func (tr *taskRepository) CreateBatch(c context.Context, tasks []domain.Task) error {
	var docs []interface{}
	for _, task := range tasks {
		docs = append(docs, task)
	}

	collection := tr.database.Collection(tr.collection)
	_, err := collection.InsertMany(c, docs)
	return err
}

func (tr *taskRepository) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	collection := tr.database.Collection(tr.collection)

	var tasks []domain.Task
	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"userID": idHex}
	cursor, err := collection.Find(c, filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(c, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tr *taskRepository) FetchTasks(c context.Context, filter bson.M, projection bson.M, limit, skip int64) ([]domain.Task, error) {
	collection := tr.database.Collection(tr.collection)

	// Configurando opções de projeção e paginação
	opts := options.Find().SetProjection(projection).SetLimit(limit).SetSkip(skip)

	var tasks []domain.Task
	cursor, err := collection.Find(c, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	if err = cursor.All(c, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
