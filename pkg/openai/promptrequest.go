package openai

import "github.com/nicksedov/sbconn-bot/pkg/settings"

func GetMessageByPrompt(promptRef string, args... string) string {
	prompt, altText := settings.GetPromptMessage(promptRef, args...)
	
	if prompt != "" {
		resp := SendRequest(0, prompt)
		if len(resp.Choices) > 0 {
			return resp.Choices[0].Message.Content
		} else {
			return altText
		}
	} else {
		return altText
	}
}