package main

import (
	"io/ioutil"
	"log"
	"github.com/go-yaml/yaml"
)

type Settings struct {
    Aliases struct {
        Chats []struct {
            Alias       string `yaml:"alias"`
            ChatId      int64  `yaml:"chatid"`
        } `yaml:"chats"`
		Emails []struct{
			Alias       string `yaml:"alias"`
			Address     string `yaml:"address"`
		} `yaml:"emails"`
    } `yaml:"aliases"`
}

var settings Settings = Settings{}
var initialized bool = false

func GetSettings() Settings {
	if !initialized {
		if *flagConfig != "" {
			yfile, ioErr := ioutil.ReadFile(*flagConfig)
			if ioErr == nil {
				ymlErr := yaml.Unmarshal(yfile, &settings)
				if ymlErr != nil {
					log.Fatal(ymlErr)
				}
			}
		}
		initialized = true
	}
	return settings
}