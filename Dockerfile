FROM golang:1.18-alpine AS builder
WORKDIR /getblockio
COPY . .
RUN go build -o .bin/main cmd/main.go
ENTRYPOINT [ "/getblockio/.bin/main" ]