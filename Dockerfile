FROM mihaildemidoff/tdlib-go:latest
COPY . /go/src/github.com/marperia/fwdbot
WORKDIR /go/src/github.com/marperia/fwdbot
RUN ["go", "get", "github.com/Arman92/go-tdlib"]
RUN ["go", "build", "main.go"]
CMD ["./main"]