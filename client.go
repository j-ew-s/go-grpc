package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/j-ew-s/go-grpc/chat"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial failed : %v", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "CLIENT MESSSAGE",
	}

	resp, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when running SayHello : %v", err)
	}

	fmt.Println(&resp.Body)

}
