FROM golang:1.17-alpine

RUN apk update && apk add --no-cache musl-dev gcc git build-base

RUN go get github.com/cosmtrek/air

WORKDIR /app

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]
