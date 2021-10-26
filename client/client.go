package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	t "time"

	"github.com/Philtoft/DIS-mini-project-1/time"
	"google.golang.org/grpc"
)

func main() {

	// Creates a virtual RPC Client Connection on port  9080 WithInsecure (because of http)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	defer conn.Close()

	//  Create new Client from generated gRPC code from proto
	c := time.NewGetCurrentTimeClient(conn)

	fmt.Println("Chat room")

	for {
		SendGetTimeRequest(c)
		t.Sleep(1 * t.Second)
	}
}

func SendGetTimeRequest(c time.GetCurrentTimeClient) {
	// Between the curly brackets are nothing, because the .proto file expects no input.
	message := time.GetTimeRequest{}

	response, err := c.GetTime(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetTime: %s", err)
	}

	fmt.Printf("Current time right now: %s\n", response.Reply)
}

func SendChatMessage(msg string, c time) {

}

func getChatInput() string {
	reader := bufio.NewReader(os.Stdin)
	chatMsg, _ := reader.ReadString('\n')
	chatMsg = strings.Replace(chatMsg, "\n", "", -1)

	return chatMsg
}

func GetMessages() {}
