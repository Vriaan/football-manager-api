# API test service name within docker-compose.yml
DOCKER_API_TEST_SERVICE := football-manager-api-test
# API service name within docker-compose.yml
DOCKER_API_SERVICE := football-manager-api
# Database service name within docker-compose.yml
DOCKER_DATABASE_SERVICE := football-manager-db

# API directory within docker container (Todo: put just the binary file)
CONTAINER_API_DIR := /usr/project/api

# api executable absolute path
API_EXECUTABLE_NAME := $(CONTAINER_API_DIR)/bin/api
# api entrypoint file
API_MAIN_FILE := $(CONTAINER_API_DIR)/main.go

run_docker_test:
	@echo "RUNNING Tests within Docker container $(DOCKER_API_TEST_SERVICE)\n"
	-docker-compose run --rm -w $(CONTAINER_API_DIR) $(DOCKER_API_TEST_SERVICE) make test
	docker-compose rm -fs
	docker container prune -f

run_docker_benchmark:
	@echo "RUNNING Benchmarks within Docker container $(DOCKER_API_TEST_SERVICE)\n"
	-docker-compose run --rm -w $(CONTAINER_API_DIR) $(DOCKER_API_TEST_SERVICE) make benchmark
	docker-compose rm -fs
	docker container prune -f

start_docker_api:
	@echo "STARTING API environment\n"
	docker-compose rm -fs
	docker container prune -f
	docker-compose up  -d --force-recreate $(DOCKER_DATABASE_SERVICE)
	sleep 5
	docker-compose up --no-recreate $(DOCKER_API_SERVICE)

stop_docker_api:
	@echo "STOPS API environment\n"
	docker-compose down
	docker container prune -f

test:
	sleep 2
	go clean -cache -testcache
	go test -cover ./...

benchmark:
	sleep 2
	go clean -cache -testcache
	go test  ./... -run=XXX -bench=.

build:
	go build -o $(API_EXECUTABLE_NAME) $(API_MAIN_FILE)

# Clean golang cache & containers
clean:
	go clean -cache -testcache
	docker-compose rm -fs
	docker container prune -f
