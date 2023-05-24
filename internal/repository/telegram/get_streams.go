package telegram

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Command: /uuid:321
func (t *Tg) getStreams(uuid string, callback tgbotapi.CallbackQuery) {
	response := t.ctrl.GetByUuid(context.TODO(), uuid)
	str := []string{}

	if len(response.Data.Streams) > 0 {
		for _, stream := range response.Data.Streams {
			str = append(str, fmt.Sprintf("%s", stream.Link))
		}
	} else {
		str = append(str, "no streams yet")
	}

	callbackMsg := tgbotapi.CallbackConfig{
		CallbackQueryID: callback.ID,
		Text:            strings.Join(str, "\n"),
	}

	t.bot.Send(callbackMsg)

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, strings.Join(str, "\n"))
	msg.ReplyToMessageID = callback.Message.MessageID
	t.bot.Send(msg)
}
