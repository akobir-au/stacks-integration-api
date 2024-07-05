FROM golang:1.22.5-alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN go build -o main cmd/app/main.go


FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/main .

COPY config ./config

EXPOSE 8080

ENTRYPOINT ["/app/main"]
