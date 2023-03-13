package openai

import (
	"fmt"
	"testing"
)

func TestSendRequest(t *testing.T) {
	//Must initialize *cli.FlagOpenAIToken
	resp := SendRequest(5093432423, "Hello buddy!")
	choices := resp.Choices
	if len(choices) > 0 {
		fmt.Printf("%s answered: %s", choices[0].Message.Role, choices[0].Message.Content)
	} else {
		fmt.Println("Test failed")
	}
}
