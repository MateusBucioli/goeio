# build
FROM golang:1.20.6-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
COPY rabbitmq ./rabbitmq
COPY consumer ./consumer

RUN go build -ldflags="-s -w" -o consumer ./consumer/main.go

# publish
FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/consumer .

ENV GIN_MODE=release

ENTRYPOINT ["/app/consumer"]