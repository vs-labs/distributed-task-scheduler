package main

import (
	"context"
	"log"
	"net/http"

	worker "github.com/vs-labs/distributed-task-scheduler/worker/job"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8081", nil)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	c := worker.NewWorkerClient(conn)
	mes, err := c.SayHello(context.Background(), &worker.Message{Name: "Hello"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", mes.Name)
}
