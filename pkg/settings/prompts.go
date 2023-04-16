package settings

import (
	"fmt"
)

var prompts map[string]string
var altTexts map[string]string

func readPromptResource() {
	cfg := GetSettings()

	prompts = make(map[string]string)
	altTexts = make(map[string]string)

	for _, msg := range cfg.Prompts {
		prompts[msg.Id] = msg.Prompt
		altTexts[msg.Id] = msg.AltText
	}
}

func GetPromptMessage(msgId string, args ...string) (string, string) {
	if props == nil {
		readPromptResource()
	}
	promptPattern := prompts[msgId]
	// Generate text for fallback 
	var altText string
	if len(args) == 0 {
		altText = altTexts[msgId]
	} else {
		altText = fmt.Sprintf(promptPattern, args)
	}
	if promptPattern == "" {
		return "", altText
	}
	// Generate prompt text
	if len(args) == 0 {
		return promptPattern, altText
	} else {
		return fmt.Sprintf(promptPattern, args), altText
	}
}