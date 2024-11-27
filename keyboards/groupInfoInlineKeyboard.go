package keyboards

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateGroupInfoKeyboard(groupId primitive.ObjectID) *telego.InlineKeyboardMarkup {
	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Все задачи").WithCallbackData("showGroupTasks&"+groupId.Hex()),
			tu.InlineKeyboardButton("Выдать таску").WithCallbackData("createTask&"+groupId.Hex()),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Пользователи").WithCallbackData("showUsers&"+groupId.Hex()),
			tu.InlineKeyboardButton("Админы").WithCallbackData("showAdmins&"+groupId.Hex()),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Переименовать").WithCallbackData("renameGroup&"+groupId.Hex()),
			tu.InlineKeyboardButton("Удалить").WithCallbackData("deleteGroup&"+groupId.Hex()),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Пригласить пользователя").WithCallbackData("inviteUser&"+groupId.Hex()),
			// tu.InlineKeyboardButton("").WithCallbackData("showAdmins&"+groupId.Hex()),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("◀️ Назад").WithCallbackData("showGroups"),
		),
	)
}
