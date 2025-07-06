# Project variables

APP_NAME=FyrnoApi
DB_NAME=ai_financial_app
DB_USER=root
DB_PASS=yourpassword
DB_HOST=localhost
DB_PORT=3306
MIGRATIONS_DIR=./migrations
DSN="mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?multiStatements=true"

# Default

.PHONY: help
help:
@echo "Makefile commands:"
@echo " run - Run the Go server"
@echo " build - Build the Go binary"
@echo " tidy - Run go mod tidy"
@echo " format - Format Go code"
@echo " migrate-up - Apply all up migrations"
@echo " migrate-down - Rollback all migrations"
@echo " migrate-create - Create a new migration file"
@echo " create-db - Create MySQL database if not exists"
@echo " docker-build - Build Docker image"
@echo " docker-up - Start Docker containers"
@echo " docker-down - Stop Docker containers"

run:
go run main.go

build:
go build -o $(APP_NAME) .

tidy:
go mod tidy

format:
go fmt ./...

create-db:
mysql -u$(DB_USER) -p$(DB_PASS) -e "CREATE DATABASE IF NOT EXISTS $(DB_NAME) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

migrate-up:
migrate -path $(MIGRATIONS_DIR) -database $(DSN) up

migrate-down:
migrate -path $(MIGRATIONS_DIR) -database $(DSN) down

migrate-create:
@read -p "Enter migration name: " name; \
 migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name

docker-build:
docker build -t $(APP_NAME):latest .

docker-up:
docker-compose up -d

docker-down:
docker-compose down
