package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	*flagConfig = "sbconn-settings.yaml"
	*flagBotToken = "5730532194:AAFPluN4ENc64MuiftC076WKcQmUmMH9iBA"
	settings := GetSettings()
	settings.Schedule.Once[0].Moment = time.Now().Add(time.Second)
	fmt.Printf("Time = %s", settings.Schedule.Once[0].Moment.Local().Format(time.RFC3339))
	Schedule()
}