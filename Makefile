# Project Makefile
TIMEOUT       = 5s
COVERAGE_FILE = /tmp/coverage.out

# TODO: Create examples and live testing directory.

all: test lint vet fmt

test:
	@go test -count=1 -timeout=${TIMEOUT} -race ./...

test-coverage-report:
	@go test -timeout=${TIMEOUT} ./... -coverprofile=${COVERAGE_FILE}
	@sed -i '/no_test_coverage.go/d' ${COVERAGE_FILE}
	@go tool cover -html=${COVERAGE_FILE}

fmt:
	@go fmt ./...

vet:
	@go vet ./...

lint:
	@golint ./...

.PHONY: all test test-coverage-report fmt vet lint