package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Philtoft/DISYS-Mini-Project-2-Chitty-Chat/time"
	"google.golang.org/grpc"
)

// På sigt: lav om til channels

func main() {
	// Creat a virtual RPC Client Connection on port  9080 WithInsecure (because  of http)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	defer conn.Close()

	//  Create new Client from generated gRPC code from proto
	// c := time.NewGetCurrentTimeClient(conn)
	c := time.NewChatClient(conn)

	for {
		// Needs to send message over channel
		// Needs to print message it receives over the channel
		SendChat(c)

	}
}

func SendChat(c time.ChatClient) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert message")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	message := time.Message{Message: text}

	_, err := c.SendChat(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling SendMessage: %s", err)
	}
}
