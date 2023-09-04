# Set up database
setup_db:
	./bin/init_db.sh

# Migrate ent schema to database
migrate_schema:
	go run ./cmd/migration/main.go

# Start dev server
start:
	air

# Generate code
generate:
	go generate .

# Run test
setup_test_db:
	./bin/init_db_test.sh

test_repository:
	go test ./pkg/adapter/repository/...


.PHONY: setup_db migrate_schema start generate