package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	wearablepb "github.com/MarioCarrion/grpc-microservice-example/gen/go/wearable/v1"
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

	//-

	client := wearablepb.NewWearableServiceClient(conn)
	stream, err := client.CalculateBeatsPerMinute(context.Background())
	if err != nil {
		log.Fatalln("Opening stream", err)
	}

	for i := 0; i < 10; i++ {
		err := stream.Send(&wearablepb.CalculateBeatsPerMinuteRequest{Uuid: "mario", Value: uint32(i), Minute: uint32(i)})
		if err != nil {
			log.Fatalln("Sending value", err)
		}
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalln("CloseSend", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("Closing", err)
		}

		fmt.Println("Total average", res.GetAverage())
	}
}
