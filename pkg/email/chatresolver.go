package email

import (
	"strconv"
	"strings"

	"github.com/nicksedov/sbconn-bot/pkg/settings"
)

func lookupChat(addr []*EmailAddress) (int64, bool) {
	for _, a := range addr {
		tokens := strings.Split(a.Address, "@")
		if strings.HasPrefix(tokens[0], "chatid") {
			chatId, err := strconv.ParseInt(strings.TrimPrefix(tokens[0], "chatid"), 10, 64)
			if err == nil {
				return chatId, true
			}
		}
		chatIdByAlias, needsCaption := getChatIdByAlias(tokens[0])
		if chatIdByAlias != 0 {
			return chatIdByAlias, needsCaption
		}
	}
	return 0, false
}

func getChatIdByAlias(token string) (int64, bool) {
	var settings = settings.GetSettings()
	aliases := settings.Aliases.Chats
	for _, chat := range aliases {
		if chat.Alias == token {
			return chat.ChatId, chat.Caption
		}
	}
	return 0, false
}
