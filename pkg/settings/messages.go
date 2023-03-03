package settings

import (
	"fmt"
)

var props map[string]string

func readMessageResource() {
	cfg := GetSettings()

	props = make(map[string]string)
	for _, msg := range cfg.Messages {
		props[msg.Id] = msg.Text
	}
}

func GetMessage(msgId string, args ...string) (string, error) {
	if props == nil {
		readMessageResource()
	}
	msgPattern := props[msgId]
	if msgPattern == "" {
		return "", fmt.Errorf("invalid message resource: %s", msgId) 
	}
	if len(args) == 0 {
		return msgPattern, nil
	} else {
		return fmt.Sprintf(msgPattern, args), nil
	}
}