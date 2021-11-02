package main

import (
	"context"
	"fmt"
	"log"
	"net"
	t "time"

	"github.com/Philtoft/DISYS-Mini-Project-2-Chitty-Chat/time"
	"google.golang.org/grpc"
)

/*
* TODO: Make connection into channels
* TODO: Apply Lamport timestamp
 */
type Server struct {
	time.UnimplementedGetCurrentTimeServer
}

type Server1 struct {
	time.UnimplementedChatServer
}

func main() {
	// Create listener tcp on port 9080
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()
	time.RegisterGetCurrentTimeServer(grpcServer, &Server{})
	time.RegisterChatServer(grpcServer, &Server1{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

func (s *Server) GetTime(ctx context.Context, in *time.GetTimeRequest) (*time.GetTimeReply, error) {
	fmt.Printf("Received GetTime request\n")
	return &time.GetTimeReply{Reply: t.Now().String()}, nil
}

// Overvej om den skal have et bedre navn
func (s *Server1) SendChat(ctx context.Context, in *time.Message) (*time.Message, error) {
	// Hvordan får jeg vist beskeden, der kommer ind?
	fmt.Println("Received message:", in.GetMessage())
	return &time.Message{Message: in.GetMessage()}, nil
}
