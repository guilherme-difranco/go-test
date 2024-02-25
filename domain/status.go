package domain

import (
	"context"
)

const (
	CollectionStatus = "status"
)

type Status struct {
	ID   string `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `bson:"name" json:"name"`
}

type StatusRepository interface {
	FetchByName(c context.Context, name string) (*Status, error)
	FetchAll(c context.Context) ([]Status, error)
}
