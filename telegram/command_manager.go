package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

type MessageHandlerFunc func(update *tgbotapi.Update) *ResponseError

type CommandManager struct {
	handlers map[string]MessageHandlerFunc
	tapir *TapirBot
}

func NewCommandManager(tapir *TapirBot) *CommandManager {
	return &CommandManager{
		handlers: make(map[string]MessageHandlerFunc),
		tapir: tapir,
	}
}

func (m *CommandManager) ProcessCommand(update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if !(strings.HasPrefix(update.Message.Text, "/") && strings.HasSuffix(update.Message.Text, "@" + m.tapir.api.Self.UserName)) {
		return
	}

	parts := strings.Split(update.Message.Text, "@")
	if len(parts) < 2 {
		_ = m.tapir.ReplyMessage(update.Message, CommandProcessingError)
		return
	}

	handlerFunc, ok := m.handlers[parts[0]]
	if !ok {
		_ = m.tapir.ReplyMessage(update.Message, CommandNotFoundError)
		return
	}

	err := handlerFunc(update)
	if err != nil {
		if err.error != nil {
			log.Println(err.error)
		}
		if err.ResponseMessage() != "" {
			_ = m.tapir.ReplyMessage(update.Message, err.ResponseMessage())
		}
	}
}


func (m *CommandManager) RegisterHandler(command string, handler MessageHandlerFunc) {
	m.handlers[command] = handler
}
