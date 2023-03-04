package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Println("UnaryClientInterceptor PRE", method)

	ctx = metadata.AppendToOutgoingContext(ctx, "hello", "world")

	err := invoker(ctx, method, req, reply, cc, opts...)

	log.Printf("UnaryClientInterceptor POST %v\n", reply)

	return err
}
