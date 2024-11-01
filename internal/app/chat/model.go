package chat

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    int                `bson:"user_id" json:"user_id"`
	Messages  []Message          `bson:"messages" json:"messages"`
	StartedAt time.Time          `bson:"started_at" json:"started_at"`
	EndedAt   time.Time          `bson:"endedAt,omitempty" json:"ended_at,omitempty"`
}

type Message struct {
	Message         string `bson:"message" json:"message"`
	ResponseMessage string `bson:"response_message" json:"reponse_message"`

	SentAt time.Time `bson:"sent_at" json:"sent_at"`
}

type ReceiveMessage struct {
	UserID  int       `bson:"user_id" json:"user_id"`
	Message string    `bson:"message" json:"message"`
	SentAt  time.Time `bson:"sent_at" json:"sent_at"`
}
