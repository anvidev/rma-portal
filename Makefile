include .envrc

MIGRATIONS_PATH = ./cmd/migrate/migrations
SERVER_ADDR = postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost/$(POSTGRES_DB)?sslmode=disable

.PHONY: run
run:
	@go run cmd/api/*.go

.PHONY: test
test:
	@go test -v ./...

.PHONY: docker
docker:
	@docker-compose -f docker-compose.dev.yml up --build --remove-orphans

.PHONY: docker-prod
docker-prod:
	@docker-compose -f docker-compose.prod.yml up --build --remove-orphans

.PHONY: migration
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(SERVER_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(SERVER_ADDR) down $(filter-out $@,$(MAKECMDGOALS))
