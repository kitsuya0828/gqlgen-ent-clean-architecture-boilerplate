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
	go run -mod=mod ./ent/entc.go
	go run -mod=mod github.com/99designs/gqlgen

.PHONY: setup_db start generate