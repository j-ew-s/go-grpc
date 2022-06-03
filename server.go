package main

import (
	"log"
	"net"

	"github.com/j-ew-s/go-grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Error when starting net listen on port 9000 : %v", err)
	}

	grpcServer := grpc.NewServer()

	chatService := chat.Server{}

	chat.RegisterChatServiceServer(grpcServer, &chatService)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Error start serving GRPC serve  : %v", err)
	}
}
