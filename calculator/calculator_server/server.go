package main

import (
	"context"
	"goGrpc-tutorial/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.Request) (*calculatorpb.Response, error) {
	numOne := req.GetNumOne()
	numTwo := req.GetNumTwo()
	resp := &calculatorpb.Response{
		Result: numOne + numTwo,
	}
	return resp, nil
}

func (*server) Minus(ctx context.Context, req *calculatorpb.Request) (*calculatorpb.Response, error) {
	numOne := req.GetNumOne()
	numTwo := req.GetNumTwo()
	resp := &calculatorpb.Response{
		Result: numOne - numTwo,
	}
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
