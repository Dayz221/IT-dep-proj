package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {
	ID     primitive.ObjectID   `bson:"_id,omitempty"`
	Users  []primitive.ObjectID `bson:"users"`
	Admins []primitive.ObjectID `bson:"admins"`
	Tasks  []primitive.ObjectID `bson:"tasks"`
}
