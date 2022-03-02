package main

import (
	"log"
	"os"

	userpb "github.com/MarioCarrion/grpc-microservice-example/gen/go/user/v1"
	"google.golang.org/protobuf/proto"
)

func main() {
	u := userpb.User{
		Uuid:      "1-2-3-4",
		FullName:  "Wario",
		BirthYear: 2900,
	}

	b, err := proto.Marshal(&u)
	if err != nil {
		log.Fatalln("Marshaling error", err)
	}

	if err := os.WriteFile("user.bin", b, 0644); err != nil {
		log.Fatalln("Writing error", err)
	}
}
