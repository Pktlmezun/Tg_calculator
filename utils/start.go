package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tg_calculator/config"
)

func Load() (tgbotapi.UpdatesChannel, *tgbotapi.BotAPI) {
	conf := config.LoadConfig()
	bot, err := tgbotapi.NewBotAPI(conf.TelegramToken)

	if err != nil {
		fmt.Println("panicking")
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return bot.GetUpdatesChan(u), bot
}
