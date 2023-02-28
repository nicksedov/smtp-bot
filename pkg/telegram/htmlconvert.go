package telegram

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// TransformHtmlToTelegramCompatible converts an HTML string to a Telegram-compatible HTML string
func TransformHtmlToTelegramCompatible(html string) string {
	// parse the input HTML string into a DOM tree
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	// iterate over all HTML elements in the tree and convert them to their
	// Telegram-compatible equivalents
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		switch s.Get(0).Data {
		case "p", "div":
			html, _ = s.Html()
			s.ReplaceWithHtml("\n" + html + "\n")
		case "span":
			html, _ = s.Html()
			s.ReplaceWithHtml(html)
		case "ul":
			s.Find("li").Each(func(i int, li *goquery.Selection) {
				li.ReplaceWithHtml(fmt.Sprintf("&bull; %s", strings.TrimLeft(li.Text(), " ")))
			})
			html, _ := s.Html()
			s.ReplaceWithHtml(html)
		case "ol":
			s.Find("li").Each(func(i int, li *goquery.Selection) {
				li.ReplaceWithHtml(fmt.Sprintf("%d. %s", i + 1, strings.TrimLeft(li.Text(), " ")))
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
