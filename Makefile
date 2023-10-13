CMD_SEED_PATH=./cmd/seeds/main.go

seed:
	@echo "Seeding database..."
	go run ${CMD_SEED_PATH}

migrate:
	@echo "Migrating database..."
	sql-migrate up

rollback:
	@echo "Rollback database by 1 step..."
	sql-migrate down 1 

