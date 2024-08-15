.PHONY: all
all: lint test

.PHONY: tidy
tidy:
	go mod tidy
	go mod -C internal/tools tidy

.PHONY: tools
tools:
	go install -C internal/tools \
		github.com/bufbuild/buf/cmd/buf \
		github.com/golangci/golangci-lint/cmd/golangci-lint \
		github.com/google/yamlfmt/cmd/yamlfmt \
		github.com/grpc-ecosystem/grpc-health-probe \
		golang.org/x/vuln/cmd/govulncheck \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		google.golang.org/protobuf/cmd/protoc-gen-go

# Formatting

.PHONY: fmt
fmt:
	go fmt ./...
	yamlfmt .

.PHONY: dirty
dirty:
	@status=$$(git status --untracked-files=no --porcelain); \
	if [ ! -z "$${status}" ]; \
	then \
		echo "ERROR: Working directory contains modified files"; \
		git status --untracked-files=no --porcelain; \
		exit 1; \
	fi

# Generate

.PHONY: generate
generate:
	go generate ./...
	buf generate

# Lint

.PHONY: lint
lint: tidy tools fmt security
	golangci-lint run ./...
	go vet ./...

# Security

.PHONY: security
security: 
	govulncheck ./...

# Test

.PHONY: test
test:
	go test -shuffle=on -race -coverprofile=coverage.txt -covermode=atomic $$(go list ./... | grep -v /cmd/)
