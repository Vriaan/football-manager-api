# API test service name within docker-compose.yml
DOCKER_API_TEST_SERVICE := football-manager-api-test
# API test service name within docker-compose.yml
DOCKER_DATABASE_TEST_SERVICE := football-manager-db-test
# API service name within docker-compose.yml
DOCKER_API_SERVICE := football-manager-api
# Database service name within docker-compose.yml
DOCKER_DATABASE_SERVICE := football-manager-db

# API directory within docker container (Todo: put just the binary file)
CONTAINER_API_DIR := /usr/project/football-manager-api

API_DIR := .

# api executable absolute path
API_EXECUTABLE_NAME := $(API_DIR)/bin/footballmanagerapi
# api entrypoint file
API_MAIN_FILE := $(API_DIR)/main.go

.PHONY: test

run_docker_test:
	@echo "RUNNING Tests within Docker container $(DOCKER_API_TEST_SERVICE)\n"
	-docker-compose run --rm -w $(CONTAINER_API_DIR) $(DOCKER_API_TEST_SERVICE) make test
	docker-compose rm -fs $(DOCKER_DATABASE_TEST_SERVICE) $(DOCKER_API_TEST_SERVICE)

run_docker_benchmark:
	@echo "RUNNING Benchmarks within Docker container $(DOCKER_API_TEST_SERVICE)\n"
	-docker-compose run --rm -w $(CONTAINER_API_DIR) $(DOCKER_API_TEST_SERVICE) make benchmark
	docker-compose rm -fs $(DOCKER_DATABASE_TEST_SERVICE) $(DOCKER_API_TEST_SERVICE)

start_docker_api:
	@echo "STARTING API environment\n"
	docker-compose rm -fs $(DOCKER_DATABASE_SERVICE) $(DOCKER_API_SERVICE)
	docker-compose up  -d --force-recreate $(DOCKER_DATABASE_SERVICE)
	sleep 5
	-docker-compose run --rm -w $(CONTAINER_API_DIR) $(DOCKER_API_SERVICE) make exec

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

exec: build
	$(API_EXECUTABLE_NAME)

clean_docker:
	docker-compose rm -fs
	docker container prune -f
	docker volume prune -f
