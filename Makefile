build:
	@go build -o onms cmd/shell/main.go

test:
	@go test ./...