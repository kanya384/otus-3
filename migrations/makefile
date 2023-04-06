include .env
export

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

run: ## runs app localy
	@go run cmd/app/main.go
.PHONY: run

up: ## testing infra
	@docker-compose -f deployments/docker-compose.yaml up -d
.PHONY: up

commit: ## build and push container to docker hub| provide tag=(...)
	@docker build -t laurkan/otus-3-migrations:$(tag) . && docker push laurkan/otus-3-migrations:$(tag)
.PHONY: commit

gen-migrations-file: ##creates new migrations files
	@migrate create -ext sql -dir migrations -seq $(name)
.PHONY: gen-migrations-file
