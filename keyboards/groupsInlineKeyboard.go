package keyboards

import (
	"itproj/models"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateGroupsInlineKeyboard(lst []models.Group, prefix string) *telego.InlineKeyboardMarkup {
	btns := make([]telego.InlineKeyboardButton, 0, 1)

	for _, el := range lst {
		btns = append(btns, tu.InlineKeyboardButton(
			el.Name,
		).WithCallbackData(prefix+"&"+el.ID.Hex()),
		)
	}

	return tu.InlineKeyboardGrid(
		tu.InlineKeyboardCols(2, btns...),
	)
}
