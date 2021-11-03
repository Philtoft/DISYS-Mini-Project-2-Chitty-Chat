package main

import (
	"log"
	"net"
	"os"

	service "github.com/Philtoft/DISYS-Mini-Project-2-Chitty-Chat/service"
	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
)

/*
* TODO: Make connection into channels
* TODO: Apply Lamport timestamp
 */

type Server struct {
	service.UnimplementedBroadcastServer
	// Lav en key-value pair, hvor hver key har en række af channels som value, der sender data af typen *service.Message. Hvorfor?
	// Repræsenterer åbenbart hvert chatrum
	rooms map[string][]chan *service.Message
}

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

func main() {

	listener, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	grpcLog.Info("Starting server at port :9080")

	service.RegisterBroadcastServer(grpcServer, &Server{
		rooms: make(map[string][]chan *service.Message),
	})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}

func (s *Server) CreateStream() {}

func (s *Server) LeaveStream() {}

func (s *Server) BroadcastMessage() {}

// func (s *Server) CreateStream(pconn *service.Connect, stream service.Broadcast_CreateStreamServer) error {
// 	conn := &Connection{
// 		stream: stream,
// 		id:     pconn.User.Id,
// 		active: true,
// 		error:  make(chan error),
// 	}

// 	s.Connection = append(s.Connection, conn)

// 	return <-conn.error

// }

// func (s *Server) BroadcastMessage(ctx context.Context, msg *service.Message) (*service.Close, error) {

// 	wait := sync.WaitGroup{}
// 	done := make(chan int)

// 	for _, conn := range s.Connection {
// 		wait.Add(1)

// 		go func(msg *service.Message, conn *Connection) {
// 			defer wait.Done()

// 			if conn.active {
// 				err := conn.stream.Send(msg)
// 				grpcLog.Info("Sending message to: ", conn.stream)

// 				if err != nil {
// 					grpcLog.Errorf("Error with Stream: %v - Error: %v", conn.stream, err)
// 					conn.active = false
// 					conn.error <- err
// 				}
// 			}
// 		}(msg, conn)
// 	}

// 	go func() {
// 		wait.Wait()
// 		close(done)
// 	}()

// 	<-done
// 	return &service.Close{}, nil
// }
