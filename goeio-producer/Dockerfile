# build
FROM golang:1.20.6-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
COPY goeio-rabbitmq ./goeio-rabbitmq
COPY goeio-producer ./goeio-producer

RUN go build -ldflags="-s -w" -o producer ./goeio-producer/main.go

# publish
FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/producer .

ENV GIN_MODE=release

ENTRYPOINT ["/app/producer"]