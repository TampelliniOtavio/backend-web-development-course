DB_URL=postgresql://username:password@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=username --owner=username simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
ifdef function
ifdef package
	go test -v -cover ./${package}/... -run ${function}
else
	go test -v -cover ./... -run ${function}
endif
else
ifdef package
	go test -v -cover ./${package}/...
else
	go test -v -cover ./...
endif
endif

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/TampelliniOtavio/backend-web-development-course/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migratedown1 migrateup1
