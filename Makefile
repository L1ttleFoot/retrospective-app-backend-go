include .env
export

export PROJECT_ROOT=$(CURDIR)

run: 
	@go run cmd/server/main.go  

run2: 
	@set RABBITMQ_HOST=localhost&& go run cmd/notification/main.go

up:
	@docker compose up -d --build

up-go:
	@docker compose up -d --build go-app

up-react:
	@docker compose up -d --build retro-frontend

down:
	@docker compose down

clean:
	@echo "Очистка окружения"
	@read -p "Очистить volumes? [y/N]" ans \
	if ["$$ans"="y"];then \
		docker compose down -v --remove-orphans \
	fi 

migrate:
	@docker compose run --rm db-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"


migrate-action:
	@docker compose run --rm db-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres_db:5432/${POSTGRES_DB}?sslmode=disable \
		"${action}"


migrate-force:
	@docker compose run --rm db-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres_db:5432/${POSTGRES_DB}?sslmode=disable \
		force "${v}"

migrate-v:
	@docker compose run --rm db-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres_db:5432/${POSTGRES_DB}?sslmode=disable \
		version

migrate-up:
	make migrate-action action=up

migrate-down:
	make migrate-action action=down

test:
	@go test ./...