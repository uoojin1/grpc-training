package main

import (
	"context"
	pb "grpc-training/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with: %v\n", req)
	return &pb.SumResponse{
		Sum: req.Number1 + req.Number2,
	}, nil
}
