package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tg_calculator/bot"
)

func main() {
	updates, bot := utils.Load()

	for update := range updates {
		if update.Message == nil {
			continue
		}
		var msg tgbotapi.MessageConfig

		if !update.Message.IsCommand() {
			msg = utils.EvaluateExpression(update)
		} else {
			msg = utils.Command(update)
		}
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
