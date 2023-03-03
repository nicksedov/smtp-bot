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
	for i := 0; i < len(settings.Schedule.Once); i++ {
		if i == 1 {
			settings.Schedule.Once[i].Moment = time.Now().Add(time.Second)
		} else {
			settings.Schedule.Once[i].Moment = time.Now().Add(-time.Second)
		}
		settings.Schedule.Once[i].Destination = "testgroup"
	}
	fmt.Printf("Time = %s", settings.Schedule.Once[0].Moment.Local().Format(time.RFC3339))
	Schedule()
}
