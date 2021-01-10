package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"regexp"
)

func NewMessageHandler() *MessageHandler {
	h := &MessageHandler{}
	h.store = make(store)
	h.dataText = make(dataText)
	h.dataCMD = make(dataCMD)

	//set default
	h.SetStorageCleanInterval(60 * 60 * 24)

	//run cleaner
	go h.StorageCleaner()
	return h
}

func (h *MessageHandler) AddText(path Path, s S) {
	h.dataText[path] = s
}

func (h *MessageHandler) AddCommands(s S) {
	h.dataCMD = dataCMD(s)
}

func (h *MessageHandler) Handle(u tgbotapi.Update) {
	setPath := func(path Path) {
		h.SetPath(u.Message.Chat.ID, path)
	}

	if string(u.Message.Text[0]) == "/" {
		if fn, ok := h.dataCMD[u.Message.Text]; ok {
			go fn(u, setPath)
		}
	} else {
		for template, fn := range h.dataText[h.GetPath(u.Message.Chat.ID)] {
			if ok, _ := regexp.MatchString(template, u.Message.Text); ok {
				go fn(u, setPath)
				break
			}
		}
	}
}
