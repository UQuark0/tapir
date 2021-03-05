package telegram

import (
	"fmt"
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

	config := m.tapir.configManager.GetConfig(update.Message.Chat.UserName)

	if config.NthMediaRule.Count <= 0 {
		return
	}

	if update.Message.Sticker != nil || update.Message.Photo != nil || update.Message.Video != nil || update.Message.Animation != nil {
		m.mediaCount[update.Message.Chat.ID]++
		if m.mediaCount[update.Message.Chat.ID] >= config.NthMediaRule.Count {
			m.mediaCount[update.Message.Chat.ID] = 0

			err := m.tapir.SetReadonly(update.Message.From.ID, update.Message.Chat.ID, time.Duration(config.NthMediaRule.ReadonlyDuration) * time.Hour)
			if err != nil {
				if err.Error() == "Bad Request: can't remove chat owner" || err.Error() == "Bad Request: user is an administrator of the chat" {
					_ = m.tapir.ReplyMessage(update.Message, SetReadonlyFailedAdmin)
				} else {
					_ = m.tapir.ReplyMessage(update.Message, SetReadonlyFailedUnknown)
				}
			} else {
				_ = m.tapir.ReplyMessage(update.Message, fmt.Sprintf(SetReadonlyNthMediaRule, config.NthMediaRule.ReadonlyDuration))
			}
		}
	} else {
		m.mediaCount[update.Message.Chat.ID] = 0
	}

	return
}