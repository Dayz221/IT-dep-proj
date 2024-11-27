package keyboards

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUserInfoKeyboard(groupId primitive.ObjectID, userId primitive.ObjectID) *telego.InlineKeyboardMarkup {
	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Задачи").WithCallbackData("showUserTasks&"+groupId.Hex()+"&"+userId.Hex()),
			tu.InlineKeyboardButton("Выдать таску").WithCallbackData("createTask&"+groupId.Hex()+"&"+userId.Hex()),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Похвалить").WithCallbackData("praiseUser&"+userId.Hex()),
			tu.InlineKeyboardButton("Наказать").WithCallbackData("punishUser&"+userId.Hex()),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Выдать админку").WithCallbackData("newAdmin&"+groupId.Hex()+"&"+userId.Hex()),
			tu.InlineKeyboardButton("Удалить из группы").WithCallbackData("deleteUser&"+groupId.Hex()+"&"+userId.Hex()),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("◀️ Назад").WithCallbackData("showGroups"),
		),
	)
}
