package handlers

import (
	"itproj/models"
	"log"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func MessageHandler(bot *telego.Bot, message telego.Message) {
	log.Println(message.Text)

	user, err := models.GetUserById(message.From.ID)
	if err != nil {
		log.Printf("Ошибка в CreateGroupHandler: %s\n", err)
	}

	if user.CurState == 1 {
		EnterGroupNameHandler(bot, message)
	} else {
		bot.SendMessage(
			tu.Message(
				message.Chat.ChatID(),
				"Я тебя не понимаю 😐",
			),
		)
	}
}
