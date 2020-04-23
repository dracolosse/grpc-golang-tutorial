package main

import (
	"context"
	"fmt"
	"go-grpc-tutorial/calculation/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Multiply(ctx context.Context, req *proto.FactorsRequest) (*proto.FactorsResponse, error)  {
	firstFactor := req.GetFactors().GetFirstFactor()
	secondFactor := req.GetFactors().GetSecondFactor()

	res := proto.FactorsResponse{Product: firstFactor*secondFactor}
	return &res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen to the port 50051 due to: ", err)
	}
	fmt.Println("Server started successfully")

	s := grpc.NewServer()
	proto.RegisterCalculationServiceServer(s, &server{})

	fmt.Println("Greeting service registered to gRPC server successfully")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
