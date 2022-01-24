postgres:
	docker run --name database -p 5432:5432 -e POSTGRES_PASSWORD=nothing -e POSTGRES_USER=nothing -d postgres:12-alpine

createdb:
	docker exec -it database createdb --username=nothing --owner=nothing bank

dropdb:
	docker exec -it database dropdb bank
up:
	migrate -path db/migration -database "postgresql://nothing:nothing@localhost:5432/bank?sslmode=disable" -verbose up
down:
	migrate -path db/migration -database "postgresql://nothing:nothing@localhost:5432/bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate

.PHONY: createdb dropdb postgres up down sqlc