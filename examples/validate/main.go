package main

import (
	"fmt"

	"github.com/bufbuild/protovalidate-go"

	userpb "github.com/MarioCarrion/grpc-microservice-example/gen/go/user/v1"
)

func main() {
	user := &userpb.User{}

	v, err := protovalidate.New()
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	if err = v.Validate(user); err != nil {
		fmt.Println("validation failed:", err)
	} else {
		fmt.Println("validation succeeded")
	}
}
