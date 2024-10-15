FROM alpine:latest

WORKDIR /app

COPY  brokerApp .

CMD ["./brokerApp"]
