package telegram

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nicksedov/sbconn-bot/pkg/cli"
	"github.com/nicksedov/sbconn-bot/pkg/openai"
	"github.com/nicksedov/sbconn-bot/pkg/settings"
)

var bot *tgbotapi.BotAPI

func InitBot() error {
	var err error
	bot, err = tgbotapi.NewBotAPI(*cli.FlagBotToken)
	if err != nil {
		return fmt.Errorf("cannot create bot API: %w", err)
	}
	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 60
	go updatesListener(bot.GetUpdatesChan(upd))
	return nil
}

func SendMessageToChat(mc tgbotapi.Chattable) error {
	bot, err := getOrInitBot()
	if err == nil {
		_, err = bot.Send(mc)
	}
	return err
}

func getOrInitBot() (*tgbotapi.BotAPI, error) {
	if bot == nil {
		err := InitBot()
		if err != nil {
			return nil, fmt.Errorf("cannot create bot API: %w", err)
		}
	}
	return bot, nil
}

func updatesListener(updates tgbotapi.UpdatesChannel) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		// execution thread locks until event received
		update := <-updates
		if update.Message != nil {
			handleMessage(update.Message)
		}
	}
}

func handleMessage(message *tgbotapi.Message) {
	text := message.Text
	chatId := message.Chat.ID
	if chatId == 0 {
		return
	}
	text = strings.TrimSpace(text)
	if strings.HasPrefix(text, "/") {
		handleCommand(chatId, text)
	} else {
		processChat(chatId, text)
	}
}

// When we get a command, we react accordingly
func handleCommand(chatId int64, command string) {
	instruction, args, argsFound := cut(command, " ")
	instruction = strings.ToLower(instruction)
	if !argsFound {
		switch instruction {
		case "/help", "/about":
			processHelp(chatId)
		default:
			msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("Unsupported command '%s'", command))
			bot.Send(msg)
		}
	} else {
		switch instruction {
		case "/draw":
			processDraw(chatId, args)
		case "/chat":
			processChat(chatId, args)
		default:
			msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("Unsupported command '%s'", command))
			bot.Send(msg)
		}
	}
}

func processChat(chatId int64, prompt string) {
	resp := openai.SendRequest(chatId, prompt)
	if len(resp.Choices) > 0 {
		msg := tgbotapi.NewMessage(chatId, resp.Choices[0].Message.Content)
		bot.Send(msg)
	}
}

func processDraw(chatId int64, prompt string) {
	resp := openai.SendImageRequest(prompt)
	if len(resp.Data) > 0 {
		url := resp.Data[0].Url
		msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("<a href='%s'>&#8205;</a>%s", url, prompt))
		msg.ParseMode = "HTML"
		msg.DisableWebPagePreview = false
		bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("Got empty response with HTTP status %d", resp.HttpStatus))
		bot.Send(msg)
	}
}

func processHelp(chatId int64) {
	tags := []string{
		"about.me", 
		"about.groupchat.details", 
		"about.privatechat.details",
		"about.draw.details"}
	for _, msgRef := range tags {
		text, err := settings.GetMessage(msgRef)
		if err != nil {
			text = err.Error()
		}
		msg := tgbotapi.NewMessage(chatId, text)
		bot.Send(msg)
	}
}

func cut(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 && i+len(sep) < len(s) {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}