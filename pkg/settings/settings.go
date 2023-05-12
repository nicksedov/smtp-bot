package settings

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/go-yaml/yaml"
	"github.com/nicksedov/sbconn-bot/pkg/cli"
)

type Settings struct {
	Aliases struct {
		Chats []struct {
			Alias   string `yaml:"alias"`
			ChatId  int64  `yaml:"chatid"`
			Caption bool   `yaml:"caption"`
		} `yaml:"chats"`
		Emails []struct {
			Alias   string `yaml:"alias"`
			Address string `yaml:"address"`
		} `yaml:"emails"`
	} `yaml:"aliases"`
	Content struct {
		Replace []struct {
			Pattern     string `yaml:"pattern"`
			Replacement string `yaml:"replacement"`
		} `yaml:"replace"`
	} `yaml:"content"`
	Events struct {
		Once []Event `yaml:"once"`
	} `yaml:"events"`
	Messages []struct {
		Id   string `yaml:"id"`
		Text string `yaml:"text"`
	} `yaml:"messages"`
	Prompts []struct {
		Id      string `yaml:"id"`
		Prompt  string `yaml:"prompt"`
		AltText string `yaml:"altText"`
	} `yaml:"prompts"`
}

type Event struct {
	Moment      time.Time `yaml:"moment"`
	PromptRef   string    `yaml:"promptRef"`
	MessageRef  string    `yaml:"messageRef"`
	MessageArgs string    `yaml:"messageArgs"`
	Destination string    `yaml:"destination"`
}

var settings Settings = Settings{}
var initialized bool = false

func GetSettings() *Settings {
	if !initialized {
		if *cli.FlagConfig != "" {
			yfile, ioErr := ioutil.ReadFile(*cli.FlagConfig)
			if ioErr == nil {
				ymlErr := yaml.Unmarshal(yfile, &settings)
				if ymlErr != nil {
					log.Fatal(ymlErr)
				}
			}
		}
		initialized = true
	}
	return &settings
}
