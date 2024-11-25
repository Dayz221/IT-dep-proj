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
			"Введи имя новой группы. Имя должно быть больше 3 символов:",
		).WithReplyMarkup(
			tu.InlineKeyboard(
				[]telego.InlineKeyboardButton{
					tu.InlineKeyboardButton(
						"Отмена",
					).WithCallbackData("cancel"),
				},
			),
		),
	)

	bot.AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{
		CallbackQueryID: query.ID,
	})
}

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

	newGroup := models.Group{
		ID:     primitive.NewObjectID(),
		Name:   name,
		Admins: []primitive.ObjectID{user.ID},
		Users:  []primitive.ObjectID{user.ID},
	}

	_, err = users.UpdateByID(context.Background(), user.ID, bson.D{{
		Key:   "$set",
		Value: bson.D{{Key: "cur_state", Value: 0}},
	}, {
		Key:   "$push",
		Value: bson.D{{Key: "groups", Value: newGroup.ID}},
	}})
	if err != nil {
		log.Printf("Ошибка в EnterGroupNameHandler: %s\n", err)
		return
	}

	_, err = groups.InsertOne(context.Background(), newGroup)
	if err != nil {
		log.Printf("Ошибка в EnterGroupNameHandler: %s\n", err)
		return
	}

	me, _ := bot.GetMe()

	_, err = bot.SendMessage(
		tu.Message(
			message.Chat.ChatID(),
			"Группа \""+name+"\" успешно создана\\!\n"+
				"Ты можешь пригласить команду по [ссылке](https://t.me/"+me.Username+"?start=invite="+newGroup.ID.Hex()+")\\.",
		).WithParseMode(telego.ModeMarkdownV2),
	)

	if err != nil {
		log.Printf("Ошибка в EnterGroupNameHandler: %s\n", err)
	}
}
