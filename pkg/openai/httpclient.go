package openai

import (
	"net/http"
	"net/url"

	"github.com/nicksedov/sbconn-bot/pkg/cli"
)

var httpClient *http.Client

func GetClient() *http.Client {
	if httpClient == nil {
		transport := &http.Transport{}
		if *cli.ProxyHost != "" {
			proxyUrl := &url.URL{
				Scheme: "http",
				User:   url.UserPassword(*cli.ProxyUser, *cli.ProxyPassword),
				Host:   *cli.ProxyHost,
		  	}
			transport.Proxy = http.ProxyURL(proxyUrl) // set proxy 
		}
		httpClient = &http.Client{Transport: transport}
	}
	return httpClient
}