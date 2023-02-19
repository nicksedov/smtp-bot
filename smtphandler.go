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
	chatId := lookupChatId(to)
	if chatId == 0 {
		cc := transformStdAddressToEmailAddress(msg.Cc)
		chatId = lookupChatId(cc)
	}
	if chatId == 0 {
		bcc := transformStdAddressToEmailAddress(msg.Bcc)
		chatId = lookupChatId(bcc)
	}

	from := strings.Join(extractEmails(msg.From), ", ")
	if msg.HTMLBody != "" {
		sendHtml(from, msg.Subject, msg.HTMLBody, chatId)
	} else if msg.TextBody != "" {
		text := strings.Split(msg.TextBody, "<!--- END OF DOCUMENT --->")[0]
		sendText(from, msg.Subject, text, chatId)
	}
	return nil
}

func sendHtml(from string, subj string, htmlDoc string, chatId int64) {
	htmlBody := getHtmlBody(htmlDoc)
	if isTelegramCompatibleHtml(htmlBody) {
		caption := fmt.Sprintf("<b>Сообщение от:</b> %s\n<b>Тема:</b> %s\n", html.EscapeString(from), html.EscapeString(subj))
		text := caption + htmlBody
		textMsg := tgbotapi.NewMessage(chatId, text)
		textMsg.ParseMode = "HTML"
		SendMessageToChat(textMsg)
	} else {
		file := tgbotapi.FileBytes{
			Name:  "Сообщение.html",
			Bytes: []byte(htmlDoc),
		}
		doc := tgbotapi.NewDocument(chatId, file)
		doc.Caption = fmt.Sprintf("*Сообщение от:* %s\n*Тема:* %s\n", from, subj)
		doc.ParseMode = "markdown"
		SendDocumentToChat(doc)
	}
}

func sendText(from string, subj string, text string, chatId int64) {
	caption := fmt.Sprintf("*Сообщение от:* %s\n*Тема:* %s\n", from, subj)
	textMsg := tgbotapi.NewMessage(chatId, caption+text)
	textMsg.ParseMode = "markdown"
	SendMessageToChat(textMsg)
}
