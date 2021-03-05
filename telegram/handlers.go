package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

func (bot *TapirBot) HandlePing(update *tgbotapi.Update) *ResponseError {
	err := bot.ReplyMessage(update.Message, PingPong)
	if err != nil {
		return NewResponseError(SendError, err)
	}
	return nil
}

func (bot *TapirBot) HandleSetReadonly24h(update *tgbotapi.Update) *ResponseError {
	if update.Message.ReplyToMessage == nil {
		err := bot.ReplyMessage(update.Message, MustBeAReply)
		if err != nil {
			return NewResponseError(SendError, err)
		}

		return nil
	}
	admins, err := bot.api.GetChatAdministrators(tgbotapi.ChatConfig{
		ChatID: update.Message.Chat.ID,
	})

	if err != nil {
		return NewResponseError(GettingAdminsError, err)
	}

	allowed := false
	for _, admin := range admins {
		if update.Message.From.ID == admin.User.ID && (admin.IsCreator() || admin.CanRestrictMembers) {
			allowed = true
			break
		}
	}

	if !allowed {
		err := bot.ReplyMessage(update.Message, RestrictPermissionRequired)
		if err != nil {
			return NewResponseError(SendError, err)
		}
		return nil
	}

	err = bot.SetReadonly(update.Message.ReplyToMessage.From.ID, update.Message.ReplyToMessage.Chat.ID, 24 * time.Hour)
	if err != nil {
		if err.Error() == "Bad Request: can't remove chat owner" || err.Error() == "Bad Request: user is an administrator of the chat" {
			_ = bot.ReplyMessage(update.Message, SetReadonlyFailedAdmin)
		} else {
			_ = bot.ReplyMessage(update.Message, SetReadonlyFailedUnknown)
		}
	} else {
		_ = bot.ReplyMessage(update.Message.ReplyToMessage, fmt.Sprintf(SetReadonlyHour, 24))
	}

	return nil
}