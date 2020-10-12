package main

import (
	"context"
	"log"
	"net"

	pb "./proto"
	grpc "google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreetingServer
}

func (s *server) SayHello(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	log.Printf("Request Received: %v", request.GetReqeustMessage())
	return &pb.Response{ReplyMessage: "Hello Client! I received your request of \"" + request.GetReqeustMessage() + "\""}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("ERROR: failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterGreetingServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("ERROR: failed to serve: %v", err)
		return
	}
}
