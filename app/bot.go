package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)
import . "tegra/library/checkError"

var Bot *tgbotapi.BotAPI
var UpdateConfig tgbotapi.UpdateConfig
var UpdatesChannel tgbotapi.UpdatesChannel

func init() {
	var err error

	Bot, err = tgbotapi.NewBotAPI(BotToken)
	CheckErrPanic(err, "Auth")

	UpdateConfig = tgbotapi.NewUpdate(0)
	UpdateConfig.Timeout = 60

	UpdatesChannel, err = Bot.GetUpdatesChan(UpdateConfig)
	CheckErrPanic(err, "Get Update Channel")
}

func Run(fn func(tgbotapi.Update)) {
	for upd := range UpdatesChannel {
		fn(upd)
	}
}