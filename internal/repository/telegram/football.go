package telegram

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ibra-bybuy/go-wsports-telegram/pkg/constants"
	"github.com/ibra-bybuy/go-wsports-telegram/pkg/model"
)

// Command: /football
func (t *Tg) footballCommand(update tgbotapi.Update, page int) {
	str, markups := t.getByPagination(constants.FOOTBALL_TYPE, page)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, str)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = markups
	t.bot.Send(msg)
}

func (t *Tg) getByPagination(sport string, page int) (string, tgbotapi.InlineKeyboardMarkup) {
	response := t.ctrl.GetBySport(context.TODO(), sport, 8, page)
	str := []string{}

	allItems := model.EventList(response.Items)
	markups := []tgbotapi.InlineKeyboardButton{}
	paginationMarkups := []tgbotapi.InlineKeyboardButton{}

	if len(allItems) < 1 {
		str = append(str, "no events yet")
	}

	for i, event := range allItems {
		str = append(str, fmt.Sprintf("%d. %s", i+1, event.GetTeamsName()))
		markups = append(markups, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%d", i+1), fmt.Sprintf("uuid:%s", event.Uuid)))
	}
	if page > 1 {
		data := fmt.Sprintf("sport:%s,page:%d", sport, page-1)
		paginationMarkups = append(paginationMarkups, tgbotapi.NewInlineKeyboardButtonData("⬅️", data))
	}

	if page < response.Pagination.TotalPages {
		data := fmt.Sprintf("sport:%s,page:%d", sport, page+1)
		paginationMarkups = append(paginationMarkups, tgbotapi.NewInlineKeyboardButtonData("➡️", data))
	}

	return strings.Join(str, "\n"), tgbotapi.NewInlineKeyboardMarkup(markups, paginationMarkups)
}
