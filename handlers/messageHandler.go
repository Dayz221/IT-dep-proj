package handlers

import (
	"itproj/models"
	"log"

	"github.com/mymmrac/telego"
)

func MessageHandler(bot *telego.Bot, message telego.Message) {
	log.Println(message.Text)

	user, err := models.GetUserById(message.From.ID)
	if err != nil {
		log.Printf("Ошибка в CreateGroupHandler: %s\n", err)
	}

	if user.CurState == 1 {
		EnterGroupNameHandler(bot, message)
	}
}
