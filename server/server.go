package main

import (
	"log"
	"net"
	"sync"
	utils "utils"

	service "service"

	"github.com/Philtoft/DISYS-Mini-Project-2-Chitty-Chat/service"
	"google.golang.org/grpc"
)

var grpcServer *grpc.Server
var serv *server
var done chan int

// Lamport clock
var t = &utils.Lamport{T: 0}

const (
	participated string = "%v joined the chat at %v"
	left         string = "%v has left the chat at %v"
)

type Connection struct {
	stream service.Broadcast_CreateStreamServer
	id     string
	user   *service.User
	active bool
	error  chan error
}

type Server struct {
	service.UnimplementedBroadcastServer
	connections []*Connection
	mu          sync.Mutex
}

func main() {

	listener, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Error creating the server %v", err)
	}

	done = make(chan int)

	grpcServer = grpc.NewServer()

	grpcServer = grpc.NewServer()

	serv = &Server{connections: make([]*Connection, 0)}

	service.RegisterBroadcastServer(grpcServer, serv)

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalf("Unknown error. :: %v", err)
		}
	}(listener)

	go func(listener net.Listener) {
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Fatalf("Failted to start gRPC server: %v", err)
		}
	}(listener)

	<-done

}

func (s *Server) CreateStream(cr *service.ConnectRequest, stream service.Broadcast_CreateStreamServer) error {
	t.MaxAndIncrement(1)

	connection := newConnection(cr.User.Id, stream, cr.User)

	s.addConnection(connection)
}

func (s *Server) DisconnectStream() {}

func newConnection(id string, stream service.Broadcast_CreateStreamServer, user *service.User) *Connection {
	return &Connection{
		id:     id,
		stream: stream,
		user:   user,
		active: false,
		error:  make(chan error),
	}
}

func (s *Server) addConnection(c *Connection) {
	s.mu.Lock()
	defer s.mu.Unlock()
}
