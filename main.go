package main

import (
	"fmt"
	"time"

	"github.com/alash3al/go-smtpsrv"
)

func main() {
	cfg := smtpsrv.ServerConfig{
		ReadTimeout:     time.Duration(*flagReadTimeout) * time.Second,
		WriteTimeout:    time.Duration(*flagWriteTimeout) * time.Second,
		ListenAddr:      *flagListenAddr,
		MaxMessageBytes: int(*flagMaxMessageSize),
		BannerDomain:    *flagServerName,
		Handler:         smtpsrv.HandlerFunc(telegramHandler),
	}

	fmt.Println(smtpsrv.ListenAndServe(&cfg))
}
