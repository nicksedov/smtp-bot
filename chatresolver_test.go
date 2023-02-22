package main

import (
	"fmt"
	"testing"
)

func TestLookupChatId(t *testing.T) {
	*flagConfig = "sbconn-settings.yaml"
	// Check existent aliases
	settings := GetSettings()
	for _, alias := range settings.Aliases.Chats {
		emails := make([]*EmailAddress, 1)
		emails[0] = &EmailAddress{Name: "Has alias " + alias.Alias, Address: alias.Alias + "@mail.com"}
		fmt.Printf("alias='%s': %s\n", alias.Alias, toString(emails))
	}
	// Check non-existent alias
	emails := make([]*EmailAddress, 1)
	emails[0] = &EmailAddress{Name: "Use alias", Address: "fakegroup@mail.com"}
	fmt.Printf("alias='fakegroup': %s\n", toString(emails))
	// Check email without alias
	emails[0] = &EmailAddress{Name: "Use chatId", Address: "chatid-999@mail.com"}
	fmt.Printf("alias='chatid-999': %s\n", toString(emails))
}

func toString(addr []*EmailAddress) string {
	id, capt := lookupChat(addr)
	return fmt.Sprintf("{chatId=%d, caption=%t}", id, capt)
}