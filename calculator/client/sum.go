package main

import (
	"context"
	pb "grpc-training/calculator/proto"
	"log"
)

func sum(c pb.CalculatorServiceClient) {
	log.Println("sum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Number1: 15,
		Number2: 12,
	})
	if err != nil {
		log.Fatalf("Could not get sum: %v", err)
	}
	log.Printf("Sum: %d", res.Sum)
}
