package email

import (
	"net/mail"

	"github.com/nicksedov/smtp-bot/pkg/settings"
)

func getEmailAliases(addr []*mail.Address, _ ...error) []string {
	ret := []string{}
	var settings = settings.GetSettings()
	for _, e := range addr {
		if e.Name != "" {
			ret = append(ret, e.Name)
		} else {
			aliases := settings.Aliases.Emails
			address := e.Address
			for _, alias := range aliases {
				if e.Address == alias.Address {
					address = alias.Alias
					break
				}

			}
			ret = append(ret, address)
		}
	}

	return ret
}

func transformStdAddressToEmailAddress(addr []*mail.Address) []*EmailAddress {
	ret := []*EmailAddress{}

	for _, e := range addr {
		ret = append(ret, &EmailAddress{
			Address: e.Address,
			Name:    e.Name,
		})
	}

	return ret
}
