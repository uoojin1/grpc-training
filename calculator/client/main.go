package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-training/calculator/proto"
)

const addr string = "localhost:50051"

func main() {
	// establish connection to server
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	// get sum;
	sum(c)
}
