package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type wrappedStream struct {
	grpc.ClientStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Received %T - %v\n", m, m)

	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Sent %T - %v\n", m, m)

	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

func StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Println("StreamClientInterceptor PRE:", method)

	s, err := streamer(ctx, desc, cc, method, opts...)

	log.Println("StreamClientInterceptor POST")

	return newWrappedStream(s), err
}
