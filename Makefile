.PHONY: all lint test build
.SILENT:

lint:
	golangci-lint run

swagger:
	swag init -g ./internal/app/idler-service/app.go

tidy:
	go mod tidy

clean:
	go clean -modcache

proto-gen:
	protoc  --go_out=internal/transport/grpc --go_opt=paths=import \
	 --go-grpc_out=internal/transport/grpc --go-grpc_opt=paths=import \
	  internal/transport/grpc/proto/idler-email.proto

build: test
	go build -o ./.bin/app ./cmd/idler-service/main.go

test:
	go test -v ./test/...

docker:
	docker compose up