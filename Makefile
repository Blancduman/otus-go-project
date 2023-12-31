install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.52.2

lint: install-lint-deps
	golangci-lint run ./...

gen-grpc:
	buf generate -v --template api/buf.gen.yaml api/grpc

test:
	go test -tags mock,integration -race -cover ./...