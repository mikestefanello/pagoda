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
	go run entgo.io/ent/cmd/ent new $(name)

# Run the application
.PHONY: run
run:
	clear
	go run cmd/web/main.go

# Run all tests
.PHONY: test
test:
	go test ./...

# Check for direct dependency updates
.PHONY: check-updates
check-updates:
	go list -u -m -f '{{if not .Indirect}}{{.}}{{end}}' all | grep "\["
