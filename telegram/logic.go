package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

func (bot *TapirBot) SetReadonly(userID int, chatID int64, duration time.Duration) error {
	f := false
	rcmConfig := tgbotapi.RestrictChatMemberConfig{
		ChatMemberConfig:      tgbotapi.ChatMemberConfig{
			ChatID:             chatID,
			UserID:             userID,
		},
		UntilDate:             time.Now().UTC().Add(duration).Unix(),
		CanSendMessages:       &f,
		CanSendMediaMessages:  &f,
		CanSendOtherMessages:  &f,
		CanAddWebPagePreviews: &f,
	}
	_, err := bot.api.RestrictChatMember(rcmConfig)
	return err
}

func (bot *TapirBot) ReplyMessage(replyTo *tgbotapi.Message, text string) error {
	msg := tgbotapi.NewMessage(replyTo.Chat.ID, text)
	msg.ReplyToMessageID = replyTo.MessageID
	_, err := bot.api.Send(msg)
	return err
}

func (bot *TapirBot) SendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.api.Send(msg)
	return err
}