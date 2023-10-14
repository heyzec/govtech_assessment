CMD_SEED_PATH=./cmd/seeds/main.go
CMD_SERVER_PATH=./cmd/server/main.go
CMD_MIGRATE_PATH=./cmd/migration/main.go

HAS_REFLEX := $(shell command -v reflex 2> /dev/null)

run:
ifdef HAS_REFLEX
	@echo "Starting server with reflex..."
	@GO_ENV=development reflex -r '\.go$$' -s go run ${CMD_SERVER_PATH}
else
	@echo "Starting server..."
	@GO_ENV=development go run ${CMD_SERVER_PATH}
endif

seed:
	@echo "Seeding database..."
	@GO_ENV=development go run ${CMD_SEED_PATH}

seed-tests:
	@echo "Seeding testing database..."
	@GO_ENV=testing go run ${CMD_SEED_PATH}

migrate:
	@echo "Migrating database..."
	@GO_ENV=development go run ${CMD_MIGRATE_PATH}
migrate-tests:
	@echo "Migrating database..."
	@GO_ENV=testing go run ${CMD_MIGRATE_PATH}

rollback:
	@echo "Rollback database by 1 step..."
	sql-migrate down 1

lint:
	@echo "Running linters..."
	golangci-lint run --fix --disable-all

