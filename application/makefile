include .env
export

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

run: ## runs app localy
	@go run cmd/app/main.go
.PHONY: run

run_test: ## runs app localy
	@go run cmd/test/main.go
.PHONY: run_test

create-user: ## creates new user, provie name=(...) login=(...) and pass=(...)
	@go run cmd/console/main.go -name=$(name) -login=$(login) -pass=$(pass)
.PHONY: create-user

up: ## testing infra
	@docker-compose -f deployments/docker-compose.yaml up -d
.PHONY: up

commit: ## testing infra
	@docker-compose -f deployments/docker-compose-prod.yaml up -d
.PHONY: up

gen-migrations-file: ##creates new migrations files
	@migrate create -ext sql -dir migrations -seq $(name)
.PHONY: gen-migrations-file
