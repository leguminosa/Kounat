.PHONY: help
help: # List all available make commands.
	@printf "\033[33m%s\033[0m\n" "Usage: make [target]"
	@grep -E '^[A-Za-z_-]+:.*?# .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-20s\033[0m%s\n", $$1, $$2}'

.PHONY: wire
wire: # Generate wire_gen files.
	@wire ./internal/app/kounatapi

.PHONY: mockgen
mockgen: # Generate mock files for all files with public interfaces.
	@./scripts/mockgen.sh

.PHONY: test
test: # Run unit tests for the whole repository.
	@go test -timeout 30s -short -count=1 -race -cover -coverprofile coverage.out -v ./...
	@go tool cover -func coverage.out

.PHONY: build # No need to actually build the binary here since we're running it using go run command anyway.
build: wire # Prepare to run your application here.

.PHONY: docker-start
docker-start: # Start docker-compose in detached mode.
	@docker-compose up -d

.PHONY: docker-stop
docker-remove: # Remove docker-compose.
	@docker-compose down -v

.PHONY: start
start: build docker-start # Run the application.
	@go run cmd/kounatapi/main.go

.PHONY: stop
stop: # Stop the application.
	@docker-compose stop