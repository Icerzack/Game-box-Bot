FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o bot cmd/server/main.go


FROM alpine

WORKDIR /app

COPY --from=builder /app/bot .
COPY --from=builder /app/.env .

CMD ["./bot"]
