build:
	@go build -o sample_server cmd/shell/main.go

run: build
	@SAMPLE_SERVER_LOG_LEVEL=TRACE ./sample_server

help_env : build
	@./sample_server --help

test:
	@go test -cover ./...