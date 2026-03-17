tidy:
	go mod tidy

run: tidy
	go run ./cmd/http

test:
	go test ./internal/service
