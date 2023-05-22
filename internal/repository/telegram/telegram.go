package telegram

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ibra-bybuy/go-wsports-telegram/internal/controller"
	"github.com/ibra-bybuy/go-wsports-telegram/internal/repository/dotenv"
)

type Tg struct {
	bot  *tgbotapi.BotAPI
	ctrl *controller.Controller
}

func New(ctrl *controller.Controller) *Tg {
	token := dotenv.Get("TG_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Tg{bot, ctrl}
}

func (t *Tg) Listen() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := t.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			txt := update.CallbackQuery.Data

			if strings.Contains(txt, "id:") {
				id := strings.Split(txt, ":")[1]
				t.getStreams(id, *update.CallbackQuery)
			}
		}

		if update.Message != nil {
			txt := update.Message.Text
			log.Printf("[%s] %s", update.Message.From.UserName, txt)

			if txt == "/app" {
				t.appCommand(update)
			} else if txt == "/mma" {
				t.mmaCommand(update)
			}
		}
	}
}
