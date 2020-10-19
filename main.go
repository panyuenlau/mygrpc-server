package main

import (
	"context"
	"log"
	"net"

	pb "github.com/panyuenlau/mygrpc-server/proto"
	grpc "google.golang.org/grpc"
	peer "google.golang.org/grpc/peer"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreetingServer
}

func (s *server) SayHello(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	p, _ := peer.FromContext(ctx) // retreieve the peer information in ctx if it exists
	clientNetwork := p.Addr.Network()
	clientIP := p.Addr.String()
	requestMsg := request.GetReqeustMessage()

	log.Printf("Request Received: \"%v\" from client: %s through %s", request.GetReqeustMessage(), clientIP, clientNetwork)

	return &pb.Response{
		ReplyMessage: "Hello Client! " + "I received your message of \"" + requestMsg + "\"",
	}, nil
}

func main() {
	log.Println("Waiting for requests from clients...")
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
