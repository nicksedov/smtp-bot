package main

import (
	"strconv"
	"strings"
)

func lookupChatId(addr []*EmailAddress) int64 {
	for _, a := range addr {
		tokens := strings.Split(a.Address, "@")
		if strings.HasPrefix(tokens[0], "chatid") {
			chatId, err := strconv.ParseInt(strings.TrimPrefix(tokens[0], "chatid"), 10, 64)
			if err == nil {
				return chatId
			}
		}
		chatIdByAlias := getChatIdByAlias(tokens[0])
		if chatIdByAlias != 0 {
			return chatIdByAlias
		}
	}
	return 0
}

func getChatIdByAlias(token string) int64 {
	var settings = GetSettings()
	aliases := settings.Aliases.Chats
	for _, chat := range aliases {
		if chat.Alias == token {
			return chat.ChatId
		}
	}
	return 0
}

