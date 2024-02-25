package domain

import (
	"context"
)

const (
	CollectionPriorities = "priorities"
)

type Priority struct {
	ID   string `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `bson:"name" json:"name"`
}

type PriorityRepository interface {
	FetchByName(c context.Context, name string) (*Priority, error)
	FetchAll(c context.Context) ([]Priority, error)
}
