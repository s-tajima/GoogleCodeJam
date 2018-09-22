run:
	go run main.go solver.go

build:
	go build -o solver

test:
	go test

test/coverage:
	go test -coverprofile=/tmp/cover.out
	go tool cover -html=/tmp/cover.out -o /tmp/cover.html
