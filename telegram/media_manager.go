package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

type MediaManager struct {
	tapir *TapirBot
	mediaCount map[int64]int
}

func NewMediaManager(tapir *TapirBot) *MediaManager {
	return &MediaManager{
		tapir: tapir,
		mediaCount: make(map[int64]int),
	}
}

func (m *MediaManager) ProcessMedia(update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.Sticker != nil || update.Message.Photo != nil || update.Message.Video != nil || update.Message.Animation != nil {
		m.mediaCount[update.Message.Chat.ID]++
		if m.mediaCount[update.Message.Chat.ID] >= 4 {
			m.mediaCount[update.Message.Chat.ID] = 0

			err := m.tapir.SetReadonly(update.Message.From.ID, update.Message.Chat.ID, 24 * time.Hour)
			if err != nil {
				if err.Error() == "Bad Request: can't remove chat owner" || err.Error() == "Bad Request: user is an administrator of the chat" {
					_ = m.tapir.ReplyMessage(update.Message, SetReadonlyFailedAdmin)
				} else {
					_ = m.tapir.ReplyMessage(update.Message, SetReadonlyFailedUnknown)
				}
			} else {
				_ = m.tapir.ReplyMessage(update.Message, SetReadonly4thMediaRule)
			}
		}
	} else {
		m.mediaCount[update.Message.Chat.ID] = 0
	}

	return
}