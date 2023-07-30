# build
FROM golang:1.20.6-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
COPY rabbitmq ./rabbitmq
COPY producer ./producer

RUN go build -ldflags="-s -w" -o producer ./producer/main.go

# publish
FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/producer .

ENV GIN_MODE=release

ENTRYPOINT ["/app/producer"]