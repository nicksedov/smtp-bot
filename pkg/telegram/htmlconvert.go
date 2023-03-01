package telegram

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// TryMakeHtmlTelegramCompatible converts an HTML string to a Telegram-compatible HTML string
func TryMakeHtmlTelegramCompatible(html string) string {
	// parse the input HTML string into a DOM tree
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	// iterate over all HTML elements in the tree and convert them to their
	// Telegram-compatible equivalents
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		node := s.Get(0)
		switch node.Data {
		case "p", "div", "br":
			html, _ = s.Html()
			s.ReplaceWithHtml("\n" + html + "\n")
		case "h1", "h2", "h3", "h4", "h5", "h6", "big":
			html, _ = s.Html()
			s.ReplaceWithHtml("<b>" + html + "</b>")
		case "span":
			html, _ = s.Html()
			s.ReplaceWithHtml(html)
		case "ul":
			s.Find("li").Each(func(i int, li *goquery.Selection) {
				html, _ = li.Html()
				li.ReplaceWithHtml(fmt.Sprintf("&bull; %s", html))
			})
			html, _ := s.Html()
			s.ReplaceWithHtml(html)
		case "ol":
			s.Find("li").Each(func(i int, li *goquery.Selection) {
				html, _ = li.Html()
				li.ReplaceWithHtml(fmt.Sprintf("%d. %s", i+1, html))
			})
			html, _ := s.Html()
			s.ReplaceWithHtml(html)
		}
	})

	// return the converted HTML string
	newHtml, _ := doc.Html()
	re := regexp.MustCompile(`\n\s*\n`)
	newHtml = re.ReplaceAllString(newHtml, "\n")
	return newHtml
}
