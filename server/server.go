package main

import (
	"context"
	"fmt"
	"log"
	"net"
	t "time"

	"github.com/Philtoft/DIS-mini-project-1/time"
	"github.com/Philtoft/DISYS-Mini-Project-2-Chitty-Chat/time"

	"google.golang.org/grpc"
)

type Server struct {
	time.UnimplementedGetCurrentTimeServer
}

func main() {
	// Create listener tcp on port 9080
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()

	time.RegisterGetMessageServer(grpcServer, &Server{})
	// time.RegisterGetCurrentTimeServer(grpcServer, &Server{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}

	// if err := grpcServer.Serve(list); err != nil {
	// 	log.Fatalf("failed to server %v", err)
	// }
}

func (s *Server) GetTime(ctx context.Context, in *time.GetTimeRequest) (*time.GetTimeReply, error) {
	fmt.Printf("Received GetTime request\n")
	return &time.GetTimeReply{Reply: t.Now().String()}, nil
}

func (s *Server) Broadcast(ctx context.Context, in *time.ChatRequest) (*time.ChatResponse, error) {
	fmt.Println("Message Received", in.GetMessage())
	return &time.ChatResponse{Message: in.GetMessage()}, nil
}
