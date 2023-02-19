package main

import (
	"fmt"
	"net/mail"
)

func extractEmails(addr []*mail.Address, _ ...error) []string {
	ret := []string{}
	var settings = GetSettings()
	for _, e := range addr {
		if e.Name != "" {
			ret = append(ret, fmt.Sprintf("%s <%s>", e.Name, e.Address))
		} else {
			aliases := settings.Aliases.Emails
			address := e.Address
			for _, alias := range aliases {
				if e.Address == alias.Address {
					address = fmt.Sprintf("%s <%s>", alias.Alias, alias.Address)
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
