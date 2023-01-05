package main

import (
	"context"
	"io"
	"testing"

	wearablepb "github.com/MarioCarrion/grpc-microservice-example/gen/go/wearable/v1"
	"google.golang.org/grpc"
)

func TestWearableService_BeatsPerMinute(t *testing.T) {
	svc := wearableService{}
	conn := newServer(t, func(srv *grpc.Server) {
		wearablepb.RegisterWearableServiceServer(srv, &svc)
	})

	client := wearablepb.NewWearableServiceClient(conn)
	stream, err := client.BeatsPerMinute(context.Background(), &wearablepb.BeatsPerMinuteRequest{})
	if err != nil {
		t.Fatalf("client.BeatsPerMinuteRequest %v", err)
	}

	index := 0
	expected := []uint32{5, 10, 15, 20, 25}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}

		if index >= len(expected) {
			t.Fatalf("expected five responses")
		}

		val := expected[index]
		if resp.Value != val+30 {
			t.Fatalf("got %d expected %d", resp.Value, val+30)
		}

		if resp.Minute != val {
			t.Fatalf("got %d expected %d", resp.Minute, val)
		}

		index++
	}

	if index != len(expected) {
		t.Fatalf("unexpectede length %d", index)
	}
}

func TestWearableService_ConsumeBeatsPerMinute(t *testing.T) {
	svc := wearableService{}
	conn := newServer(t, func(srv *grpc.Server) {
		wearablepb.RegisterWearableServiceServer(srv, &svc)
	})

	client := wearablepb.NewWearableServiceClient(conn)
	stream, err := client.ConsumeBeatsPerMinute(context.Background())
	if err != nil {
		t.Fatalf("client.ConsumeBeatsPerMinute %v", err)
	}

	for i := 0; i < 5; i++ {
		if err1 := stream.Send(&wearablepb.ConsumeBeatsPerMinuteRequest{}); err1 != nil {
			t.Fatalf("stream.Send %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("stream.CloseAndRecv %v", err)
	}

	if res.Total != 5 {
		t.Fatalf("expected %d got %d", 5, res.Total)
	}
}

func TestWearableService_CalculateBeatsPerMinute(t *testing.T) {
	svc := wearableService{}
	conn := newServer(t, func(srv *grpc.Server) {
		wearablepb.RegisterWearableServiceServer(srv, &svc)
	})

	client := wearablepb.NewWearableServiceClient(conn)
	stream, err := client.CalculateBeatsPerMinute(context.Background())
	if err != nil {
		t.Fatalf("client.CalculateBeatsPerMinute %v", err)
	}

	for i := 0; i < 5; i++ {
		if err1 := stream.Send(&wearablepb.CalculateBeatsPerMinuteRequest{Value: 10}); err != nil {
			t.Fatalf("stream.Send %v", err1)
		}
	}

	if err := stream.CloseSend(); err != nil {
		t.Fatalf("stream.CloseSend %v", err)
	}

	var result float32

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}

		result = resp.Average
	}

	if result != 10.0 {
		t.Fatalf("expected 10.0 got %f", result)
	}
}
