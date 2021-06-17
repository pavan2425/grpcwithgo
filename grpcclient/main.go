package main

import (
	pb "grpc/protofiles"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net error")
	}

	// Create a client
	c := pb.NewMoneyTransactionClient(conn)
	from := "1234"
	to := "5678"
	amount := float32(1250.75)
	// Make a server request.
	r, err := c.MakeTransaction(context.Background(),
		&pb.TransactionRequest{From: from,
			To: to, Amount: amount})
	log.Println("data", r)

}
