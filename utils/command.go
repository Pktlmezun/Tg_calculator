package utils

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("help"),
	),
)

func Command(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Command() {
	case "start":
		msg.Text = StartMsg
		msg.ReplyMarkup = numericKeyboard
	case "help":
		msg.Text = HelpMsg
		msg.ParseMode = "MarkdownV2"
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	default:
		msg.Text = DefaultMsg
		msg.ReplyMarkup = numericKeyboard
	}
	return msg
}
