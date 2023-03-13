postgres:
    docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

migrateup:
    migrate -path db/migration -database "migrate -verbose -source file://database/migration -database postgres://postgres:@127.0.0.1:5432/postgres?sslmode=disable up 1" -verbose up

migratedown:
    migrate -path db/migration -database "migrate -verbose -source file://database/migration -database postgres://postgres:@127.0.0.1:5432/postgres?sslmode=disable up 1" -verbose down

hello:
	echo "Hello"
.PHONY: postgres migrateup migratedown