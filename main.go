package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"tg_calculator/utils"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("help"),
	),
)
var validExpr = regexp.MustCompile(`^[\d\s+\-*/()]+$`)

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
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			var ans string

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if update.Message.Text == "help" {
				ans = utils.HelpMsg
			} else if strings.TrimSpace(update.Message.Text) == "" || !validExpr.MatchString(update.Message.Text) {
				ans = utils.InvalidSyntaxMsg
			} else {
				ans = strconv.Itoa(utils.Calculate(update.Message.Text))
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, ans)
			msg.ParseMode = "MarkdownV2"
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
			msg.Text = utils.StartMsg
			msg.ReplyMarkup = numericKeyboard
		case "help":
			msg.Text = utils.HelpMsg
			msg.ParseMode = "MarkdownV2"
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		default:
			msg.Text = utils.DefaultMsg
			msg.ReplyMarkup = numericKeyboard
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
