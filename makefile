
default: help

.PHONY: help
help:
	@echo 'local-storage'
	@echo 'usage: make [target] ...'

.PHONY: install-tool
install-tool:
	go get -u github.com/golang/mock/gomock
	go get -u github.com/golang/mock/mockgen

.PHONY: install-dependency
install-dependency:
	go mod tidy
	go mod verify
	go mod vendor

.PHONY: clean-dependency
clean-dependency:
	rm -f go.sum
	rm -rf vendor
	go clean -modcache

.PHONY: install
install:
	go install -v ./...

.PHONY: test
test:
	go test ./... -coverprofile coverage.out
	go tool cover -func coverage.out | grep ^total:

.PHONY: generate-mock
generate-mock:
	mockgen -package=mock -source internal/logging/log.go -destination=internal/mock/logging_log_mock.go
	mockgen -package=mock -source internal/serialization/serializer.go -destination=internal/mock/serialization_serializer_mock.go
	mockgen -package=mock -source internal/healthcheck/health.go -destination=internal/mock/healthcheck_health_mock.go
	mockgen -package=mock -source internal/explorer/file.go -destination=internal/mock/explorer_file_mock.go
	mockgen -package=mock -source internal/repository/file.go -destination=internal/mock/repository_file_mock.go

.PHONY: run-grpc-app
run-grpc-app:
	go run cmd/grpc-app/main.go

.PHONY: run-rest-app
run-rest-app:
	go run cmd/rest-app/main.go

.PHONY: build-grpc-app
build-grpc-app:
	go build -o ./build/grpc-app/ ./cmd/grpc-app/main.go

.PHONY: build-rest-app
build-rest-app:
	go build -o ./build/rest-app/ ./cmd/rest-app/main.go

.PHONY: run-docker-dev
run-docker-dev:
	docker-compose up -d
