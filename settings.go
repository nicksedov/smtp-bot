package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/go-yaml/yaml"
)

type Settings struct {
    Aliases struct {
        Chats []struct {
            Alias       string `yaml:"alias"`
            ChatId      int64  `yaml:"chatid"`
			Caption     bool   `yaml:"caption"`
        } `yaml:"chats"`
		Emails []struct{
			Alias       string `yaml:"alias"`
			Address     string `yaml:"address"`
		} `yaml:"emails"`
    } `yaml:"aliases"`
	Schedule struct {
		Once []struct {
			Moment 		time.Time `yaml:"moment"`
			Message     string `yaml:"message"`
			Destination string `yaml:"destination"`
		} `yaml:"once"`
	} `yaml:"schedule"`
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