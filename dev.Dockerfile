FROM golang:1.19-alpine

RUN apk update && apk add --no-cache musl-dev gcc git build-base

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
