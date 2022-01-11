# Connect to the primary database
.PHONY: db
db:
	 psql postgresql://admin:admin@localhost:5432/app

# Connect to the test database
.PHONY: db-test
db-test:
	 psql postgresql://admin:admin@localhost:5432/app_test

# Connect to the cache
.PHONY: cache
cache:
	 redis-cli

# Install Ent code-generation module
.PHONY: ent-install
ent-install:
	go get -d entgo.io/ent/cmd/ent

# Generate Ent code
.PHONY: ent-gen
ent-gen:
	go generate ./ent

# Create a new Ent entity
.PHONY: ent-new
ent-new:
	go run entgo.io/ent/cmd/ent init $(name)

# Start the Docker containers
.PHONY: up
up:
	docker-compose up -d
	sleep 3

# Rebuild Docker containers to wipe all data
.PHONY: reset
reset:
	docker-compose down
	make up

# Run the application
.PHONY: run
run:
	clear
	go run main.go

# Run all tests
.PHONY: test
test:
	go test -p 1 ./...

# Run the application using cosmtrek/air for live-reloading
.PHONY: dev
dev:
	docker-compose -f docker-compose.dev.yml up -d
	make -s dev-logs

# Stop the docker-compose resources
.PHONY: dev-down
dev-down:
	docker-compose -f docker-compose.dev.yml down

# Follow the logs for the specified docker-compose services. Additional services can be added after api if desired.
.PHONY: dev-logs
dev-logs:
	docker-compose -f docker-compose.dev.yml logs -f api
