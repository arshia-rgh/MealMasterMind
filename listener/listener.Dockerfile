FROM alpine:latest

WORKDIR /app

COPY listenerApp .

CMD ["./listenerApp"]