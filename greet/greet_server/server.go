package main

import (
	"context"
	"fmt"
	"go-grpc-tutorial/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Greet(ctx context.Context,req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	greeting := fmt.Sprintf("Hello %s, %s", lastName, firstName)
	res := greetpb.GreetResponse{Result: greeting}
	return &res, nil
}

func (*server) Multiply(ctx context.Context, req *greetpb.FactorsRequest) (*greetpb.FactorsResponse, error)  {
	firstFactor := req.GetFactors().GetFirstFactor()
	secondFactor := req.GetFactors().GetSecondFactor()

	res := greetpb.FactorsResponse{Product: firstFactor*secondFactor}
	return &res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen to the port 50051 due to: ", err)
	}
	fmt.Println("Server started successfully")

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	fmt.Println("Greeting service registered to gRPC server successfully")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}


