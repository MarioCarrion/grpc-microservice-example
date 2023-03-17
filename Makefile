tools:
	go install -C internal/tools \
		github.com/bufbuild/buf/cmd/buf \
		github.com/grpc-ecosystem/grpc-health-probe \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		google.golang.org/protobuf/cmd/protoc-gen-go
