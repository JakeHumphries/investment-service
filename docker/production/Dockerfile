FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 \
    GO111MODULE=on

WORKDIR /service

COPY ./ /service

RUN go build -o main /service/cmd/main.go

FROM alpine:latest

WORKDIR /service

COPY --from=builder /service/main /service/main

EXPOSE 8080

CMD ["/service/main"]
