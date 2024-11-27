package handlers

import (
	"fmt"
	"itproj/keyboards"
	"itproj/utils"
	"strings"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ShowGroupInfoHandler(bot *telego.Bot, query telego.CallbackQuery) {
	group, err := utils.GetGroupInfoByID(strings.Split(query.Data, "&")[1])
	if err != nil {
		fmt.Printf("Что-то наебнулось в ShowGroupInfoHandler: %s\n", err)
		return
	}

	user, err := utils.GetUserByTgId(query.From.ID)
	if err != nil {
		fmt.Printf("Что-то наебнулось в ShowGroupInfoHandler: %s\n", err)
		return
	}

	keyboard := keyboards.CreateGroupInfoKeyboard(group.ID)
	if !utils.CheckAdmin(strings.Split(query.Data, "&")[1], user.ID.Hex()) {
		keyboard = keyboards.CreateGroupInfoKeyboardForUser(group.ID, user.ID)
	}

	bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:      tu.ID(query.Message.GetChat().ID),
		MessageID:   query.Message.GetMessageID(),
		Text:        "Действия с группой \"" + group.Name + "\":",
		ReplyMarkup: keyboard,
	})
}
