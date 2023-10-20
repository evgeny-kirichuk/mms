# Build the manager binary
FROM golang:1.13 as builder

COPY . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /build/goapp ./cmd/goapp

FROM alpine:3.10

COPY --from=builder /build/goapp /usr/local/bin/goapp
RUN chmod +x /usr/local/bin/goapp

CMD tail -f /dev/null
