package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "messenger/grpc-server"
	"net"
)

const (
	port = ":7777"
)

type server struct {
	pb.UnimplementedPingServer
}

func (s *server) YesHello(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &pb.PingMessage{Greeting: in.Greeting}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
