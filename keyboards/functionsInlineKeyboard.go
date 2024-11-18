package keyboards

import tu "github.com/mymmrac/telego/telegoutil"

var FunctionsInlineKeyboard = tu.InlineKeyboard(
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("Создать группу").WithCallbackData("createGroup"),
		tu.InlineKeyboardButton("Показать группы").WithCallbackData("showGroups"),
	),
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("Показать все мои задания").WithCallbackData("showAllTasks"),
	),
)
