test:
	@go test -v -cover -covermode=atomic ./tests

engine:
	@go build -o engine main.go

run:
	@go run main.go