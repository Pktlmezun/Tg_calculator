package utils

import (
	"fmt"
	"github.com/Knetic/govaluate"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func EvaluateExpression(update tgbotapi.Update) tgbotapi.MessageConfig {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	ans, parseMode := calculate(update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, ans)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	if parseMode {
		msg.ParseMode = "MarkdownV2"
	}
	return msg
}

func calculate(s string) (string, bool) {
	var ans string
	parseMode := false
	if s == "help" {
		return helpMsg, true
	}
	expr, err := govaluate.NewEvaluableExpression(s)
	if err != nil {
		log.Printf("syntax error: %s\n", err.Error())
		ans = invalidSyntaxMsg
		parseMode = true
	} else {
		parameters := make(map[string]interface{})
		res, err := expr.Evaluate(parameters)
		if err != nil {
			log.Printf("evaluate error: %s\n", err.Error())
			ans = invalidSyntaxMsg
			parseMode = true
		} else {
			ans = fmt.Sprintf("%v", res)
		}
	}
	return ans, parseMode
}
