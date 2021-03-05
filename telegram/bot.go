package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"log"
)

type TapirBot struct {
	api            *tgbotapi.BotAPI
	commandManager *CommandManager
	mediaManager   *MediaManager
	configManager  *TapirConfigManager
	pipeline       *Pipeline
}

func NewTapirBot(token string, configReader io.Reader) (*TapirBot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot := TapirBot{
		api:      api,
		pipeline: NewPipeline(),
	}
	bot.commandManager = NewCommandManager(&bot)
	bot.mediaManager = NewMediaManager(&bot)
	configManager, err := NewTapirConfigManager(configReader)
	if err != nil {
		return nil, err
	}
	bot.configManager = configManager
	return &bot, nil
}

func (bot *TapirBot) Init() {
	bot.pipeline.AddProcessors(bot.mediaManager.ProcessMedia, bot.commandManager.ProcessCommand)
	bot.commandManager.RegisterHandler("/ping", bot.HandlePing)
}

func (bot *TapirBot) Run() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updatesChan, err := bot.api.GetUpdatesChan(updateConfig)
	if err != nil {
		return err
	}

	for update := range updatesChan {
		responseError := bot.pipeline.Process(&update)
		if responseError != nil {
			if responseError.responseMessage != "" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseError.ResponseMessage())
				msg.ReplyToMessageID = update.Message.MessageID
				_, _ = bot.api.Send(msg)
			}
			if responseError.error != nil {
				log.Println(responseError.error)
			}
		}
	}

	return nil
}