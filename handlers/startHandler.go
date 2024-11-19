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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func inviteHandler(bot *telego.Bot, message telego.Message, groupIdHex string) {
	log.Println("/invite handler to group: " + groupIdHex)

	groupCollection := mongodb.GetGroupCollection()
	userCollection := mongodb.GetUserCollection()

	var group models.Group
	groupId, err := primitive.ObjectIDFromHex(groupIdHex)
	if err != nil {
		log.Printf("Ошибка в inviteHandler: %s\n", err)
		return
	}

	err = groupCollection.FindOne(context.Background(), bson.D{{Key: "_id", Value: groupId}}).Decode(&group)
	if err != nil {
		log.Printf("Ошибка в inviteHandler: %s\n", err)
		return
	}

	user, err := models.GetUserById(message.From.ID)
	if err != nil {
		log.Printf("Ошибка в inviteHandler: %s\n", err)
		return
	}

	for _, el := range group.Users {
		if el == user.ID {
			bot.SendMessage(
				tu.Message(
					message.Chat.ChatID(),
					"Ты уже состоишь в группе \""+group.Name+"\"!",
				),
			)
			return
		}
	}

	_, err = groupCollection.UpdateByID(context.Background(), groupId, bson.M{
		"$push": bson.M{"users": user.ID},
	})
	if err != nil {
		log.Printf("Ошибка в inviteHandler: %s\n", err)
		return
	}

	_, err = userCollection.UpdateByID(context.Background(), user.ID, bson.M{
		"$push": bson.M{"groups": groupId},
	})
	if err != nil {
		log.Printf("Ошибка в inviteHandler: %s\n", err)
		return
	}

	bot.SendMessage(
		tu.Message(
			message.Chat.ChatID(),
			"Ты успешно вступил в группу \""+group.Name+"\"!",
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
			log.Printf("Ошибка в StartHandler: %s\n", err)
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
			log.Printf("Ошибка в StartHandler: %s\n", err)
			return
		}
	}

	args := strings.Split(message.Text, " ")
	if len(args) == 2 && strings.Split(args[1], "=")[0] == "invite" {
		inviteHandler(bot, message, strings.Split(args[1], "=")[1])
	} else {
		defaultHandler(bot, message)
	}
}
