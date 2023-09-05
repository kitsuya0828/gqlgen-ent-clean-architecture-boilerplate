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

# E2E
setup_e2e_db:
	./bin/init_db_e2e.sh

e2e:
	go test ./test/e2e/...

.PHONY: setup_db migrate_schema start generate setup_test_db test_repository setup_e2e_db e2e