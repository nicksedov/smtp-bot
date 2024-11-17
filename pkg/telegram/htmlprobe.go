package telegram

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func GetHtmlBodyContent(htmlDoc string) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(htmlDoc))
	sel := doc.Find("BODY")
	html, _ := sel.Html()
	return html
}

func IsHtmlAdaptedForTelegram(text string) bool {

	tkn := html.NewTokenizer(strings.NewReader(text))

	var acceptedVals []string = []string{"b", "strong", "i", "em", "u", "ins", "s", "strike", "del",
		"tg-spoiler", "a", "code", "pre"}

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
