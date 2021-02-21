package main

import (
	"context"
	"fmt"
	"io"
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

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Println("Calculator function was invoked with %v", req)

	number := req.GetNumber()
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			log.Printf("Devisor has been increased to %v", divisor)
		}

	}
	return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Println("Calculator function was invoked ")
	sum := int32(0)
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float64(sum) / float64(count)
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				NumberAverage: average,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		sum += req.GetNumber()
		count++
	}
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
