package main

import (
	"net/mail"
	"fmt"
	"testing"
)

func TestExtractEmails(t *testing.T) {
	*flagConfig = "sbconn-settings.yaml"
	emails := make([]*mail.Address, 1)
	emails[0] = &mail.Address { Name: "", Address: "devtools@nsedov.com" } 
	res := extractEmails(emails)
	fmt.Printf("%s", res[0])
}