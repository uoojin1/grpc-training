package main

import (
	pb "grpc-training/calculator/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const addr string = "localhost:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("listening on %s\n", addr)

	// create new grpc server
	s := grpc.NewServer()
	// register calculator service server
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
