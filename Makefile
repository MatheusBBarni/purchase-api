api:
	go run cmd/api/main.go

test:
	go test -v ./...

coverage:
	go test -coverprofile=coverage.out ./...

coverage_page:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out