up:
	docker compose up -d

test:
	go test -timeout=10m ./...

test_v:
	go test -v ./...

test_short:
	go test -short ./...

test_race:
	go test -short -race ./...

test_stress:

test_codecov:

test_coverage:
	go test -timeout=10m -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html

test_reconnect:
	go test -tags=reconnect ./...

wait:
	go run ./internal/wait-for

build:
	go build ./...

fmt:
	go fmt ./...
	goimports -l -w .

fmt-check:
	@test -z "$$(gofmt -l .)" || (echo "Files need formatting:"; gofmt -l .; exit 1)
	@test -z "$$(goimports -l .)" || (echo "Imports need organizing:"; goimports -l .; exit 1)

lint:
	golangci-lint run --timeout=5m

update_watermill:
	go get -u github.com/ThreeDotsLabs/watermill
	go mod tidy

	sed -i '\|go 1\.|d' go.mod
	go mod edit -fmt

pgcli:
	@pgcli postgres://watermill:password@localhost:5432/watermill?sslmode=disable

ci: fmt-check lint build test test_race
