package handlers

// Хэндленры затычки

import (
	"log"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func TODO_HANDLER(bot *telego.Bot, message telego.Message) {
	log.Println("СДЕЛАТЬ ХЭНДЛЕР! " + message.Text)
}

func TODO_CALLBACK(bot *telego.Bot, query telego.CallbackQuery) {
	log.Println("СДЕЛАТЬ ХЭНДЛЕР! " + query.Data)
}

func GET_MOTIVATION(bot *telego.Bot, query telego.CallbackQuery) {
	stk := tu.FileFromID("CAACAgIAAxkBAAEKcjhnRgzTfXBoQVTB0z_7JmhiCsom5wAC7mAAAuJqIEiCrizr2MtgwDYE")
	bot.SendSticker(&telego.SendStickerParams{
		ChatID:  tu.ID(query.Message.GetChat().ID),
		Sticker: stk,
	})
	bot.AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{
		CallbackQueryID: query.ID,
	})
}
