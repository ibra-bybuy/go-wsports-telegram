package telegram

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ibra-bybuy/go-wsports-telegram/pkg/constants"
	"github.com/ibra-bybuy/go-wsports-telegram/pkg/model"
)

// Command: /mma
func (t *Tg) mmaCommand(update tgbotapi.Update) {
	response := t.ctrl.GetBySport(context.TODO(), constants.MMA_TYPE, 50, 1)
	str := []string{}

	allItems := model.EventList(response.Items)
	filterItems := allItems.GetUniqueByName()
	markups := []tgbotapi.InlineKeyboardButton{}

	for i, event := range filterItems {
		str = append(str, fmt.Sprintf("%d. %s", i+1, event.Name))
		markups = append(markups, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%d", i+1), fmt.Sprintf("id:%s", event.ID)))
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, strings.Join(str, "\n"))
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(markups)
	t.bot.Send(msg)
}
