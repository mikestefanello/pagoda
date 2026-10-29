FROM golang:1.25.0-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=1 go build -o pagoda ./cmd/web/main.go

FROM alpine:3.19 AS runner

WORKDIR /app

COPY --from=builder /app/pagoda /pagoda
COPY /public /public
COPY /config /config

RUN addgroup -S pagoda && adduser -S pagoda -G pagoda && chown -R pagoda:pagoda /app

USER pagoda

EXPOSE 8000

CMD ["/pagoda"]   