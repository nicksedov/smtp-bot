package scheduler

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/go-yaml/yaml"
	"github.com/nicksedov/sbconn-bot/pkg/cli"
	"github.com/nicksedov/sbconn-bot/pkg/settings"
)

type Secrets struct {
	BotToken string    `yaml:"BotToken"`
	OpenAIToken string `yaml:"OpenAIToken"`
}

func TestSchedule(t *testing.T) {
	*cli.FlagConfig = "../../sbconn-settings.yaml"
	secrets := getSecrets()
	*cli.FlagBotToken = secrets.BotToken
	*cli.FlagOpenAIToken = secrets.OpenAIToken
	settings := settings.GetSettings()
	now := time.Now()
	for i := 0; i < len(settings.Events.Once); i++ {
		if i == 0 {
			settings.Events.Once[i].Moment = now.Add(500 * time.Millisecond)
		} else if i == 1 {
			settings.Events.Once[i].Moment = now.Add(2000 * time.Millisecond)
		} else {
			settings.Events.Once[i].Moment = now
		}
		settings.Events.Once[i].Destination = "testgroup-prompt"
	}
	fmt.Printf("Time = %s", settings.Events.Once[0].Moment.Local().Format(time.RFC3339))
	Schedule()
}

func getSecrets() Secrets {
	secrets := Secrets{}
	yfile, ioErr := ioutil.ReadFile("../../secrets.yaml")
	if ioErr == nil {
		ymlErr := yaml.Unmarshal(yfile, &secrets)
		if ymlErr != nil {
			log.Fatal(ymlErr)
		}
	}
	return secrets
}
