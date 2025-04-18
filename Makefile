postgres:
	docker run -- name postgreslatest -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=incorrect -d postgres:latest
createdb:
	docker exec -it postgreslatest createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgreslatest dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:incorrect@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:incorrect@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
