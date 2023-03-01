package telegram

import (
	"regexp"

	"github.com/nicksedov/sbconn-bot/pkg/settings"
)

func ContentPreprocessor(source string) string {
	var settings = settings.GetSettings()
	for _, t := range settings.Content.Replace {
		re := regexp.MustCompile(t.Pattern)
		source = re.ReplaceAllString(source, t.Replacement)
	}
	return source
}
