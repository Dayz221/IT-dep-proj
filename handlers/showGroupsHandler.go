package handlers

import (
	"itproj/keyboards"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ShowGroupsHandler(bot *telego.Bot, query telego.CallbackQuery) {
	keyboard := keyboards.CreateGroupsInlineKeyboard(
		GetListOfGroups(bot, query),
		"showGroup",
	)

	keyboards.WithButton(
		keyboard,
		tu.InlineKeyboardButton("◀️ Назад").WithCallbackData("backToFunctions"),
	)

	bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:      tu.ID(query.Message.GetChat().ID),
		MessageID:   query.Message.GetMessageID(),
		Text:        "Вот твои группы, кусок говна ❤️:",
		ReplyMarkup: keyboard,
	})

	bot.AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{
		CallbackQueryID: query.ID,
	})
}