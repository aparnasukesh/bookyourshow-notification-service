package chat

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Database
}

type Repository interface {
	CreateMessage(ctx context.Context, msg Chat) error
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateMessage(ctx context.Context, msg Chat) error {
	coll := r.db.Collection("chat")
	if _, err := coll.InsertOne(ctx, msg); err != nil {
		return err
	}
	return nil
}
