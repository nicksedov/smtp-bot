package main

import (
	"fmt"
	"testing"
)

func TestLookupChatId(t *testing.T) {
	*flagConfig = "sbconn-settings.yaml"
	emails := make([]*EmailAddress, 1)
	emails[0] = &EmailAddress { Name: "Use alias", Address: "botgroup@mail.com" } 
	fmt.Printf("alias='botgroup', chatId=%d\n", lookupChatId(emails))
	emails[0] = &EmailAddress { Name: "Use chatId", Address: "chatid-999@mail.com" } 
	fmt.Printf("alias='chatid-999', chatId=%d\n", lookupChatId(emails))
}