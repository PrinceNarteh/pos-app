include .env
MIGRATION_DIR=./cmd/migrations/

.PHONY: dev
dev:
	@air

.PHONY: create-migration
create-migration:
	@read -p "Enter migration name: " name; \
	goose -s create $$name sql; \
	echo "Created migration: $$name"

.PHONY: migration-up
migration-up:
	goose up

.PHONY: migration-down
migration-down:
	goose down

.PHONY: migration-reset
migration-clean:
	@goose to 0

.PHONY: migration-clean
migration-clean: migration-reset migration-up
	@echo "Database cleaned successfully"
