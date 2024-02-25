package repository

import (
	"context"

	"github.com/guilherme-difranco/go-test/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type statusRepository struct {
	database   *mongo.Database
	collection string
}

func NewStatusRepository(database *mongo.Database, collection string) domain.StatusRepository {
	return &statusRepository{database: database, collection: collection}
}

func (r *statusRepository) FetchByName(c context.Context, name string) (*domain.Status, error) {
	var status domain.Status
	collection := r.database.Collection(r.collection)
	err := collection.FindOne(c, bson.M{"name": name}).Decode(&status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}
func (r *statusRepository) FetchAll(c context.Context) ([]domain.Status, error) {
	var statuses []domain.Status
	collection := r.database.Collection(r.collection)
	cursor, err := collection.Find(c, bson.D{{}}) // Busca todos os documentos
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)
	if err = cursor.All(c, &statuses); err != nil {
		return nil, err
	}
	return statuses, nil
}
