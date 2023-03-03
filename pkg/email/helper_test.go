package email

import (
	"fmt"
	"net/mail"
	"testing"

	"github.com/nicksedov/sbconn-bot/pkg/cli"
)

func TestExtractEmails(t *testing.T) {
	*cli.FlagConfig = "../../sbconn-settings.yaml"
	emails := make([]*mail.Address, 1)
	emails[0] = &mail.Address{Name: "", Address: "devtools@nsedov.com"}
	res := getEmailAliases(emails)
	fmt.Println(res[0])
}
