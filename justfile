set shell := ["powershell.exe", "-c"]


postgres:
    docker run --name pg-local -p 5433:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres:14.24-alpine3.23

createdb:
    docker exec -it pg-local createdb --username=postgres --owner=postgres omar_bank

dropdb:
    docker exec -it pg-local dropdb -U postgres omar_bank

migrate-up:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5433/omar_bank?sslmode=disable" -verbose up

migrate-up1:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5433/omar_bank?sslmode=disable" -verbose up 1

migrate-down:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5433/omar_bank?sslmode=disable" -verbose down

migrate-down1:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5433/omar_bank?sslmode=disable" -verbose down 1

sqlc:
    docker run --rm -v "${PWD}:/src" -w /src sqlc/sqlc generate

test:
    go test -v -cover ./...  

server:
    go run main.go

mock:
    mockgen -package mockdb -destination db/mock/store.go omarbank/db/sqlc Store