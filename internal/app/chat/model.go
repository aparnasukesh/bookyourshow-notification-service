package chat

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	UserID       primitive.ObjectID  `bson:"userId" json:"user_id"`
	AdminID      *primitive.ObjectID `bson:"adminId,omitempty" json:"admin_id,omitempty"`
	SuperAdminID *primitive.ObjectID `bson:"superAdminId,omitempty" json:"super_admin_id,omitempty"`
	Messages     []Message           `bson:"messages" json:"messages"`
	StartedAt    time.Time           `bson:"startedAt" json:"started_at"`
	EndedAt      *time.Time          `bson:"endedAt,omitempty" json:"ended_at,omitempty"`
}

type Message struct {
	MessageID primitive.ObjectID `bson:"messageId,omitempty" json:"message_id"`
	SenderID  primitive.ObjectID `bson:"senderId" json:"sender_id"`
	Message   string             `bson:"message" json:"message"`
	SentAt    time.Time          `bson:"sentAt" json:"sent_at"`
}
