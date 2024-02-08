build:
	@go build -o bin/blocker

run: build
	@./bin/blocker

test:
	@go test -v ./...

.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/types.proto