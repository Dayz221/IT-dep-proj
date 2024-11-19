package handlers

import (
	"context"
	"itproj/models"
	"itproj/mongodb"
	"log"
	"strings"
	"unicode/utf8"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EnterGroupNameHandler(bot *telego.Bot, message telego.Message) {
	users := mongodb.GetUserCollection()
	groups := mongodb.GetGroupCollection()

	name := strings.Trim(message.Text, " ")
	if utf8.RuneCountInString(name) <= 3 {
		bot.SendMessage(
			tu.Message(
				message.Chat.ChatID(),
				"Имя группы должно быть больше 3 символов.\nВведи другое название:",
			),
		)
		return
	}

	user, err := models.GetUserById(message.From.ID)
	if err != nil {
		log.Printf("Ошибка в EnterGroupNameHandler: %s\n", err)
		return
	}

	_, err = users.UpdateByID(context.Background(), user.ID, bson.D{{
		Key:   "$set",
		Value: bson.D{{Key: "cur_state", Value: 0}},
	}})
	if err != nil {
		log.Printf("Ошибка в EnterGroupNameHandler: %s\n", err)
		return
	}

	newGroup := models.Group{
		Name:   name,
		Admins: []primitive.ObjectID{user.ID},
		Users:  []primitive.ObjectID{user.ID},
	}

	_, err = groups.InsertOne(context.Background(), newGroup)
	if err != nil {
		log.Printf("Ошибка в EnterGroupNameHandler: %s\n", err)
		return
	}

	bot.SendMessage(
		tu.Message(
			message.Chat.ChatID(),
			"Группа \""+name+"\" успешно создана!",
		),
	)
}

func CreateGroupHandler(bot *telego.Bot, query telego.CallbackQuery) {
	users := mongodb.GetUserCollection()
	user, err := models.GetUserById(query.From.ID)
	if err != nil {
		log.Printf("Ошибка в EnterGroupNameHandler: %s\n", err)
		return
	}
	_, err = users.UpdateByID(context.Background(), user.ID, bson.D{{
		Key:   "$set",
		Value: bson.D{{Key: "cur_state", Value: 1}},
	}})
	if err != nil {
		log.Printf("Ошибка в CreateGroupHandler: %s\n", err)
		return
	}

	bot.SendMessage(
		tu.Message(
			tu.ID(query.Message.GetChat().ID),
			"Введите имя новой группы",
		),
	)
}
