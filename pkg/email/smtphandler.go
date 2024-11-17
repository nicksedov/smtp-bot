package email

import (
	"fmt"
	"html"
	"strings"

	"github.com/alash3al/go-smtpsrv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nicksedov/smtp-bot/pkg/telegram"
)

const (
	// Envelope sign
	from_unicode string = "\u2709"  
	from_html string = "&#x2709;"
	// Double quote sign
	subj_unicode string = "\u275D"
	subj_html string = "&#x275D;"
)

func SmtpHandler(c *smtpsrv.Context) error {
	msg, err := c.Parse()
	if err != nil {
		return fmt.Errorf("cannot read your message: %w", err)
	}
	to := transformStdAddressToEmailAddress(msg.To)
	cc := transformStdAddressToEmailAddress(msg.Cc)
	bcc := transformStdAddressToEmailAddress(msg.Bcc)
	addresses := append(append(to, cc...), bcc...)
	chatId, needsCaption := lookupChat(addresses)
	if chatId == 0 {
		return nil
	}

	from := decodeRFC2047(strings.Join(getEmailAliases(msg.From), "; "))
	subj := decodeRFC2047(msg.Subject)
	if msg.HTMLBody != "" {
		sendHtml(from, subj, msg.HTMLBody, chatId)
	} else if msg.TextBody != "" {
		text := msg.TextBody
		if needsCaption {
			sendTextWithCaption(from, subj, text, chatId)
		} else {
			sendText(text, chatId)
		}
	}
	return nil
}

func sendHtml(from string, subj string, htmlDoc string, chatId int64) {
	tgCompatibleHtmlDoc := telegram.TryAdaptHtmlForTelegram(htmlDoc)
	htmlBody := telegram.GetHtmlBodyContent(tgCompatibleHtmlDoc)
	if telegram.IsHtmlAdaptedForTelegram(htmlBody) {
		htmlFrom := html.EscapeString(from)
		htmlSubj := html.EscapeString(subj)
		msgText := fmt.Sprintf("%s %s\n %s  %s\n%s", from_html, htmlFrom, subj_html, htmlSubj, htmlBody)
		chattable := tgbotapi.NewMessage(chatId, msgText)
		chattable.ParseMode = "HTML"
		telegram.SendMessageToChat(chattable)
	} else {
		file := tgbotapi.FileBytes{
			Name:  "Сообщение.html",
			Bytes: []byte(htmlDoc),
		}
		chattable := tgbotapi.NewDocument(chatId, file)
		chattable.Caption = fmt.Sprintf("%s %s\n %s  %s\n", from_unicode, from, subj_unicode, subj)
		chattable.ParseMode = "markdown"
		telegram.SendMessageToChat(chattable)
	}
}

func sendText(content string, chatId int64) {
	chattable := tgbotapi.NewMessage(chatId, content)
	chattable.ParseMode = "markdown"
	telegram.SendMessageToChat(chattable)
}

func sendTextWithCaption(from string, subj string, content string, chatId int64) {
	msgText := fmt.Sprintf("%s %s\n %s  %s\n%s", from_unicode, from, subj_unicode, subj, content)
	sendText(msgText, chatId)
}
