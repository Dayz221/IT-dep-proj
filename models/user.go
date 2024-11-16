package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty"`
	UserId   int64                `bson:"user_id"`
	Username string               `bson:"username"`
	Groups   []primitive.ObjectID `bson:"groups"`
}

func NewUser() User {
	return User{UserId: 0, Username: "", Groups: []primitive.ObjectID{}}
}
