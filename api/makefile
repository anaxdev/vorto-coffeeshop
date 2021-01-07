.PHONY: build

build:
	go fmt ./...
	@mkdir -p ./bin/
	CGO_ENABLED=0 go build -o ./bin/server server.go
	CGO_ENABLED=0 go build -o ./bin/client client.go
