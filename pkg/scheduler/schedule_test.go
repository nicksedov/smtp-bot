package scheduler

import (
	"fmt"
	"testing"
	"time"

	"github.com/nicksedov/sbconn-bot/pkg/cli"
	"github.com/nicksedov/sbconn-bot/pkg/settings"
)

func TestSchedule(t *testing.T) {
	*cli.FlagConfig = "../../sbconn-settings.yaml"
	*cli.FlagBotToken = "5730532194:AAFPluN4ENc64MuiftC076WKcQmUmMH9iBA"
	settings := settings.GetSettings()
	now := time.Now()
	for i := 0; i < len(settings.Events.Once); i++ {
		if i == 0 {
			settings.Events.Once[i].Moment = now.Add(500 * time.Millisecond)
		} else if i == 1 {
			settings.Events.Once[i].Moment = now.Add(1000 * time.Millisecond)
		} else {
			settings.Events.Once[i].Moment = now
		}
		settings.Events.Once[i].Destination = "testgroup"
	}
	fmt.Printf("Time = %s", settings.Events.Once[0].Moment.Local().Format(time.RFC3339))
	Schedule()
}
