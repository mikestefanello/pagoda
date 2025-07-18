# Get the OS name in lowercase (linux, darwin)
OS_SYSNAME := $(shell uname -s | tr A-Z a-z)
# Get the machine architecture (x86_64, arm64)
OS_MACHINE := $(shell uname -m)

# If mac OS, use `macos-arm64` or `macos-x64`
ifeq ($(OS_SYSNAME),darwin)
	OS_SYSNAME = macos
	ifneq ($(OS_MACHINE),arm64)
		OS_MACHINE = x64
	endif
endif

# If Linux, use `linux-x64`
ifeq ($(OS_SYSNAME),linux)
	OS_MACHINE = x64
endif

TAILWIND_PACKAGE = tailwindcss-$(OS_SYSNAME)-$(OS_MACHINE)

.PHONY: help
help: ## Print make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: ent-install air-install tailwind-install ## Install all dependencies

.PHONY: tailwind-install
tailwind-install: ## Install the Tailwind CSS CLI
	curl -sLo tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/latest/download/$(TAILWIND_PACKAGE)
	chmod +x tailwindcss
	curl -sLO https://github.com/saadeghi/daisyui/releases/latest/download/daisyui.js
	curl -sLO https://github.com/saadeghi/daisyui/releases/latest/download/daisyui-theme.js

.PHONY: ent-install
ent-install: ## Install Ent code-generation module
	go get entgo.io/ent/cmd/ent

.PHONY: air-install
air-install: ## Install air
	go install github.com/air-verse/air@latest

.PHONY: ent-gen
ent-gen: ## Generate Ent code
	go generate ./ent

.PHONY: ent-new
ent-new: ## Create a new Ent entity (ie, make ent-new name=MyEntity)
	go run entgo.io/ent/cmd/ent new $(name)

.PHONY: admin
admin: ## Create a new admin user (ie, make admin email=myemail@web.com)
	go run cmd/admin/main.go --email=$(email)

.PHONY: run
run: ## Run the application
	clear
	go run cmd/web/main.go

.PHONY: watch
watch: ## Run the application and watch for changes with air to automatically rebuild
	clear
	air

.PHONY: test
test: ## Run all tests
	go test ./...

.PHONY: check-updates
check-updates: ## Check for direct dependency updates
	go list -u -m -f '{{if not .Indirect}}{{.}}{{end}}' all | grep "\["

.PHONY: css
css: ## Build and minify Tailwind CSS
	./tailwindcss -i tailwind.css -o public/static/main.css -m

.PHONY: build
build: css ## Build CSS and compile the application binary
	go build -o ./tmp/main ./cmd/web
