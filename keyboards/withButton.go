package keyboards

import "github.com/mymmrac/telego"

func WithButton(keyboard *telego.InlineKeyboardMarkup, button telego.InlineKeyboardButton) *telego.InlineKeyboardMarkup {
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []telego.InlineKeyboardButton{button})
	return keyboard
}
