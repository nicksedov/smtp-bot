sbconn-bot (email-to-web)
========================
sbconn-bot is a simple smtp server that resends the incoming email to the telegram chat.

Dev 
===
- `go mod vendor`
- `go build`

Production
==========
using go default compiler
- `go build -ldflags="-s -w"`

using gccgo
- `go build -compiler gccgo -gccgoflags "-march=native -O3" main.go`

Dev with Docker
==============
Locally :
- `go mod vendor`
- `docker build -f Dockerfile.dev -t sbconn-bot-dev .`
- `docker run -p 25:25 sbconn-bot-dev --timeout.read=50 --timeout.write=50 --bot.token=<token>`

Or build it as it comes from the repo :
- `docker build -t sbconn-bot .`
- `docker run -p 25:25 sbconn-bot --timeout.read=50 --timeout.write=50 --bot.token=<token>`

The `timeout` options are of course optional but make it easier to test in local with `telnet localhost 25`
Here is a telnet example payload : 
```
HELO zeus
# smtp answer

MAIL FROM:<email@from.com>
# smtp answer

RCPT TO:<youremail@example.com>
# smtp answer

DATA
your mail content
.

```

Docker (production)
=====
**Docker images arn't available online for now**
**See "Dev with Docker" above**
- `docker run -p 25:25 sbconn-bot --bot.token=<token>`

Native usage
=====
`sbconn-bot --listen=:25 --bot.token=<token>`
`sbconn-bot --help`

Contribution
============
Based on original repo from @alash3al
Thanks to @aranajuan


