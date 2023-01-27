package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func getHtmlBody(t string) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(t))
	sel := doc.Find("BODY")
	html, _ := sel.Html()
	return html
}

func isSupportedMarkdown(text string) bool {

	tkn := html.NewTokenizer(strings.NewReader(text))

	var acceptedVals []string = []string{"b", "strong", "i", "em", "u", "ins", "s", "strike", "del",
		"span", "tg-spoiler", "a", "code", "pre"}

	for {
		tt := tkn.Next()
		if tt == html.ErrorToken {
			return true
		} else if tt == html.StartTagToken {
			t := tkn.Token()
			tag := strings.ToLower(t.Data)
			unsupported := true
			for _, accepted := range acceptedVals {
				if tag == accepted {
					unsupported = false
					break
				}
			}
			if unsupported {
				return false
			}
		}
	}
}
