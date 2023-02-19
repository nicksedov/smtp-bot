package main

import (
	"strconv"
	"strings"
)

func lookupChatId(addr []*EmailAddress) int64 {
	var settings = GetSettings()
	for _, a := range addr {
		tokens := strings.Split(a.Address, "@")
		if strings.HasPrefix(tokens[0], "chatid") {
			chatId, err := strconv.ParseInt(strings.TrimPrefix(tokens[0], "chatid"), 10, 64)
			if err == nil {
				return chatId
			}
		}
		aliases := settings.Aliases.Chats
		for _, chat := range aliases {
			if chat.Alias == tokens[0] {
				return chat.ChatId
			}
		}
	}
	return 0
}
