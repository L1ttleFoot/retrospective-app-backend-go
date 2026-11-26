include .env
export

export PROJECT_ROOT=$(CURDIR)

run: 
	@go run cmd/server/main.go  

run2: 
	@set RABBITMQ_HOST=localhost&& go run cmd/notification/main.go

up:
	@docker compose up -d --build

up-back:
	@docker compose up -d --build --no-deps go-app notification-app

up-front:
	@docker compose up -d --build --no-deps retro-frontend

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

migrate-action-one:
	@docker compose run --rm db-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres_db:5432/${POSTGRES_DB}?sslmode=disable \
		"${action}" 1


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

migrate-up-one:
	make migrate-action-one action=up

migrate-down-one:
	make migrate-action-one action=down

test:
	@go test ./...

sqlc-gen:
	@go tool sqlc generate

KAFKA_CONTAINER ?= kafka
TOPIC ?= board-events

.PHONY: kafka-read

kafka-read:
	@echo Reading from Kafka
	@docker exec -it $(KAFKA_CONTAINER) /opt/kafka/bin/kafka-console-consumer.sh \
		--bootstrap-server localhost:9092 \
		--topic $(TOPIC) \
		--from-beginning \
		--property print.key=true \
		--property print.partition=true \
		--property print.offset=true

kafka-topics:
	@echo Listing Kafka topics
	@docker exec -it $(KAFKA_CONTAINER) /opt/kafka/bin/kafka-topics.sh \
		--bootstrap-server localhost:9092 \
		--list