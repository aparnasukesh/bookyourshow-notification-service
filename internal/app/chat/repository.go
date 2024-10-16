package chat

import "go.mongodb.org/mongo-driver/mongo"

type repository struct {
	db *mongo.Database
}

type Repository interface {
}

func NewRepository(db *mongo.Database) Repository {
	return repository{
		db: db,
	}
}
