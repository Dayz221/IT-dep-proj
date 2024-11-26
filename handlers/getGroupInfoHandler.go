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

	bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:      tu.ID(query.Message.GetChat().ID),
		MessageID:   query.Message.GetMessageID(),
		Text:        "Действия с группой \"" + group.Name + "\":",
		ReplyMarkup: keyboards.CreateGroupInfoKeyboard(group.ID),
	})
}
