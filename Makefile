BROKER_BINARY=brokerApp
LISTENER_BINARY=listenerApp


up:
	@echo "Starting Docker images..."
	sudo docker compose up
	@echo "Docker images started!"

up_build: build_broker
	@echo "Stopping docker images (if running...)"
	sudo docker compose down
	@echo "Building (when required) and starting docker images..."
	sudo docker compose up --build
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	sudo docker compose down
	@echo "Done!"

build_broker:
	@echo "Building broker binary..."
	cd broker && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

build_listener:
	@echo "Building listener binary..."
	cd listener && env GOOS=linux CGO_ENABLED=0 go build -o ${LISTENER_BINARY} ./event
	@echo "Done!"