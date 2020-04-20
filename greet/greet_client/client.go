package main

import (
	"context"
	"fmt"
	"go-grpc-tutorial/greet/greetpb"
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

	client := greetpb.NewGreetServiceClient(conn)

	greetByUnary(client)

	multiplyByUnary(client)
}

func greetByUnary(client greetpb.GreetServiceClient) {
	greeting := greetpb.Greeting{
		FirstName: "Truc",
		LastName:  "Nguyen",
	}
	req := greetpb.GreetRequest{Greeting: &greeting}
	res, err := client.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func multiplyByUnary(client greetpb.GreetServiceClient) {
	firstFactor, secondFactor := 50.55, 100.55

	factors := greetpb.Factors{
		FirstFactor: 50.55,
		SecondFactor: 100.55,
	}
	req := greetpb.FactorsRequest{Factors: &factors}

	res, err := client.Multiply(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling Multiply RPC: %v", err)
	}
	log.Printf("Result of multilication of %f * %f is %f", firstFactor, secondFactor, res.Product)
}
