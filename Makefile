lint: tools generate bufgenerate golangci govulncheck vet dirty

tools:
	go install -C internal/tools \
		github.com/bufbuild/buf/cmd/buf \
		github.com/golangci/golangci-lint/cmd/golangci-lint \
		github.com/grpc-ecosystem/grpc-health-probe \
		golang.org/x/vuln/cmd/govulncheck \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		google.golang.org/protobuf/cmd/protoc-gen-go

generate:
	go generate ./...

bufgenerate:
	buf generate

golangci:
	golangci-lint run ./...

govulncheck:
	govulncheck ./...

dirty:
	@status=$$(git status --untracked-files=no --porcelain); \
	if [ ! -z "$${status}" ]; \
	then \
		echo "ERROR: Working directory contains modified files"; \
		git status --untracked-files=no --porcelain; \
		exit 1; \
	fi

vet:
	go vet ./...

test:
	go test -shuffle=on -race -coverprofile=coverage.txt -covermode=atomic $$(go list ./... | grep -v /cmd/)
