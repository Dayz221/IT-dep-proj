package handlers

import (
	"itproj/keyboards"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func FunctionsHandler(bot *telego.Bot, message telego.Message) {
	bot.SendMessage(
		tu.Message(
			message.Chat.ChatID(),
			"Что ты хочешь сделать? 🥺\nВот, что я умею:",
		).WithReplyMarkup(keyboards.FunctionsInlineKeyboard),
	)
}

func BackToFunctionsHandler(bot *telego.Bot, query telego.CallbackQuery) {
	bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:      tu.ID(query.Message.GetChat().ID),
		MessageID:   query.Message.GetMessageID(),
		Text:        "Что ты хочешь сделать? 🥺\nВот, что я умею:",
		ReplyMarkup: keyboards.FunctionsInlineKeyboard,
	})
}
