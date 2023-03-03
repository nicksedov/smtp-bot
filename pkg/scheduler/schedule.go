package scheduler

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/nicksedov/sbconn-bot/pkg/settings"
	"github.com/nicksedov/sbconn-bot/pkg/email"
)

func Schedule() {
	var schedWaitGroup sync.WaitGroup
	var cfg = settings.GetSettings()
	for _, t := range cfg.Schedule.Once {
		chatId, needsCaption := email.GetChatIdByAlias(t.Destination)
		if chatId != 0 {
			duration := time.Until(t.Moment)
			if duration > 0 {
				schedWaitGroup.Add(1)
				go func(wg *sync.WaitGroup, msgRef string, msgArgs string) {
					defer wg.Done()
					time.Sleep(duration)
					varArgs := []string{}
					if msgArgs != "" {
						varArgs = strings.Split(msgArgs, ",")
					}
					message, error := settings.GetMessage(msgRef, varArgs...)
					if error != nil {
						log.Println(error)
					} else {
						if (needsCaption) {
							email.SendTextWithCaption("Jenkins", "Напоминание", message, chatId)
						} else {
							email.SendText(message, chatId)
						}
					}
				}(&schedWaitGroup, t.MessageRef, t.MessageArgs)
			}
		}
	}
	schedWaitGroup.Wait()
}
