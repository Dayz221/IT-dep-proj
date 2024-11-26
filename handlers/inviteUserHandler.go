package handlers

import (
	"itproj/utils"
	"log"
	"strings"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func InviteUserHandler(bot *telego.Bot, query telego.CallbackQuery) {
	group, err := utils.GetGroupInfoByID(strings.Split(query.Data, "&")[1])
	if err != nil {
		log.Printf("Что-то наебнулось в InviteUserHandler: %s\n", err)
		return
	}

	me, _ := bot.GetMe()

	bot.SendMessage(
		tu.Message(
			tu.ID(query.Message.GetChat().ID),
			"Ты можешь пригласить команду по [этой ссылке](https://t.me/"+me.Username+"?start=invite="+group.ID.Hex()+")\\.",
		).WithParseMode(telego.ModeMarkdownV2),
	)

	bot.AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{
		CallbackQueryID: query.ID,
	})
}
