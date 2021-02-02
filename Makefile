
DATABASE_CONTAINER=testdb

createdb:
	docker exec -it $(DATABASE_CONTAINER) createdb --username=root --owner=root postgres

dropdb:
	docker exec -it $(DATABASE_CONTAINER) dropdb postgres

migrateup:
	migrate -path schema -database "postgresql://postgres:example@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path schema -database "postgresql://postgres:example@localhost:5432/postgres?sslmode=disable" -verbose down

rebuild_app:
	docker-compose up -d --no-deps --build app

run:
	docker-compose up app

test:
	go test -tags all_tests ./...

.PHONY: postgres createdb dropdb migrateup migratedown

.PHONY: postgres createdb dropdb