package openai

import (
	"fmt"
	"testing"
)

func TestSendImageRequest(t *testing.T) {
	//Must initialize *cli.FlagOpenAIToken
	prompt := "Ugly cow is screaming"
	resp := SendImageRequest(prompt)
	choices := resp.Data
	if len(choices) > 0 {
		fmt.Printf("Picture %s available by url: [%s]", prompt, choices[0].Url)
	} else {
		fmt.Println("Test failed")
	}
}