package email

import (
	"fmt"
	"net/mail"
	"testing"

	"github.com/nicksedov/smtp-bot/pkg/cli"
)

func TestExtractEmails(t *testing.T) {
	*cli.FlagConfig = "../../settings.yaml"
	emails := make([]*mail.Address, 1)
	emails[0] = &mail.Address{Name: "", Address: "omv-admin@sbcon.asuscomm.com"}
	res := getEmailAliases(emails)
	fmt.Println(res[0])
}
