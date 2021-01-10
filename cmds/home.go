package cmds

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"tegra/app"
	"tegra/library/handler"
)

func Home(u tgbotapi.Update, setPath handler.FnPath) {
	t := "Hello"
	m := tgbotapi.NewMessage(u.Message.Chat.ID, t)
	app.Bot.Send(m)
}
