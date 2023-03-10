package telegram

import (
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nicksedov/sbconn-bot/pkg/cli"
)

var (
	// Menu texts
	firstMenu  = "<b>Menu 1</b>\n\nA beautiful menu with a shiny inline button."
	secondMenu = "<b>Menu 2</b>\n\nA better menu with even more shiny inline buttons."

	// Button texts
	nextButton     = "Next"
	backButton     = "Back"
	tutorialButton = "Tutorial"

	// Store bot screaming status
	screaming = false
	bot       *tgbotapi.BotAPI

	// Keyboard layout for the first menu. One button, one row
	firstMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(nextButton, nextButton),
		),
	)

	// Keyboard layout for the second menu. Two buttons, one per row
	secondMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(backButton, backButton),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(tutorialButton, "https://core.telegram.org/bots/api"),
		),
	)
)

func SendMessageToChat(mc tgbotapi.Chattable) error {
	bot, err := getBotAPI()
	if err == nil {
		_, err = bot.Send(mc)
	}
	return err
}

func getBotAPI() (*tgbotapi.BotAPI, error) {
	if bot == nil {
		var err error
		bot, err = tgbotapi.NewBotAPI(*cli.FlagBotToken)
		if err != nil {
			return nil, fmt.Errorf("cannot create bot API: %w", err)
		}
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := bot.GetUpdatesChan(u)
	
		// Pass cancellable context to goroutine
		go receiveUpdates(updates)
	}
	return bot, nil
}


func receiveUpdates(updates tgbotapi.UpdatesChannel) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		// receive update from channel and then handle it
		update := <-updates
		handleUpdate(update)
		time.Sleep(3 * time.Second)
	}
}

func handleUpdate(update tgbotapi.Update) {
	switch {
	// Handle messages
	case update.Message != nil:
		handleMessage(update.Message)
	// Handle button clicks
	case update.CallbackQuery != nil:
		handleButton(update.CallbackQuery)
	}
}

func handleMessage(message *tgbotapi.Message) {
	user := message.From
	text := message.Text

	if user == nil {
		return
	}

	// Print to console
	log.Printf("%s wrote %s", user.FirstName, text)

	var err error
	if strings.HasPrefix(text, "/") {
		err = handleCommand(message.Chat.ID, text)
	} else if strings.HasPrefix(text, "Re*") {
		if screaming && len(text) > 0 {
			msg := tgbotapi.NewMessage(message.Chat.ID, strings.ToUpper(text))
			// To preserve markdown, we attach entities (bold, italic..)
			msg.Entities = message.Entities
			_, err = bot.Send(msg)
		} else {
			// This is equivalent to forwarding, without the sender's name
			copyMsg := tgbotapi.NewCopyMessage(message.Chat.ID, message.Chat.ID, message.MessageID)
			_, err = bot.CopyMessage(copyMsg)
		}
	}
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}
}

// When we get a command, we react accordingly
func handleCommand(chatId int64, command string) error {
	var err error

	switch command {
	case "/scream":
		screaming = true
	case "/whisper":
		screaming = false
	case "/menu":
		err = sendMenu(chatId)
	}
	return err
}

func handleButton(query *tgbotapi.CallbackQuery) {
	var text string

	markup := tgbotapi.NewInlineKeyboardMarkup()
	message := query.Message

	if query.Data == nextButton {
		text = secondMenu
		markup = secondMenuMarkup
	} else if query.Data == backButton {
		text = firstMenu
		markup = firstMenuMarkup
	}

	callbackCfg := tgbotapi.NewCallback(query.ID, "")
	bot.Send(callbackCfg)

	// Replace menu text and keyboard
	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, markup)
	msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}

func sendMenu(chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, firstMenu)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = firstMenuMarkup
	_, err := bot.Send(msg)
	return err
}

