# Connect to the primary database
.PHONY: db
db:
	psql postgresql://admin:admin@localhost:5432/app

# Connect to the test database
.PHONY: db-test
db-test:
	psql postgresql://admin:admin@localhost:5432/app_test

# Connect to the primary cache
.PHONY: cache
cache:
	redis-cli

 # Connect to the test cache
.PHONY: cache-test
cache-test:
	redis-cli -n 1

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