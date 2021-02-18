package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-udemy/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculator(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Println("Calculator function was invoked with %v", req)
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()

	result := firstNumber + secondNumber

	res := &calculatorpb.SumResponse{
		SumResult: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello RPC Server calculaor")
	lis, err := net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		log.Fatalf("failed to listen %v", err)

	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
