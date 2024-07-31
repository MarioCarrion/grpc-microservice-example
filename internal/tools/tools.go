package tools

import (
	_ "github.com/bufbuild/buf/cmd/buf"                     // Buf CLI
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint" // Linter
	_ "github.com/grpc-ecosystem/grpc-health-probe"         // gRPC Health Check
	_ "golang.org/x/vuln/cmd/govulncheck"                   // Official Go vulnerability checks
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"       // Protocol Buffer Plugin
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"        // gRPC Plugin
)
