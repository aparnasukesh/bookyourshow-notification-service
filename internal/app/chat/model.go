package chat

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       int                `bson:"user_id" json:"user_id"`
	AdminID      int                `bson:"admin_id,omitempty" json:"admin_id,omitempty"`
	SuperAdminID int                `bson:"super_admin_id,omitempty" json:"super_admin_id,omitempty"`
	Messages     []Message          `bson:"messages" json:"messages"`
	StartedAt    time.Time          `bson:"started_at" json:"started_at"`
	EndedAt      *time.Time         `bson:"endedAt,omitempty" json:"ended_at,omitempty"`
}

type Message struct {
	MessageID primitive.ObjectID `bson:"message_id,omitempty" json:"message_id"`
	SenderID  int                `bson:"sender_id" json:"sender_id"`
	Message   string             `bson:"message" json:"message"`
	SentAt    time.Time          `bson:"sent_at" json:"sent_at"`
}
