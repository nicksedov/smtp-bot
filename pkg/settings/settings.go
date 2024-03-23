package settings

import (
	"log"

	"os"

	"github.com/go-yaml/yaml"
	"github.com/nicksedov/smtp-bot/pkg/cli"
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
}

var settings Settings = Settings{}
var initialized bool = false

func GetSettings() *Settings {
	if !initialized {
		if *cli.FlagConfig != "" {
			yfile, ioErr := os.ReadFile(*cli.FlagConfig)
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
