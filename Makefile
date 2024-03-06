postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

migrateup:
	migrate -path database/migration -database "postgres://postgres:@127.0.0.1:5432/example?sslmode=disable" -verbose up 1
migratedown:
	migrate -path database/migration -database "postgres://postgres:@127.0.0.1:5432/example?sslmode=disable" -verbose down 1
    
hello:
	echo "Hello"

build_linux_amd64:
	go build -v -a -o \ release/${GOOS}/${GOARCH}/helloworld

temp:
	migrate -path database/migration -database "postgresql://postgres:kyS5OVaURtdWnW0UdyE6@drone2.cilgsltl3kau.us-east-1.rds.amazonaws.com:5432/Drone2" -verbose up 1 

.PHONY:postgres migrateup migratedown hello

