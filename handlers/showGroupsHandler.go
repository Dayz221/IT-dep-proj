package handlers

import (
	"itproj/keyboards"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ShowGroupsHandler(bot *telego.Bot, query telego.CallbackQuery) {
	bot.DeleteMessage(&telego.DeleteMessageParams{
		ChatID:    tu.ID(query.Message.GetChat().ID),
		MessageID: query.Message.GetMessageID(),
	})

	bot.SendMessage(&telego.SendMessageParams{
		ChatID: tu.ID(query.Message.GetChat().ID),
		Text:   "Вот твои группы, кусок говна ❤️:",
		ReplyMarkup: keyboards.CreateGroupsInlineKeyboard(
			GetListOfGroups(bot, query),
			"groupClick",
		),
	})

	bot.AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{
		CallbackQueryID: query.ID,
	})
}
