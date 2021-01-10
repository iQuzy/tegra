package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

type Path *interface{}
type FnPath func(Path)
type SFunc func(tgbotapi.Update, FnPath)
type S map[string]SFunc

type StoreData struct {
	UserPath   Path
	CreateTime int64
}

type store map[int64]*StoreData
type dataText map[Path]S
type dataCMD S

type MessageHandler struct {
	store                        store
	dataText                     dataText
	dataCMD                      dataCMD
	storageCleanInterval         int64
	storageCleanIntervalDuration time.Duration
	DefaultPath                  Path
}
