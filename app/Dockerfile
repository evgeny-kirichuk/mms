# # Compile stage
# FROM golang:1.17 AS build-env
# ADD . /dockerdev
# WORKDIR /dockerdev
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /server ./cmd/goapp

# # Final stage
# FROM debian:buster
# EXPOSE 8000
# WORKDIR /
# COPY --from=build-env /server /
# CMD ["/server"]

FROM golang:1.18

# Copy application data into image
COPY . /go/src/app
WORKDIR /go/src/app

# COPY go.mod ./
# COPY go.sum ./
RUN go mod download

# Copy only `.go` files, if you want all files to be copied then replace `with `COPY . .` for the code below.
# COPY . .

# Install our third-party application for hot-reloading capability.
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "install", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -polling -log-prefix=false -build="go build ." -command="./app" -directory="./"