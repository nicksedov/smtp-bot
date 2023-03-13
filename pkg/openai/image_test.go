package openai

import (
	"fmt"
	"testing"

	"github.com/nicksedov/sbconn-bot/pkg/cli"
)

func TestSendImageRequest(t *testing.T) {
	*cli.FlagOpenAIToken = "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"
	prompt := "Ugly cow is screaming"
	resp := SendImageRequest(prompt)
	choices := resp.Data
	if len(choices) > 0 {
		fmt.Printf("Picture %s available by url: [%s]", prompt, choices[0].Url)
	} else {
		fmt.Println("Test failed")
	}
}