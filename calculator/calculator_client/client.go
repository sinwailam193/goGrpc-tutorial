package main

import (
	"context"
	"fmt"
	"goGrpc-tutorial/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Unary RPC ...")
	req := &calculatorpb.Request{
		NumOne: 3,
		NumTwo: 4,
	}
	sum, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Calculator RPC: %v", err)
	}
	log.Printf("Sum response from Calculator: %v", sum.Result)

	req = &calculatorpb.Request{
		NumOne: 3,
		NumTwo: 4,
	}
	minus, err := c.Minus(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Calculator RPC: %v", err)
	}
	log.Printf("Minus response from Calculator: %v", minus.Result)
}
