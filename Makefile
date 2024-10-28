LOGGER_BINARY=loggerApp
MEAL_BINARY=mealApp


up:
	@echo "Starting Docker images..."
	sudo docker compose up
	@echo "Docker images started!"

up_build: build_logger build_meal
	@echo "Stopping docker images (if running...)"
	sudo docker compose down
	@echo "Building (when required) and starting docker images..."
	sudo docker compose up --build
	@echo "Docker images built and started!"

build_meal:
	@echo "Building meal binary..."
	cd mealPlanning && env CGO_ENABLED=0 go build -o ${MEAL_BINARY} ./cmd/api
	@echo "Done!"

build_logger:
	@echo "Building logger binary..."
	cd logger && env CGO_ENABLED=0 go build -o ${LOGGER_BINARY} .
	@echo "Done!"

down:
	@echo "Stopping docker compose..."
	sudo docker compose down
	@echo "Done!"

