package openai

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/go-yaml/yaml"
	"github.com/nicksedov/sbconn-bot/pkg/cli"
)

type Secrets struct {
	BotToken string      `yaml:"BotToken"`
	OpenAIToken string   `yaml:"OpenAIToken"`
	ProxyHost string     `yaml:"ProxyHost"`
	ProxyUser string     `yaml:"ProxyUser"`
	ProxyPassword string `yaml:"ProxyPassword"`

}

func TestSendRequest(t *testing.T) {
	*cli.FlagConfig = "../../sbconn-settings.yaml"
	secrets := getSecrets()
	*cli.FlagBotToken = secrets.BotToken
	*cli.FlagOpenAIToken = secrets.OpenAIToken
	*cli.ProxyHost = secrets.ProxyHost
	*cli.ProxyUser = secrets.ProxyUser
	*cli.ProxyPassword = secrets.ProxyPassword

	//Must initialize *cli.FlagOpenAIToken
	resp := SendRequest(5093432423, "Hello buddy!")
	fmt.Printf("Response ID is %s\n", resp.ID)
	choices := resp.Choices
	if len(choices) > 0 {
		fmt.Printf("%s answered:\n - %s", choices[0].Message.Role, choices[0].Message.Content)
	} else {
		fmt.Println("Test failed, ")
	}
}

func getSecrets() Secrets {
	secrets := Secrets{}
	yfile, ioErr := os.ReadFile("../../secrets.yaml")
	if ioErr == nil {
		ymlErr := yaml.Unmarshal(yfile, &secrets)
		if ymlErr != nil {
			log.Fatal(ymlErr)
		}
	}
	return secrets
}
