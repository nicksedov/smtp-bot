package main

import (
	"errors"
	"log"

	"github.com/alash3al/go-smtpsrv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func telegramHandler(c *smtpsrv.Context) error {
	_, err := c.Parse()
	if err != nil {
		return errors.New("Cannot read your message: " + err.Error())
	}
	bot, err := tgbotapi.NewBotAPI(*flagBotToken)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	resp := tgbotapi.NewMessage(*flagBotChatId, "Получено сообщение")
	bot.Send(resp)
	return nil
}
