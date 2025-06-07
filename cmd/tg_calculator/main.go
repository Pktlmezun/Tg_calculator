package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tg_calculator/utils"
)

func main() {
	updates, tgBot := bot.Load()

	for update := range updates {
		if update.Message == nil {
			continue
		}
		var msg tgbotapi.MessageConfig

		if !update.Message.IsCommand() {
			msg = bot.EvaluateExpression(update)
		} else {
			msg = bot.Command(update)
		}
		if _, err := tgBot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
