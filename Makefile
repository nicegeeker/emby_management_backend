postgres:
	docker run --name postgres-emby -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.5-alpine

createdb:
	docker exec -it postgres-emby createdb --username=root --owner=root emby_management

dropdb:
	docker exec -it postgres-emby dropdb emby_management

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/emby_management?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/emby_management?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/emby_management?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/emby_management?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/nicegeeker/emby_management_backend/db/sqlc Store

.PHONY:postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock