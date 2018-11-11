FROM mihaildemidoff/tdlib-go:latest
COPY . /go/src/tgShareBot
WORKDIR /go/src/tgShareBot
RUN ["go", "get", "github.com/Arman92/go-tdlib"]
RUN ["go", "build", "main.go"]
CMD ["./main"]