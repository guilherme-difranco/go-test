package repository

import (
	"context"

	"github.com/guilherme-difranco/go-test/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type priorityRepository struct {
	database   *mongo.Database
	collection string
}

func NewPriorityRepository(database *mongo.Database, collection string) domain.PriorityRepository {
	return &priorityRepository{database: database, collection: collection}
}

func (r *priorityRepository) FetchByName(c context.Context, name string) (*domain.Priority, error) {
	var priority domain.Priority
	collection := r.database.Collection(r.collection)
	err := collection.FindOne(c, bson.M{"name": name}).Decode(&priority)
	if err != nil {
		return nil, err
	}
	return &priority, nil
}

func (r *priorityRepository) FetchAll(c context.Context) ([]domain.Priority, error) {
	var priorities []domain.Priority
	collection := r.database.Collection(r.collection)
	cursor, err := collection.Find(c, bson.D{{}}) // Busca todos os documentos
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)
	if err = cursor.All(c, &priorities); err != nil {
		return nil, err
	}
	return priorities, nil
}
