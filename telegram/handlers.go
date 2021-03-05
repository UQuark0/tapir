package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (bot *TapirBot) HandlePing(update *tgbotapi.Update) *ResponseError {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, PingPong)
	msg.ReplyToMessageID = update.Message.MessageID
	_, err := bot.api.Send(msg)
	if err != nil {
		return NewResponseError(SendError, err)
	}
	return nil
}
