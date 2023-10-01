run:
	docker-compose up --build

walk:
	go run cmd/integration-test-exercises/main.go

test:
	go test ./pkg/converter/...