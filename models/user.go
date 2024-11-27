package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	UserId    int64                `bson:"user_id"`
	Username  string               `bson:"username"`
	Groups    []primitive.ObjectID `bson:"groups"`
	CurState  int                  `bson:"cur_state"` // 0 - ничего не ожидаем, 1 - ожидаем имя группы, 2 - ожидаем данные для создания задания
	TaskState TaskState            `bson:"task_state"`
}

func NewUser() User {
	return User{
		UserId:    0,
		Username:  "",
		Groups:    []primitive.ObjectID{},
		CurState:  0,
		TaskState: NewTaskState(),
	}
}
