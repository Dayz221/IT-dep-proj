package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty"`
	ChatId   int64                `bson:"chat_id"`
	Username string               `bson:"username"`
	Groups   []primitive.ObjectID `bson:"groups,omitempty"`
}
