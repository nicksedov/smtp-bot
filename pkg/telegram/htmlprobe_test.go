package telegram

import (
	"fmt"
	"strconv"
	"testing"
)

var (
	customHtml = "<html>\n" +
		"  <head></head>\n" +
		"  <body>\n" +
		"    <content>value\n" +
		"      <in tag=\"1\">inner</in>\n" +
		"      <in tag=\"2\">inner</in>\n" +
		"    </content>\n" +
		"  </body>\n" +
		"</html>"

	tgHtml = "<html>\n" +
		"  <head></head>\n" +
		"  <body>\n" +
		"    <b>value\n" +
		"      <a href=\"#\">link 1</a>\n" +
		"      <a href=\"http://example.com\">link 2</a>\n" +
		"    </b>\n" +
		"    <strong>value</strong>\n" +
		"  </body>\n" +
		"</html>"
)

func TestGetHtmlBodyContent(t *testing.T) {
	html := GetHtmlBodyContent(customHtml)
	fmt.Printf("HTML <body> content: %s", html)
}

func TestIsHtmlAdaptedForTelegram(t *testing.T) {
	customHtmlSupported := IsHtmlAdaptedForTelegram(GetHtmlBodyContent(customHtml))
	tgHtmlSupported := IsHtmlAdaptedForTelegram(GetHtmlBodyContent(tgHtml))
	fmt.Printf("%s, %s", strconv.FormatBool(customHtmlSupported), strconv.FormatBool(tgHtmlSupported))
}
