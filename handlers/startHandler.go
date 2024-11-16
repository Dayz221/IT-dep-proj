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

	var candidate models.User
	filter := bson.D{{Key: "chat_id", Value: message.Chat.ID}}
	err := users.FindOne(context.Background(), filter).Decode(&candidate)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Printf("Ошибка в handler /start: %s\n", err)
			return
		}

		newUser := models.NewUser()
		newUser.Username = message.From.Username
		newUser.UserId = message.From.ID

		_, err = users.InsertOne(
			context.Background(),
			newUser,
		)

		if err != nil {
			log.Printf("Ошибка в handler /start: %s\n", err)
		}
	}

}
