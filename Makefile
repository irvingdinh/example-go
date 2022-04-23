install:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/vektra/mockery/v2@latest

goose.create:
	@read -p "Migration filename: " filename; \
	cd ./database/migrations; \
	goose create $$filename sql;

goose.up:
	@read -p "Database DSN: " dsn; \
	cd ./database/migrations; \
	goose mysql "$$dsn" up

goose.down:
	@read -p "Database DSN: " dsn; \
	cd ./database/migrations; \
	goose mysql "$$dsn" down

goose.status:
	@read -p "Database DSN: " dsn; \
	cd ./database/migrations; \
	goose mysql "$$dsn" status
