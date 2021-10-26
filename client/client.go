package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	t "time"

	"github.com/Philtoft/DISYS-Mini-Project-2-Chitty-Chat/time"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	defer conn.Close()

	// c := time.NewGetCurrentTimeClient(conn)
	c := time.NewChatClient(conn)

	// Contact the server and print out its response
	msg := ""

	ctx, cancel := context.WithTimeout(context.Background(), t.Second)
	defer cancel()
	r, err := c.Broadcast(ctx, &time.ChatRequest{Message: msg})

	if err != nil {
		log.Fatalf("could not get chat message: %v", err)
	}
	log.Printf(r.GetMessage())
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

// func SendChatMessage(c time.ChatRequest) {
// 	message := time.Broadcast{}
// }

func getChatInput() string {
	reader := bufio.NewReader(os.Stdin)
	chatMsg, _ := reader.ReadString('\n')
	chatMsg = strings.Replace(chatMsg, "\n", "", -1)

	return chatMsg
}

func GetMessages() {}
