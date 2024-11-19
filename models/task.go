package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	User        primitive.ObjectID `bson:"user"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Deadline    time.Time          `bson:"deadline"`
}

func NewTask() Task {
	return Task{Title: "", Description: "", Deadline: time.Now()}
}

type TaskState struct {
	Stage int  `bson:"stage"` // 0 - ввод ID пользователя, 1 - ввод заголовка, 2 - ввод описания, 3 - ввод дедлайна
	task  Task `bson:"task"`
}

func NewTaskState() TaskState {
	return TaskState{0, NewTask()}
}

func (taskState TaskState) GetTask() Task {
	taskState.task.ID = primitive.NewObjectID()
	return taskState.task
}
