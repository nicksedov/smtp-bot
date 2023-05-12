package scheduler

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/nicksedov/sbconn-bot/pkg/email"
	"github.com/nicksedov/sbconn-bot/pkg/openai"
	"github.com/nicksedov/sbconn-bot/pkg/settings"
)

func Schedule() {
	var schedWaitGroup sync.WaitGroup
	var cfg = settings.GetSettings()
	for _, t := range cfg.Events.Once {
		chatId, needsCaption := email.GetChatIdByAlias(t.Destination)
		if chatId != 0 {
			duration := time.Until(t.Moment)
			if duration > 0 {
				schedWaitGroup.Add(1)
				go func(wg *sync.WaitGroup, t settings.Event) {
					defer wg.Done()
					time.Sleep(duration)
					varArgs := []string{}
					if t.MessageArgs != "" {
						varArgs = strings.Split(t.MessageArgs, ",")
					}
					var message string
					var err error
					if t.PromptRef != "" {
						message = openai.GetMessageByPrompt(t.PromptRef, varArgs...)
					} else {
						message, err = settings.GetMessage(t.MessageRef, varArgs...)
					}
					if err != nil {
						log.Println(err)
					} else {
						if needsCaption {
							email.SendTextWithCaption("Jenkins", "Напоминание", message, chatId)
						} else {
							email.SendText(message, chatId)
						}
					}
				}(&schedWaitGroup, t)
			}
		}
	}
	schedWaitGroup.Wait()
}