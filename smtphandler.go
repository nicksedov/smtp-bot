package main

import (
    "errors"
    "strconv"
    "strings"
    "github.com/alash3al/go-smtpsrv"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func smtpHandler(c *smtpsrv.Context) error {
    msg, err := c.Parse()
    if err != nil {
	return errors.New("Cannot read your message: " + err.Error())
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

func lookupChatId(addr []*EmailAddress) int64 {
    var chatId int64
    var err error
    for _, a := range addr {
	tokens := strings.Split(a.Address, "@")
	if strings.HasPrefix(tokens[0], "chatid") {
	    chatId, err = strconv.ParseInt(strings.TrimPrefix(tokens[0], "chatid"), 10, 64)
	    if err == nil {
		break
	    }
	}
    }
    return chatId
}

func sendHtml(from string, subj string, html string, chatId int64) {
    file := tgbotapi.FileBytes{
	Name:  "Сообщение.html",
	Bytes: []byte(html),
    }
    doc := tgbotapi.NewDocument(chatId, file)
    doc.Caption = "**Сообщение от:** " + from + "\n**Тема:** " + subj + "\n***"
    doc.ParseMode = "MarkdownV2"
    SendDocumentToChat(doc)
}

func sendText(from string, subj string, text string, chatId int64) {
    caption := "**Сообщение от:** " + from + "\n**Тема:** " + subj + "\n***"
    textMsg := tgbotapi.NewMessage(chatId, caption+text)
    textMsg.ParseMode = "MarkdownV2"
    SendMessageToChat(textMsg)
}
