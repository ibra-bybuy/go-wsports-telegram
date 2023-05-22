package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ibra-bybuy/go-wsports-telegram/internal/repository/dotenv"
)

// Command: /app
func (t *Tg) appCommand(update tgbotapi.Update) {
	appURL := dotenv.Get("APP_URL")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Visit or website to download an app to watch sport event %s", appURL))
	t.bot.Send(msg)
}
