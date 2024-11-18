package handlers

import (
	"log"

	"github.com/mymmrac/telego"
)

func TODO_HANDLER(bot *telego.Bot, message telego.Message) {
	log.Println("СДЕЛАТЬ ХЭНДЛЕР! " + message.Text)
}

func TODO_CALLBACK(bot *telego.Bot, query telego.CallbackQuery) {
	log.Println("СДЕЛАТЬ ХЭНДЛЕР! " + query.Data)
}
