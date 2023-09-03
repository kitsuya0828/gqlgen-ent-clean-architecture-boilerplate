# Set up database
setup_db:
	./bin/init_db.sh
# Start dev server
start:
	air

.PHONY: setup_db start