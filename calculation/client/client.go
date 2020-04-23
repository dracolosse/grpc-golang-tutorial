package main

import (
	"context"
	"fmt"
	"go-grpc-tutorial/calculation/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Starting a client...")
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())// remove in prod code
	if err != nil {
		log.Fatal("Could not connect to server at port 50051 due to: ", err)
	}
	defer conn.Close()

	client := proto.NewCalculationServiceClient(conn)

	multiplyByUnary(client)
}

func multiplyByUnary(client proto.CalculationServiceClient) {
	firstFactor, secondFactor := 50.55, 100.55

	factors := proto.Factors{
		FirstFactor: 50.55,
		SecondFactor: 100.55,
	}
	req := proto.FactorsRequest{Factors: &factors}

	res, err := client.Multiply(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling Multiply RPC: %v", err)
	}
	log.Printf("Result of multilication of %f * %f is %f", firstFactor, secondFactor, res.Product)
}
