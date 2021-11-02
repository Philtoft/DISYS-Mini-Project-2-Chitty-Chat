package main

import (
	"context"
	"log"
	"net"
	"os"
	"sync"

	service "github.com/Philtoft/DISYS-Mini-Project-2-Chitty-Chat/service"
	grpc "google.golang.org/grpc"
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
	error  chan error
}

type Server struct {
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
		error:  make(chan error),
	}

	s.Connection = append(s.Connection, conn)

	return <-conn.error

}

func (s *Server) BroadcastMessage(ctx context.Context, msg *service.Message) (*service.Close, error) {

	wait := sync.WaitGroup{}
	done := make(chan int)

	for _, conn := range s.Connection {
		wait.Add(1)

		go func(msg *service.Message, conn *Connection) {
			defer wait.Done()

			if conn.active {
				err := conn.stream.Send(msg)
				grpcLog.Info("Sending message to: ", conn.stream)

				if err != nil {
					grpcLog.Errorf("Error with Stream: %v - Error: %v", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)
	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return &service.Close{}, nil
}

func main() {
	var connections []*Connection

	server := &Server{connections}

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}

	grpcLog.Info("Starting server at port :8080")

	service.RegisterBroadcastServer(grpcServer, server)
	grpcServer.Serve(listener)

}
