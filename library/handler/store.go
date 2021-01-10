package handler

import (
	"log"
	"time"
)

func NewStoreData(path Path) *StoreData {
	return &StoreData{path, time.Now().Unix()}
}

func (h *MessageHandler) GetPath(uid int64) Path {
	if data, ok := h.store[uid]; ok {
		return data.UserPath
	}
	h.SetPath(uid, h.DefaultPath)
	return h.DefaultPath
}

func (h *MessageHandler) SetPath(uid int64, path Path) {
	h.store[uid] = NewStoreData(path)
}

func (h *MessageHandler) StorageCleaner() {
	for {
		time.Sleep(h.storageCleanIntervalDuration)
		println("\n----------------------------")
		log.Println("Store clean..")

		now, count := time.Now().Unix(), 0
		for uid, data := range h.store {
			if now-data.CreateTime >= h.storageCleanInterval {
				log.Println("delete", uid)
				delete(h.store, uid)
				count++
			}
		}

		log.Printf("Clean finish(deleted %d)\n", count)
		print("----------------------------\n\n")
	}
}

func (h *MessageHandler) SetStorageCleanInterval(seconds int64) {
	h.storageCleanInterval = seconds
	h.storageCleanIntervalDuration = time.Second * time.Duration(seconds)
}
