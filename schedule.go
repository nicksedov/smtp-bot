package main

import (
	"sync"
	"time"
)

func Schedule() {
	var schedWaitGroup sync.WaitGroup
	var settings = GetSettings()
	for _, t := range settings.Schedule.Once {
		chatId := getChatIdByAlias(t.Destination)
		if chatId != 0 {
			duration := time.Until(t.Moment)
			if duration > 0 {
				schedWaitGroup.Add(1)
				message := t.Message
        		go func(wg *sync.WaitGroup) {
            		defer wg.Done()
					time.Sleep(duration)
					sendText("Календарь", "Событие", message, chatId)
        		}(&schedWaitGroup)
			}
		}
	}
	schedWaitGroup.Wait()
}