package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/alash3al/go-smtpsrv"
	"github.com/nicksedov/smtp-bot/pkg/cli"
	"github.com/nicksedov/smtp-bot/pkg/email"
	"github.com/nicksedov/smtp-bot/pkg/telegram"
)

func main() {
	flag.Parse()
	// Startup telegram bot process in background
	err := telegram.InitBot()
	if err != nil {
		panic(err)
	}
	// Run SMTP server process
	cfg := smtpsrv.ServerConfig{
		ReadTimeout:     time.Duration(*cli.FlagReadTimeout) * time.Second,
		WriteTimeout:    time.Duration(*cli.FlagWriteTimeout) * time.Second,
		ListenAddr:      *cli.FlagListenAddr,
		MaxMessageBytes: int(*cli.FlagMaxMessageSize),
		BannerDomain:    *cli.FlagServerName,
		Handler:         smtpsrv.HandlerFunc(email.SmtpHandler),
	}
	fmt.Println(smtpsrv.ListenAndServe(&cfg))
}
