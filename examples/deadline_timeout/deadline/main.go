package main

import (
	"context"
	"fmt"
	"log"
	"time"

	userpb "github.com/MarioCarrion/grpc-microservice-example/gen/go/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9879", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()
	client := userpb.NewUserServiceClient(conn)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	res, err := client.GetUser(ctx, &userpb.GetUserRequest{
		Uuid: "1-2-3",
	})
	if err != nil {
		log.Fatalf("Failed to GetUser: %v", err)
	}

	fmt.Printf("%+v\n", res)
}
