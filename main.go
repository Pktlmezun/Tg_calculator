package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"tg_calculator/utils"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TELEGRAM_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		fmt.Println("panicking")
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			ans := strconv.Itoa(utils.Calculate(update.Message.Text))
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, ans)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
