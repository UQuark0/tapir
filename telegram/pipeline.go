package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type MessageProcessorFunc func(update *tgbotapi.Update) (*ResponseError, bool)

type Pipeline struct {
	processors []MessageProcessorFunc
}

func NewPipeline() *Pipeline {
	return &Pipeline{processors: make([]MessageProcessorFunc, 0)}
}

func (p *Pipeline) AddProcessors(processors ...MessageProcessorFunc) {
	p.processors = append(p.processors, processors...)
}

func (p *Pipeline) Process(update *tgbotapi.Update) *ResponseError {
	for _, processorFunc := range p.processors {
		responseError, processed := processorFunc(update)
		if processed {
			return responseError
		}
	}

	return nil
}