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
				go func(wg *sync.WaitGroup, promptRef string, msgRef string, msgArgs string) {
					defer wg.Done()
					time.Sleep(duration)
					varArgs := []string{}
					if msgArgs != "" {
						varArgs = strings.Split(msgArgs, ",")
					}
					var message string
					var err error
					if promptRef != "" {
						message = openai.GetMessageByPrompt(promptRef, varArgs...)
					} else {
						message, err = settings.GetMessage(priorityChoice(promptRef, msgRef), varArgs...)
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
				}(&schedWaitGroup, t.PromptRef, t.MessageRef, t.MessageArgs)
			}
		}
	}
	schedWaitGroup.Wait()
}

func priorityChoice(primary string, secondary string) string {
	if primary != "" { 
		return primary 
	} else {
		return secondary
	}
}

func processResponse(resp *openai.ChatResponse, altText string) string {
	choices := resp.Choices
	if len(choices) > 0 {
		return choices[0].Message.Content
	}
	return altText
}