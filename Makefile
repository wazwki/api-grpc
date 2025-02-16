.PHONY: lint test build up

DEBUG ?= false
SWAGGER_URL="swagger.json"
export DEBUG
export SWAGGER_URL

lint:
	golangci-lint run
	@echo "Линтер пройден"

test:
	go test ./...
	@echo "Тесты пройдены"

build:
	docker compose -f deployments/docker/name-docker-compose.yml build

up:
	docker compose -f deployments/docker/name-docker-compose.yml up -d

stop:
	docker compose -f deployments/docker/name-docker-compose.yml down

run: lint test build up
	@echo "Сервис успешно запущен!"
	
proto:
	mkdir -p third_party/google/api/
	curl -o third_party/google/api/annotations.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto
	curl -o third_party/google/api/http.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto
	protoc \
		-I api/proto/ \
		-I third_party/google/api \
		-I third_party \
		--go_out api/proto \
		--go-grpc_out api/proto \
		--grpc-gateway_out api/proto \
		api/proto/name.proto

debug:
	DEBUG=true
	docker compose -f deployments/docker/debug-docker-compose.yml up -d

debug-stop:
	docker compose -f deployments/docker/debug-docker-compose.yml down
	DEBUG=false

swagger:
	rm -rf third_party/swagger-ui
	mkdir -p third_party/swagger-ui

	git clone --depth=1 https://github.com/swagger-api/swagger-ui.git third_party/swagger-tmp
	mv third_party/swagger-tmp/dist/* third_party/swagger-ui/
	rm -rf third_party/swagger-tmp

	perl -pi -e 's/url:\s*"[^"]*"/url: $(SWAGGER_URL)/g' third_party/swagger-ui/swagger-initializer.js

	protoc \
		-I api/proto \
		-I api/proto/google/api \
		--openapiv2_out api/docs \
		--openapiv2_opt logtostderr=true,allow_merge=true \
		api/proto/name.proto
