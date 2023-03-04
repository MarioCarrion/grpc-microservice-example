package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("UnaryServerInterceptor PRE", info.FullMethod)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("no metadata")
	} else {
		log.Printf("Metadata %v+\n", md)
	}

	m, err := handler(ctx, req)

	log.Println("UnaryServerInterceptor POST", info.FullMethod)

	return m, err
}

type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Received %T - %v\n", m, m)

	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Sent %T - %v\n", m, m)

	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("StreamServerInterceptor PRE", info.FullMethod)

	err := handler(srv, newWrappedStream(ss))

	log.Println("StreamServerInterceptor POST")

	return err
}
