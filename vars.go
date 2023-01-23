package main

import "flag"

var (
	flagServerName     = flag.String("name", "sbconn-bot", "the server name")
	flagListenAddr     = flag.String("listen", ":smtp", "the smtp address to listen on")
	flagMaxMessageSize = flag.Int64("msglimit", 1024*1024*2, "maximum incoming message size")
	flagReadTimeout    = flag.Int("timeout.read", 5, "the read timeout in seconds")
	flagWriteTimeout   = flag.Int("timeout.write", 5, "the write timeout in seconds")
	flagAuthUSER       = flag.String("user", "", "user for smtp client")
	flagAuthPASS       = flag.String("pass", "", "pass for smtp client")
	flagDomain         = flag.String("domain", "", "domain for recieving mails")
	flagBotToken       = flag.String("bot.token", "", "telegram bot token")
)

func init() {
	flag.Parse()
}
