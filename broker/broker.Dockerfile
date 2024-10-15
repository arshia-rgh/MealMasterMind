FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x brokerApp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/brokerApp .

CMD ["./brokerApp"]
