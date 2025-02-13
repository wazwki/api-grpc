.PHONY: lint test build up

DEBUG ?= false

export DEBUG

lint:
	golangci-lint run
	@echo "Линтер пройден"

test:
	go test ./...
	@echo "Тесты пройдены"

build:
	docker compose build

up:
	docker compose up -d

run: lint test build up
	@echo "Сервис успешно запущен!"
	
proto:
	mkdir -p api/proto/google/api
	curl -o api/proto/google/api/annotations.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto
	curl -o api/proto/google/api/http.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto
	protoc \
		-I api/proto \
		-I api/proto/google/api \
		--go_out api/proto \
		--go-grpc_out api/proto \
		--grpc-gateway_out api/proto \
		--openapiv2_out api/docs \
		--openapiv2_opt logtostderr=true,allow_merge=true \
		api/proto/name.proto
