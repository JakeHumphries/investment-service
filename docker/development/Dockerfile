FROM golang:1.23-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0

RUN apk add git \
    && go install github.com/githubnemo/CompileDaemon@1.3.0

WORKDIR /service

COPY ./ /service

ENTRYPOINT CompileDaemon \
    -build="go build -o /service/main /service/cmd/main.go" \
    -command="/service/main" \
    -directory="cmd" \
    -directory="service" \
    -directory="investment" \
    -directory="database" \
    -directory="graph" \
    -directory="models" \
    -directory="config" \
    -exclude-dir=".git" \
    -exclude-dir=".idea"



