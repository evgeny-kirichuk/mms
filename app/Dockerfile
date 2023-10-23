FROM golang:1.21.3-alpine3.18

RUN ["mkdir", "/app"]
ADD . /app
WORKDIR /app
# VOLUME /app

RUN ["go", "install", "github.com/cosmtrek/air@latest"]

CMD ["air"]