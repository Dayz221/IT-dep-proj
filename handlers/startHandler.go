package handlers

import (
	"context"
	"itproj/keyboards"
	"itproj/models"
	"itproj/mongodb"
	"log"
	"strings"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func defaultHandler(bot *telego.Bot, message telego.Message) {
	log.Println("/start handler from " + message.From.Username)
	bot.SendMessage(
		tu.Message(
			message.Chat.ChatID(),
			"Бу! Испугался? Не бойся! Я бот от команды Pupupu!\n"+
				"Я помогу с оптимизацией задач для твоей группы.\n"+
				"Давай создавать задачи, пока мы не устанем!\n",
		).WithReplyMarkup(keyboards.StartInlineKeyboard),
	)
}

func inviteHandler(bot *telego.Bot, message telego.Message, groupId string) {
	log.Println("/invite handler to group: " + groupId)

	// TODO: сделать вход в гуппу

	bot.SendMessage(
		tu.Message(
			message.Chat.ChatID(),
			"Ты вступил в группу TODO!\n",
		),
	)
}

func StartHandler(bot *telego.Bot, message telego.Message) {
	users := mongodb.GetUserCollection()

	var candidate models.User
	filter := bson.D{{Key: "user_id", Value: message.Chat.ID}}
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

	args := strings.Split(message.Text, " ")
	if len(args) == 2 && strings.Split(args[1], "=")[0] == "invite" {
		inviteHandler(bot, message, strings.Split(args[1], "=")[1])
	} else {
		defaultHandler(bot, message)
	}
}
