package chat

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Database
}

type Repository interface {
	CreateMessage(ctx context.Context, msg Chat) error
	AppendMessage(userID int, message Message) error
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

func (r *repository) AppendMessage(userID int, message Message) error {
	coll := r.db.Collection("chats")

	filter := bson.M{"user_id": userID}

	update := bson.M{
		"$push": bson.M{
			"messages": message,
		},
	}

	updateResult, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}

	if updateResult.MatchedCount == 0 {
		newChat := Chat{
			UserID:    userID,
			Messages:  []Message{message},
			StartedAt: time.Now(),
			EndedAt:   time.Now(),
		}

		_, err := coll.InsertOne(context.TODO(), newChat)
		if err != nil {
			return fmt.Errorf("error inserting new document: %v", err)
		}
	}

	return nil
}
