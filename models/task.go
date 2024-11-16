package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	User        primitive.ObjectID  `bson:"user"`
	Title       string              `bson:"title"`
	Description string              `bson:"description"`
	Deadline    primitive.Timestamp `bson:"deadline"`
}
