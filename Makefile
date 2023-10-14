CMD_SEED_PATH=./cmd/seeds/main.go
CMD_SERVER_PATH=./cmd/server/main.go

HAS_REFLEX := $(shell command -v reflex 2> /dev/null)

run:
ifdef HAS_REFLEX
	@echo "Starting server with reflex..."
	reflex -r '\.go$$' -s go run ${CMD_SERVER_PATH}
else
	@echo "Starting server..."
	go run ${CMD_SERVER_PATH}
endif

seed:
	@echo "Seeding database..."
	go run ${CMD_SEED_PATH}

migrate:
	@echo "Migrating database..."
	sql-migrate up

rollback:
	@echo "Rollback database by 1 step..."
	sql-migrate down 1 

lint:
	@echo "Running linters..."
	golangci-lint run --fix --disable-all

