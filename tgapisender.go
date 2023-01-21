package main

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessageToChat(mc tgbotapi.MessageConfig) error {
	bot, err := tgbotapi.NewBotAPI(*flagBotToken)
	if err != nil {
		return errors.New("Cannot deliver the message: " + err.Error())
	}
	bot.Send(mc)
	return nil
}
