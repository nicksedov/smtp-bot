package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func GetBotAPI() (*tgbotapi.BotAPI, error) {
	if bot == nil {
		var err error
		bot, err = tgbotapi.NewBotAPI(*flagBotToken)
		if err != nil {
			return nil, fmt.Errorf("cannot create bot API: %w", err)
		}
	}
	return bot, nil
}

func SendMessageToChat(mc tgbotapi.Chattable) error {
	bot, err := GetBotAPI()
	if err == nil {
		_, err = bot.Send(mc)
	}
	return err
}