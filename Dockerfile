FROM golang:1.14.4 as builder
RUN git clone https://github.com/alash3al/sbconn-bot /go/src/build
WORKDIR /go/src/build
RUN go mod vendor
ENV CGO_ENABLED=0
RUN GOOS=linux go build -mod vendor -a -o sbconn-bot .

FROM golang:1.16.2  
WORKDIR /root/
COPY --from=builder /go/src/build/sbconn-bot /usr/bin/sbconn-bot
ENTRYPOINT ["sbconn-bot"]
