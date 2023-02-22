package main

import (
	"fmt"
	"net/mail"
	"testing"
)

func TestExtractEmails(t *testing.T) {
	*flagConfig = "sbconn-settings.yaml"
	emails := make([]*mail.Address, 1)
	emails[0] = &mail.Address{Name: "", Address: "devtools@nsedov.com"}
	res := getEmailAliases(emails)
	fmt.Println(res[0])
}
