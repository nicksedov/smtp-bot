package main

import (
	"errors"
	"log"
	"strings"

	"github.com/alash3al/go-smtpsrv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SmtpHandler(c *smtpsrv.Context) error {
	msg, err := c.Parse()
	if err != nil {
		return errors.New("Cannot read your message: " + err.Error())
	}
	var content string
	var parseMode string
	if strings.HasPrefix(msg.ContentType, "text/plain") {
		content = string(msg.TextBody)
	} else if strings.HasPrefix(msg.ContentType, "text/html") {
		content = string(msg.HTMLBody)
		parseMode = "HTML"
	} else {
		log.Printf("Unknown content type %s", msg.ContentType)
		return nil
	}
	tgMsg := tgbotapi.NewMessage(*flagBotChatId, content)
	if parseMode != "" {
		tgMsg.ParseMode = parseMode
	}
	SendMessageToChat(tgMsg)
	return nil
}
