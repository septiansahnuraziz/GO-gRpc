package main

import (
	"context"
	"fmt"
	"grpc-udemy/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am Clien of Calculator")

	cc, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	// doUnary(c)
	doServerStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do unary RPC ...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  100,
		SecondNumber: 50,
	}
	res, err := c.Calculator(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling calculaor PRC: %v", err)
	}
	log.Printf("Response devide is %v", res.SumResult)
}
