package cli

import (
	"flag"
)

var (
	FlagServerName     = flag.String("name", "sbconn-bot", "the server name")
	FlagListenAddr     = flag.String("listen", ":smtp", "the smtp address to listen on")
	FlagMaxMessageSize = flag.Int64("msglimit", 1024*1024*2, "maximum incoming message size")
	FlagReadTimeout    = flag.Int("timeout.read", 5, "the read timeout in seconds")
	FlagWriteTimeout   = flag.Int("timeout.write", 5, "the write timeout in seconds")
	FlagBotToken       = flag.String("bot.token", "", "telegram bot token")
	FlagConfig         = flag.String("config", "sbconn-settings.yaml", "configuration YAML file")
	FlagOpenAIToken    = flag.String("openai.token", "", "OpenAI token")
)
