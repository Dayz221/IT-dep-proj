package handlers

import (
	"context"
	"itproj/models"
	"itproj/mongodb"
	"log"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.mongodb.org/mongo-driver/bson"
)

func CancelHandler(bot *telego.Bot, query telego.CallbackQuery) {
	users := mongodb.GetUserCollection()
	user, err := models.GetUserById(query.From.ID)
	if err != nil {
		log.Printf("Ошибка в EnterGroupNameHandler: %s\n", err)
		return
	}
	_, err = users.UpdateByID(context.Background(), user.ID, bson.D{{
		Key:   "$set",
		Value: bson.D{{Key: "cur_state", Value: 0}},
	}})
	if err != nil {
		log.Printf("Ошибка в CreateGroupHandler: %s\n", err)
		return
	}
	bot.EditMessageText(
		&telego.EditMessageTextParams{
			ChatID:    tu.ID(query.Message.GetChat().ID),
			MessageID: query.Message.GetMessageID(),
			Text:      "Создание группы отменено.",
		},
	)
}
