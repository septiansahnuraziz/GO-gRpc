package main

import (
	"context"
	"fmt"
	"grpc-udemy/calculator/calculatorpb"
	"io"
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
	// doServerStreaming(c)
	doClientStreaming(c)
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
func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do PrimeDecomposition server streaming RPC ...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 12,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling calculaor PRC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("someting wrong %v", err)
		}

		fmt.Println(res.GetPrimeFactor())
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do ComputeAverage server streaming RPC ...")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while opneing stream; %v", err)
	}
	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		fmt.Printf("sanding number: %v", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %v", err)
	}

	log.Printf("The average is: %v\n", res.GetNumberAverage())
}
