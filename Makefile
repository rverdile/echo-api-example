POSTGRESQL_URL := 'postgres://testuser:password@127.0.0.1:5432/testdb?sslmode=disable'

build:
	go build -o main.go

run:
	go run main.go

test:
	go test ./...

clean:
	go clean

migrate-up:
	./migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	./migrate -database ${POSTGRESQL_URL} -path db/migrations down