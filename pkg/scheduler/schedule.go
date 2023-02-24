package scheduler

import (
	"sync"
	"time"

	"github.com/nicksedov/sbconn-bot/pkg/settings"
	"github.com/nicksedov/sbconn-bot/pkg/email"

)

func Schedule() {
	var schedWaitGroup sync.WaitGroup
	var settings = settings.GetSettings()
	for _, t := range settings.Schedule.Once {
		chatId, needsCaption := email.GetChatIdByAlias(t.Destination)
		if chatId != 0 {
			duration := time.Until(t.Moment)
			if duration > 0 {
				schedWaitGroup.Add(1)
				message := t.Message
				go func(wg *sync.WaitGroup) {
					defer wg.Done()
					time.Sleep(duration)
					if (needsCaption) {
						email.SendTextWithCaption("Jenkins", "Напоминание", message, chatId)
					} else {
						email.SendText(message, chatId)
					}
				}(&schedWaitGroup)
			}
		}
	}
	schedWaitGroup.Wait()
}
