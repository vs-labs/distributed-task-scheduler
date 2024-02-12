package main

import (
	"log"
	"net"

	worker "github.com/vs-labs/distributed-task-scheduler/worker/job" // Replace with the correct import path for the worker package
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	workerServer := worker.Server{}

	grpcServer := grpc.NewServer()

	worker.RegisterWorkerServer(grpcServer, &workerServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
