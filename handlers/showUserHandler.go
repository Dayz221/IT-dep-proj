package handlers

import (
	"fmt"
	"itproj/keyboards"
	"itproj/utils"
	"strings"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ShowUserInfoHandler(bot *telego.Bot, query telego.CallbackQuery) {
	group, err := utils.GetGroupInfoByID(strings.Split(query.Data, "&")[1])
	if err != nil {
		fmt.Printf("Что-то наебнулось в ShowGroupInfoHandler: %s\n", err)
		return
	}

	user, err := utils.GetUserById(strings.Split(query.Data, "&")[2])
	if err != nil {
		fmt.Printf("Что-то наебнулось в ShowGroupInfoHandler: %s\n", err)
		return
	}

	// user, err := models.GetUserById()

	bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:      tu.ID(query.Message.GetChat().ID),
		MessageID:   query.Message.GetMessageID(),
		Text:        "Действия с пользователем @" + user.Username + " в группе \"" + group.Name + "\":",
		ReplyMarkup: keyboards.CreateUserInfoKeyboard(group.ID, user.ID),
	})
}

func ShowAdminInfoHandler(bot *telego.Bot, query telego.CallbackQuery) {

}
