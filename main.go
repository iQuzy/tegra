package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"tegra/app"
	"tegra/cmds"
	"tegra/cmds/paths"
	"tegra/library/handler"
)

func main() {
	//add message handler
	hl := handler.NewMessageHandler()

	hl.DefaultPath = &paths.HOME
	hl.SetStorageCleanInterval(60 * 60 * 24 * 3)

	hl.AddCommands(handler.S{
		"/start": cmds.Home,
	})

	hl.AddText(&paths.HOME, handler.S{
		"привет": cmds.Home,
	})

	//run bot listen
	app.Run(func(upd tgbotapi.Update) {
		log.Printf("[@%s] %s\n", upd.Message.Chat.UserName, upd.Message.Text)
		hl.Handle(upd)
	})
}