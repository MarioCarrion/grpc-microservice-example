# "User Management" gRPC Microservice Example

## Introduction

Welcome! 👋

This is an educational repository that includes a small microservice written in Go using gRPC, this is equivalent to the [REST Example](https://github.com/MarioCarrion/todo-api-microservice-example), however this project is the principal example of my video series: [System Design using gRPC with Go](https://www.youtube.com/playlist?list=PL7yAAGMOat_EX1nv8fgltlm0CnJTH8Nwg).

Join the fun at [https://youtube.com/@MarioCarrion](https://www.youtube.com/@MarioCarrion).

## Prerequisites

* **required** Go **1.19**, _and_
* **recommended** `direnv`, to allow all Go-based binaries to be local to this folder and not installed globally. For more details you can refer to [this post](https://mariocarrion.com/2020/11/20/golang-go-tool-direnv.html).

## Tools

Install the following tools:

* **required** Protocol Buffers Compiler, `protoc` (`3.21.9`, version to date):
    * Homebrew: `brew install protobuf`
    * Alpine 3.15: `apk add protobuf-dev protobuf`
    * Ubuntu 21.10: `apt-get install protobuf-compiler libprotobuf-dev`
* **required** `buf` CLI for linting and compiling:
    * `go install github.com/bufbuild/buf/cmd/buf@v1.9.0`
* **required** Protocol Buffer Plugin for Go:
    * `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1`
* **required** gRPC Plugin for Go:
    * `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0`
* **recommended** Code Formatting, `clang-format`, you can use `find . -name '*.proto' | xargs clang-format -i`
    * Homebrew: `brew install clang-format` (`14.0.6`, version to date):
    * Alpine 3.15: `apk add clang-extra-tools`
    * Ubuntu 21.10: `apt-get install clang-format`
* **recommended** gRPC Health Check tester using `grpc-health-probe`:
    * `go install github.com/grpc-ecosystem/grpc-health-probe@v0.4.14`
