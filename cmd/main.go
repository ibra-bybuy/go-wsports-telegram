package main

import (
	"github.com/ibra-bybuy/go-wsports-telegram/internal/controller"
	"github.com/ibra-bybuy/go-wsports-telegram/internal/repository/dotenv"
	"github.com/ibra-bybuy/go-wsports-telegram/internal/repository/rest"
	"github.com/ibra-bybuy/go-wsports-telegram/internal/repository/telegram"
)

func main() {
	// Load .env vars
	dotenv.Load()

	// Repository
	rep := rest.New(dotenv.Get("BASE_URL"))

	// Controller
	ctrl := controller.New(rep)

	// Loading tg
	tg := telegram.New(ctrl)
	tg.Listen()
}
