package keyboards

import (
	tu "github.com/mymmrac/telego/telegoutil"
)

var StartInlineKeyboard = tu.InlineKeyboard(
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("Создать новую группу").WithCallbackData("createGroup"),
	),
)