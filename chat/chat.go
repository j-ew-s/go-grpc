package chat

import (
	context "context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(c context.Context, m *Message) (*Message, error) {
	log.Printf("Received message from Client : %v", m.Body)
	return &Message{Body: "Its a message from Server"}, nil
}
