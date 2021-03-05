package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

type MessageHandlerFunc func(update *tgbotapi.Update) *ResponseError

type CommandManager struct {
	handlers map[string]MessageHandlerFunc
	api *tgbotapi.BotAPI
}

func NewCommandManager(api *tgbotapi.BotAPI) *CommandManager {
	return &CommandManager{
		handlers: make(map[string]MessageHandlerFunc),
		api: api,
	}
}

func (m *CommandManager) ProcessCommand(update *tgbotapi.Update) (*ResponseError, bool) {
	if update.Message == nil {
		return nil, true
	}

	if !(strings.HasPrefix(update.Message.Text, "/") && strings.HasSuffix(update.Message.Text, "@" + m.api.Self.UserName)) {
		return nil, true
	}

	parts := strings.Split(update.Message.Text, "@")
	if len(parts) < 2 {
		return NewResponseError(CommandProcessingError, errors.New("len(parts) < 2")), false
	}

	handlerFunc, ok := m.handlers[parts[0]]
	if !ok {
		return NewResponseError(CommandNotFoundError, errors.New("command not found")), false
	}

	return handlerFunc(update), false
}


func (m *CommandManager) RegisterHandler(command string, handler MessageHandlerFunc) {
	m.handlers[command] = handler
}
