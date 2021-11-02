package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	t "time"

	service "github.com/Philtoft/DISYS-Mini-Project-2-Chitty-Chat/service"
	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
)

/*
* TODO: Make connection into channels
* TODO: Apply Lamport timestamp
 */

type Connection struct {
	// Fatter ikke denne del
	stream service.Broadcast_CreateStreamServer
	id     string
	active bool
}

type Server struct {
	// array of connections with a reference to the connection location in memory
	Connection []*Connection
}

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

func (s *Server) CreateStream(pconn *service.Connect, stream service.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     pconn.User.Id,
		active: true,
		// TODO: Reasearch if this error: make(chan error),
	}

	s.Connection = append(s.Connection, conn)
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
