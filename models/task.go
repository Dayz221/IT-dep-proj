package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	User        primitive.ObjectID  `bson:"user"`
	Title       string              `bson:"title"`
	Description string              `bson:"description"`
	Deadline    primitive.Timestamp `bson:"deadline"`
}

func NewTask() Task {
	return Task{Title: "", Description: "", Deadline: primitive.Timestamp{T: uint32(time.Now().Add(24 * 7 * time.Hour).Unix()), I: 0}}
}
