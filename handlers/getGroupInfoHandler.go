package handlers

import (
	"context"
	"fmt"
	"itproj/keyboards"
	"itproj/models"
	"itproj/mongodb"
	"strings"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ShowGroupInfoHandler(bot *telego.Bot, query telego.CallbackQuery) {
	groupId, err := primitive.ObjectIDFromHex(strings.Split(query.Data, "&")[1])
	if err != nil {
		fmt.Printf("Что-то наебнулось в ShowGroupInfoHandler: %s\n", err)
		return
	}

	var group models.Group
	groupCollection := mongodb.GetGroupCollection()
	if err := groupCollection.FindOne(context.Background(), bson.D{{Key: "_id", Value: groupId}}).Decode(&group); err != nil {
		fmt.Printf("Что-то наебнулось в ShowGroupInfoHandler: %s\n", err)
		return
	}

	bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:      tu.ID(query.Message.GetChat().ID),
		MessageID:   query.Message.GetMessageID(),
		Text:        "Действия с группой \"" + group.Name + "\":",
		ReplyMarkup: keyboards.CreateGroupInfoKeyboard(group.ID),
	})
}
