BIN := "./bin/twirler"
DOCKER_IMG="twirler:develop"

compose:
	docker-compose -f ./docker-compose.yaml up -d

build:
	go build -v -o $(BIN) .

run: build
	$(BIN)

build-img:
	docker build \
		-t $(DOCKER_IMG) \
		-f Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

test-unit:
	go test -race -count 100 ./internal/...

test-integration:
	go test ./... -tags integration

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.52.2

lint: install-lint-deps
	golangci-lint run ./...

migrate-up:
	go run main.go migrate up

migrate-down:
	go run main.go migrate down

.PHONY: build run build-img run-img test lint
