# Determine if you have docker-compose or docker compose installed locally
# If this does not work on your system, just set the name of the executable you have installed
DCO_BIN := $(shell { command -v docker-compose || command -v docker compose; } 2>/dev/null)

# Connect to the primary database
.PHONY: db
db:
	docker exec -it pagoda_db psql postgresql://admin:admin@localhost:5432/app

# Connect to the test database (you must run tests first before running this)
.PHONY: db-test
db-test:
	docker exec -it pagoda_db psql postgresql://admin:admin@localhost:5432/app_test

# Connect to the primary cache
.PHONY: cache
cache:
	docker exec -it pagoda_cache redis-cli

 # Connect to the test cache
.PHONY: cache-test
cache-test:
	docker exec -it pagoda_cache redis-cli -n 1

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
	$(DCO_BIN) up -d
	sleep 3

# Rebuild Docker containers to wipe all data
.PHONY: reset
reset:
	$(DCO_BIN) down
	make up

# Run the application
.PHONY: run
run:
	clear
	go run cmd/web/main.go

# Run all tests
.PHONY: test
test:
	go test -count=1 -p 1 ./...

# Run the worker
.PHONY: worker
worker:
	clear
	go run cmd/worker/main.go

# Check for direct dependency updates
.PHONY: check-updates
check-updates:
	go list -u -m -f '{{if not .Indirect}}{{.}}{{end}}' all | grep "\["
