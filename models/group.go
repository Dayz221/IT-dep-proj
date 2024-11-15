package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {
	ID     primitive.ObjectID   `bson:"_id,omitempty"`
	Users  []primitive.ObjectID `bson:"users,omitempty"`
	Admins []primitive.ObjectID `bson:"admins,omitempty"`
	Tasks  []primitive.ObjectID `bson:"tasks,omitempty"`
}
