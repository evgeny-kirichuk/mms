FROM golang:latest
RUN ["go", "install", "github.com/cosmtrek/air@latest"]
WORKDIR /cmd/server
ENTRYPOINT ["air"]