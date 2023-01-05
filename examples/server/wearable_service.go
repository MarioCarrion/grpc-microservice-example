package main

import (
	"fmt"
	"io"

	wearablepb "github.com/MarioCarrion/grpc-microservice-example/gen/go/wearable/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type wearableService struct {
	wearablepb.UnimplementedWearableServiceServer
}

func (w *wearableService) BeatsPerMinute(
	req *wearablepb.BeatsPerMinuteRequest,
	stream wearablepb.WearableService_BeatsPerMinuteServer) error {
	results := []uint32{5, 10, 15, 20, 25}

	for i := 0; i < len(results); i++ {
		select {
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "Stream has ended")
		default:
			value := results[i]

			if err := stream.SendMsg(&wearablepb.BeatsPerMinuteResponse{
				Value:  uint32(value + 30),
				Minute: uint32(value),
			}); err != nil {
				return status.Error(codes.Canceled, "Stream has ended")
			}
		}
	}

	return nil
}

func (w *wearableService) ConsumeBeatsPerMinute(stream wearablepb.WearableService_ConsumeBeatsPerMinuteServer) error {
	var total uint32

	for {
		value, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wearablepb.ConsumeBeatsPerMinuteResponse{
				Total: total,
			})
		}

		if err != nil {
			return err
		}

		fmt.Println(value.GetUuid(), value.GetMinute(), value.GetValue())
		total++
	}
}

func (w *wearableService) CalculateBeatsPerMinute(stream wearablepb.WearableService_CalculateBeatsPerMinuteServer) error {
	var count, total uint32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		total += req.GetValue()

		fmt.Println("Received", req.GetValue())

		count++

		if count%5 == 0 {
			fmt.Println("Total", total, "Sending", float32(total)/5)

			if err := stream.Send(&wearablepb.CalculateBeatsPerMinuteResponse{
				Average: float32(total) / 5,
			}); err != nil {
				return err
			}

			total = 0
		}
	}
}
