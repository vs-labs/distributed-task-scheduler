package worker

import (
	"context"
	"log"
)

type Server struct{}

// mustEmbedUnimplementedWorkerServer implements WorkerServer. so that the compiler will complain if we don't implement all the methods
func (*Server) mustEmbedUnimplementedWorkerServer() {
	panic("unimplemented")
}

func (*Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	msg := in.GetName()
	log.Printf("Received: %v", msg)
	return &Message{Name: msg}, nil
}
