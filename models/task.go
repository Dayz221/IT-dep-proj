package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	User        primitive.ObjectID  `bson:"user,omitempty"`
	Title       string              `bson:"title,omitempty"`
	Description string              `bson:"description,omitempty"`
	Deadline    primitive.Timestamp `bson:"deadline,omitempty"`
}
