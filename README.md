# "User Management" gRPC Microservice Example

## Introduction

Welcome! 👋

This is an educational repository that includes a small microservice written in Go using gRPC, this is equivalent to the [REST Example](https://github.com/MarioCarrion/todo-api-microservice-example), however this project is the principal example of my video series: [System Design using gRPC with Go](https://www.youtube.com/playlist?list=PL7yAAGMOat_EX1nv8fgltlm0CnJTH8Nwg).

Join the fun at [https://youtube.com/@MarioCarrion](https://www.youtube.com/@MarioCarrion).

## Prerequisites

* **required** Go **1.21**, _and_
* **recommended** `direnv`, to allow all Go-based binaries to be local to this folder and not installed globally. For more details you can refer to [this post](https://mariocarrion.com/2020/11/20/golang-go-tool-direnv.html).

## Tools

Make sure the following tools are installed:

* **required** Protocol Buffers Compiler, `protoc` (`23.4`, version to date):
    * Homebrew: `brew install protobuf`
    * Alpine 3.17: `apk add protobuf-dev protobuf`
    * Ubuntu 21.10: `apt-get install protobuf-compiler libprotobuf-dev`
* **recommended** Code Formatting, `clang-format`, you can use `find . -name '*.proto' | xargs clang-format -i`
    * Homebrew: `brew install clang-format` (`16.0.6`, version to date):
    * Alpine 3.17: `apk add clang-extra-tools`
    * Ubuntu 21.10: `apt-get install clang-format`
* The run `make tools` or if don't have Makefile copy/paste the `go install` instructions defined in the [Makefile](`Makefile#L1`)
