package main

import (
	"fmt"
	"html"
	"strings"

	"github.com/alash3al/go-smtpsrv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func smtpHandler(c *smtpsrv.Context) error {
	msg, err := c.Parse()
	if err != nil {
		return fmt.Errorf("cannot read your message: %w", err)
	}
	to := transformStdAddressToEmailAddress(msg.To)
	chatId, needsCaption := lookupChat(to)
	if chatId == 0 {
		cc := transformStdAddressToEmailAddress(msg.Cc)
		chatId, needsCaption = lookupChat(cc)
	}
	if chatId == 0 {
		bcc := transformStdAddressToEmailAddress(msg.Bcc)
		chatId, needsCaption = lookupChat(bcc)
	}

	from := decodeRFC2047(strings.Join(getEmailAliases(msg.From), "; "))
	subj := decodeRFC2047(msg.Subject)
	if msg.HTMLBody != "" {
		sendHtml(from, subj, msg.HTMLBody, chatId)
	} else if msg.TextBody != "" {
		text := strings.Split(msg.TextBody, "<!--- END OF DOCUMENT --->")[0]
		if needsCaption {
			sendTextWithCaption(from, subj, text, chatId)
		} else {
			sendText(text, chatId)
		}
	}
	return nil
}

func sendHtml(from string, subj string, htmlDoc string, chatId int64) {
	htmlBody := getHtmlBody(htmlDoc)
	if isTelegramCompatibleHtml(htmlBody) {
		htmlFrom := html.EscapeString(from)
		htmlSubj := html.EscapeString(subj)
		msgText := fmt.Sprintf("<b>Сообщение от:</b> %s\n<b>Тема:</b> %s\n%s", htmlFrom, htmlSubj, htmlBody)
		chattable := tgbotapi.NewMessage(chatId, msgText)
		chattable.ParseMode = "HTML"
		SendMessageToChat(chattable)
	} else {
		file := tgbotapi.FileBytes{
			Name:  "Сообщение.html",
			Bytes: []byte(htmlDoc),
		}
		chattable := tgbotapi.NewDocument(chatId, file)
		chattable.Caption = fmt.Sprintf("*Сообщение от:* %s\n*Тема:* %s\n", from, subj)
		chattable.ParseMode = "markdown"
		SendMessageToChat(chattable)
	}
}

func sendText(content string, chatId int64) {
	chattable := tgbotapi.NewMessage(chatId, content)
	chattable.ParseMode = "markdown"
	SendMessageToChat(chattable)
}

func sendTextWithCaption(from string, subj string, content string, chatId int64) {
	msgText := fmt.Sprintf("*Сообщение от:* %s\n*Тема:* %s\n%s", from, subj, content)
	sendText(msgText, chatId)
}
