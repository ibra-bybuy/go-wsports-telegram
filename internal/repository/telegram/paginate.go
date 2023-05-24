package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Command: /sport:footba,page=1
func (t *Tg) paginate(callback tgbotapi.CallbackQuery, sport string, page int) {
	str, markups := t.getByPagination(sport, page)

	callbackMsg := tgbotapi.CallbackConfig{
		CallbackQueryID: callback.ID,
		Text:            "success",
	}
	t.bot.Send(callbackMsg)

	msg := tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, str)
	msg.ReplyMarkup = &markups
	t.bot.Send(msg)
}
