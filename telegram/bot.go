package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type TapirBot struct {
	api *tgbotapi.BotAPI
	cm *CommandManager
	mm *MediaManager
	pl *Pipeline
}

func NewTapirBot(token string) (*TapirBot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot := TapirBot{
		api: api,
		pl: NewPipeline(),
	}
	bot.cm = NewCommandManager(&bot)
	bot.mm = NewMediaManager(&bot)
	return &bot, nil
}

func (bot *TapirBot) Init() {
	bot.pl.AddProcessors(bot.mm.ProcessMedia, bot.cm.ProcessCommand)
	bot.cm.RegisterHandler("/ping", bot.HandlePing)
}

func (bot *TapirBot) Run() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updatesChan, err := bot.api.GetUpdatesChan(updateConfig)
	if err != nil {
		return err
	}

	for update := range updatesChan {
		responseError := bot.pl.Process(&update)
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