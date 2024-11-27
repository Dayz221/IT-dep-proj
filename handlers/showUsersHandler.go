package handlers

import (
	"itproj/keyboards"
	"itproj/utils"
	"log"
	"strings"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ShowGroupUsersHandler(bot *telego.Bot, query telego.CallbackQuery) {
	groupId := strings.Split(query.Data, "&")[1]

	group, err := utils.GetGroupInfoByID(groupId)
	if err != nil {
		log.Printf("Че та наебнулось к чертям в ShowGroupUsersHandler: %s\n", err)
	}

	keyboard := keyboards.CreateUsersInlineKeyboard(
		utils.GetListOfUsers(groupId),
		"showUserInfo&"+groupId,
	)

	keyboards.WithButton(
		keyboard,
		tu.InlineKeyboardButton("◀️ Назад").WithCallbackData("showGroup&"+groupId),
	)

	bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:      tu.ID(query.Message.GetChat().ID),
		MessageID:   query.Message.GetMessageID(),
		Text:        "Пользователи группы \"" + group.Name + "\":",
		ReplyMarkup: keyboard,
	})

	bot.AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{
		CallbackQueryID: query.ID,
	})
}

func ShowGroupAdminsHandler(bot *telego.Bot, query telego.CallbackQuery) {
	groupId := strings.Split(query.Data, "&")[1]

	group, err := utils.GetGroupInfoByID(groupId)
	if err != nil {
		log.Printf("Че та наебнулось к чертям в ShowGroupAdminsHandler: %s\n", err)
	}

	keyboard := keyboards.CreateUsersInlineKeyboard(
		utils.GetListOfAdmins(groupId),
		"showAdminInfo&"+groupId,
	)

	keyboards.WithButton(
		keyboard,
		tu.InlineKeyboardButton("◀️ Назад").WithCallbackData("showGroup&"+groupId),
	)

	bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:      tu.ID(query.Message.GetChat().ID),
		MessageID:   query.Message.GetMessageID(),
		Text:        "Админы группы \"" + group.Name + "\":",
		ReplyMarkup: keyboard,
	})

	bot.AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{
		CallbackQueryID: query.ID,
	})
}
