package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/katoozi/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculateServiceClient(cc)

	// do unary services
	// doUnary(c)

	// do server streaming rpc
	doServerStreaming(c)
}

func doServerStreaming(c calculatorpb.CalculateServiceClient) {
	fmt.Println("starting to do server streaming rpc...")

	req := &calculatorpb.PrimeNumberCompositionRequest{
		Number: 40,
	}
	stream, err := c.PrimeNumberComposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While call PrimeNumberComposition: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while Read Stream: %v", err)
		}
		log.Printf("Response From Server. Prime Number: %d", msg.GetResult())
	}
}

func doUnary(c calculatorpb.CalculateServiceClient) {
	fmt.Println("starting to do unary rpc...")

	doSumCall(c)
	doMulCall(c)
	doDivCall(c)
	doSubCall(c)
}

func doSumCall(c calculatorpb.CalculateServiceClient) {
	request := &calculatorpb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	}
	res, err := c.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling Sum service: %v", err)
	}
	log.Println("Result: ", res.GetResult())
}

func doDivCall(c calculatorpb.CalculateServiceClient) {
	request := &calculatorpb.DivRequest{
		FirstNumber:  20,
		SecondNumber: 5,
	}
	res, err := c.Div(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling Div service: %v", err)
	}
	log.Println("Result: ", res.GetResult())
}

func doMulCall(c calculatorpb.CalculateServiceClient) {
	request := &calculatorpb.MultiplyRequest{
		FirstNumber:  20,
		SecondNumber: 5,
	}
	res, err := c.Multiply(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling Multiply service: %v", err)
	}
	log.Println("Result: ", res.GetResult())
}

func doSubCall(c calculatorpb.CalculateServiceClient) {
	request := &calculatorpb.SubRequest{
		FirstNumber:  20,
		SecondNumber: 5,
	}
	res, err := c.Sub(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling Multiply service: %v", err)
	}
	log.Println("Result: ", res.GetResult())
}