package keyboards

import (
	"itproj/models"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

var FunctionsInlineKeyboard = tu.InlineKeyboard(
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("Создать группу").WithCallbackData("createGroup"),
		tu.InlineKeyboardButton("Мои группы").WithCallbackData("showGroups"),
	),
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("Показать все мои задания").WithCallbackData("showAllTasks"),
		tu.InlineKeyboardButton("Выдать таску").WithCallbackData("createTask"),
	),
)

func CreateInlineKeyboard(lst []models.Group) telego.InlineKeyboardMarkup {

	rows := make([][]telego.InlineKeyboardButton, 0, 1)
	count := len(lst) / 2
	for i := 0; i < count; i += 2 {
		var row []telego.InlineKeyboardButton
		for j := i; (j < i+2) && j < len(lst); j++ {
			row = append(row, telego.InlineKeyboardButton{
				Text:         lst[j].Name,
				CallbackData: lst[j].ID.Hex(),
			})
		}
		rows = append(rows, row)
	}
	return telego.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
}
