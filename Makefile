include .envrc

MIGRATIONS_PATH = ./cmd/migrate/migrations
SERVER_ADDR = postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost/$(POSTGRES_DB)?sslmode=disable

run:
	@go run cmd/api/*.go

test:
	@go test -v ./...

docker:
	@docker-compose -f docker-compose.dev.yml up --build --remove-orphans

docker-prod:
	@docker-compose -f docker-compose.prod.yml up --build --remove-orphans

migrate:
	@read -p "Enter the sequence name: " SEQ; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $${SEQ}

migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(SERVER_ADDR) up

migrate-down:
	@read -p "Number of migrations you want to rollback (default: 1): " NUM; NUM=$${NUM:-1}; \
	migrate -path=$(MIGRATIONS_PATH) -database=$(SERVER_ADDR) down $${NUM}

migrate-force: 
	@read -p "Enter the version to force: " VERSION; \
	@migrate -path=$(MIGRATIONS_PATH) -database=$(SERVER_ADDR) force $${VERSION}
