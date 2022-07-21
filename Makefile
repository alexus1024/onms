build:
	@go build -o onms cmd/shell/main.go

run: build
	./onms

test:
	@go test ./...