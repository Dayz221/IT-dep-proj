package handlers

import (
	"context"
	"itproj/models"
	"itproj/mongodb"
	"log"

	"github.com/mymmrac/telego"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartHandler(bot *telego.Bot, message telego.Message) {
	users := mongodb.GetUserCollection()

	filter := bson.D{{Key: "chat_id", Value: message.Chat.ID}}
	err := users.FindOne(context.Background(), filter).Decode(nil)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return
		}
	}

	newUser := models.User{
		Username: message.From.Username,
		ChatId:   message.Chat.ID,
	}

	_, err = users.InsertOne(
		context.Background(),
		newUser,
	)

	if err != nil {
		log.Printf("Ошибка в handler /start: %s", err)
	}
}
