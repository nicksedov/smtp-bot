package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/alash3al/go-smtpsrv"
)

var (
	flagServerName     = flag.String("name", "sbconn-bot", "the server name")
	flagListenAddr     = flag.String("listen", ":smtp", "the smtp address to listen on")
	flagMaxMessageSize = flag.Int64("msglimit", 1024*1024*2, "maximum incoming message size")
	flagReadTimeout    = flag.Int("timeout.read", 5, "the read timeout in seconds")
	flagWriteTimeout   = flag.Int("timeout.write", 5, "the write timeout in seconds")
	flagBotToken       = flag.String("bot.token", "", "telegram bot token")
)

func main() {
	flag.Parse()

	cfg := smtpsrv.ServerConfig{
		ReadTimeout:     time.Duration(*flagReadTimeout) * time.Second,
		WriteTimeout:    time.Duration(*flagWriteTimeout) * time.Second,
		ListenAddr:      *flagListenAddr,
		MaxMessageBytes: int(*flagMaxMessageSize),
		BannerDomain:    *flagServerName,
		Handler:         smtpsrv.HandlerFunc(smtpHandler),
	}

	fmt.Println(smtpsrv.ListenAndServe(&cfg))
}
