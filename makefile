# Makefile

# Build the Docker images
build:
	docker-compose build

# Run the Docker containers
run:
	docker-compose up -d

# Stop the Docker containers
stop:
	docker-compose down

# Run the tests
test:
	go test -v ./...

swag:
	swag init --dir cmd,controllers,models,errors --output docs
