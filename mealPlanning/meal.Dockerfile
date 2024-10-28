FROM alpine:latest

WORKDIR /app

COPY mealApp /app

CMD ["/app/mealApp"]
