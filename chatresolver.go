package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/go-yaml/yaml"
)

func lookupChatId(addr []*EmailAddress) int64 {
	var aliases map[string]string
	if (*flagConfig != "") {
		yfile, ioErr := ioutil.ReadFile(*flagConfig)
		if ioErr == nil {
			var yamlTree map[string]map[string]string
			ymlErr := yaml.Unmarshal(yfile, &yamlTree)
			if ymlErr != nil {
				log.Fatal(ymlErr)
			}
			aliases = yamlTree["aliases"]
		}
	}
	for _, a := range addr {
		tokens := strings.Split(a.Address, "@")
		if strings.HasPrefix(tokens[0], "chatid") {
			chatId, err := strconv.ParseInt(strings.TrimPrefix(tokens[0], "chatid"), 10, 64)
			if err == nil {
				return chatId
			}
		}
		if (&aliases != nil) {
			for k, v := range aliases {
				if k == tokens[0] {
					chatId, err := strconv.ParseInt(v, 10, 64)
					if err == nil {
						return chatId
					}
				}
			}
		}
	}
	return 0
}

