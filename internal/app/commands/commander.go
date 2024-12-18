package commands

import (
	"github.com/ArthurBJ/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registeredCommand = map[string]func(c *Commander, inputMessage *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	default:
		c.Default(update.Message)
	}
}
