.PHONY: db-schema run-server run-frontend

db-schema:
	if [ -f schema/schema.sql ]; then rm schema/schema.sql; fi;
	sqlite3 watchman.db .schema > schema/schema.sql;

run-server:
	go mod tidy
	go run main.go

run-frontend:
	cd frontend && git pull origin master
	bun run dev
